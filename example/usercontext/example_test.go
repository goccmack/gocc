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

package example

import (
	"testing"

	"github.com/goccmack/gocc/example/usercontext/ast"
	"github.com/goccmack/gocc/example/usercontext/lexer"
	"github.com/goccmack/gocc/example/usercontext/parser"
)

func parseCode(context *ast.Context, code []byte) (interface{}, error) {
	l := lexer.NewLexer(code)
	p := parser.NewParser()
	p.UserContext = context
	return p.Parse(l)
}

func testEval(t *testing.T, code string, names []string) {
	context := &ast.Context{Path: ".", Filename: code}

	st, err := parseCode(context, []byte(code))
	if err != nil {
		panic(err)
	}
	if names == nil {
		if st != nil {
			t.Fatalf("produced import list where none was expected")
		}
		return
	}
	if st == nil {
		t.Fatalf("failed to produce import list")
		return
	}
	importList := st.(ast.ImportList)
	if len(names) != len(importList) {
		t.Fatalf("len(st) should be %d; got %d", len(names), len(importList))
	}
	for idx, imported := range importList {
		if imported.Context != context {
			t.Fatalf("missing context in ast entry")
		}
		if string(imported.Token.Lit) != names[idx] {
			t.Fatalf("wrong imports")
		}
	}
}

func TestUserContext(t *testing.T) {
	testCases := []struct {
		name  string
		code  string
		names []string
	}{
		{"footer only", "...42...", nil},
		{"single import", "import question;\n...42...\n", []string{"question"}},
		{"multi import", "import life; import universe; import everything; ...42...",
			[]string{"life", "universe", "everything"}},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			testEval(t, tc.code, tc.names)
		})
	}
}
