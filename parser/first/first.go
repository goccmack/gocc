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

package first

import (
	"bytes"
	"code.google.com/p/gocc/ast"
	"code.google.com/p/gocc/parser/symbols"
	"fmt"
)

type First struct {
	// key: symbol
	symbol map[string]FirstSet

	// key: prod id
	ntMap  map[string][]*ast.SyntaxBasicProd
	ntList []string

	tList []string
}

func New(symbols *symbols.Symbols, prods []*ast.SyntaxBasicProd) *First {
	first := &First{
		symbol: make(map[string]FirstSet),
		ntMap:  make(map[string][]*ast.SyntaxBasicProd),
		ntList: symbols.ListNonTerminals(),
		tList:  symbols.ListTerminals(),
	}
	first.AddTerminals()
	first.AddNonTerminals(prods)
	return first
}

func (this *First) AddNonTerminals(prods []*ast.SyntaxBasicProd) {
	for again := true; again; {
		again = false
		for _, prod := range prods {
			changed := false
			if this.symbol[prod.Id], changed = this.symbol[prod.Id].Add(this.firstTerms(prod.Terms)...); changed {
				again = true
			}
		}
	}
}

func (this *First) AddTerminals() {
	for _, s := range this.tList {
		if _, exist := this.symbol[s]; exist {
			panic(fmt.Sprintf("Duplicate terminal: %s", s))
		}
		this.symbol[s] = FirstSet{s}
	}
}

func (this *First) FirstSymbol(sym string) FirstSet {
	return this.symbol[sym]
}

func (this *First) firstTerms(terms ast.SyntaxTerms) (first FirstSet) {
	for t := 0; t < len(terms); t++ {
		switch term := terms[t].(type) {
		case ast.SyntaxTokId:
			first, _ = first.Add(string(term))
			return
		case ast.SyntaxProdId:
			if f := this.symbol[string(term)]; !f.Contain("ℇ") {
				first, _ = first.Add(f...)
				return
			} else {
				for _, sym := range f {
					if sym != "ℇ" {
						first, _ = first.Add(sym)
					}
				}
			}
		case ast.SyntaxStringLit:
			first, _ = first.Add(string(term))
			return
		case ast.SyntaxError:
			first, _ = first.Add("error")
			return
		}
	}
	first, _ = first.Add("ℇ")
	return
}

func (this *First) FirstString(s []string, context ...string) (fs FirstSet) {
	if len(s) == 0 {
		fs, _ = fs.Add(context...)
		return
	}
	frst := make(FirstSet, 0, 4)
	deriveEmpty := true
	for i := 0; deriveEmpty && i < len(s); deriveEmpty, i = frst.Contain("ℇ"), i+1 {
		frst = this.FirstSymbol(s[i])
		fs, _ = fs.Add(frst.Min("ℇ")...)
	}
	if deriveEmpty {
		fs, _ = fs.Add(context...)
	}
	return
}

func (this *First) String() string {
	w := new(bytes.Buffer)
	fmt.Fprintf(w, "Terminals:\n")
	for _, t := range this.tList {
		fmt.Fprintf(w, "\t%s\t%s\n", t, this.symbol[t])
	}
	fmt.Fprintf(w, "\nNon-terminals:\n")
	for _, nt := range this.ntList {
		fmt.Fprintf(w, "\t%s\t%s\n", nt, this.symbol[nt])
	}
	return w.String()
}
