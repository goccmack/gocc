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

package check

import (
	"code.google.com/p/gocc/ast"
	"code.google.com/p/gocc/frontend/lexer"
	"code.google.com/p/gocc/frontend/parser"
	"testing"
)

func Test1(t *testing.T) {
	l := lexer.NewLexer([]byte(`A : "b" ; A : "c" ;`))
	g, err := parser.NewParser().Parse(l)
	if err != nil {
		t.Fatalf("%s", err)
	}
	errs := DuplicateProductions(g.(*ast.Grammar))
	if errs == nil {
		t.Fatal("Expected duplicate production, A")
	}
}
