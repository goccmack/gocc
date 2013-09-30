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
	actions, conflicts := Actions(pgmStates, symbols)
	handleConflicts(conflicts, pgmStates, cfg)

	if cfg.Verbose2() {
		writeBasicProductions(cfg, prods)
		io.WriteFileString(path.Join(cfg.OutDir(), "CFG_items.txt"), items.String())
		writeCFGSymbols(cfg, symbols)
		io.WriteFileString(path.Join(cfg.OutDir(), "first.txt"), first.String())
		io.WriteFile(path.Join(cfg.OutDir(), "PGM_states.txt"), statesString(pgmStates, actions, symbols))
		if cfg.Knuth() {
			io.WriteFileString(path.Join(cfg.OutDir(), "knuth_lr1_states.txt"), knuthStates.String())
		}
	}

}

func handleConflicts(conflicts [][]*Conflict, states *states.States, cfg config.Config) {
	if len(conflicts) <= 0 {
		return
	}
	fmt.Printf("LR(1) conflicts\n")
	for si, sc := range conflicts {
		for _, c := range sc {
			fmt.Printf("\tS%d: %s\n", si, c)
		}
	}
	// if !cfg.AutoResolveLRConf() {
	// 	error1(fmt.Sprintf("Error: %s\n", conflictString(conflicts, numSets)), nil)
	// }
}

func numConflicts(conflicts [][]*Conflict) (num int) {
	for _, sc := range conflicts {
		for _, c := range sc {
			if c != nil {
				num += c.NumConflicts()
			}
		}
	}
	return
}

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

func statesString(states *states.States, actions []map[string]Action, symbols *symbols.Symbols) []byte {
	w := new(bytes.Buffer)
	for si, state := range states.List {
		fmt.Fprintf(w, "%s", state)
		fmt.Fprintf(w, "Actions:\n")
		for _, symT := range symbols.ListTerminals() {
			if action := actions[si][symT]; action != nil {
				fmt.Fprintf(w, "\t%s: %s\n", symT, action)
			}
		}
		fmt.Fprintln(w)
	}
	return w.Bytes()
}

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
