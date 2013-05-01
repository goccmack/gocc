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
	allowUnreachable  *bool
	autoLRConfResolve *bool
	genScanner        *bool
	srcFile           string
	srcDir            string
	pkg               string
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
	initDecl, prods, tm := parser.Parse()
	g, err := ast.NewAugmentedGrammar(prods, tm)
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

	gto := lr1.NewGotoTable(len(sets), g, trans)
	act, srConflicts, rrConflicts := sets.ActionTable(g, *autoLRConfResolve)
	switch {
	case *autoLRConfResolve && (srConflicts > 0 || rrConflicts > 0):
		fmt.Fprintf(os.Stdout, "Resolved %d shift/reduce, %d reduce/reduce conflicts\n", srConflicts, rrConflicts)
	case !*autoLRConfResolve && (srConflicts > 0 || rrConflicts > 0):
		fmt.Fprintf(os.Stderr, "ABORTING: %d shift/reduce, %d reduce/reduce conflicts\n", srConflicts, rrConflicts)
		os.Exit(1)
	}

	if err := gen.WriteFiles(srcDir, pkg, prjName(pkg), lr1.GenGo(act, gto, g), initDecl, g.TokenMap, *genScanner); err != nil {
		error1("Error generating files:", err)
	}
}

func checkGrammar(g *ast.Grammar) {
	checkFirstSets(g)
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

func checkFirstSets(g *ast.Grammar) {
	prods := ast.SymbolS{}
	for _, p := range g.Prod {
		prods = prods.AddSetElement(p.Head)
	}
	missingFirstSets := ast.SymbolS{}
	for _, p := range prods {
		if g.FirstSets().GetSet(p) == nil {
			missingFirstSets = missingFirstSets.AddSetElement(p)
		}
	}
	if len(missingFirstSets) > 0 {
		fmt.Fprintln(os.Stderr, "Error - Missing first sets:")
		for i, p := range missingFirstSets {
			fmt.Fprintf(os.Stderr, "%d: %s\n", i, p.TokLit)
		}
		os.Exit(1)
	}
}

func getArgs() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot get working directory. Error: ", err.Error())
		os.Exit(1)
	}

	allowUnreachable = flag.Bool("u", false, "allow unreachable productions")
	autoLRConfResolve = flag.Bool("a", false, "automatically resolve LR(1) conflicts")
	genScanner = flag.Bool("s", false, "generate a scanner")
	flag.StringVar(&srcDir, "o", wd, "output dir.")
	flag.StringVar(&pkg, "p", defaultPackage(srcDir), "package")
	flag.Parse()
	if len(flag.Args()) != 1 {
		errorMsg("too few arguments")
	}
	srcFile = flag.Arg(0)
}

func defaultPackage(wd string) string {
	srcPath := path.Join(os.Getenv("GOPATH"), "src")
	pkg := strings.Replace(wd, srcPath, "", -1)
	if strings.HasPrefix(pkg, "/") {
		pkg = pkg[1:]
	}
	return pkg
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

func prjName(pkg string) string {
	_, file := path.Split(pkg)
	return file
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
