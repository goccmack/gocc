package pgm

import (
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
			merged := false
			for _, snum := range symSuccessors[trans.Sym] {
				if weaklyCompatibleStates(s.List[snum], trans.State) {
					merged = true
					newStates := mergeStates(s.List[snum], trans.State)
					for sym, states := range newStates {
						for _, state := range states {
							s.List = append(s.List, state)
							state.Number = len(s.List) - 1
							symSuccessors[sym] = append(symSuccessors[sym], state.Number)
						}
					}

					st_trans = append(st_trans, states.Transition{trans.Sym, s.List[snum]})
				}
			}
			if !merged {
				st_trans = append(st_trans, trans)
				s.List = append(s.List, trans.State)
				trans.State.Number = len(s.List) - 1
				// fmt.Printf("New S%d\n", trans.State.Number)
				symSuccessors[trans.Sym] = append(symSuccessors[trans.Sym], trans.State.Number)
			}
		}
		s.List[si].Transitions = states.NewTransitionsList(st_trans, symbols)
	}
	// s.GenActions()
	return s
}

func makeContext(from *states.ConfigGroupSet, to *states.ConfigGroupSet) *states.ConfigGroupSet {
	newNucleus := to.CloneFromCore()
	for _, cfg := range from.List() {
		if toCfg := newNucleus.GetGroupByCore(cfg.Item.NextItem); toCfg != nil {
			toCfg.AddContext(cfg.ContextSet.List...)
		}
	}
	return newNucleus
}

func mergeStates(this *states.State, that *states.State) (newStates map[string][]*states.State) {
	oldNucleus := this.Nucleus.Clone()
	mergeCfgGrpSet(this.Nucleus, that.Nucleus)
	this.Closure()
	cfgDiff := this.Nucleus.ContextDiff(oldNucleus)
	cfgDiff.Add(this.NonNucleus.List()...)
	return propagateContext(this, cfgDiff)
}

/*
The context of all config groups in that, which give context to the nucleus group of this,
is merged into the nucleus of this.
*/
func mergeContext(this *states.State, thatNucleus *states.ConfigGroupSet) {
	for _, grp := range thatNucleus.List() {
		this.Nucleus.GetGroup(grp).AddContext(grp.ContextSet.List...)
	}
	this.Closure()
}

/*
Return a map of (symbol:*states.State) of newly generated states. Each new state has minus the same number as the one it replaces.
For example: The existing set has number 10 and the new state -10. The caller of this function must resolve
the state numbers as other sets may still have valid transitions to the original.
*/
func propagateContext(this *states.State, expandedNucleusCfGrps *states.ConfigGroupSet) (newStates map[string][]*states.State) {
	newStates = make(map[string][]*states.State)
	if this.Transitions == nil {
		return
	}
	for _, t := range this.Transitions.List() {
		sym, nextState := t.Sym, t.State
		newNucleus := makeContext(expandedNucleusCfGrps, nextState.Nucleus)
		if weaklyCompatibleCfgGrpSet(nextState.Nucleus, newNucleus) {
			mergeContext(nextState, newNucleus)
		} else {
			newState := this.NextSymbol(sym)
			newState.Number = -nextState.Number
			this.Transitions.Replace(sym, newState)
			newStates[sym] = append(newStates[sym], newState)
		}
	}
	return
}

func weaklyCompatibleStates(this *states.State, that *states.State) bool {
	return weaklyCompatibleCfgGrpSet(this.Nucleus, that.Nucleus)
}

func mergeCfgGrpSet(this *states.ConfigGroupSet, that *states.ConfigGroupSet) {
	for _, cfGrp := range this.List() {
		if diffContext := that.GetGroup(cfGrp).ContextSet.Diff(cfGrp.ContextSet); diffContext.Size() > 0 {
			cfGrp.ContextSet.Add(diffContext.List...)
		}
	}
}

func weaklyCompatibleCfgGrpSet(this *states.ConfigGroupSet, that *states.ConfigGroupSet) bool {
	if !this.CoreEqual(that) {
		return false
	}
	if contextMutuallyDisjunct(this, that) {
		return true
	}
	return contextInternallyDisjunct(this) || contextInternallyDisjunct(that)
}

func contextMutuallyDisjunct(this *states.ConfigGroupSet, that *states.ConfigGroupSet) bool {
	for _, thisCfgGrp := range this.List() {
		for _, thatCfgGrp := range that.List() {
			if thisCfgGrp.HashKey() != thatCfgGrp.HashKey() && len(thisCfgGrp.ContextSet.Intersection(thatCfgGrp.ContextSet)) > 0 {
				return false
			}
		}
	}
	return true
}

func contextInternallyDisjunct(this *states.ConfigGroupSet) bool {
	for i := 0; i < len(this.List())-1; i++ {
		for j := i + 1; j < len(this.List()); j++ {
			if len(this.List()[i].ContextSet.Intersection(this.List()[j].ContextSet)) > 0 {
				return false
			}
		}
	}
	return true
}
