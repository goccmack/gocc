package ebnf

import (
	"code.google.com/p/gocc/test/ebnf/t2/lexer"
	"code.google.com/p/gocc/test/ebnf/t2/parser"
	"code.google.com/p/gocc/test/ebnf/t2/token"
	"testing"
)

func Test1(t *testing.T) {
	lex := lexer.NewLexer([]byte("a"))
	attrib, err := parser.NewParser().Parse(lex)
	if err != nil {
		t.Fatalf("%s", err)
	}
	a, ok := attrib.([]interface{})
	if !ok {
		t.Fatalf("Expected []interface{}, got %T", attrib)
	}
	if len(a) != 1 {
		t.Fatalf("Expected len(a) == 1, got %d", len(a))
	}
	a1, ok := a[0].(*token.Token)
	if !ok {
		t.Fatalf("Expected a1.(*token.Token), got %T", a[0])
	}
	if string(a1.Lit) != "a" {
		t.Fatalf("Expected `a`, got `%s`", a1.Lit)
	}
}

func Test2(t *testing.T) {
	lex := lexer.NewLexer([]byte(""))
	attrib, err := parser.NewParser().Parse(lex)
	if err != nil {
		t.Fatalf("%s", err)
	}
	if attrib != nil {
		t.Fatalf("Expected nil, got %v", attrib)
	}
}

func Test3(t *testing.T) {
	lex := lexer.NewLexer([]byte("a a a a a"))
	attrib, err := parser.NewParser().Parse(lex)
	if err != nil {
		t.Fatalf("%s", err)
	}
	a, ok := attrib.([]interface{})
	if !ok {
		t.Fatalf("Expected []interface{}, got %T", attrib)
	}
	if len(a) != 5 {
		t.Fatalf("Expected len(a) == 5, got %d", len(a))
	}
	for i, ai := range a {
		if string(ai.(*token.Token).Lit) != "a" {
			t.Fatalf("Expected a[%d] = `a`, got `%s`", i, string(ai.(*token.Token).Lit))
		}
	}
}
