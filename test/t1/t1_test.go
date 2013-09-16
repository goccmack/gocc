package t1

import (
	"code.google.com/p/gocc/test/t1/lexer"
	"code.google.com/p/gocc/test/t1/parser"
	"testing"
)

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
			t.Fatal()
		}
		if str, ok := slice[1].(string); !ok {
			t.Fatalf("%T", slice[1])
		} else {
			if str != "c" {
				t.Fatal()
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
			t.Fatal()
		} else {
			if str != "b" {
				t.Fatal()
			}
		}
		if str, ok := slice[1].(string); !ok {
			t.Fatalf("%T", slice[1])
		} else {
			if str != "c" {
				t.Fatal()
			}
		}
	}
}
