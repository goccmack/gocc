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

//The lr1 package generates the tables for the lr1 grammar.
package lr1

import (
	"code.google.com/p/gocc/ast"
	"strconv"
)

//An LR1 Item.
type Item struct {
	ProdIdx   int // index in list of productions in Grammar.Prod
	Pos       int
	NextToken *ast.Symbol
	Grammar   *ast.Grammar
}

//Creates a clone of the Item.
func (this *Item) Clone() *Item {
	return &Item{this.ProdIdx, this.Pos, this.NextToken, this.Grammar}
}

//Returns whether two Items are equal based on their ProdIdx, Pos and NextToken.
func (this *Item) Equals(that *Item) bool {
	if that == nil {
		return false
	}

	return (this.ProdIdx == that.ProdIdx &&
		this.Pos == that.Pos &&
		this.NextToken.Equals(that.NextToken))
}

func (this *Item) Symbol(i int) *ast.Symbol {
	return this.Grammar.Prod[this.ProdIdx].Body.Symbols[i]
}

// returns number of symbols in body of item
func (this *Item) Len() int {
	return len(this.Grammar.Prod[this.ProdIdx].Body.Symbols)
}

const DOT = "•"

//Returns a string representing the Item.
func (this *Item) String() string {
	prod := this.Grammar.Prod[this.ProdIdx]
	res := prod.Head.TokLit + " : "
	if prod.Body == ast.NULL_BODY {
		res += "𝞊"
	} else {
		for i, s := range prod.Body.Symbols {
			if this.Pos == i {
				res += DOT
			}
			res += s.TokLit
			if i < len(prod.Body.Symbols)-1 {
				res += " "
			}
		}
	}
	if this.Pos == len(prod.Body.Symbols) {
		res += DOT
	}
	res += " «"
	res += this.NextToken.TokLit
	res += "»"
	return res
}

//A list of Items.
type Items []*Item

//Creates a new list of items.
func NewItems() Items {
	return make(Items, 0, 10)
}

//Adds an Item to the list of items if it is not already contained in the list.
func (this Items) AddSetElement(i *Item) Items {
	if this.Contains(i) {
		return this
	}

	return append(this, i)
}

//Returns the union of the two lists of items.
func (this Items) AddSet(that Items) Items {
	res := this
	for _, i := range that {
		res = res.AddSetElement(i)
	}
	return res
}

//Returns whether an Item is contained in the list of Items.
func (this Items) Contains(i *Item) bool {
	for _, i1 := range this {
		if i1.Equals(i) {
			return true
		}
	}
	return false
}

//Returns a string representing the Items in the graphviz DOT format.
func (this Items) Dot(id string) string {
	res := id + "["
	res += "label=\"" + id + "\\n"
	for i, item := range this {
		res += item.String()
		if i < len(this)-1 {
			res += "\\n"
		}
	}
	res += "\""
	res += "]\n"
	return res
}

//Returns whether two lists of Items are equal.
func (this Items) Equals(that Items) bool {
	if that == nil || len(this) != len(that) {
		return false
	}

	for _, i := range this {
		if !that.Contains(i) {
			return false
		}
	}

	return true
}

//Returns a string representing the list of Items.
func (this Items) String() string {
	res := "{\n"
	for _, item := range this {
		res += "  "
		res += item.String()
		res += "\n"
	}
	res += "}"
	return res
}

// Dragon book, 2nd ed, section 4.7.2, p261
func Closure(i Items, fs ast.FirstSets) (c Items) {
	if len(i) == 0 {
		return NewItems()
	}
	c = NewItems().AddSet(i)
	for again := true; again; {
		again = false
		for _, i := range c {
			// fmt.Println("items.Closure, i -", i)
			prod := i.Grammar.Prod[i.ProdIdx]
			if i.Pos >= len(prod.Body.Symbols) || prod.Body.Symbols[i.Pos].IsTerminal() {
				continue
			}
			for pi, p := range i.Grammar.Prod {
				if p.Head.Equals(prod.Body.Symbols[i.Pos]) {
					first := First(i.Grammar, prod.Body.Symbols[i.Pos+1:], i.NextToken)
					for _, t := range first {
						// Why won't the if compile with the simple statement?
						item := &Item{pi, 0, t, i.Grammar}
						if !c.Contains(item) {
							c, again = c.AddSetElement(item), true
						}
					}
				}
			}
		}
	}
	return
}

//Returns the first symbols of a grammar given symbols and the nextToken.
func First(g *ast.Grammar, symbols ast.SymbolS, nextToken *ast.Symbol) ast.SymbolS {
	return g.FirstS(append(symbols, nextToken))
}

// Dragon book, 2nd ed, section 4.7.2, p261
func Goto(I Items, X *ast.Symbol, fs ast.FirstSets) Items {
	J := NewItems()
	for _, item := range I {
		if prod := item.Grammar.Prod[item.ProdIdx]; item.Pos < len(prod.Body.Symbols) && prod.Body.Symbols[item.Pos].Equals(X) {

			nextItem := item.Clone()
			nextItem.Pos += 1
			J = J.AddSetElement(nextItem)
		}
	}
	itemS := Closure(J, fs)

	return itemS
}

//Returns the inital Item of a Grammar.
func InitialItem(g *ast.Grammar) *Item {
	return &Item{
		ProdIdx:   0,
		Pos:       0,
		NextToken: ast.EOF_SYMBOL,
		Grammar:   g,
	}
}

//A list of a list of Items.
type ItemSets []Items

//Returns whether the list of a list of items contains the list of items.
func (this ItemSets) Contains(I Items) bool {
	for _, i := range this {
		if i.Equals(I) {
			return true
		}
	}
	return false
}

//Returns the index of the list of items.
func (this ItemSets) GetIndex(I Items) int {
	if I == nil || len(I) == 0 {
		return -1
	}

	for i, items := range this {
		if items.Equals(I) {
			return i
		}
	}
	return -1
}

//Returns a string representing the list of the list of items.
func (this ItemSets) String() string {
	res := ""
	for i, is := range this {
		res += "S" + strconv.Itoa(i) + " "
		res += is.String() + "\n"
	}
	return res
}

// g is a BNF grammar. Items returns the sets of Items of the grammar g.
func GetItemSets(g *ast.Grammar) (ItemSets, TransitionTable) {
	C := ItemSets{Closure(Items{InitialItem(g)}, g.FirstSets())}
	trans := make(TransitionTable, 0, 50)
	included := -1
	for again := true; again; {
		again = false
		for i, I := range C {
			if i > included {
				for _, X := range g.Symbols {
					gto := Goto(I, X, g.FirstSets())
					if len(gto) > 0 {
						idx := C.GetIndex(gto)
						if idx == -1 {
							C, again = append(C, gto), true
							idx = len(C) - 1
						}
						trans = append(trans, &Transition{i, idx, X})
					}
				}
				included = i
			}
		}
	}
	return C, trans
}
