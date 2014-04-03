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
	"code.google.com/p/gocc/parser/first"
	"code.google.com/p/gocc/parser/lr1/action"
	"code.google.com/p/gocc/parser/symbols"
	"fmt"
)

type ItemSet struct {
	SetNo int
	imap  map[string]*Item
	Items []*Item
	// string: symbol, int: nextState
	Transitions map[string]int

	Symbols *symbols.Symbols
	Prods   ast.SyntaxProdList
	FS      *first.FirstSets
}

//Creates a set of items.
func NewItemSet(symbols *symbols.Symbols, prods ast.SyntaxProdList, fs *first.FirstSets) *ItemSet {
	return &ItemSet{
		SetNo:       -1,
		imap:        make(map[string]*Item),
		Items:       make([]*Item, 0, 16),
		Transitions: make(map[string]int),
		Symbols:     symbols,
		Prods:       prods,
		FS:          fs,
	}
}

func (this *ItemSet) Action(symbol string) (act1 action.Action, conflicts []action.Action) {
	conflictMap := make(map[string]action.Action)
	act1 = action.ERROR
	for _, item := range this.Items {
		act2 := item.action(symbol, this.Transitions[symbol])
		switch {
		case act2.Equal(action.ERROR):
			// ignore
		case act1.Equal(action.ERROR):
			act1 = act2
		case !act1.Equal(act2):
			conflictMap[act1.String()] = act1
			conflictMap[act2.String()] = act2
			act1 = act1.ResolveConflict(act2)
		default:
			// act1 == act2. Do nothing
		}
	}
	for _, act := range conflictMap {
		conflicts = append(conflicts, act)
	}

	return
}

func (this *ItemSet) AddItem(items ...*Item) {
	for _, i := range items {
		if _, contain := this.imap[i.str]; !contain {
			this.imap[i.str] = i
			this.Items = append(this.Items, i)
		}
	}

}

func (this *ItemSet) AddTransition(symbol string, nextSet int) {
	if _, exist := this.Transitions[symbol]; exist {
		panic(fmt.Sprintf("Transition %s -> %d already exists", symbol, nextSet))
	}
	this.Transitions[symbol] = nextSet
}

func (this *ItemSet) CanRecover() bool {
	for _, item := range this.Items {
		if item.canRecover() {
			return true
		}
	}
	return false
}

func (this *ItemSet) NextSetIndex(symbol string) int {
	if nextSet, exist := this.Transitions[symbol]; exist {
		return nextSet
	}
	return -1
}

//TODO: optimise loop
/*
Dragon book, 2nd ed, section 4.7.2, p261
	Closure(I)
	repeat
		for each item [A->x•By, a] in I
			for each production B -> z in G'
				for each terminal b in FIRST(ya)
					add [B -> •z, b] to I
	until no more items are added to I
*/
func (this *ItemSet) Closure() (c *ItemSet) {
	if this.Size() == 0 {
		return NewItemSet(this.Symbols, this.Prods, this.FS)
	}
	c = NewItemSet(this.Symbols, this.Prods, this.FS)
	c.AddItem(this.Items...)
	included := -1
	for again := true; again; {
		again = false
		for idx, i := range c.Items {
			if idx > included {
				if i.Pos >= i.Len || this.Symbols.IsTerminal(i.ExpectedSymbol) {
					continue
				}
				for pi, prod := range this.Prods {
					if prod.Id == i.ExpectedSymbol {
						first := first1(this.FS, i.Body[i.Pos+1:], i.FollowingSymbol)
						for t := range first {
							if item := NewItem(pi, prod, 0, t); !c.Contain(item) {
								c.AddItem(item)
								again = true
							}
						}
					}
				}

				included = idx
			}
		}
	}
	return
}

func (this *ItemSet) Contain(item *Item) bool {
	if _, contain := this.imap[item.str]; contain {
		return true
	}
	return false
}

func (this *ItemSet) ContainString(item string) bool {
	if _, contain := this.imap[item]; contain {
		return true
	}
	return false
}

//Returns whether two lists of Items are equal.
func (this *ItemSet) Equal(that *ItemSet) bool {
	if that == nil || len(this.Items) != len(that.Items) {
		return false
	}
	for k := range this.imap {
		if _, contain := that.imap[k]; !contain {
			return false
		}
	}

	return true
}

func first1(firstSets *first.FirstSets, symbols []string, following string) first.SymbolSet {
	return first.FirstS(firstSets, append(symbols, following))
}

// Dragon book, 2nd ed, section 4.7.2, p261
func (I *ItemSet) Goto(X string) *ItemSet {
	J := NewItemSet(I.Symbols, I.Prods, I.FS)
	for _, item := range I.Items {
		if item.Pos < item.Len && X == item.ExpectedSymbol {
			nextItem := item.Move()
			J.AddItem(nextItem)
		}
	}
	if J.Size() > 0 {
		J = J.Closure()
	}
	return J
}

/*
Returns the number of items in the set.
*/
func (this *ItemSet) Size() int {
	return len(this.Items)
}

func (this *ItemSet) String() string {
	buf := new(bytes.Buffer)
	buf.WriteString("{\n")
	for _, item := range this.Items {
		fmt.Fprintf(buf, "\t%s\n", item)
	}
	buf.WriteString("}\n")
	fmt.Fprintf(buf, "Transitions:\n")
	for sym, set := range this.Transitions {
		fmt.Fprintf(buf, "\t%s -> %d\n", sym, set)
	}
	fmt.Fprintf(buf, "\n")
	return buf.String()
}
