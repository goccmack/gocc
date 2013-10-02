package ebnf

import (
	"code.google.com/p/gocc/test/ebnf/t1/lexer"
	"code.google.com/p/gocc/test/ebnf/t1/parser"
	"code.google.com/p/gocc/test/ebnf/t1/token"
	"testing"
)

func Test1(t *testing.T) {
	lex := lexer.NewLexer([]byte("a"))
	attrib, err := parser.NewParser().Parse(lex)
	if err != nil {
		t.Fatalf("%s", err)
	}
	a, ok := attrib.(*token.Token)
	if !ok {
		t.Fatalf("Expected *token.Token, got %T", attrib)
	}
	if string(a.Lit) != "a" {
		t.Fatalf("Expected `a`, got `%s`", a.Lit)
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
