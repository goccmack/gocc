package states

import (
	"bytes"
	"fmt"
	"code.google.com/p/gocc/parser/first"
	"code.google.com/p/gocc/parser/lrk/items"
	"code.google.com/p/gocc/parser/symbols"
)

type States struct {
	List []*State
}

func (this *States) NewState0(symbols *symbols.Symbols, items *items.Items, first *first.First, cg ...*ConfigGroup) *States {
	s0 := NewState(symbols, items, first).Add(cg...).Closure()
	s0.Number = 0
	this.List = append(this.List, s0)
	return this
}

func (this *States) String() string {
	w := new(bytes.Buffer)
	for _, s := range this.List {
		fmt.Fprintf(w, "%s\n\n", s)
	}
	return w.String()
}
