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

package calc

import (
	"code.google.com/p/gocc/example/calc/lexer"
	"code.google.com/p/gocc/example/calc/parser"
	"fmt"
	"testing"
)

type TI struct {
	src    string
	expect int64
}

var testData = []*TI{
	{"1 + 1", 2},
	{"1 * 1", 1},
	{"1 + 2 * 3", 7},
}

func Test1(t *testing.T) {
	p := parser.NewParser()
	pass := true
	for _, ts := range testData {
		s := lexer.NewLexer([]byte(ts.src))
		sum, err := p.Parse(s)
		if err != nil {
			pass = false
			t.Log(err.Error())
		}
		if sum != ts.expect {
			pass = false
			t.Log(fmt.Sprintf("Error: %s = %d. Got %d\n", ts.src, sum, ts.expect))
		}
	}
	if !pass {
		t.Fail()
	}
}
