package sr

import (
	"testing"

	"github.com/maxcalandrelli/gocc/example/sr/ast"
	"github.com/maxcalandrelli/gocc/example/sr/sr.grammar/sr"
)

func parse(src string) (stmt ast.Stmt, err error) {
	lex := sr.NewLexerString(src)
	p := sr.NewParser()
	if res, err, _ := p.Parse(lex); err == nil {
		stmt = res.(ast.Stmt)
	}
	return
}

func Test1(t *testing.T) {
	stmt, err := parse("if c1 then s1")
	if err != nil {
		t.Fatal(err.Error())
	}
	ifs, ok := stmt.(*ast.If)
	if !ok {
		t.Fatalf("sr_test.Test1: stmt is not *ast.If")
	}
	if !ifs.MatchIf("c1", ast.IdStmt("s1")) {
		t.Fail()
	}
}

func Test2(t *testing.T) {
	stmt, err := parse("if c1 then s1 else s2")
	if err != nil {
		t.Fatal(err.Error())
	}
	ifes, ok := stmt.(*ast.IfElse)
	if !ok {
		t.Fatalf("stmt is not *ast.IfElse")
	}
	if !ifes.MatchIfElse("c1", ast.IdStmt("s1"), ast.IdStmt("s2")) {
		t.Fail()
	}
}

func Test3(t *testing.T) {
	stmt, err := parse("if c1 then if c2 then s2 else s3")
	if err != nil {
		t.Fatal(err.Error())
	}
	ifs, ok := stmt.(*ast.If)
	if !ok {
		t.Fatalf("stmt is not *ast.IfElse")
	}
	ife2 := &ast.IfElse{C: "c2", S1: ast.IdStmt("s2"), S2: ast.IdStmt("s3")}
	if !ifs.MatchIf("c1", ife2) {
		t.Fail()
	}
}
