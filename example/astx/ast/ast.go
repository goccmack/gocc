package ast

import (
	"github.com/goccmack/gocc/example/astx/token"
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

// stmtList is passed using $T0.
func NewStmt(stmtList *token.Token) (Stmt, error) {
	return Stmt(stmtList.Lit), nil
}
