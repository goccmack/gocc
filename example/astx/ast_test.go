package astx

import (
	"code.google.com/p/gocc/example/astx/ast"
	"code.google.com/p/gocc/example/astx/parser"
	"code.google.com/p/gocc/example/astx/scanner"
	"code.google.com/p/gocc/example/astx/token"
	"fmt"
	"testing"
)

func TestPass(t *testing.T) {
	sml, err := test([]byte("a b c d e f"))
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Printf("output: %s\n", sml)
}

func TestFail(t *testing.T) {
	sml, err := test([]byte("a b ; d e f"))
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(sml)
}

func test(src []byte) (astree ast.StmtList, err error) {
	fmt.Printf("input: %s\n", src)
	s := &scanner.Scanner{}
	s.Init([]byte(src), token.ASTXTokens)
	p := parser.NewParser(parser.ActionTable, parser.GotoTable, parser.ProductionsTable, token.ASTXTokens)
	a, err := p.Parse(s)
	if err == nil {
		astree = a.(ast.StmtList)
	}
	return
}
