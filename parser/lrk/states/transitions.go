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

package states

import (
	"bytes"
	"code.google.com/p/gocc/parser/symbols"
	"fmt"
)

/*
The set of transitions of a state.
*/
type Transitions struct {
	//key: symbol
	transitions map[string]*State

	symbols *symbols.Symbols

	list []Transition
}

func NewTransitions(s *symbols.Symbols) *Transitions {
	return &Transitions{
		transitions: make(map[string]*State),
		symbols:     s,
	}
}

func NewTransitionsList(trans []Transition, s *symbols.Symbols) *Transitions {
	transitions := NewTransitions(s)
	transitions.list = trans
	for _, t := range trans {
		transitions.transitions[t.Sym] = t.State
	}
	return transitions
}

func (this *Transitions) List() []Transition {
	if len(this.list) != len(this.transitions) {
		this.list = make([]Transition, 0, len(this.transitions))
		for _, sym := range this.symbols.List() {
			if st, exist := this.transitions[sym]; exist {
				this.list = append(this.list, Transition{State: st, Sym: sym})
			}
		}
	}
	return this.list
}

func (this *Transitions) Replace(sym string, state *State) {
	this.transitions[sym] = state
	this.list = nil
}

func (this *Transitions) String() string {
	w := new(bytes.Buffer)
	for _, t := range this.List() {
		fmt.Fprintf(w, "%s: S%d\n", t.Sym, t.State.Number)
	}
	return w.String()
}

func (this *Transitions) Transition(sym string) *State {
	return this.transitions[sym]
}
