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
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

func ActionExpressionVal(lit []byte) string {
	var res string
	if match := matchIndividualParams(lit); len(match) > 0 {
		res = individualParams(lit, match)
	} else {
		if match := matchVariadicParams(lit); len(match) > 0 {
			res = variadicParams(lit, match)
		} else {
			res = string(lit)
		}
	}
	return strings.TrimSpace(res[2 : len(res)-2])
}

func individualParams(lit []byte, match [][]int) string {
	w := new(bytes.Buffer)
	for i, loc := range match {
		if loc[0] > 0 {
			if i > 0 {
				fmt.Fprintf(w, "%s", lit[match[i-1][1]:loc[0]])
			} else {
				fmt.Fprintf(w, "%s", lit[0:loc[0]])
			}
			fmt.Fprintf(w, "X[%s]", lit[loc[0]+1:loc[1]])
		}
	}
	if match[len(match)-1][1] < len(lit) {
		fmt.Fprintf(w, "%s", lit[match[len(match)-1][1]:])
	}
	return w.String()
}

func matchIndividualParams(lit []byte) [][]int {
	rex, err := regexp.Compile("\\$[0-9]+")
	if err != nil {
		panic(err)
	}
	return rex.FindAllIndex(lit, -1)
}

func matchVariadicParams(lit []byte) [][]int {
	rex, err := regexp.Compile("\\$\\.\\.\\.")
	if err != nil {
		panic(err)
	}
	return rex.FindAllIndex(lit, -1)
}

func variadicParams(lit []byte, match [][]int) string {
	w := new(bytes.Buffer)
	for i, loc := range match {
		if loc[0] > 0 {
			if i > 0 {
				fmt.Fprintf(w, "%s", lit[match[i-1][1]:loc[0]])
			} else {
				fmt.Fprintf(w, "%s", lit[0:loc[0]])
			}
			fmt.Fprintf(w, "X...")
			if match[len(match)-1][1] < len(lit) {
				fmt.Fprintf(w, "%s", lit[match[len(match)-1][1]:])
			}
		}
	}
	return w.String()
}
