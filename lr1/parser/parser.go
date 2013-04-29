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

//the parser package helps with lr1 table generation.
package parser

import (
	"code.google.com/p/gocc/token"
	"strconv"
)

type (
	// ActionTab []ActionRow
	// ActionRow map[token.Type]Action
	ActionTab []*ActionRow
	ActionRow struct {
		CanRecover bool
		Actions    Actions
	}
	Actions map[token.Type]Action
)

type (
	Accept int
	Shift  State
	Reduce int

	Action interface {
		Act()
		String() string
	}
)

func (this Accept) Act() {}
func (this Shift) Act()  {}
func (this Reduce) Act() {}

func (this Accept) String() string { return "Accept(0)" }
func (this Shift) String() string  { return "Shift(" + strconv.Itoa(int(this)) + ")" }
func (this Reduce) String() string { return "Reduce(" + strconv.Itoa(int(this)) + ")" }

type (
	GotoTab []GotoRow
	GotoRow map[NT]State
	State   int
	NT      string
)

type (
	ProdTab      []ProdTabEntry
	ProdTabEntry struct {
		String     string
		Head       NT
		NumSymbols int
		ReduceFunc func([]Attrib) (Attrib, error)
	}
	Attrib interface {
		// String() string
	}
)
