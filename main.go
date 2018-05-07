//Copyright 2013 Vastech SA (PTY) LTD
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

//Gocc is LR1 parser generator for go written in go. The generator uses a BNF with very easy to use SDT rules.
//Please see https://github.com/goccmack/gocc/ for more documentation.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/goccmack/gocc/internal/ast"
	"github.com/goccmack/gocc/internal/config"
	"github.com/goccmack/gocc/internal/frontend/parser"
	"github.com/goccmack/gocc/internal/frontend/scanner"
	"github.com/goccmack/gocc/internal/frontend/token"
	"github.com/goccmack/gocc/internal/io"
	genLexer "github.com/goccmack/gocc/internal/lexer/gen/golang"
	lexItems "github.com/goccmack/gocc/internal/lexer/items"
	"github.com/goccmack/gocc/internal/parser/first"
	"github.com/goccmack/gocc/internal/parser/follow"
	lr1Action "github.com/goccmack/gocc/internal/parser/lr1/action"
	lr1Items "github.com/goccmack/gocc/internal/parser/lr1/items"
	"github.com/goccmack/gocc/internal/parser/symbols"
	outToken "github.com/goccmack/gocc/internal/token"
	genToken "github.com/goccmack/gocc/internal/token/gen"
	genUtil "github.com/goccmack/gocc/internal/util/gen"
	// "runtime/pprof"
	"time"

	genParser "github.com/goccmack/gocc/internal/parser/gen"
)

func main() {
	flag.Usage = usage
	cfg, err := config.New()
	if err != nil {
		fmt.Printf("Error reading configuration: %s\n", err)
		flag.Usage()
	}

	if cfg.Verbose() {
		cfg.PrintParams()
	}

	if cfg.Help() {
		flag.Usage()
	}

	scanner := &scanner.Scanner{}
	srcBuffer, err := ioutil.ReadFile(cfg.SourceFile())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	scanner.Init(srcBuffer, token.FRONTENDTokens)
	parser := parser.NewParser(parser.ActionTable, parser.GotoTable, parser.ProductionsTable, token.FRONTENDTokens)
	grammar, err := parser.Parse(scanner)
	if err != nil {
		fmt.Printf("Parse error: %s\n", err)
		os.Exit(1)
	}

	g := grammar.(*ast.Grammar)

	gSymbols := symbols.NewSymbols(g)
	if cfg.Verbose() {
		writeTerminals(gSymbols, cfg)
	}

	var tokenMap *outToken.TokenMap

	gSymbols.Add(g.LexPart.TokenIds()...)
	g.LexPart.UpdateStringLitTokens(gSymbols.ListStringLitSymbols())
	lexSets := lexItems.GetItemSets(g.LexPart)
	if cfg.Verbose() {
		io.WriteFileString(path.Join(cfg.OutDir(), "lexer_sets.txt"), lexSets.String())
	}
	tokenMap = outToken.NewTokenMap(gSymbols.ListTerminals())
	if !cfg.NoLexer() {
		genLexer.Gen(cfg.Package(), cfg.OutDir(), g.LexPart.Header.SDTLit, lexSets, tokenMap, cfg)
	}

	if g.SyntaxPart != nil {
		firstSets := first.GetFirstSets(g, gSymbols)
		if cfg.Verbose() {
			io.WriteFileString(path.Join(cfg.OutDir(), "first.txt"), firstSets.String())
		}

		followSets := follow.GetFollowSets(g, gSymbols, firstSets)
		if cfg.Verbose() {
			io.WriteFileString(path.Join(cfg.OutDir(), "follow.txt"), followSets.String())
		}

		lr1Sets := lr1Items.GetItemSets(g, gSymbols, firstSets)
		if cfg.Verbose() {
			io.WriteFileString(path.Join(cfg.OutDir(), "LR1_sets.txt"), lr1Sets.String())
		}

		conflicts := genParser.Gen(cfg.Package(), cfg.OutDir(), g.SyntaxPart.Header.SDTLit, g.SyntaxPart.ProdList, gSymbols, lr1Sets, tokenMap, cfg)
		handleConflicts(conflicts, lr1Sets.Size(), cfg, g.SyntaxPart.ProdList)
	}

	genToken.Gen(cfg.Package(), cfg.OutDir(), tokenMap)
	genUtil.Gen(cfg.OutDir())

}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: gocc flags bnf_file\n\n")
	fmt.Fprintf(os.Stderr, "  bnf_file: contains the BNF grammar\n\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
	os.Exit(1)
}

func errorMsg(msg string) {
	fmt.Println("Error:", msg)
	flag.Usage()
}

func error1(msg string, err error) {
	fmt.Println(msg, err)
	os.Exit(1)
}

func handleConflicts(conflicts map[int]lr1Items.RowConflicts, numSets int, cfg config.Config, prods ast.SyntaxProdList) {
	if len(conflicts) <= 0 {
		return
	}
	fmt.Printf("%d LR-1 conflicts \n", len(conflicts))
	if cfg.Verbose() {
		io.WriteFileString(path.Join(cfg.OutDir(), "LR1_conflicts.txt"), conflictString(conflicts, numSets, prods))
	}
	if !cfg.AutoResolveLRConf() {
		os.Exit(1)
	}
}

func conflictString(conflicts map[int]lr1Items.RowConflicts, numSets int, prods ast.SyntaxProdList) string {
	w := new(bytes.Buffer)
	fmt.Fprintf(w, "%d LR-1 conflicts: \n", len(conflicts))
	for i := 0; i < numSets; i++ {
		if cnf, exist := conflicts[i]; exist {
			fmt.Fprintf(w, "\tS%d\n", i)
			for sym, conflicts := range cnf {
				fmt.Fprintf(w, "\t\tsymbol: %s\n", sym)
				for _, cflct := range conflicts {
					switch c := cflct.(type) {
					case lr1Action.Reduce:
						fmt.Fprintf(w, "\t\t\tReduce(%d:%s)\n", c, prods[c])
					case lr1Action.Shift:
						fmt.Fprintf(w, "\t\t\t%s\n", cflct)
					default:
						panic(fmt.Sprintf("unexpected type of action: %s", cflct))
					}
				}
			}
		}
	}
	return w.String()
}

func printTime(from time.Time, msg string) {
	fmt.Printf("%s: elapsed time%s\n", msg, time.Since(from))
}

func writeTerminals(gSymbols *symbols.Symbols, cfg config.Config) {
	buf := new(bytes.Buffer)
	for _, t := range gSymbols.ListTerminals() {
		fmt.Fprintf(buf, "%s\n", t)
	}
	io.WriteFile(path.Join(cfg.OutDir(), "terminals.txt"), buf.Bytes())
}
