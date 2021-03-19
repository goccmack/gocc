package t1

import (
	"testing"

	"github.com/goccmack/gocc/internal/test/t1/errors"
	"github.com/goccmack/gocc/internal/test/t1/lexer"
	"github.com/goccmack/gocc/internal/test/t1/parser"
	"github.com/goccmack/gocc/internal/test/t1/token"
)

type Equaler interface {
	Equal(interface{}) bool
}

func Test1(t *testing.T) {
	ast, err := parser.NewParser().Parse(lexer.NewLexer([]byte(`c`)))
	if err != nil {
		t.Fail()
	}
	if slice, ok := ast.([]interface{}); !ok {
		t.Fail()
	} else {
		if len(slice) != 2 {
			t.Fatalf("len(slice)==%d", len(slice))
		}
		if slice[0] != nil {
			t.Fatal("slice[0] != nil")
		}
		if str, ok := slice[1].(string); !ok {
			t.Fatalf("%T", slice[1])
		} else {
			if str != "c" {
				t.Fatal(`str != "c"`)
			}
		}
	}
}

func Test2(t *testing.T) {
	ast, err := parser.NewParser().Parse(lexer.NewLexer([]byte(`b c`)))
	if err != nil {
		t.Fatal(err)
	}
	if slice, ok := ast.([]interface{}); !ok {
		t.Fail()
	} else {
		if len(slice) != 2 {
			t.Fatalf("len(slice)==%d", len(slice))
		}
		if str, ok := slice[0].(string); !ok {
			t.Fatal(`str, ok := slice[0].(string); !ok`)
		} else {
			if str != "b" {
				t.Fatal(`str != "b"`)
			}
		}
		if str, ok := slice[1].(string); !ok {
			t.Fatalf("%T", slice[1])
		} else {
			if str != "c" {
				t.Fatal(`str != "c"`)
			}
		}
	}
}

func Test3(t *testing.T) {
	toks := lexer.NewLexer([]byte(`c b`))
	_, err := parser.NewParser().Parse(toks)
	if err == nil {
		t.Fatal("No error for erronous input.")
	}
	if errs, ok := err.(*errors.Error); !ok {
		t.Fatal("Incompatible error type for erronous input.")
	} else {
		if errs.ErrorToken.Column != 3 {
			t.Fatal("errs.ErrorToken.Column = ", errs.ErrorToken.Column, " != 3")
		}
	}
}

// Tests that require access to the generated go, but interact directly
// with components.

func mockToken(name string, tokenType token.Type, line, column int) *token.Token {
	return &token.Token{
		Type: tokenType,
		Lit:  []byte(name),
		Pos:  token.Pos{Offset: (line-1)*80 + column, Line: line, Column: column},
	}
}

// assertEqualityIs will verify that lhs.Equal(rhs) and rhs.Equal(lhs) is true/false
// depending on 'equality'.
func assertEqualityIs(t *testing.T, lhs, rhs Equaler, equality bool) {
	t.Helper()
	if lhs.Equal(rhs) != equality {
		t.Fatalf("expected lhs.Equal(rhs) to be %v: %v vs %v", equality, lhs, rhs)
	}
	if rhs.Equal(lhs) != equality {
		t.Fatalf("expected rhs.Equal(lhs) to be %v: %v vs %v", equality, rhs, lhs)
	}
}

func assertEqual(t *testing.T, lhs, rhs Equaler) {
	t.Helper()
	assertEqualityIs(t, lhs, rhs, true)
}

func assertNotEqual(t *testing.T, lhs, rhs Equaler) {
	t.Helper()
	assertEqualityIs(t, lhs, rhs, false)
}

func Test_Token(t *testing.T) {
	var (
		t1 = mockToken("1st", 101, 7, 11)
		t2 = mockToken("2nd", 9010, 15, 22)
		t3 = mockToken("3rd", 101, 7, 15)
		t4 = mockToken("bad", token.INVALID, 1, 1)
		t5 = mockToken("bad", token.EOF, 1, 1)
	)
	// none should be equal: assertNotEqual tests symterrically, so we only need
	// to compare one way.
	assertNotEqual(t, t1, t2)
	assertNotEqual(t, t1, t3)
	assertNotEqual(t, t1, t4)
	assertNotEqual(t, t2, t5)
	assertNotEqual(t, t2, t3)
	assertNotEqual(t, t2, t4)
	assertNotEqual(t, t2, t5)
	assertNotEqual(t, t3, t4)
	assertNotEqual(t, t3, t5)
	assertNotEqual(t, t4, t5)

	// They should self-equal, because of the short-circuit
	assertEqual(t, t1, t1)
	assertEqual(t, t2, t2)
	assertEqual(t, t3, t3)
	assertEqual(t, t4, t4)
	assertEqual(t, t5, t5)

	// Test an explicit value match by creating a new token that
	// does not match t1 and incrementally bringing it closer.
	altT1 := mockToken("x", 1000, 100, 10)
	assertNotEqual(t, t1, altT1)

	altT1.Type = t1.Type
	assertNotEqual(t, t1, altT1)

	altT1.Lit = []byte("1st")
	altT1.Type = 999
	assertNotEqual(t, t1, altT1)

	// meet the equality criteria.
	altT1.Type = t1.Type
	assertEqual(t, t1, altT1)
}
