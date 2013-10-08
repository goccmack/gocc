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

package lrk

import (
	"bytes"
	"code.google.com/p/gocc/ast"
	"code.google.com/p/gocc/config"
	"code.google.com/p/gocc/io"
	"code.google.com/p/gocc/parser/first"
	"code.google.com/p/gocc/parser/lrk/action"
	"code.google.com/p/gocc/parser/lrk/codegen"
	"code.google.com/p/gocc/parser/lrk/items"
	"code.google.com/p/gocc/parser/lrk/knuth"
	"code.google.com/p/gocc/parser/lrk/pgm"
	"code.google.com/p/gocc/parser/lrk/states"
	"code.google.com/p/gocc/parser/symbols"
	"errors"
	"fmt"
	"os"
	"path"
)

func Gen(cfg config.Config, prods []*ast.SyntaxBasicProd, symbols *symbols.Symbols, header string) (errs []error) {
	removeOldFiles(cfg)

	items := items.NewItems(prods)
	first := first.New(symbols, prods)
	var states *states.States
	if cfg.Knuth() {
		states = knuth.States(symbols, items, first)
	} else {
		states = pgm.States(symbols, items, first)
	}

	actions, conflicts := action.GetActions(states, symbols)
	if err := handleConflicts(conflicts, states, cfg); err != nil {
		errs = append(errs, err)
	}

	if cfg.Verbose2() {
		writeBasicProductions(cfg, prods)
		io.WriteFileString(path.Join(cfg.OutDir(), "CFG_items.txt"), items.String())
		writeCFGSymbols(cfg, symbols)
		io.WriteFileString(path.Join(cfg.OutDir(), "first.txt"), first.String())
		io.WriteFile(path.Join(cfg.OutDir(), "LR1_states.txt"), statesString(states, actions, symbols))
		io.WriteFileString(path.Join(cfg.OutDir(), "LR_states.dot"), statesDotString(states, actions, prods))
	}
	codegen.Gen(cfg, header, prods, symbols, states, actions)
	return
}

func handleConflicts(conflicts [][]*action.Conflict, states *states.States, cfg config.Config) error {
	if numConflicts(conflicts) <= 0 {
		return nil
	}
	if !cfg.AutoResolveLRConf() || cfg.Verbose2() {
		writeConflicts(cfg, conflicts)
	}
	if !cfg.AutoResolveLRConf() {
		return errors.New(fmt.Sprintf("%d LR(1) conflicts. See LR1_conflicts.txt", numConflicts(conflicts)))
	}
	return nil
}

func numConflicts(conflicts [][]*action.Conflict) (num int) {
	for _, sc := range conflicts {
		for _, c := range sc {
			if c != nil {
				num++
			}
		}
	}
	return
}

func writeConflicts(cfg config.Config, conflicts [][]*action.Conflict) {
	w := new(bytes.Buffer)
	fmt.Fprintf(w, "%d LR(1) Conflicts:\n", numConflicts(conflicts))
	conflictNo := 1
	for si, sc := range conflicts {
		for _, c := range sc {
			fmt.Fprintf(w, "%4d) S%d: %s\n", conflictNo, si, c)
			conflictNo++
		}
	}
	io.WriteFile(path.Join(cfg.OutDir(), "LR1_conflicts.txt"), w.Bytes())
}

func reduceLabel(state *states.State, actions map[string]action.Action, prods []*ast.SyntaxBasicProd) string {
	w := new(bytes.Buffer)
	numRedux, accept := 0, false
	for sym, a := range actions {
		switch red := a.(type) {
		case action.Accept:
			accept = true
		case action.Reduce:
			if numRedux > 0 {
				fmt.Fprintf(w, "\\n")
			}
			fmt.Fprintf(w, "%s {%s}", prods[red].DotString(), sym)
			numRedux++
		}
	}
	switch {
	case numRedux > 0:
		return fmt.Sprintf("S%d [shape=box,label=\"S%d\\n%s\"]\n", state.Number, state.Number, w.String())
	case accept:
		return fmt.Sprintf("S%d [shape=doublecircle]\n", state.Number)
	default:
		return ""
	}
}

func statesDotString(states *states.States, actions action.Actions, prods []*ast.SyntaxBasicProd) string {
	w := new(bytes.Buffer)
	fmt.Fprintf(w, "digraph{\n")
	for si, state := range states.List {
		fmt.Fprintf(w, "%s", reduceLabel(state, actions[si], prods))
		for _, t := range state.Transitions.List() {
			fmt.Fprintf(w, "\tS%d ->S%d [label = \"%s\"]\n", state.Number, t.State.Number, t.Sym)
		}
	}
	fmt.Fprintf(w, "}\n")
	return w.String()
}

func statesString(states *states.States, actions action.Actions, symbols *symbols.Symbols) []byte {
	w := new(bytes.Buffer)
	for si, state := range states.List {
		fmt.Fprintf(w, "%s", state)
		fmt.Fprintf(w, "Actions:\n")
		if actions != nil {
			for _, symT := range symbols.ListTerminals() {
				if action := actions[si][symT]; action != nil {
					fmt.Fprintf(w, "\t%s: %s\n", symT, action)
				}
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
	os.Remove(path.Join(cfg.OutDir(), "LR1_states.txt"))
	os.Remove(path.Join(cfg.OutDir(), "LR1_conflicts.txt"))
}

func writeBasicProductions(cfg config.Config, prods []*ast.SyntaxBasicProd) {
	w := new(bytes.Buffer)
	for i, prod := range prods {
		fmt.Fprintf(w, "%4d: %s\n\n", i, prod)
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
