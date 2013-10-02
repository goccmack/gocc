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

package items

import (
	"bytes"
	"code.google.com/p/gocc/ast"
	"fmt"
)

/*
A general CFG item without context. Context is the set of symbols which may follow an item.
*/
type Item struct {
	BasicProdIdx int
	Id           string
	Symbols      []string
	Position     int
	NextItem     *Item
	hashKey      string
	str          string
}

/*
index: index of prod in list of basic productions after rewrite from parsed productions.
*/
func NewItem(prod *ast.SyntaxBasicProd, index int) []*Item {
	items := make([]*Item, 0, 8)
	item0 := &Item{
		BasicProdIdx: index,
		Id:           prod.Id,
		Symbols:      itemSymbols(prod.Terms),
		Position:     0,
		hashKey:      itemHashKey(index, 0),
	}
	item0.str = makeItemString(item0.Id, item0.Symbols, item0.Position)

	item := item0
	for next := item.nextItem(); next != nil; item, next = next, next.nextItem() {
		item.NextItem = next
		items = append(items, item)
	}
	items = append(items, item)
	return items
}

func (this *Item) Equal(that *Item) bool {
	if this == nil || that == nil {
		return false
	}
	return this.hashKey == that.hashKey
}

func (this *Item) ExpectedSymbol() string {
	if this.Position >= len(this.Symbols) {
		return ""
	} else {
		return this.Symbols[this.Position]
	}
}

func (this *Item) HashKey() string {
	return this.hashKey
}

/*
Return true iff this is of the form: u -> α•
*/
func (this *Item) Reduce() bool {
	return this.Position >= len(this.Symbols)
}

/*
Return true iff this is of the form: u -> α•β and len(β) > 0
*/
func (this *Item) Shift() bool {
	return this.Position < len(this.Symbols)
}

func (this *Item) String() string {
	return this.str
}

/*
Return the symbols following the expected symbol.
Possibly empty
*/
func (this *Item) TailString() []string {
	start := this.Position + 1
	if start >= len(this.Symbols) {
		return []string{}
	}
	return this.Symbols[start:]
}

func itemSymbols(terms ast.SyntaxTerms) []string {
	symbols := make([]string, len(terms))
	for i, term := range terms {
		switch t := term.(type) {
		case ast.SyntaxStringLit:
			symbols[i] = string(t)
		case ast.SyntaxTokId:
			symbols[i] = string(t)
		case ast.SyntaxProdId:
			symbols[i] = string(t)
		case ast.SyntaxError:
			symbols[i] = "error"
		default:
			panic(fmt.Sprintf("Unexpected type of term: %t", term))
		}
	}
	return symbols
}

func (this *Item) nextItem() *Item {
	if this.Position >= len(this.Symbols) {
		return nil
	}
	return &Item{
		BasicProdIdx: this.BasicProdIdx,
		Id:           this.Id,
		Symbols:      this.Symbols,
		Position:     this.Position + 1,
		hashKey:      itemHashKey(this.BasicProdIdx, this.Position+1),
		str:          makeItemString(this.Id, this.Symbols, this.Position+1),
	}
}

func itemHashKey(pIndex, pos int) string {
	return fmt.Sprintf("%d:%d", pIndex, pos)
}

func makeItemString(prodId string, symbols []string, pos int) string {
	w := new(bytes.Buffer)
	fmt.Fprintf(w, "%s :", prodId)
	for i, sym := range symbols {
		if pos == i {
			fmt.Fprintf(w, " •%s", sym)
		} else {
			fmt.Fprintf(w, " %s", sym)
		}
	}
	if pos >= len(symbols) {
		fmt.Fprintf(w, "•")
	}
	return w.String()
}
