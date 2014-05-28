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
