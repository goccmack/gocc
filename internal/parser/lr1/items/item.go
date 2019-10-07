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

	"github.com/maxcalandrelli/gocc/internal/config"

	"github.com/maxcalandrelli/gocc/internal/ast"
	"github.com/maxcalandrelli/gocc/internal/parser/lr1/action"
)

//An LR1 Item.
type Item struct {
	ProdIdx         int             // index in list of productions in Grammar.Prod
	Prod            *ast.SyntaxProd // The syntax production of this item
	Id              ast.SyntaxSymbol
	Body            ast.SyntaxSymbols
	Pos             int              // position of • in the item
	ExpectedSymbol  ast.SyntaxSymbol // the next expected symbol in the item, if this isn't a reduce item.
	FollowingSymbol ast.SyntaxSymbol // The next expected symbol after the item has been recognised
	Len             int              // the number of symbols making up the body
	str             string
}

/*
following symbol: the symbol expected after this item has been reduced
*/
func NewItem(prodIdx int, prod *ast.SyntaxProd, pos int, followingSymbol ast.SyntaxSymbol) *Item {
	item := &Item{
		ProdIdx:         prodIdx,
		Prod:            prod,
		Id:              prod.Id,
		Pos:             pos,
		FollowingSymbol: followingSymbol,
	}
	if prod.Body.Empty() {
		item.Len = 0
	} else {
		item.Len = len(prod.Body.Symbols)
	}
	if pos > item.Len {
		panic(fmt.Sprintf("%s : %s, pos=%d", item.Id, item.Body, item.Pos))
	}
	item.Body = make(ast.SyntaxSymbols, item.Len)
	if item.Len > 0 {
		for i, sym := range prod.Body.Symbols {
			item.Body[i] = sym
		}
	}
	if item.Len > 0 && item.Pos < item.Len {
		item.ExpectedSymbol = item.Body[item.Pos]
	} else {
		item.ExpectedSymbol = nil
	}
	item.str = item.getString()
	return item
}

func (this *Item) accept(sym ast.SyntaxSymbol) bool {
	return this.ProdIdx == 0 && this.Pos >= this.Len && ast.EofSymbol == this.FollowingSymbol && ast.EofSymbol == sym
}

/*
If the action is shift the next state is nextState
*/
func (this *Item) action(sym ast.SyntaxSymbol, nextState int) action.Action {
	switch {
	case ast.InvalidSyntaxSymbol{} == sym:
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
	return this.Len > 0 && ast.ErrorSymbol == this.Body[0]
}

//Returns whether two Items are equal based on their ProdIdx, Pos and NextToken.
func (this *Item) Equals(that *Item) bool {
	if that == nil {
		return false
	}

	return this.ProdIdx == that.ProdIdx &&
		this.Pos == that.Pos &&
		this.FollowingSymbol == that.FollowingSymbol
}

func (this *Item) Move() (next *Item) {
	return NewItem(this.ProdIdx, this.Prod, this.Pos+1, this.FollowingSymbol)
}

/*
Returns true if this is a reduce item
*/
func (this *Item) reduce() bool {
	return this.Len == 0 || this.Pos >= this.Len
}

func (this *Item) Symbol(i int) ast.SyntaxSymbol {
	return this.Body[i]
}

func (this *Item) getString() string {
	buf := new(strings.Builder)
	fmt.Fprintf(buf, "%s : ", this.Id)
	if this.Len == 0 {
		fmt.Fprintf(buf, config.SYMBOL_EMPTY)
	} else {
		for i, s := range this.Body {
			if this.Pos == i {
				fmt.Fprintf(buf, "•")
			}
			fmt.Fprintf(buf, s.SymbolName())
			if i < this.Len-1 {
				fmt.Fprintf(buf, " ")
			}
		}
	}
	if this.Pos == this.Len {
		fmt.Fprintf(buf, "•")
	}
	fmt.Fprintf(buf, " «%s»", this.FollowingSymbol)
	return buf.String()
}

func (this *Item) String() string {
	return this.str
}
