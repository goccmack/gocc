package ast

import (
	"fmt"

	"github.com/maxcalandrelli/gocc/example/calc/calc.grammar/calc"
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
	calc_parser.NewParser().Parse(calc.NewLexerString())
	return fmt.Sprintf("%s ")
}
