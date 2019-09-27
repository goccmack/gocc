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

	"github.com/maxcalandrelli/gocc/internal/ast"
)

// set is kept sorted
type DisjunctRangeSet struct {
	set      []CharRange
	MatchAny bool
}

func NewDisjunctRangeSet() *DisjunctRangeSet {
	return &DisjunctRangeSet{
		set:      make([]CharRange, 0, 16),
		MatchAny: false,
	}
}

func (this *DisjunctRangeSet) AddLexTNode(sym ast.LexTNode) {
	switch s := sym.(type) {
	case *ast.LexCharRange:
		if s.Negate {
			this.SubtractRange(s.From.Val, s.To.Val)
		} else {
			this.AddRange(s.From.Val, s.To.Val)
		}
	case *ast.LexCharLit:
		if s.Negate {
			this.SubtractRange(s.Val, s.Val)
		} else {
			this.AddRange(s.Val, s.Val)
		}
	case *ast.LexDot:
		this.MatchAny = true
	case *ast.LexRegDefId:
		// ignore
	default:
		panic(fmt.Sprintf("Unexpected type of symbol: %T", sym))
	}
}

func (this *DisjunctRangeSet) addRange(rng ...CharRange) {
	for _, r := range rng {
		this.AddRange(r.From, r.To)
	}
}

/*
			|===========|
|-------|	|===========| 1
|-----------+---|-------+ 2
|-----------+-----------| 3
|-----------+-----------+---| 4
			|---|-------+ 5
			|===========| 6
			|-----------+---| 7
			+===========+	|----| 8
			+---|----|--+ 9
			+---|-------| 10
			+---|-------+---| 11
*/
func (this *DisjunctRangeSet) AddRange(from, to rune) {
	for i := 0; i < len(this.set) && from <= to; {
		rng := this.set[i]
		switch {
		case from < rng.From:
			switch {
			case to < rng.From: // (1)
				this.insertRange(i, from, to)
				i += 2
			case to < rng.To: // but >= set[i].From (2)
				this.set[i] = CharRange{from, rng.From - 1}
				this.insertRange(i+1, rng.From, to)
				this.insertRange(i+2, to+1, rng.To)
				i += 3
			case to == rng.To: //(3)
				this.insertRange(i, from, rng.From-1)
				i += 2
			case to > rng.To: // (4)
				this.insertRange(i, from, rng.From-1)
				i += 2
			}
			from = rng.To + 1
		case from == rng.From:
			switch {
			case to < rng.To: // (5)
				this.set[i].To = to
				this.insertRange(i+1, to+1, rng.To)
				i += 2
			case to > rng.To: // (7)
				i++ // do nothing
			case to == rng.To: // (6)
				i++ // do nothing
			}
			from = rng.To + 1
		case from > rng.To: // (8)
			i++ // do nothing
		default: // from > rng.From && from <= rng.To:
			switch {
			case to < rng.To: // 9
				this.set[i].To = from - 1
				this.insertRange(i+1, from, to)
				this.insertRange(i+2, to+1, rng.To)
				i += 3
			case to == rng.To: // (10)
				this.set[i].To = from - 1
				this.insertRange(i+1, from, to)
				i += 2
			case to > rng.To: // (11)
				this.set[i].To = from - 1
				this.insertRange(i+1, from, rng.To)
				i += 2
			}
			from = rng.To + 1
		}
	}
	if from <= to {
		this.set = append(this.set, CharRange{from, to})
	}
}

func (this *DisjunctRangeSet) subtractRange(index int, from, to rune) int {
	//fmt.Printf("subtracting <%c-%c> at index %d<%c-%c>\n", from, to, index, this.set[index].From, this.set[index].To)
	rng := this.set[index]
	if to >= from && to >= rng.From {
		if from > rng.To || to < rng.From {
			return index + 1
		}
		if from < rng.From {
			return this.subtractRange(index, rng.From, to)
		}
		if to > rng.To {
			this.subtractRange(index, from, rng.To)
			return this.subtractRange(index+1, rng.To+1, to)
		}
		//
		//  from >= rng.to && to <= rng.to
		//
		//fmt.Printf("actually deleting range <%c-%c> at index %d<%c-%c>\n", from, to, index, this.set[index].From, this.set[index].To)
		switch {
		case from == rng.From && to == rng.To:
			copy(this.set[index:], this.set[index+1:])
			this.set = this.set[:len(this.set)-1]
		case from == rng.From:
			this.set[index].From = to + 1
		case to == rng.To:
			this.set[index].To = from - 1
		default:
			this.set[index].From = this.set[index].From
			this.set[index].To = from - 1
			this.AddRange(to+1, rng.To)
		}
	}
	return len(this.set)
}

func (this *DisjunctRangeSet) SubtractRange(from, to rune) {
	for index := 0; index < len(this.set); index = this.subtractRange(index, from, to) {
	}
}

func (this *DisjunctRangeSet) Range(i int) CharRange {
	return this.set[i]
}

func (this *DisjunctRangeSet) insertRange(at int, from, to rune) {
	mvSegSize := len(this.set) - at
	this.set = append(this.set, CharRange{from, to})
	if mvSegSize > 0 {
		copy(this.set[at+1:], this.set[at:at+mvSegSize])
		this.set[at].From, this.set[at].To = from, to

	}
}

func (this *DisjunctRangeSet) List() []CharRange {
	return this.set
}

func (this *DisjunctRangeSet) Size() int {
	return len(this.set)
}

func (this *DisjunctRangeSet) String() string {
	w := new(strings.Builder)
	fmt.Fprintf(w, "{")
	for i, rng := range this.set {
		if i > 0 {
			fmt.Fprintf(w, ", %s", rng.String())
		} else {
			fmt.Fprintf(w, "%s", rng.String())
		}
	}
	fmt.Fprintf(w, "}")
	return w.String()
}
