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
	"fmt"
	"strings"

	"github.com/johnkerl/gocc/internal/ast"
	"github.com/johnkerl/gocc/internal/parser/first"
	"github.com/johnkerl/gocc/internal/parser/symbols"
)

// A list of a list of Items.
type ItemSets struct {
	sets []*ItemSet
	// Map from canonical key to set index for O(1) lookup
	keyMap map[string]int
}

// g is a BNF grammar. Items returns the sets of Items of the grammar g.
func GetItemSets(g *ast.Grammar, s *symbols.Symbols, firstSets *first.FirstSets) *ItemSets {
	initialSet := InitialItemSet(g, s, firstSets).Closure()
	S := &ItemSets{
		sets:   []*ItemSet{initialSet},
		keyMap: make(map[string]int),
	}
	// Add initial set to the map
	S.keyMap[initialSet.canonicalKey()] = 0
	initialSet.SetNo = 0

	symbols := s.List()
	included := -1
	for again := true; again; {
		again = false
		for i, I := range S.sets {
			if i > included {
				for _, X := range symbols {
					gto := I.Goto(X)
					if gto.Size() > 0 {
						idx := S.GetIndex(gto)
						if idx == -1 {
							idx = len(S.sets)
							S.sets = append(S.sets, gto)
							S.keyMap[gto.canonicalKey()] = idx
							gto.SetNo = idx
							again = true
						}
						I.AddTransition(X, idx)
					}
				}
				included = i
			}
		}
	}
	return S
}

// Returns whether the list of a list of items contains the list of items.
func (this *ItemSets) Contains(I *ItemSet) bool {
	if I == nil || I.Size() == 0 {
		return false
	}
	// Use map for O(1) lookup instead of O(n) linear search
	if this.keyMap == nil {
		// Fallback to linear search if keyMap not initialized (shouldn't happen normally)
		for _, i := range this.sets {
			if i.Equal(I) {
				return true
			}
		}
		return false
	}
	_, exists := this.keyMap[I.canonicalKey()]
	return exists
}

// Returns the index of the list of items.
func (this *ItemSets) GetIndex(I *ItemSet) int {
	if I == nil || I.Size() == 0 {
		return -1
	}
	// Use map for O(1) lookup instead of O(n) linear search
	if this.keyMap == nil {
		// Fallback to linear search if keyMap not initialized (shouldn't happen normally)
		for i, items := range this.sets {
			if items.Equal(I) {
				return i
			}
		}
		return -1
	}
	if idx, exists := this.keyMap[I.canonicalKey()]; exists {
		return idx
	}
	return -1
}

/*
Return a slice containing all the item sets in increasing order of index
*/
func (this *ItemSets) List() []*ItemSet {
	return this.sets
}

/*
return set[SetNo]
*/
func (this *ItemSets) Set(SetNo int) *ItemSet {
	return this.sets[SetNo]
}

/*
Returns the number of item sets
*/
func (this *ItemSets) Size() int {
	return len(this.sets)
}

// Returns a string representing the list of the list of items.
func (this *ItemSets) String() string {
	buf := new(strings.Builder)
	for i, is := range this.sets {
		fmt.Fprintf(buf, "S%d%s\n", i, is.String())
	}
	return buf.String()
}

// Returns the initial Item of a Grammar.
func InitialItemSet(g *ast.Grammar, symbols *symbols.Symbols, fs *first.FirstSets) *ItemSet {
	set := NewItemSet(symbols, g.SyntaxPart.ProdList, fs)
	set.SetNo = 0
	prod := g.SyntaxPart.ProdList[0]
	set.AddItem(NewItem(0, prod, 0, "␚"))
	return set
}
