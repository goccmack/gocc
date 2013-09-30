package lrk

import (
	"bytes"
	"fmt"
)

type Conflict struct {
	Symbol  string
	Actions []Action
}

func (this *Conflict) AddConflict(sym string, act1, act2 Action) *Conflict {
	if this == nil {
		return &Conflict{
			Symbol:  sym,
			Actions: []Action{act1, act2},
		}
	}

	if !this.contain(act1) {
		this.Actions = append(this.Actions, act1)
	}
	if !this.contain(act2) {
		this.Actions = append(this.Actions, act2)
	}
	return this
}

func (this *Conflict) NumConflicts() int {
	if this == nil {
		return 0
	}
	return len(this.Actions)
}

func (this *Conflict) contain(action Action) bool {
	for _, act := range this.Actions {
		if act.Equal(action) {
			return true
		}
	}
	return false
}

func (this *Conflict) String() string {
	w := new(bytes.Buffer)
	fmt.Fprintf(w, "%s: {", this.Symbol)
	for i, a := range this.Actions {
		if i > 0 {
			fmt.Fprintf(w, ",")
		}
		fmt.Fprintf(w, " %s", a)
	}
	fmt.Fprintf(w, "}")
	return w.String()
}
