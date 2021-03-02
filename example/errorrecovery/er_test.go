package astx

import (
	"fmt"
	"testing"

	"github.com/goccmack/gocc/example/errorrecovery/ast"
	"github.com/goccmack/gocc/example/errorrecovery/errors"
	"github.com/goccmack/gocc/example/errorrecovery/lexer"
	"github.com/goccmack/gocc/example/errorrecovery/parser"
)

func TestFail(t *testing.T) {
	errors.Severity = "as-expected"
	defer func() { errors.Severity = "error" }()
	sml, err := test([]byte("a b ; d e f"))
	if err != nil {
		t.Fail()
	}
	fmt.Print("output: [\n")
	for idx, s := range sml {
		switch sym := s.(type) {
		case *errors.Error:
			fmt.Printf("#%d %q\n", idx, sym)
		default:
			fmt.Printf("#%d\t%v\n", idx, sym)
		}
	}
	fmt.Println("]")
}

func test(src []byte) (astree ast.StmtList, err error) {
	fmt.Printf("input: %s\n", src)
	s := lexer.NewLexer(src)
	p := parser.NewParser()
	a, err := p.Parse(s)
	if err == nil {
		astree = a.(ast.StmtList)
	}
	return
}
