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
	"strings"
	"testing"

	"github.com/goccmack/gocc/example/usercontext/ast"
	"github.com/goccmack/gocc/example/usercontext/lexer"
	"github.com/goccmack/gocc/example/usercontext/parser"
)

func tryParse(t *testing.T, code string, lexContext, parseContext interface{}) (interface{}, error) {
	t.Helper()

	l := lexer.NewLexer([]byte(code))
	l.Context = lexContext

	p := parser.NewParser()
	p.Context = parseContext

	return p.Parse(l)
}

func TestTokenContext(t *testing.T) {
	t.Run("nil context", func(t *testing.T) {
		t.Run("successful parse", func(t *testing.T) {
			if _, err := tryParse(t, "Life Universe Everything ...42...", nil, nil); err != nil {
				t.Fatalf("not expecting an error, got: %q", err.Error())
			}
		})

		t.Run("syntax error", func(t *testing.T) {
			// missing identifiers.
			_, err := tryParse(t, "...42...", nil, nil)
			if err == nil {
				t.Fatalf("expecting an error")
			}
			const expected = `Error in S0: ...42...(2,...42...), Pos(offset=0, line=1, column=1), expected one of: capitalized lowercase `
			if err.Error() != expected {
				t.Fatalf("incorrect error:\nexpected: %q\nactual  : %q", expected, err.Error())
			}
			t.Log(err.Error())
		})

		t.Run("context type mismatch", func(t *testing.T) {
			// if we use an identifier, it requires context, and the ast code will
			// reject the context for being nil.
			_, err := tryParse(t, "sixbyseven ...42...", nil, nil)
			if err == nil {
				t.Fatalf("expecting an error")
			}
			// not sure why err.ErrorToken has 12 in its pos.Column field, but currently
			// this is correct.
			const expected = `Error in S0: ...42...(2,...42...), Pos(offset=11, line=1, column=12): sixbyseven: NewIdentifier expected ParserContext, got <nil>`
			if err.Error() != expected {
				t.Fatalf("incorrect error:\nexpected: %q\nactual  : %q", expected, err.Error())
			}
			t.Log(err.Error())
		})
	})

	t.Run("contexts", func(t *testing.T) {
		lc := &ast.LexicalContext{SourceFile: "bender.txt", ForbiddenWords: []string{"daisy"}}
		pc := &ast.ParserContext{Visitors: make([]string, 0)} // not using any other fields

		sym, err := tryParse(t, "Bite My shiny metal ...42...", lc, pc)
		if err != nil {
			t.Fatalf("not expecting an error, got: %s", err.Error())
		}
		t.Logf("returned symbol was %#+v\n", sym)

		// Because of the grammar, the ast should be just the last Identifier created by
		// "Words", which should correspond to "metal".
		switch ident := sym.(type) {
		case *ast.Identifier:
			if string(ident.Name.Lit) != "metal" {
				t.Fatalf("expected 'metal' symbol, got: %s: %v", string(ident.Name.Lit), sym)
			}
			if ident.ParserContext != pc {
				t.Fatalf("expected ident.ParserContext to be &ParserContext, got: %v", ident.ParserContext)
			}
			if ident.Name.Pos.Context != lc {
				t.Fatalf("expected ident.Name.Pos.Context to be &LexicalContext, got: %v", ident.Name.Pos.Context)
			}

		default:
			t.Fatalf("first argument of parse is not the identifier we expected: got %#+v", sym)
		}

		// Verify that ParserContext.Visitors was populated.
		if len(pc.Visitors) != 2 {
			t.Fatalf("wrong count for pc.Visitors, expected: %d, got: %d", 2, len(pc.Visitors))
		}
		expected := "shiny metal"
		actual := strings.Join(pc.Visitors, " ")
		t.Logf("visitors: %s", actual)
		if actual != expected {
			t.Fatalf("wrong Visitor list, expected: %q, got: %q", expected, actual)
		}
	})
}
