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

package ast

import (
	"code.google.com/p/gocc/token"
	"fmt"
	"os"
)

type Grammar struct {
	Prod         ProdS
	firstSets    FirstSets
	Symbols      SymbolS
	NonTerminals SymbolS
	Terminals    SymbolS
	TokenMap     *token.TokenMap
}

func NewAugmentedGrammar(prods ProdS, tm *token.TokenMap) (g *Grammar, err error) {
	head1 := prods[0].Head.Clone()
	head1.TokLit = "S!"
	S0 := &Production{head1, &Body{Symbols: SymbolS{prods[0].Head}}}
	prods = append(ProdS{S0}, prods...)
	g = &Grammar{Prod: prods}
	g.generateSymbols()
	g.TokenMap = tm
	return
}

func (this *Grammar) GetProductions(name *Symbol) (res []*Production) {
	res = make([]*Production, 0, 10)
	for _, prod := range this.Prod {
		if name.TokLit == prod.Head.TokLit {
			res = append(res, prod)
		}
	}
	return
}

func (this *Grammar) CheckProductions() (unreachable, missing SymbolS) {
	missing = make(SymbolS, 0, 2)
	unreachable = make(SymbolS, 0, 2)
	reachable := SymbolS{this.Prod[0].Head}

	for i, again := 0, true; again; i++ {
		prods := this.GetProductions(reachable[i])
		for _, p := range prods {
			for _, s := range p.Body.Symbols {
				if !s.IsTerminal() {
					if !this.ContainsProdName(s) {
						missing = missing.AddSetElement(s)
					} else if !reachable.Contains(s) {
						reachable = reachable.AddSetElement(s)
					}
				}
			}
		}
		again = len(reachable) > i+1
	}

	for _, p := range this.Prod {
		if !reachable.Contains(p.Head) {
			unreachable = unreachable.AddSetElement(p.Head)
		}
	}

	return
}

func (this *Grammar) Contains(production *Production) bool {
	for i := 0; i < len(this.Prod); i++ {
		if this.Prod[i].Equals(production) {
			return true
		}
	}
	return false
}

func (this *Grammar) ContainsProdName(name *Symbol) bool {
	for _, p := range this.Prod {
		if p.Head.Equals(name) {
			return true
		}
	}
	return false
}

const _SYMBOLS_START_SIZE = 50

func (G *Grammar) generateSymbols() {
	G.NonTerminals = make(SymbolS, 0, _SYMBOLS_START_SIZE)
	G.Symbols = make(SymbolS, 0, _SYMBOLS_START_SIZE)

	for _, p := range G.Prod {
		G.NonTerminals = G.NonTerminals.AddSetElement(p.Head)
		G.Symbols = G.Symbols.AddSetElement(p.Head)
		for _, s := range p.Body.Symbols {
			G.Symbols = G.Symbols.AddSetElement(s)
		}
	}

}

func (this *Grammar) String() string {
	res := ""
	for _, prod := range this.Prod {
		res += prod.String() + "\n"
	}
	return res
}

func (this *Grammar) DotToFile(filename string) bool {
	success := true
	file, err := os.Create(filename)
	if file == nil {
		fmt.Printf("can't open file; err=%s\n", err.Error())
		success = false
	}
	file.Write([]byte(this.Dot()))
	file.Close()
	return success
}

func (this *Grammar) Dot() string {
	res := "digraph \"\" { \n"
	for i, prod := range this.Prod {
		res += "Grammar" + " -> " + prod.Dot(fmt.Sprint(i))
	}
	res += "}\n"
	return res
}

func (this *Grammar) Equals(that *Grammar) bool {
	if that == nil || len(this.Prod) != len(that.Prod) {
		return false
	}

	for i, p := range this.Prod {
		if !p.Equals(that.Prod[i]) {
			return false
		}
	}
	return true
}

type FirstSets map[string]SymbolS

func NewFirstSets() FirstSets {
	return make(FirstSets)
}

func (this FirstSets) AddSet(prodName *Symbol, terminals SymbolS) (symbolsAdded bool) {
	for _, t := range terminals {
		if this.AddToken(prodName, t) {
			symbolsAdded = true
		}
	}
	return
}

func (this FirstSets) AddToken(prodName *Symbol, terminal *Symbol) (symbolAdded bool) {
	idx := prodName.TokLit
	set, ok := this[idx]
	if !ok {
		set = make(SymbolS, 0, 10)
	}
	if !set.Contains(terminal) {
		set = append(set, terminal)
		symbolAdded = true
	}
	this[idx] = set
	return
}

func (this FirstSets) GetSet(prodName *Symbol) SymbolS {
	if set, ok := this[prodName.TokLit]; ok {
		return set
	}
	return nil
}

func (this FirstSets) String() string {
	str := ""
	for nt, set := range this {
		str += nt
		str += ": "
		str += set.String()
		str += "\n"
	}
	return str
}

func (this *Grammar) FirstSets() FirstSets {
	if this.firstSets == nil {
		this.firstSets = NewFirstSets()

		for i, again := 1, true; again; i++ {
			again = false
			for _, prod := range this.Prod {
				if prod.Body.Symbols[0].IsTerminal() {
					if this.firstSets.AddToken(prod.Head, prod.Body.Symbols[0]) {
						again = true
					}
				} else {
					for i, doNext := 0, true; i < len(prod.Body.Symbols) && doNext; i++ {
						if this.firstSets.AddSet(prod.Head, this.firstSets.GetSet(prod.Body.Symbols[i])) {
							again = true
						}
						doNext = this.firstSets.GetSet(prod.Body.Symbols[i]).
							Contains(EPSILON_SYM)
						doNext = false
					}
				}
			}
		}
	}

	return this.firstSets
}

func (G *Grammar) First(sym *Symbol) SymbolS {
	if sym.IsTerminal() {
		return SymbolS{sym}
	}
	return G.FirstSets().GetSet(sym)
}

func (this *Grammar) FirstS(symbols SymbolS) (first SymbolS) {
	first = make(SymbolS, 0, 10)
	fst := this.First(symbols[0])
	first = first.AddSet(fst)
	for i := 1; i < len(symbols) && fst.Contains(EPSILON_SYM); i++ {
		fst = this.First(symbols[i])
		first = first.AddSet(fst)
	}
	if !fst.Contains(EPSILON_SYM) {
		first = first.Min(EPSILON_SYM)
	}
	return
}
