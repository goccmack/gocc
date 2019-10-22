package ast

import (
	"fmt"
)

type (
	StmtList []Stmt
	Stmt     string
)

func NewStmtList(stmt string) (StmtList, error) {
	return StmtList{Stmt(stmt)}, nil
}

func AppendStmt(stmtList StmtList, stmt string) (StmtList, error) {
	return append(stmtList, Stmt(stmt)), nil
}

func NewStmt(stmt string) (Stmt, error) {
	return Stmt(stmt), nil
}

func CalcResult(result interface{}) (Stmt, error) {
	res := result.(int64)
	rstr := fmt.Sprintf("%d", res)
	fmt.Printf("result:  %s\n", rstr)
	return Stmt(rstr), nil
}
