package ctx

import (
	"fmt"
	"testing"

	"github.com/maxcalandrelli/gocc/example/ctx/ast"
	"github.com/maxcalandrelli/gocc/example/ctx/ctx.grammar/ctx"
	"github.com/maxcalandrelli/gocc/example/ctx/ctx1.grammar/ctx1"
)

func TestPass(t *testing.T) {
	//sml, err := test("\\u0022AA\\u0022 other a b c calc 12 * 6 + 4 \\u03b3k\\u03b5 d e  \\u03b3_\\u03b5 \\u03b3\\u03b5 f")
	sml0, sml1, err0, err1 := test("β1α αβ1α β11β1β11 αβ33 αβ0β9 ")
	if err0 != nil {
		t.Fatal(err0.Error())
	}
	if err1 != nil {
		t.Fatal(err1.Error())
	}
	fmt.Printf("output0: %s\n", sml0)
	fmt.Printf("output1: %s\n", sml1)
}

func TestFail(t *testing.T) {
	_, _, err0, err1 := test("a b 22 c d e f")
	if err0 == nil {
		t.Fatal("expected parse error (0)")
	} else {
		fmt.Printf("Parsing failed as expected (0): %v\n", err0)
	}
	if err1 == nil {
		t.Fatal("expected parse error (1)")
	} else {
		fmt.Printf("Parsing failed as expected (1): %v\n", err0)
	}
}

func test(src string) (astree0, astree1 ast.StmtList, err0, err1 error) {
	var a0, a1 interface{}
	a0, err0 = ctx.ParseText(src)
	a1, err1 = ctx1.ParseText(src)
	fmt.Printf("input: %s\n", src)
	if err0 == nil {
		astree0 = a0.(ast.StmtList)
	}
	if err1 == nil {
		astree1 = a1.(ast.StmtList)
	}
	return
}
