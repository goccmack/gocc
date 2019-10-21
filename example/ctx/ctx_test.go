package ctx

import (
	"fmt"
	"testing"

	"github.com/maxcalandrelli/gocc/example/ctx/ast"
	ctx0 "github.com/maxcalandrelli/gocc/example/ctx/ctx0.grammar/ctx0"
	ctx1 "github.com/maxcalandrelli/gocc/example/ctx/ctx1.grammar/ctx1"
	ctx2 "github.com/maxcalandrelli/gocc/example/ctx/ctx2.grammar/ctx2"
)

func TestPass(t *testing.T) {
	//sml, err := test("\\u0022AA\\u0022 other a b c calc 12 * 6 + 4 \\u03b3k\\u03b5 d e  \\u03b3_\\u03b5 \\u03b3\\u03b5 f")
	sml, err := test("β1α αβ1α 6 + 4 β11β1β11 αβ33 αβ0β9 ")
	for _ix := 0; _ix < len(sml); _ix++ {
		if err[_ix] != nil {
			t.Log(err[_ix].Error())
			t.Fail()
		}
		fmt.Printf("output #%d: %s\n", _ix+1, sml[_ix])
	}
}

func TestFail(t *testing.T) {
	_, err := test("αβ1α β0 22 β9 α a α")
	for _ix := 0; _ix < len(err); _ix++ {
		if err[_ix] == nil {
			t.Fatal("expected parse error")
		} else {
			fmt.Printf("Parsing failed as expected: %v\n", err[_ix])
		}
	}
}

func test1(src string, astree []ast.StmtList, errors []error, p func(string) (interface{}, error)) ([]ast.StmtList, []error) {
	at := ast.StmtList(nil)
	a, err := p(src)
	at, _ = a.(ast.StmtList)
	return append(astree, at), append(errors, err)
}

func test(src string) (astree []ast.StmtList, err []error) {
	fmt.Printf("input: %s\n", src)
	astree, err = test1(src, astree, err, ctx0.ParseText)
	astree, err = test1(src, astree, err, ctx1.ParseText)
	astree, err = test1(src, astree, err, ctx2.ParseText)
	return
}
