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
	"code.google.com/p/gocc/parser/first"
	"code.google.com/p/gocc/parser/lrk/items"
	"code.google.com/p/gocc/parser/symbols"
	"fmt"
)

type States struct {
	List []*State
}

func (this *States) NewState0(symbols *symbols.Symbols, items *items.Items, first *first.First, cg ...*ConfigGroup) *States {
	s0 := NewState(symbols, items, first).Add(cg...).Closure().Closure()
	s0.Number = 0
	this.List = append(this.List, s0)
	return this
}

/*
Return the number of states
*/
func (this *States) Size() int {
	return len(this.List)
}

func (this *States) String() string {
	w := new(bytes.Buffer)
	for _, s := range this.List {
		fmt.Fprintf(w, "%s\n\n", s)
	}
	return w.String()
}
