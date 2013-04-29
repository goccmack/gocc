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

package lr1

import (
	"code.google.com/p/gocc/ast"
	"code.google.com/p/gocc/lr1/parser"
	"code.google.com/p/gocc/token"
	"fmt"
	"os"
	"strconv"
)

// Grammar symbol index -> ItemSet index
type GotoTableEntry map[int]int

//The TransitionTable.
type TransitionTable []*Transition

func (T TransitionTable) Goto(fromSet int, sym *ast.Symbol) (toSet int) {
	for _, t := range T {
		if t.From == fromSet && t.Symbol.Equals(sym) {
			return t.To
		}
	}
	return -1
}

//Returns a string representing the TransitionTable.
func (T TransitionTable) String() string {
	res := ""
	for _, t := range T {
		res += t.String()
	}
	return res
}

//Do a simple validity check.
func (T TransitionTable) Check() bool {
	for _, t1 := range T {
		for _, t2 := range T {
			if t1.From == t2.From && t1.Symbol.Equals(t2.Symbol) && t1.To != t2.To {
				return false
			}
		}
	}
	return true
}

//Returns a string representing the TransitionTable in the graphviz DOT format.
func (T TransitionTable) Dot() string {
	str := "digraph {\n"
	for _, t := range T {
		str += "\t" + t.Dot() + "\n"
	}
	str += "}"
	return str
}

//The Transition.
type Transition struct {
	From, To int
	Symbol   *ast.Symbol
}

//Returns a string representing the Transition.
func (T *Transition) String() string {
	return strconv.Itoa(T.From) + " -> " + strconv.Itoa(T.To) + " <<" + T.Symbol.TokLit + ">>\n"
}

//Returns a string representing the Transition in the graphviz DOT format.
func (T *Transition) Dot() string {
	str := strconv.Itoa(T.From) + " -> " + strconv.Itoa(T.To)
	str += `[label="` + T.Symbol.TokLit + `"]`
	return str
}

//The GotoTable.
type GotoTable struct {
	nt    ast.SymbolS
	ntmap map[string]int
	GTO   []GotoTableEntry
}

//Creates a new GotoTable.
func NewGotoTable(numStates int, g *ast.Grammar, trans TransitionTable) *GotoTable {
	tab := &GotoTable{
		nt:    g.NonTerminals.Min(g.Prod[0].Head),
		ntmap: make(map[string]int),
		GTO:   make([]GotoTableEntry, numStates),
	}
	for i, sym := range tab.nt {
		tab.ntmap[sym.String()] = i
	}
	for i := range tab.GTO {
		tab.GTO[i] = make(map[int]int)
	}
	for _, t := range trans {
		if t.Symbol.IsTerminal() {
			continue
		}
		symIdx := tab.ntmap[t.Symbol.String()]
		switch to, exists := tab.GTO[t.From][symIdx]; {
		case !exists:
			tab.GTO[t.From][symIdx] = t.To
		case to != t.To:
			fmt.Println("LR(1) conflict: From", t.From, "sym", t.Symbol, "To", t.To, "tab.To", tab.GTO[t.From][symIdx])
			fmt.Println(" New transition:", t)
			fmt.Println(" Existing transition:", strconv.Itoa(t.From)+" -> "+strconv.Itoa(tab.GTO[t.From][symIdx]), t.Symbol)
			for _, t1 := range trans {
				if t.From == t1.From {
					fmt.Println(t1)
				}
			}
			os.Exit(1)
		}
	}
	return tab
}

//Walks the goto table given a state and a next symbol and returns the nextState.
func (this *GotoTable) NextState(FromState int, sym ast.Symbol) (nextState int) {
	return this.GTO[FromState][this.ntmap[sym.String()]]
}

//Returns a list reachabilities.
func (this *GotoTable) ReachableNodes() []bool {
	res := make([]bool, len(this.GTO))
	for _, e := range this.GTO {
		for _, to := range e {
			res[to] = true
		}
	}
	return res
}

//Returns a string representing the GoToTable.
func (this *GotoTable) String() string {
	str := "state," + SymbolsHeadingString(this.nt) + "\n"
	for i, ge := range this.GTO {
		str += strconv.Itoa(i) + "," + ge.GotoEntryString(this.nt) + "\n"
	}
	return str
}

//Returns a comma seperated list of symbols.
func SymbolsHeadingString(symbols ast.SymbolS) string {
	str := ""
	for i, symbol := range symbols {
		str += symbol.String()
		if i < len(symbols)-1 {
			str += ","
		}
	}

	return str
}

func (this GotoTableEntry) GotoEntryString(symbols ast.SymbolS) string {
	slice := make([]string, len(symbols))
	for sym, state := range this {
		slice[sym] = strconv.Itoa(state)
	}
	str := ""
	for i, s := range slice {
		str += s
		if i < len(symbols)-1 {
			str += ","
		}
	}
	return str
}

//The ActionTable.
type ActionTable []*ActionTableEntry

//Converts the ItemSets into an ActionTable given the Grammar.
func (this ItemSets) ActionTable(g *ast.Grammar, sr_auto_resolve bool) (actTab ActionTable, srConflicts int) {
	actTab = make([]*ActionTableEntry, len(this))
	for i := range actTab {
		// actTab[i] = make(ActionTableEntry)
		actTab[i] = &ActionTableEntry{
			CanRecover: false,
			Actions:    make(parser.Actions),
		}
	}
	for si, set := range this {
		// fmt.Println("ActionTable: Set", si)
		for _, item := range set {
			if recoveryItem(item) {
				actTab[si].CanRecover = true
			}
			prod := item.Grammar.Prod[item.ProdIdx]

			if item.Pos >= len(prod.Body.Symbols) {
				switch {
				case item.Equals(&Item{0, 1, ast.EOF_SYMBOL, item.Grammar}):
					actTab[si].Actions[item.NextToken.TokType] = parser.Accept(0)
				case prod.Head.TokLit != item.Grammar.Prod[0].Head.TokLit:
					actTab[si].Actions[item.NextToken.TokType] = parser.Reduce(item.ProdIdx)
				}
			} else if prod.Body.Symbols[item.Pos].IsTerminal() {
				if nextState := this.GetIndex(Goto(set, prod.Body.Symbols[item.Pos], g.FirstSets())); nextState >= 0 {
					action := parser.Shift(nextState)
					if act, exists := actTab[si].Actions[prod.Body.Symbols[item.Pos].TokType]; exists && act != action {
						srConflicts++
						if !sr_auto_resolve {
							fmt.Println("LR(1) conflict, I=", si, item, "sym=", prod.Body.Symbols[item.Pos])
							fmt.Println("\t", act, "!=", action)
						}
					}
					actTab[si].Actions[prod.Body.Symbols[item.Pos].TokType] = action
				}
			}
		}
	}
	return
}

func recoveryItem(i *Item) bool {
	if i.Pos > 0 || i.Len() == 0 {
		return false
	}
	return i.Symbol(0).SymType == ast.ERROR_LIT
}

//Returns an Action Table Header given the nonTerminals.
func ActionTableHeader(nonTerminals ast.SymbolS) string {
	str := "State,"
	for i, nt := range nonTerminals {
		str += nt.String()
		if i < len(nonTerminals)-1 {
			str += ","
		}
	}
	return str
}

//Returns a map of the NonTerminals.
func NonTerminalsMap(nonTerminals ast.SymbolS) map[string]int {
	res := make(map[string]int)
	for i, s := range nonTerminals {
		res[s.String()] = i
	}
	return res
}

//Returns a map of the terminals.
func symbolMap(terminals ast.SymbolS) map[string]int {
	res := make(map[string]int)
	for i, s := range terminals {
		res[s.String()] = i
	}
	return res
}

// Non-Terminal -> Action
type ActionTableEntry parser.ActionRow

//Returns the Go Code String of the action.
func GenGoForAction(a parser.Action) string {
	return a.String()
}

// Generate action table and goto table Go code
func GenGo(act ActionTable, gto *GotoTable, g *ast.Grammar) string {
	goc := genProductionsTable(g)
	goc += "\n"
	goc += act.GenGo(g.TokenMap)
	goc += "\n"
	goc += gto.GenGo()
	goc += "\n"
	return goc
}

//Returns a string of go code for the ActionTable.
func (this ActionTable) GenGo(tm *token.TokenMap) string {
	res := "var ActionTable ActionTab = ActionTab {\n"
	for i, ar := range this {
		res += ar.GenGo(i, tm)
	}
	res += "}\n"
	return res
}

//Returns a string of go code for the ActionTableEntry.
func (this ActionTableEntry) GenGo(state int, tm *token.TokenMap) string {
	res := fmt.Sprintf("\t// state %d\n", state)
	res += fmt.Sprintf("\t&ActionRow{\n")
	res += fmt.Sprintf("\t\tCanRecover: %t,\n", this.CanRecover)
	res += fmt.Sprintf("\t\tActions: Actions{\n")
	for tok, a := range this.Actions {
		res += fmt.Sprintf("\t\t\t%d: %s, // %s\n", int(tok), a, tm.TokenString(tok))
	}
	res += fmt.Sprintf("\t\t},\n")
	res += fmt.Sprintf("\t},\n\n")
	return res
}

//Returns a string of go code for the GoToTable.
func (this GotoTable) GenGo() string {
	res := "var GotoTable GotoTab = GotoTab{\n"
	for i, gr := range this.GTO {
		res += gr.GenGo(i, this.nt)
	}
	res += "}\n"
	return res
}

//Returns a string of go code for the GotoTableEntry.
func (this GotoTableEntry) GenGo(state int, nt ast.SymbolS) string {
	res := fmt.Sprintf("\t// state %d\n\tGotoRow{\n", state)
	for i, g := range this {
		res += "		\"" + nt[i].TokLit + "\": State(" + strconv.Itoa(g) + "),\n"
	}
	res += "\t},\n"
	return res
}

//Returns a string of go code for a GotoTableFunc.
func GotoTableFunc(nextState int) string {
	res := "func()State{return "
	res += strconv.Itoa(nextState)
	res += "},"
	return res
}

func genProductionsTable(g *ast.Grammar) string {
	res := "var ProductionsTable = ProdTab{\n"
	for i, p := range g.Prod {
		res += genProductionsTableEntry(i, p)
	}
	res += "}\n"
	return res
}

func genProductionsTableEntry(pidx int, p *ast.Production) string {
	res := "\t// [" + strconv.Itoa(pidx) + "]\n"
	res += "\tProdTabEntry{\n"
	res += "\t\t\"" + p.String() + "\",\n"
	res += "\t\t\"" + p.Head.TokLit + "\",\n"
	res += "\t\t" + strconv.Itoa(len(p.Body.Symbols)) + ",\n"
	res += "\t\t" + "func(X []Attrib) (Attrib, error) {\n"
	if p.Body.SDT != "" {
		res += "\t\t\t" + "return " + p.Body.SDT + "\n"
	} else if p.Body == ast.NULL_BODY {
		res += "\t\t\t" + "return nil, nil" + "\n"
	} else {
		res += "\t\t\t" + "return X[0], nil" + "\n"
	}
	res += "\t\t},\n"
	res += "\t},\n"
	return res
}
