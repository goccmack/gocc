package ast

import (
	"fmt"

	"github.com/maxcalandrelli/gocc/example/ctx/ctx.grammar/ctx/iface"
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
	return Stmt(stmtList.(*iface.Token).Lit), nil
}

func CalcResult(result interface{}) (Stmt, error) {
	res := result.(int64)
	rstr := fmt.Sprintf("%d", res)
	fmt.Printf("result:  %s\n", rstr)
	return Stmt(rstr), nil
}
