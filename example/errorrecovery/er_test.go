package astx

import (
	"code.google.com/p/gocc/example/errorrecovery/ast"
	"code.google.com/p/gocc/example/errorrecovery/errors"
	"code.google.com/p/gocc/example/errorrecovery/parser"
	"code.google.com/p/gocc/example/errorrecovery/scanner"
	"code.google.com/p/gocc/example/errorrecovery/token"
	"fmt"
	"testing"
)

func TestFail(t *testing.T) {
	sml, err := test([]byte("a b ; d e f"))
	if err != nil {
		t.Fail()
	}
	fmt.Print("output: [\n")
	for _, s := range sml {
		switch sym := s.(type) {
		case *errors.Error:
			fmt.Printf("\terror:\n")
			if sym.Err != nil {
				fmt.Printf("\t\tERR: %v\n", sym.Err.Error())
			} else {
				fmt.Printf("\t\tERR: nil\n")
			}
			fmt.Printf("\t\tErrorToken: %v\n", sym.ErrorToken)
			fmt.Printf("\t\tErrorPos: %v\n", sym.ErrorPos)
			fmt.Printf("\t\tErrorSymbols: %v\n", sym.ErrorSymbols)
			fmt.Printf("\t\tExpectedTokens: %v\n", sym.ExpectedTokens)
		default:
			fmt.Printf("\t%v\n", sym)
		}
	}
	fmt.Println("]")
}

func test(src []byte) (astree ast.StmtList, err error) {
	fmt.Printf("input: %s\n", src)
	s := &scanner.Scanner{}
	s.Init([]byte(src), token.ERRORRECOVERYTokens)
	p := parser.NewParser(parser.ActionTable, parser.GotoTable, parser.ProductionsTable, token.ERRORRECOVERYTokens)
	a, err := p.Parse(s)
	if err == nil {
		astree = a.(ast.StmtList)
	}
	return
}
