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

/*
Support for the symbols of the language defined by the input grammar, G. This package supports code generation.
*/
package symbols

import (
	"fmt"
	"strings"

	"github.com/maxcalandrelli/gocc/internal/ast"
)

type Symbols struct {
	tSymbols        ast.SyntaxSymbols
	ntSymbols       ast.SyntaxSymbols
	symbols         map[ast.SyntaxSymbol]struct{}
	literalType     map[string]int
	nonTerminalType map[string]int
}

func NewSymbols(grammar *ast.Grammar) *Symbols {
	symbols := &Symbols{
		symbols:         make(map[ast.SyntaxSymbol]struct{}),
		tSymbols:        make(ast.SyntaxSymbols, 0, 16),
		ntSymbols:       make(ast.SyntaxSymbols, 0, 16),
		literalType:     make(map[string]int),
		nonTerminalType: make(map[string]int),
	}

	symbols.Add(ast.InvalidSyntaxSymbol{})
	symbols.Add(ast.EofSymbol)

	if grammar.SyntaxPart == nil {
		return symbols
	}

	for _, p := range grammar.SyntaxPart.ProdList {
		symbols.Add(p.Id)
		for _, sym := range p.Body.Symbols {
			symbols.Add(sym)
		}
	}
	return symbols
}

func (this *Symbols) Add(symbols ...ast.SyntaxSymbol) {
	for _, sym := range symbols {
		if _, exist := this.symbols[sym]; !exist {
			if sym.IsTerminal() {
				id := len(this.tSymbols)
				this.symbols[sym] = struct{}{}
				this.tSymbols = append(this.tSymbols, sym)
				this.literalType[sym.SymbolString()] = id
			} else if sym.IsNonTerminal() {
				id := len(this.ntSymbols)
				this.symbols[sym] = struct{}{}
				this.ntSymbols = append(this.ntSymbols, sym)
				this.nonTerminalType[sym.SymbolString()] = id
			}
		}
	}
}

func (this *Symbols) List() ast.SyntaxSymbols {
	return append(this.ntSymbols, this.tSymbols...)
}

func (this *Symbols) listMatchingSymbols(match func(ast.SyntaxSymbol) bool) ast.SyntaxSymbols {
	res := make(ast.SyntaxSymbols, 0, 16)
	for sym := range this.symbols {
		if match(sym) {
			res = append(res, sym)
		}
	}
	return res
}

func (this *Symbols) ListContextDependentTokenSymbols() ast.SyntaxSymbols {
	return this.listMatchingSymbols(func(s ast.SyntaxSymbol) bool {
		_, r := s.(ast.SyntaxContextDependentTokId)
		return r
	})
}

/*
Return a slice containing the ids of all symbols declared as string literals in the grammar.
*/
func (this *Symbols) ListStringLitSymbols() ast.SyntaxSymbols {
	return this.listMatchingSymbols(func(s ast.SyntaxSymbol) bool {
		_, r := s.(ast.SyntaxStringLit)
		return r
	})
}

func (this *Symbols) ListTerminals() ast.SyntaxSymbols {
	ret := ast.SyntaxSymbolsByName(this.tSymbols)
	//sort.Sort(&ret)
	return ast.SyntaxSymbols(ret)
}

func (this *Symbols) StringLitType(id string) int {
	if typ, exist := this.literalType[id]; exist {
		return typ
	}
	return -1
}

/*
Return the id of the NT with index idx, or "" if there is no NT symbol with index, idx.
*/
func (this *Symbols) NTId(idx int) string {
	if idx < 0 || idx >= len(this.nonTerminalType) {
		return ""
	}
	return this.ntSymbols[idx].SymbolName()
}

/*
Return the number of NT symbols in the grammar
*/
func (this *Symbols) NumNTSymbols() int {
	return len(this.ntSymbols)
}

/*
Returns a slice containing all the non-terminal symbols of the grammar.
*/
func (this *Symbols) NTList() ast.SyntaxSymbols {
	ret := ast.SyntaxSymbolsByName(this.ntSymbols)
	//sort.Sort(&ret)
	return ast.SyntaxSymbols(ret)
}

/*
Returns the NT index of a symbol (index in 0..|NT|-1) or -1 if the symbol is not in NT.
*/
func (this *Symbols) NTType(symbol string) int {
	if idx, exist := this.nonTerminalType[symbol]; exist {
		return idx
	}
	return -1
}

/*
Returns the total number of symbols in grammar: the sum of the terminals and non-terminals.
*/
func (this *Symbols) NumSymbols() int {
	return len(this.symbols)
}

func (this *Symbols) String() string {
	w := new(strings.Builder)
	for i, sym := range this.List() {
		fmt.Fprintf(w, "%3d: %s\n", i, sym.SymbolName())
	}
	return w.String()
}

/*
func (this *Symbols) Type(id ast.SyntaxSymbol) int {
	if typ, ok := this.idMap[id.SymbolName()]; ok {
		return typ
	}
	return -1
}
*/
