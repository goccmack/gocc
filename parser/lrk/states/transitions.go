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
		this.list = make([]Transition, len(this.transitions))
		i := 0
		for _, sym := range this.symbols.List() {
			if st, exist := this.transitions[sym]; exist {
				this.list[i].State = st
				this.list[i].Sym = sym
				i++
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
