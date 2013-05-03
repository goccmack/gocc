//Copyright 2012 Vastech SA (PTY) LTD
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

package lr1

import "code.google.com/p/gocc/lr1/parser"

type (
	Conflict struct{
		State int
		A1, A2 parser.Action
	}

	ConflictSet []*Conflict
)

func (this ConflictSet) Add(c *Conflict) (res ConflictSet) {
	if !this.Contain(c) {
		res = append(this, c)
	} else {
		res = this
	}
	return
}

func (this ConflictSet) Contain(c1 *Conflict) bool {
	for _, c2 := range this {
		if c1.Equal(c2) {
			return true
		}
	}
	return false
}

func (this *Conflict) Equal(that *Conflict) bool {
	if this.State != that.State {
		return false
	}
	if this.A1.Equal(that.A1) {
		return this.A2.Equal(that.A2)
	} else {
		if this.A1.Equal(that.A2) {
			return this.A2.Equal(that.A1)
		}
	}
	return false
}