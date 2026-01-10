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
	"github.com/johnkerl/gocc/internal/parser/lr1/action"
)

// An LR1 Item.
type Item struct {
	ProdIdx         int             // index in list of productions in Grammar.Prod
	Prod            *ast.SyntaxProd // The syntax production of this item
	Id              string
	Body            []string
	Pos             int    // position of • in the item
	ExpectedSymbol  string // the next exptect symbol in the item, if this isn't a reduce item.
	FollowingSymbol string // The next expected symbol after the item has been recognised
	Len             int    // the number of symbols making up the body
	str             string
}

// following symbol: the symbol expected after this item has been reduced.
func NewItem(prodIdx int, prod *ast.SyntaxProd, pos int, followingSymbol string) *Item {
	item := &Item{
		ProdIdx:         prodIdx,
		Prod:            prod,
		Id:              prod.Id,
		Pos:             pos,
		FollowingSymbol: followingSymbol,
	}
	if prod.Body.Symbols[0].SymbolString() == "empty" {
		item.Len = 0
	} else {
		item.Len = len(prod.Body.Symbols)
	}
	if pos > item.Len {
		panic(fmt.Sprintf("%s : %s, pos=%d", item.Id, item.Body, item.Pos))
	}
	item.Body = make([]string, item.Len)
	if item.Len > 0 {
		for i, sym := range prod.Body.Symbols {
			item.Body[i] = sym.SymbolString()
		}
	}
	if item.Len > 0 && item.Pos < item.Len {
		item.ExpectedSymbol = item.Body[item.Pos]
	} else {
		item.ExpectedSymbol = ""
	}
	item.str = item.getString()
	return item
}

func (this *Item) accept(sym string) bool {
	return this.ProdIdx == 0 &&
		this.Pos >= this.Len &&
		this.FollowingSymbol == "␚" &&
		sym == "␚"
}

// If the action is shift the next state is nextState.
func (this *Item) action(sym string, nextState int) action.Action {
	switch {
	case sym == "INVALID":
		return action.ERROR
	case this.accept(sym):
		return action.ACCEPT
	case this.reduce() && this.FollowingSymbol == sym:
		return action.Reduce(this.ProdIdx)
	case sym == this.ExpectedSymbol:
		return action.Shift(nextState)
	}
	return action.ERROR
}

func (this *Item) canRecover() bool {
	return this.Len > 0 && this.Body[0] == "error"
}

// Equals weturns whether two Items are equal based on their ProdIdx, Pos and NextToken.
func (this *Item) Equals(that *Item) bool {
	if that == nil {
		return false
	}

	return this.ProdIdx == that.ProdIdx &&
		this.Pos == that.Pos &&
		this.FollowingSymbol == that.FollowingSymbol
}

func (this *Item) Move() (next *Item) {
	// Optimize Move() by reusing Body slice to avoid allocations
	next = &Item{
		ProdIdx:         this.ProdIdx,
		Prod:            this.Prod,
		Id:              this.Id,
		Body:            this.Body, // Reuse the Body slice
		Pos:             this.Pos + 1,
		Len:             this.Len,
		FollowingSymbol: this.FollowingSymbol,
	}
	// Update ExpectedSymbol based on new position
	if next.Pos < next.Len {
		next.ExpectedSymbol = next.Body[next.Pos]
	} else {
		next.ExpectedSymbol = ""
	}
	// Recompute string representation
	next.str = next.getString()
	return next
}

// reduce returns true if this is a reduce item.
func (this *Item) reduce() bool {
	return this.Len == 0 || this.Pos >= this.Len
}

func (this *Item) Symbol(i int) string {
	return this.Body[i]
}

func (this *Item) getString() string {
	// Pre-allocate buffer with estimated capacity to avoid reallocations
	// Estimate: Id (avg 10) + ": " + body symbols (avg 5 * 5 chars) + "• " + " «»" + following (avg 5) = ~50
	estimatedCap := len(this.Id) + 2 + this.Len*8 + 10 + len(this.FollowingSymbol)
	buf := strings.Builder{}
	buf.Grow(estimatedCap)

	// Use WriteString instead of fmt.Fprintf for better performance
	buf.WriteString(this.Id)
	buf.WriteString(" : ")
	if this.Len == 0 {
		buf.WriteString("empty")
	} else {
		for i, s := range this.Body {
			if this.Pos == i {
				buf.WriteString("•")
			}
			buf.WriteString(s)
			if i < this.Len-1 {
				buf.WriteString(" ")
			}
		}
	}
	if this.Pos == this.Len {
		buf.WriteString("•")
	}
	buf.WriteString(" «")
	buf.WriteString(this.FollowingSymbol)
	buf.WriteString("»")
	return buf.String()
}

func (this *Item) String() string {
	return this.str
}
