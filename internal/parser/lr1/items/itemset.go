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
	"sort"
	"strings"

	"github.com/johnkerl/gocc/internal/ast"
	"github.com/johnkerl/gocc/internal/parser/first"
	"github.com/johnkerl/gocc/internal/parser/lr1/action"
	"github.com/johnkerl/gocc/internal/parser/symbols"
)

type ItemSet struct {
	SetNo int
	imap  map[string]*Item
	Items []*Item
	// string: symbol, int: nextState.
	Transitions map[string]int

	Symbols *symbols.Symbols
	Prods   ast.SyntaxProdList
	FS      *first.FirstSets

	// Cached canonical key to avoid recomputation
	cachedKey string
	keyDirty  bool
}

// NewItemSet returns a newly cosntructed set of items.
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
			this.keyDirty = true // Mark key as dirty when items are added
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

	// Build a map from production ID to list of production indices for O(1) lookup
	prodIdMap := make(map[string][]int)
	for pi, prod := range this.Prods {
		prodIdMap[prod.Id] = append(prodIdMap[prod.Id], pi)
	}

	included := -1
	for again := true; again; {
		again = false
		// Process only items up to the current length (newly added items will be processed next iteration)
		itemsLen := len(c.Items)
		for idx := included + 1; idx < itemsLen; idx++ {
			i := c.Items[idx]
			if i.Pos >= i.Len || this.Symbols.IsTerminal(i.ExpectedSymbol) {
				included = idx
				continue
			}
			// Use the map for O(1) lookup instead of O(n) linear search
			if prodIndices, exists := prodIdMap[i.ExpectedSymbol]; exists {
				// Pre-compute the body suffix slice once per item
				var bodySuffix []string
				if i.Pos+1 < i.Len {
					bodySuffix = i.Body[i.Pos+1:]
				}
				// Compute first set once per item instead of per production
				first := first1(this.FS, bodySuffix, i.FollowingSymbol)
				// Pre-allocate expected number of new items to reduce allocations
				for _, pi := range prodIndices {
					prod := this.Prods[pi]
					for _, t := range first {
						// Create item and check containment
						item := NewItem(pi, prod, 0, t)
						// Use direct map lookup instead of Contain() for better performance
						if _, exists := c.imap[item.str]; !exists {
							c.imap[item.str] = item
							c.Items = append(c.Items, item)
							c.keyDirty = true
							again = true
						}
					}
				}
			}
			included = idx
		}
	}
	return
}

func (this *ItemSet) Contain(item *Item) bool {
	_, exists := this.imap[item.str]
	return exists
}

func (this *ItemSet) ContainString(item string) bool {
	if _, contain := this.imap[item]; contain {
		return true
	}
	return false
}

// Equal will return whether two lists of Items are equal.
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

// first1 returns the characters contained within the first set.
// We iterate over the map directly to avoid unnecessary sorting.
func first1(firstSets *first.FirstSets, symbols []string, following string) []string {
	var symbolSeq []string
	if len(symbols) == 0 {
		// Avoid allocation for the common case of empty symbols
		symbolSeq = []string{following}
	} else {
		// Pre-allocate with exact size to avoid reallocation
		symbolSeq = make([]string, len(symbols)+1)
		copy(symbolSeq, symbols)
		symbolSeq[len(symbols)] = following
	}

	firsts := first.FirstS(firstSets, symbolSeq)
	// Pre-allocate keys slice with exact capacity to avoid reallocations
	keys := make([]string, 0, len(firsts))
	for key := range firsts {
		keys = append(keys, key)
	}
	return keys
}

// Goto implements Dragon book, 2nd ed, section 4.7.2, p261.
func (I *ItemSet) Goto(X string) *ItemSet {
	J := NewItemSet(I.Symbols, I.Prods, I.FS)
	// Pre-filter items that match X to avoid unnecessary Move() calls
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

// Size returns the number of items in the set.
func (this *ItemSet) Size() int {
	return len(this.Items)
}

// canonicalKey returns a canonical string representation of the item set
// used for efficient set equality comparisons. The result is cached.
func (this *ItemSet) canonicalKey() string {
	if len(this.imap) == 0 {
		return ""
	}
	// Return cached key if still valid
	if !this.keyDirty && this.cachedKey != "" {
		return this.cachedKey
	}
	// Recompute key - pre-allocate with exact size to avoid append reallocations
	keys := make([]string, 0, len(this.imap))
	totalLen := 0
	// Collect keys and compute total length in one pass
	for k := range this.imap {
		keys = append(keys, k)
		totalLen += len(k) + 1 // +1 for separator
	}
	sort.Strings(keys)
	// Pre-allocate builder with exact capacity to avoid reallocation
	buf := strings.Builder{}
	buf.Grow(totalLen)
	for i, k := range keys {
		if i > 0 {
			buf.WriteByte('\x00')
		}
		buf.WriteString(k)
	}
	this.cachedKey = buf.String()
	this.keyDirty = false
	return this.cachedKey
}

func (this *ItemSet) String() string {
	buf := new(strings.Builder)
	fmt.Fprintf(buf, "{\n")
	for _, item := range this.Items {
		fmt.Fprintf(buf, "\t%s\n", item)
	}
	fmt.Fprintf(buf, "}\n")
	fmt.Fprintf(buf, "Transitions:\n")
	var keys transitions
	for sym, set := range this.Transitions {
		key := transition{symbol: sym, nextState: set}
		keys = append(keys, key)
	}
	sort.Sort(keys)
	for _, key := range keys {
		sym, set := key.symbol, key.nextState
		fmt.Fprintf(buf, "\t%s -> %d\n", sym, set)
	}
	fmt.Fprintf(buf, "\n")
	return buf.String()
}

// transitions implements the sort.Sort interface, sorting transitions in
// ascending order based on the next state.
type transitions []transition

func (ts transitions) Len() int           { return len(ts) }
func (ts transitions) Swap(i, j int)      { ts[i], ts[j] = ts[j], ts[i] }
func (ts transitions) Less(i, j int) bool { return ts[i].nextState < ts[j].nextState }

// A transition represents a transition from a symbol to a given state.
type transition struct {
	symbol    string
	nextState int
}
