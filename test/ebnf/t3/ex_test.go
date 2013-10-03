package ebnf

import (
	"code.google.com/p/gocc/test/ebnf/t3/ast"
	"code.google.com/p/gocc/test/ebnf/t3/lexer"
	"code.google.com/p/gocc/test/ebnf/t3/parser"
	"testing"
)

type testDataStruct struct {
	src string
	exp *ast.Result
}

var testData = []testDataStruct{
	{
		src: "a",
		exp: &ast.Result{
			Rep: []string{"a"},
		},
	},
	{
		src: "b",
		exp: &ast.Result{
			Opt: "b",
		},
	},
	{
		src: "",
		exp: (*ast.Result)(nil),
	},
	{
		src: "b a a a",
		exp: &ast.Result{
			Opt: "b",
			Rep: []string{"a", "a", "a"},
		},
	},
	{
		src: " a a a",
		exp: &ast.Result{
			Opt: "",
			Rep: []string{"a", "a", "a"},
		},
	},
}

func Test1(t *testing.T) {
	for i, tst := range testData {
		lex := lexer.NewLexerString(tst.src)
		res, err := parser.NewParser().Parse(lex)
		if err != nil {
			t.Fatalf("%s", err)
		}
		verifyResult(i, res, tst.exp, t)
	}
}

func verifyResult(i int, res interface{}, exp *ast.Result, t *testing.T) {
	if exp == nil {
		if res != nil {
			t.Fatalf("Expected nil, got %T", res)
		} else {
			return
		}
	}
	result, ok := res.(*ast.Result)
	if !ok {
		t.Fatalf("Expected *ast.Result, got %T", res)
	}
	if result.Opt != exp.Opt {
		t.Errorf("Expected opt=%s, got %s", exp.Opt, result.Opt)
	}
	if len(result.Rep) != len(exp.Rep) {
		t.Fatalf("Expected rep len %d, got %d", len(exp.Rep), len(result.Rep))
	}
	for i, rep := range result.Rep {
		if exp.Rep[i] != rep {
			t.Errorf("Rep[%d], expected %s, got %s", i, exp.Rep[i], result.Rep[i])
		}
	}
}
