package ast

import (
	"fmt"

	calc_lexer "github.com/maxcalandrelli/gocc/example/calc/lexer"
	calc_parser "github.com/maxcalandrelli/gocc/example/calc/parser"
	"github.com/maxcalandrelli/gocc/example/ctx/token"
)

type (
	StmtList []Stmt
	Stmt     string
)

func NewStmtList(stmt interface{}) (StmtList, error) {
	return StmtList{stmt.(Stmt)}, nil
}

func AppendStmt(stmtList, stmt interface{}) (StmtList, error) {
	return append(stmtList.(StmtList), stmt.(Stmt)), nil
}

func NewStmt(stmtList interface{}) (Stmt, error) {
	return Stmt(stmtList.(*token.Token).Lit), nil
}

func Calc(verb interface{}) (Stmt, error) {
	calc_parser.NewParser().Parse(calc_lexer.NewLexer())
	return fmt.Sprintf("%s ")
}
