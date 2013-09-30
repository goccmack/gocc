//Copyright 2012 Vastech SA (PTY) LTD
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
	"code.google.com/p/gocc/ast"
	astrw "code.google.com/p/gocc/ast/rewrite"
	"code.google.com/p/gocc/config"
	"code.google.com/p/gocc/frontend/lexer"
	"code.google.com/p/gocc/frontend/parser"
	genlexer "code.google.com/p/gocc/lexer"
	genparser "code.google.com/p/gocc/parser"
	cfgsymbols "code.google.com/p/gocc/parser/symbols"
	"code.google.com/p/gocc/semantic"
	semerr "code.google.com/p/gocc/semantic/errors"
	gentokmap "code.google.com/p/gocc/token"
	gentoken "code.google.com/p/gocc/token/gen"
	genutil "code.google.com/p/gocc/util/gen"
	"flag"
	"fmt"
	"os"
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

	lex, err := lexer.NewLexerFile(cfg.SourceFile())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	parser := parser.NewParser()
	grammar, err := parser.Parse(lex)
	if err != nil {
		fmt.Printf("Parse error: %s\n", err)
		os.Exit(1)
	}
	g := grammar.(*ast.Grammar)
	basicSyntaxProds := astrw.BasicProds(g.SyntaxPart.ProdList)
	cfgSymbols := cfgsymbols.NewSymbols(basicSyntaxProds)
	g.LexPart.UpdateStringLitTokens(cfgSymbols.ListStringLitSymbols())
	tokenMap := gentokmap.NewTokenMap(cfgSymbols.ListTerminals())

	if errs := semantic.Check(g, basicSyntaxProds, cfgSymbols); errs != nil {
		reportSemanticErrors(errs)
	}

	genlexer.Gen(cfg, g.LexPart, tokenMap)
	genparser.Gen(cfg, basicSyntaxProds, cfgSymbols)

	// var tokenMap *outToken.TokenMap

	// io.WriteFileString(path.Join(cfg.OutDir(), "lex_prods.txt"), g.LexPart.String())
	// gSymbols.Add(g.LexPart.TokenIds()...)
	// g.LexPart.UpdateStringLitTokens(gSymbols.ListStringLitSymbols())
	// lexSets := lexItems.GetItemSets(g.LexPart, gSymbols.ListStringLitSymbols())
	// io.WriteFileString(path.Join(cfg.OutDir(), "lexer_sets.txt"), lexSets.String())
	// tokenMap = outToken.NewTokenMap(gSymbols.List())
	// genLexer.Gen(cfg.Package(), cfg.OutDir(), g.LexPart.Header.SDTLit, lexSets, tokenMap)

	gentoken.Gen(cfg.Package(), cfg.OutDir(), tokenMap)
	genutil.Gen(cfg.OutDir())

}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: gocc flags bnf_file\n")
	fmt.Fprintf(os.Stderr, "   - bnf_file: contains the BNF grammar\n")
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

func printTime(from time.Time, msg string) {
	fmt.Printf("%s: elapsed time%s\n", msg, time.Since(from))
}

func reportSemanticErrors(errs []*semerr.Error) {
	for _, err := range errs {
		fmt.Printf("Error: %s\n", err)
	}
}
