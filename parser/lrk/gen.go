package lrk

import (
	"bytes"
	"code.google.com/p/gocc/ast"
	"code.google.com/p/gocc/config"
	"code.google.com/p/gocc/io"
	"code.google.com/p/gocc/parser/first"
	"code.google.com/p/gocc/parser/lrk/items"
	"code.google.com/p/gocc/parser/lrk/knuth"
	"code.google.com/p/gocc/parser/lrk/pgm"
	"code.google.com/p/gocc/parser/lrk/states"
	"code.google.com/p/gocc/parser/symbols"
	"fmt"
	"os"
	"path"
)

func Gen(cfg config.Config, prods []*ast.SyntaxBasicProd, symbols *symbols.Symbols) {
	removeOldFiles(cfg)
	items := items.NewItems(prods)
	first := first.New(symbols, prods)
	pgmStates := pgm.States(symbols, items, first)
	var knuthStates *states.States
	if cfg.Knuth() {
		knuthStates = knuth.States(symbols, items, first)
	}

	if cfg.Verbose2() {
		writeBasicProductions(cfg, prods)
		io.WriteFileString(path.Join(cfg.OutDir(), "CFG_items.txt"), items.String())
		writeCFGSymbols(cfg, symbols)
		io.WriteFileString(path.Join(cfg.OutDir(), "first.txt"), first.String())
		io.WriteFileString(path.Join(cfg.OutDir(), "PGM_states.txt"), pgmStates.String())
		if cfg.Knuth() {
			io.WriteFileString(path.Join(cfg.OutDir(), "knuth_lr1_states.txt"), knuthStates.String())
		}
	}

}

// func handleConflicts(conflicts map[int]lr1Items.RowConflicts, numSets int, cfg config.Config) {
// 	if len(conflicts) <= 0 {
// 		return
// 	}
// 	if !cfg.AutoResolveLRConf() {
// 		error1(fmt.Sprintf("Error: %s\n", conflictString(conflicts, numSets)), nil)
// 	}
// }

// func conflictString(conflicts map[int]lr1Items.RowConflicts, numSets int) string {
// 	w := new(bytes.Buffer)
// 	fmt.Fprintf(w, "LR1 conflicts: \n")
// 	for i := 0; i < numSets; i++ {
// 		if cnf, exist := conflicts[i]; exist {
// 			fmt.Fprintf(w, "\tS%d\n", i)
// 			for sym, conflicts := range cnf {
// 				fmt.Fprintf(w, "\t\tsymbol: %s\n", sym)
// 				for _, cflct := range conflicts {
// 					fmt.Fprintf(w, "\t\t\t%s\n", cflct)
// 				}
// 			}
// 		}
// 	}
// 	return w.String()
// }

func removeOldFiles(cfg config.Config) {
	os.Remove(path.Join(cfg.OutDir(), "basic_productions.txt"))
	os.Remove(path.Join(cfg.OutDir(), "CFG_items.txt"))
	os.Remove(path.Join(cfg.OutDir(), "CFG_symbols.txt"))
	os.Remove(path.Join(cfg.OutDir(), "first.txt"))
	os.Remove(path.Join(cfg.OutDir(), "knuth_lr1_states.txt"))
	os.Remove(path.Join(cfg.OutDir(), "PGM_states.txt"))
}

func writeBasicProductions(cfg config.Config, prods []*ast.SyntaxBasicProd) {
	w := new(bytes.Buffer)
	for _, prod := range prods {
		fmt.Fprintf(w, "%s\n\n", prod)
	}
	io.WriteFileString(path.Join(cfg.OutDir(), "basic_productions.txt"), w.String())
}

func writeCFGSymbols(cfg config.Config, symbols *symbols.Symbols) {
	w := new(bytes.Buffer)
	fmt.Fprintf(w, "Start Symbol (S): %s\n\n", symbols.StartSymbol)
	fmt.Fprintf(w, "Terminal Symbols (T):\n")
	for _, t := range symbols.ListTerminals() {
		fmt.Fprintf(w, "\t%s\n", t)
	}
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "Non-terminal Symbols (NT):\n")
	for _, nt := range symbols.ListNonTerminals() {
		fmt.Fprintf(w, "\t%s\n", nt)
	}
	io.WriteFileString(path.Join(cfg.OutDir(), "CFG_symbols.txt"), w.String())
}
