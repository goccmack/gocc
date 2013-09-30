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

package lrk

import (
	"code.google.com/p/gocc/parser/lrk/states"
	"code.google.com/p/gocc/parser/symbols"
	"fmt"
)

/*
Returns:
actions - One entry for each state. Each entry is a map of symbol -> action
conflicts - for each state a list of conflict lists. Each conflict list contains a list of conflicts for each
conflicted symbol in the state.
*/
func Actions(states *states.States, symbols *symbols.Symbols) (actions []map[string]Action, conflicts [][]*Conflict) {
	actions = make([](map[string]Action), len(states.List))
	conflicts = make([][]*Conflict, len(states.List))
	for si, state := range states.List {
		actions[si] = make(map[string]Action)
		for _, sym := range symbols.ListTerminals() {
			act, cnf := stateAction(state, sym)
			actions[si][sym] = act
			if cnf != nil {
				conflicts[si] = append(conflicts[si], cnf)
			}
		}
	}
	return
}

func stateAction(state *states.State, nextSym string) (action Action, conflict *Conflict) {
	if nextState := state.Transitions.Transition(nextSym); nextState != nil {
		action = Shift(nextState.Number)
	}
	for _, cfgrp := range state.ConfigGroups().List() {
		if act1 := configGroupAction(cfgrp, nextSym); act1 != nil {
			switch {
			case action == nil:
				action = act1
			case !act1.Equal(action):
				conflict = conflict.AddConflict(nextSym, action, act1)
				action = action.ResolveConflict(act1)
			}
		}
	}
	return
}

func configGroupAction(cfgrp *states.ConfigGroup, nextSym string) (action Action) {
	if cfgrp.Item.Reduce() {
		if cfgrp.ContextSet.Contain[nextSym] {
			if cfgrp.Item.BasicProdIdx == 0 {
				action = ACCEPT
			} else {
				action = Reduce(cfgrp.Item.BasicProdIdx)
			}
		}
	}
	return
}

type Action interface {
	Equal(Action) bool
	ResolveConflict(that Action) Action
	String() string
}

type (
	Accept bool
	Error  bool
	Reduce int
	Shift  int
)

const (
	ACCEPT = Accept(true)
	ERROR  = Error(true)
)

func (Accept) Equal(act Action) bool {
	if _, ok := act.(Accept); ok {
		return true
	}
	return false
}

func (Error) Equal(act Action) bool {
	if _, ok := act.(Error); ok {
		return true
	}
	return false
}

func (this Reduce) Equal(act Action) bool {
	if that, ok := act.(Reduce); ok {
		return this == that
	}
	return false
}

func (this Shift) Equal(act Action) bool {
	if that, ok := act.(Shift); ok {
		return this == that
	}
	return false
}

func (this Accept) ResolveConflict(that Action) Action {
	if _, ok := that.(Error); ok {
		return this
	}
	panic(fmt.Sprintf("Cannot have LR1 conflict with Accept."))
}

func (Error) ResolveConflict(that Action) Action {
	return that
}

func (this Shift) ResolveConflict(that Action) Action {
	switch that := that.(type) {
	case Accept:
		panic(fmt.Sprintf("Impossible conflict: Shift(%d)/Accept", int(this)))
	case Shift:
		panic(fmt.Sprintf("Cannot have Shift(%d)/Shift(%d)", int(this), int(that)))
	case Error:
		return this
	case Reduce:
		return this
	}
	panic(fmt.Sprintf("Conflict with unknown type of action: %T", that))
}

func (this Reduce) ResolveConflict(that Action) Action {
	switch that := that.(type) {
	case Accept:
		panic(fmt.Sprintf("Impossible conflict: Shift(%d)/Accept", int(this)))
	case Shift:
		return that
	case Error:
		return this
	case Reduce:
		if this < that {
			return this
		}
		return that
	}
	panic(fmt.Sprintf("Conflict with unknown type of action: %T", that))
}

func (this Accept) String() string {
	return "accept"
}

func (this Error) String() string {
	return "error"
}

func (this Reduce) String() string {
	return fmt.Sprintf("Reduce(%d)", this)
}

func (this Shift) String() string {
	return fmt.Sprintf("Shift(%d)", this)
}
