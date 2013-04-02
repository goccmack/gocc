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

import (
	"strconv"
)

//Creates a string of the DFA transitions in the graphviz DOT format.
func DFADot(trans []*Transition) string {
	dot := "digraph{\n"

	for _, t := range trans {
		dot += NodeLabel(t.From) + "->" + NodeLabel(t.To)
		dot += "[label=" + t.Symbol.String() + "]"
		dot += "\n"
	}

	dot += "}"
	return dot
}

//Used by the DFADot function to create a node from a number.
func NodeLabel(idx int) string {
	return "I" + strconv.Itoa(idx)
}
