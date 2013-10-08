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

//Gocc is a compiler kit for go written in go.
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
	errs := []error{}

	g := grammar.(*ast.Grammar)
	var basicSyntaxProds *ast.SyntaxBasicProdsList
	if g.SyntaxPart != nil {
		basicSyntaxProds = astrw.BasicProds(g.SyntaxPart.ProdList)
	}

	cfgSymbols, serr := cfgsymbols.NewSymbols(g.LexPart, basicSyntaxProds)
	errs = append(errs, serr...)

	g.LexPart.UpdateStringLitTokens(cfgSymbols.ListStringLitSymbols())
	tokenMap := gentokmap.NewTokenMap(cfgSymbols.ListTerminals())

	genlexer.Gen(cfg, g.LexPart, tokenMap)

	if g.SyntaxPart != nil {
		if serrs := semantic.Check(g, basicSyntaxProds, cfgSymbols); serrs != nil {
			errs = append(errs, serrs...)
		}
		errs = append(errs,
			genparser.Gen(cfg, basicSyntaxProds, cfgSymbols, g.SyntaxPart.Header.Lit)...)
	}

	gentoken.Gen(cfg.Package(), cfg.OutDir(), tokenMap)
	genutil.Gen(cfg.OutDir())

	if len(errs) > 0 {
		report(errs)
		os.Exit(1)
	}
	os.Exit(0)
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: gocc flags bnf_file\n")
	fmt.Fprintf(os.Stderr, "   - bnf_file: contains the BNF grammar\n")
	flag.PrintDefaults()
	os.Exit(1)
}

func printTime(from time.Time, msg string) {
	fmt.Printf("%s: elapsed time%s\n", msg, time.Since(from))
}

func report(errs []error) {
	for _, err := range errs {
		fmt.Printf("Error: %s\n", err)
	}
}
