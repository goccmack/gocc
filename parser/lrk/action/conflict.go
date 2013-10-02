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

package action

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
