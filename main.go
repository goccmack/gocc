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
//Please see https://code.google.com/p/gocc/ for more documentation.
package main

import (
	"bytes"
	"code.google.com/p/gocc/ast"
	"code.google.com/p/gocc/config"
	"code.google.com/p/gocc/frontend/parser"
	"code.google.com/p/gocc/frontend/scanner"
	"code.google.com/p/gocc/frontend/token"
	"code.google.com/p/gocc/io"
	genLexer "code.google.com/p/gocc/lexer/gen/golang"
	lexItems "code.google.com/p/gocc/lexer/items"
	"code.google.com/p/gocc/parser/first"
	lr1Items "code.google.com/p/gocc/parser/lr1/items"
	"code.google.com/p/gocc/parser/symbols"
	outToken "code.google.com/p/gocc/token"
	genToken "code.google.com/p/gocc/token/gen"
	genUtil "code.google.com/p/gocc/util/gen"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	// "runtime/pprof"
	genParser "code.google.com/p/gocc/parser/gen"
	"time"
)

var cfg config.Config

func main() {
	if cfg1, err := config.New(); err != nil {
		fmt.Printf("Error reading configuration: %s\n", err)
		usage()
	} else {
		cfg = cfg1
	}

	if cfg.Verbose() {
		cfg.PrintParams()
	}

	if cfg.Help() {
		usage()
	}

	// if *profile {
	// 	startProfiler()
	// 	defer pprof.StopCPUProfile()
	// }

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

	var tokenMap *outToken.TokenMap

	gSymbols.Add(g.LexPart.TokenIds()...)
	g.LexPart.UpdateStringLitTokens(gSymbols.ListStringLitSymbols())
	lexSets := lexItems.GetItemSets(g.LexPart)
	io.WriteFileString(path.Join(cfg.OutDir(), "lexer_sets.txt"), lexSets.String())
	tokenMap = outToken.NewTokenMap(gSymbols.List())
	genLexer.Gen(cfg.Package(), cfg.OutDir(), g.LexPart.Header.SDTLit, lexSets, tokenMap, cfg)

	if g.SyntaxPart != nil {
		firstSets := first.GetFirstSets(g, gSymbols)
		io.WriteFileString(path.Join(cfg.OutDir(), "first.txt"), firstSets.String())

		lr1Sets := lr1Items.GetItemSets(g, gSymbols, firstSets)
		io.WriteFileString(path.Join(cfg.OutDir(), "LR1_sets.txt"), lr1Sets.String())

		conflicts := genParser.Gen(cfg.Package(), cfg.OutDir(), g.SyntaxPart.Header.SDTLit, g.SyntaxPart.ProdList, gSymbols, lr1Sets, tokenMap, cfg)
		handleConflicts(conflicts, lr1Sets.Size(), cfg)
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

func handleConflicts(conflicts map[int]lr1Items.RowConflicts, numSets int, cfg config.Config) {
	if len(conflicts) <= 0 {
		return
	}
	io.WriteFileString(path.Join(cfg.OutDir(), "LR1_conflicts.txt"), conflictString(conflicts, numSets))
	switch {
	case !cfg.AutoResolveLRConf():
		error1(fmt.Sprintf("Error: %d LR-1 conflicts\n", len(conflicts)), nil)
	case cfg.Verbose():
		fmt.Printf("%d LR-1 conflicts \n", len(conflicts))
	}
}

func conflictString(conflicts map[int]lr1Items.RowConflicts, numSets int) string {
	w := new(bytes.Buffer)
	fmt.Fprintf(w, "%d LR-1 conflicts: \n", len(conflicts))
	for i := 0; i < numSets; i++ {
		if cnf, exist := conflicts[i]; exist {
			fmt.Fprintf(w, "\tS%d\n", i)
			for sym, conflicts := range cnf {
				fmt.Fprintf(w, "\t\tsymbol: %s\n", sym)
				for _, cflct := range conflicts {
					fmt.Fprintf(w, "\t\t\t%s\n", cflct)
				}
			}
		}
	}
	return w.String()
}

func printTime(from time.Time, msg string) {
	fmt.Printf("%s: elapsed time%s\n", msg, time.Since(from))
}

// func startProfiler() {
// 	f, err := os.Create("cpu.prof")
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "ABORT: cannot create cpu profile file, \"%s\"\n", err)
// 		os.Exit(1)
// 	}
// 	pprof.StartCPUProfile(f)
// }
