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

package knuth

import (
	// "fmt"
	"code.google.com/p/gocc/parser/first"
	"code.google.com/p/gocc/parser/lrk/items"
	"code.google.com/p/gocc/parser/lrk/states"
	"code.google.com/p/gocc/parser/lrk/symbolsuccessors"
	"code.google.com/p/gocc/parser/symbols"
)

func States(symbols *symbols.Symbols, lr0items *items.Items, first *first.First) *states.States {
	s := &states.States{
		List: make([]*states.State, 0, 64),
	}
	s.NewState0(symbols, lr0items, first, states.NewConfigGroup(lr0items.List[0], "$"))
	symSuccessors := symbolsuccessors.NewSymbolSuccessors()
	for si := 0; si < len(s.List); si++ {
		st_trans := make([]states.Transition, 0, 4)
		for _, trans := range s.List[si].Next().List() {
			newState := true
			for _, snum := range symSuccessors[trans.Sym] {
				if s.List[snum].Equal(trans.State) {
					st_trans = append(st_trans, states.Transition{trans.Sym, s.List[snum]})
					newState = false
				}
			}
			if newState {
				st_trans = append(st_trans, trans)
				s.List = append(s.List, trans.State)
				trans.State.Number = len(s.List) - 1
				symSuccessors[trans.Sym] = append(symSuccessors[trans.Sym], trans.State.Number)
			}
		}
		s.List[si].Transitions = states.NewTransitionsList(st_trans, symbols)
	}
	return s
}
