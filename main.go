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

package main

import (
	"code.google.com/p/gocc/ast"
	"code.google.com/p/gocc/gen"
	"code.google.com/p/gocc/lr1"
	"code.google.com/p/gocc/parser"
	"code.google.com/p/gocc/scanner"
	"code.google.com/p/gocc/token"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

var (
	allowUnreachable *bool
	debug            *bool
	dumpGrammar      *bool
	dumpFirst        *bool
	dumpSets         *bool
	genScanner       *bool
	srcFile          string
	srcDir           string
	pkg              string
)

func main() {
	getArgs()

	tokenmap := token.NewMapFromStrings(token.GoccStrings)
	scanner := &scanner.Scanner{}
	srcBuffer, err := ioutil.ReadFile(srcFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	scanner.Init(srcBuffer, tokenmap)
	parser := parser.NewParser(scanner, tokenmap)
	g, err := ast.NewAugmentedGrammar(parser.Parse())
	if *dumpGrammar {
		fmt.Println("Grammar:")
		fmt.Println(g)
		fmt.Println("g.TokenMap:")
		fmt.Println(g.TokenMap)
	}
	if err != nil {
		error1("Error creating augmented grammar:", err)
	}
	checkGrammar(g)

	sets, trans := lr1.GetItemSets(g)
	if !trans.Check() {
		error1("transition table not same", nil)
	}
	writeStateMachine(srcDir, g.FirstSets(), sets, trans)
	writeFirstBodies(srcDir, g)
	if *dumpFirst {
		fmt.Println("\nFirstSets:")
		fmt.Println(g.FirstSets())
		fmt.Println()
	}
	if *dumpSets {
		fmt.Println("Sets:\n" + sets.String())
		fmt.Println("Transitions:\n" + trans.String())
	}

	gto := lr1.NewGotoTable(len(sets), g, trans)
	act := sets.ActionTable(g)

	if err := gen.WriteFiles(*debug, srcDir, pkg, prjName(srcFile), lr1.GenGo(act, gto, g), g.TokenMap, *genScanner); err != nil {
		error1("Error generating files:", err)
	}
}

func checkGrammar(g *ast.Grammar) {
	unreachable, missing := g.CheckProductions()
	if len(unreachable) > 0 {
		fmt.Println("Unreachable productions:")
		for _, p := range unreachable {
			fmt.Println("	", p.TokLit)
		}
	}
	if len(missing) > 0 {
		fmt.Println("Missing productions:", missing)
		fmt.Println()
	}
	if len(missing) > 0 || len(unreachable) > 0 && !*allowUnreachable {
		error1("Errors in grammar", nil)
	}
}

func getArgs() {
	allowUnreachable = flag.Bool("u", false, "allow unreachable productions")
	debug = flag.Bool("d", false, "debug")
	dumpGrammar = flag.Bool("g", false, "dump grammar")
	dumpFirst = flag.Bool("f", false, "write first sets")
	dumpSets = flag.Bool("s", false, "write item sets")
	genScanner = flag.Bool("scanner", false, "generate a scanner")
	flag.StringVar(&srcDir, "o", "", "src output dir")
	flag.StringVar(&pkg, "p", "", "package root")
	flag.Parse()
	if len(flag.Args()) != 1 {
		errorMsg("too few arguments")
	}
	srcFile = flag.Arg(0)
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: gocc flags bnf_file")
	fmt.Fprintf(os.Stderr, "   - bnf_file: contains the BNF grammar")
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

func prjName(bnfFile string) string {
	_, file := path.Split(bnfFile)
	return file[:strings.LastIndex(file, ".")]
}

func writeFirstBodies(prjDir string, g *ast.Grammar) {
	first := make(map[string]string)
	for _, p := range g.Prod {
		for i := range p.Body.Symbols {
			syms := p.Body.Symbols[i:]
			fst := g.FirstS(syms)
			if fst1, exists := first[syms.String()]; exists && fst1 != fst.String() {
				fmt.Println("gocc.writeFirstBodies: duplicates")
				fmt.Println("	", syms.String())
				fmt.Println("		", fst.String())
				fmt.Println("		", fst1)
				os.Exit(1)
			} else {
				first[syms.String()] = fst.String()
			}
		}
	}
	str := ""
	for k, v := range first {
		str += k + "\n"
		str += "	" + v + "\n"
	}
	writeStringToFile(prjDir, "sm_first_bodies.txt", str)
}

func writeStateMachine(prjDir string, first ast.FirstSets, sets lr1.ItemSets, trans lr1.TransitionTable) {
	writeStringToFile(prjDir, "sm_first.txt", first.String())
	writeStringToFile(prjDir, "sm_sets.txt", sets.String())
	writeStringToFile(prjDir, "sm_transitions.dot", trans.Dot())
}

func writeStringToFile(prjDir, file, data string) {
	if err := os.MkdirAll(prjDir, 0766); err != nil {
		fmt.Println("Error creating", prjDir)
		os.Exit(1)
	}
	fname := path.Join(prjDir, file)
	if err := ioutil.WriteFile(fname, []byte(data), 0666); err != nil {
		fmt.Println("Could not write", fname, err)
	}
}
