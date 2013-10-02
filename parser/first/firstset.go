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

package first

import (
	"bytes"
	"fmt"
)

type FirstSet []string

func (this FirstSet) Add(sym ...string) (fs FirstSet, new bool) {
	fs = this
	for _, s := range sym {
		if !this.Contain(s) {
			fs = append(fs, s)
			new = true
		}
	}
	return
}

func (this FirstSet) Contain(sym string) bool {
	if this == nil {
		return false
	}
	for _, s := range this {
		if s == sym {
			return true
		}
	}
	return false
}

func (this FirstSet) Size() int {
	return len(this)
}

func (this FirstSet) String() string {
	w := new(bytes.Buffer)
	fmt.Fprintf(w, "{")
	for i, sym := range this {
		if i > 0 {
			fmt.Fprintf(w, ", %s", sym)
		} else {
			fmt.Fprintf(w, "%s", sym)
		}
	}
	fmt.Fprintf(w, "}")
	return w.String()
}
