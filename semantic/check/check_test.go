package check

import (
	"code.google.com/p/gocc/ast"
	"code.google.com/p/gocc/ast/rewrite"
	"code.google.com/p/gocc/frontend/lexer"
	"code.google.com/p/gocc/frontend/parser"
	"code.google.com/p/gocc/parser/symbols"
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

func Test2(t *testing.T) {
	l := lexer.NewLexer([]byte(`A : "b" ; B : "c" ;`))
	g, err := parser.NewParser().Parse(l)
	if err != nil {
		t.Fatalf("%s", err)
	}
	s := symbols.NewSymbols(rewrite.BasicProds(g.(*ast.Grammar).SyntaxPart.ProdList))
	errs := UnreachableProductions(g.(*ast.Grammar).SyntaxPart.ProdList, s)
	if errs == nil {
		t.Fatal("Expected unreachable production, B")
	}
}
