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
	"code.google.com/p/gocc/config"
	"code.google.com/p/gocc/gen"
	"code.google.com/p/gocc/lr1"
	parserDefs "code.google.com/p/gocc/lr1/parser"
	"code.google.com/p/gocc/parser"
	"code.google.com/p/gocc/scanner"
	"code.google.com/p/gocc/token"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	// "runtime/pprof"
	"time"
)

// var (
// 	allowUnreachable  *bool
// 	autoLRConfResolve *bool
// 	profile           *bool
// 	genScanner        *bool
// 	help              *bool
// 	verbose           *bool
// 	srcFile           string
// 	outDir            string
// 	pkg               string
// )

var cfg config.Config

func main() {
	if cfg1, err := config.New(); err != nil {
		fmt.Printf("Error reading configuration: %s\n", err)
		usage()
	} else {
		cfg = cfg1
	}

	if cfg.Help() {
		usage()
	}

	if cfg.Verbose() {
		cfg.PrintParams()
	}

	// if *profile {
	// 	startProfiler()
	// 	defer pprof.StopCPUProfile()
	// }

	tokenmap := token.NewMapFromStrings(token.GoccStrings)
	scanner := &scanner.Scanner{}
	srcBuffer, err := ioutil.ReadFile(cfg.SourceFile())
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

	writeStateMachine(cfg.OutDir(), g.FirstSets(), sets, trans)

	writeFirstBodies(cfg.OutDir(), g)

	gto := lr1.NewGotoTable(len(sets), g, trans)

	act, lr1Conflicts := sets.ActionTable(trans, cfg.AutoLRConfResolve())
	if cfg.Verbose() {
		printConflicts(lr1Conflicts, g)
	}
	switch {
	case cfg.AutoLRConfResolve() && len(lr1Conflicts) > 0:
		fmt.Fprintf(os.Stdout, "Resolved %d LR(1) conflicts\n", len(lr1Conflicts))
	case !cfg.AutoLRConfResolve() && len(lr1Conflicts) > 0:
		fmt.Fprintf(os.Stderr, "ABORTING: %d LR(1) conflicts\n", len(lr1Conflicts))
		os.Exit(1)
	}

	if err := 
		gen.WriteFiles(cfg, lr1.GenGo(act, gto, g), initDecl, g.TokenMap, act, gto, g); 
		err != nil {
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
	if len(missing) > 0 || len(unreachable) > 0 && !cfg.AllowUnreachable() {
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

// func getArgs() {
// 	wd, err := os.Getwd()
// 	if err != nil {
// 		fmt.Fprintln(os.Stderr, "Cannot get working directory. Error: ", err.Error())
// 		os.Exit(1)
// 	}

// 	allowUnreachable = flag.Bool("u", false, "allow unreachable productions")
// 	autoLRConfResolve = flag.Bool("a", false, "automatically resolve LR(1) conflicts")
// 	genScanner = flag.Bool("s", false, "generate a scanner")
// 	help = flag.Bool("h", false, "help")
// 	verbose = flag.Bool("v", false, "verbose")
// 	profile = flag.Bool("prof", false, "write profile to file")
// 	flag.StringVar(&outDir, "o", wd, "output dir.")
// 	flag.StringVar(&pkg, "p", defaultPackage(outDir), "package")
// 	flag.Parse()
// 	if len(flag.Args()) != 1 {
// 		errorMsg("too few arguments")
// 	}
// 	srcFile = flag.Arg(0)
// }

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

func printConflicts(conflicts lr1.ConflictSet, g *ast.Grammar) {
	for _, c := range conflicts {
		fmt.Fprintf(os.Stderr, "LR(1) conflict: %s\n", conflictString(c, g))
	}
}

func conflictString(c *lr1.Conflict, g *ast.Grammar) (res string) {
	if _, ok := c.A1.(parserDefs.Reduce); ok {
		if _, ok := c.A2.(parserDefs.Reduce); ok {
			res = fmt.Sprintf("S%d %s / %s", c.State, actionString(c.A1, g), actionString(c.A2, g))
		} else {
			res = fmt.Sprintf("S%d %s / %s", c.State, actionString(c.A2, g), actionString(c.A1, g))
		}
	} else {
		res = fmt.Sprintf("S%d %s / %s", c.State, actionString(c.A1, g), actionString(c.A2, g))
	}
	return
}

func printTime(from time.Time, msg string) {
	fmt.Printf("%s: elapsed time%s\n", msg, time.Since(from))
}

func actionString(a parserDefs.Action, g *ast.Grammar) string {
	if act, ok := a.(parserDefs.Shift); ok {
		return fmt.Sprintf("Shift:%d", int(act))
	}
	act := a.(parserDefs.Reduce)
	return fmt.Sprintf("Reduce:%d(%s)", act, g.Prod[act].Head.TokLit)
}

// func startProfiler() {
// 	f, err := os.Create("cpu.prof")
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "ABORT: cannot create cpu profile file, \"%s\"\n", err)
// 		os.Exit(1)
// 	}
// 	pprof.StartCPUProfile(f)
// }

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
