package er

import (
	"fmt"
	"testing"

	"github.com/maxcalandrelli/gocc/example/errorrecovery/ast"
	"github.com/maxcalandrelli/gocc/example/errorrecovery/er.grammar/er"
)

func TestFail(t *testing.T) {
	sml, err := test([]byte("a b ; d e f"))
	if err != nil {
		t.Fail()
	}
	fmt.Print("output: [\n")
	for _, s := range sml {
		switch sym := s.(type) {
		case *er.Error:
			fmt.Printf("%s\n", sym)
		default:
			fmt.Printf("\t%v\n", sym)
		}
	}
	fmt.Println("]")
}

func test(src []byte) (astree ast.StmtList, err error) {
	fmt.Printf("input: %s\n", src)
	s := er.NewLexerBytes(src)
	p := er.NewParser()
	a, err, _ := p.Parse(s)
	if err == nil {
		astree = a.(ast.StmtList)
	}
	return
}
