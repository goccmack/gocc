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
	"strings"

	"github.com/maxcalandrelli/gocc/internal/ast"
	genBase "github.com/maxcalandrelli/gocc/internal/base/gen"
	"github.com/maxcalandrelli/gocc/internal/config"
	oldparser "github.com/maxcalandrelli/gocc/internal/frontend/stock/parser"
	oldscanner "github.com/maxcalandrelli/gocc/internal/frontend/stock/scanner"
	oldtoken "github.com/maxcalandrelli/gocc/internal/frontend/stock/token"
	"github.com/maxcalandrelli/gocc/internal/io"
	genIo "github.com/maxcalandrelli/gocc/internal/io/gen"
	genLexer "github.com/maxcalandrelli/gocc/internal/lexer/gen/golang"
	lexItems "github.com/maxcalandrelli/gocc/internal/lexer/items"
	"github.com/maxcalandrelli/gocc/internal/parser/first"
	genParser "github.com/maxcalandrelli/gocc/internal/parser/gen"
	lr1Action "github.com/maxcalandrelli/gocc/internal/parser/lr1/action"
	lr1Items "github.com/maxcalandrelli/gocc/internal/parser/lr1/items"
	"github.com/maxcalandrelli/gocc/internal/parser/symbols"
	outToken "github.com/maxcalandrelli/gocc/internal/token"
	genToken "github.com/maxcalandrelli/gocc/internal/token/gen"
	genUtil "github.com/maxcalandrelli/gocc/internal/util/gen"
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
		fmt.Fprintf(os.Stderr, "gocc base version\n")
		flag.Usage()
	}

	config.CurrentConfiguration = cfg

	srcBuffer, err := ioutil.ReadFile(cfg.SourceFile())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var (
		grammar interface{}
	)
	ast.StringGetter = func(v interface{}) string { return string(v.(*oldtoken.Token).Lit) }
	scanner := &oldscanner.Scanner{}
	scanner.Init(srcBuffer, oldtoken.FRONTENDTokens)
	parser := oldparser.NewParser(oldparser.ActionTable, oldparser.GotoTable, oldparser.ProductionsTable, oldtoken.FRONTENDTokens)
	grammar, err = parser.Parse(scanner)
	if err != nil {
		fmt.Printf("Parse error: %s\n", err)
		os.Exit(1)
	}

	outdir_base := cfg.OutDir()
	subpath := "internal"
	outdir_log := path.Join(outdir_base, "log")
	g := grammar.(*ast.Grammar)

	gSymbols := symbols.NewSymbols(g)
	if cfg.Verbose() {
		writeTerminals(gSymbols, cfg, outdir_log)
	}

	var tokenMap *outToken.TokenMap

	gSymbols.Add(g.LexPart.TokenIds()...)
	g.LexPart.UpdateStringLitTokens(gSymbols.ListStringLitSymbols())
	lexSets := lexItems.GetItemSets(g.LexPart)
	if cfg.Verbose() {
		io.WriteFileString(path.Join(outdir_log, "lexer_sets.txt"), lexSets.String())
	}
	tokenMap = outToken.NewTokenMap(gSymbols.ListTerminals())
	if !cfg.NoLexer() {
		genLexer.Gen(cfg.Package(), outdir_base, g.LexPart.Header.SDTLit, lexSets, tokenMap, cfg, subpath)
	}

	if g.SyntaxPart != nil {
		firstSets := first.GetFirstSets(g, gSymbols)
		if cfg.Verbose() {
			io.WriteFileString(path.Join(outdir_log, "first.txt"), firstSets.String())
		}

		lr1Sets := lr1Items.GetItemSets(g, gSymbols, firstSets)
		if cfg.Verbose() {
			io.WriteFileString(path.Join(outdir_log, "LR1_sets.txt"), lr1Sets.String())
		}

		conflicts := genParser.Gen(cfg.Package(), outdir_base, g.SyntaxPart.Header.SDTLit, g.SyntaxPart.ProdList, gSymbols, lr1Sets, tokenMap, cfg, subpath)
		handleConflicts(conflicts, lr1Sets.Size(), cfg, g.SyntaxPart.ProdList, outdir_log)
	}

	genToken.Gen(cfg.Package(), outdir_base, tokenMap, subpath)
	genUtil.Gen(outdir_base, subpath)
	genBase.Gen(cfg.Package(), outdir_base, subpath)
	genIo.Gen(cfg.Package(), outdir_base, subpath)
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: gocc flags bnf_file\n\n")
	fmt.Fprintf(os.Stderr, "  bnf_file: contains the BNF grammar\n\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
	os.Exit(1)
}

func handleConflicts(conflicts map[int]lr1Items.RowConflicts, numSets int, cfg config.Config, prods ast.SyntaxProdList, outdir string) {
	if len(conflicts) <= 0 {
		return
	}
	fmt.Printf("%d LR-1 conflicts \n", len(conflicts))
	if cfg.Verbose() {
		io.WriteFileString(path.Join(outdir, "LR1_conflicts.txt"), conflictString(conflicts, numSets, prods))
	}
	if !cfg.AutoResolveLRConf() {
		os.Exit(1)
	}
}

func conflictString(conflicts map[int]lr1Items.RowConflicts, numSets int, prods ast.SyntaxProdList) string {
	w := new(strings.Builder)
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

func writeTerminals(gSymbols *symbols.Symbols, cfg config.Config, outdir string) {
	buf := new(bytes.Buffer)
	for _, t := range gSymbols.ListTerminals() {
		fmt.Fprintf(buf, "%s\n", t)
	}
	io.WriteFile(path.Join(outdir, "terminals.txt"), buf.Bytes())
}
