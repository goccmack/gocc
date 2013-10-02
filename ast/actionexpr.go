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

package ast

import (
	"regexp"
	"strings"
)

func ActionExpressionVal(lit []byte) string {
	sdt := string(lit)
	rex, err := regexp.Compile("\\$[0-9]+")
	if err != nil {
		panic(err)
	}
	idx := rex.FindAllStringIndex(sdt, -1)
	res := ""
	if len(idx) <= 0 {
		res = sdt
	} else {
		for i, loc := range idx {
			if loc[0] > 0 {
				if i > 0 {
					res += sdt[idx[i-1][1]:loc[0]]
				} else {
					res += sdt[0:loc[0]]
				}
			}
			res += "X["
			res += sdt[loc[0]+1 : loc[1]]
			res += "]"
		}
		if idx[len(idx)-1][1] < len(sdt) {
			res += sdt[idx[len(idx)-1][1]:]
		}
	}
	return strings.TrimSpace(res[2 : len(res)-2])
}
