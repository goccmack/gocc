package states

import (
	"bytes"
	"fmt"
	"code.google.com/p/gocc/parser/first"
	lr0items "code.google.com/p/gocc/parser/lrk/items"
	"code.google.com/p/gocc/parser/symbols"
)

type State struct {
	Number      int
	Nucleus     *ConfigGroupSet
	NonNucleus  *ConfigGroupSet
	symbols     *symbols.Symbols
	items       *lr0items.Items
	first       *first.First
	Transitions *Transitions // assigned in states.NewStates()
}

func NewState(symbols *symbols.Symbols, items *lr0items.Items, first *first.First) *State {
	return &State{
		Number:     -1,
		Nucleus:    NewConfigGroupSet(),
		NonNucleus: NewConfigGroupSet(),
		symbols:    symbols,
		items:      items,
		first:      first,
	}
}

func (this *State) Add(cg ...*ConfigGroup) *State {
	for _, grp := range cg {
		if grp.IsNucleus() {
			this.Nucleus.Add(grp)
		} else {
			this.NonNucleus.Add(grp)
		}
	}
	return this
}

func (this *State) Closure() *State {
	cgList := append(this.Nucleus.List(), this.NonNucleus.List()...)
	for i := 0; i < len(cgList); i++ {
		cgi := cgList[i]
		if this.symbols.IsNT(cgi.Item.ExpectedSymbol()) {
			startItems := this.items.StartItems(cgi.Item.ExpectedSymbol())
			for _, si := range startItems {
				newCfgGrp := NewConfigGroup(si, this.first.FirstString(cgi.Item.TailString(), cgi.ContextSet.Clone().List...)...)
				if oldCfgGrp := this.GetGroup(newCfgGrp); oldCfgGrp == nil {
					this.Add(newCfgGrp)
					cgList = append(cgList, newCfgGrp)
				} else {
					oldCfgGrp.AddContext(newCfgGrp.ContextSet.List...)
				}
			}
		}
	}
	return this
}

func (this *State) ConfigGroups() *ConfigGroupSet {
	set := NewConfigGroupSet()
	set.Add(this.Nucleus.List()...)
	set.Add(this.NonNucleus.List()...)
	return set
}

func (this *State) Equal(that *State) bool {
	return this.Nucleus.Equal(that.Nucleus) && this.NonNucleus.Equal(that.NonNucleus)
}

func (this *State) GetGroup(cg *ConfigGroup) *ConfigGroup {
	if grp := this.Nucleus.GetGroup(cg); grp != nil {
		return grp
	}
	return this.NonNucleus.GetGroup(cg)
}

func (this *State) Core() []*lr0items.Item {
	return append(this.Nucleus.Core(), this.NonNucleus.Core()...)
}

func (this *State) Next() *Transitions {
	transitions := NewTransitions(this.symbols)
	this.nextConfigGroupSet(this.Nucleus, transitions)
	this.nextConfigGroupSet(this.NonNucleus, transitions)

	return transitions
}

func (this *State) NextSymbol(sym string) *State {
	nextState := NewState(this.symbols, this.items, this.first)
	for _, cfgrp := range this.Nucleus.List() {
		if cfgrp.Item.ExpectedSymbol() == sym {
			nextState.Add(NewConfigGroup(cfgrp.Item.NextItem, cfgrp.ContextSet.List...))
		}
	}
	nextState.Closure()
	return nextState
}

func (this *State) nextConfigGroupSet(cgset *ConfigGroupSet, transitions *Transitions) {
	for _, cfgGrp := range cgset.list {
		if cfgGrp.Item.ExpectedSymbol() != "" {
			state, exist := transitions.transitions[cfgGrp.Item.ExpectedSymbol()]
			if !exist {
				state = NewState(this.symbols, this.items, this.first)
				transitions.transitions[cfgGrp.Item.ExpectedSymbol()] = state
			}
			state.Add(NewConfigGroup(cfgGrp.Item.NextItem, cfgGrp.ContextSet.List...))
			state.Closure()
		}
	}
}

func (this *State) String() string {
	w := new(bytes.Buffer)
	fmt.Fprintf(w, "S%d {\n", this.Number)
	fmt.Fprintf(w, "%s", this.Nucleus)
	fmt.Fprintf(w, "%s", this.NonNucleus)
	fmt.Fprintf(w, "}\n")
	fmt.Fprintf(w, "Transitions:\n")
	if this.Transitions != nil {
		for _, t := range this.Transitions.List() {
			fmt.Fprintf(w, "\t%s\n", t)
		}
	}
	return w.String()
}
