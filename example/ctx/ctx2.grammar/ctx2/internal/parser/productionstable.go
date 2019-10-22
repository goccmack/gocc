// Code generated by gocc; DO NOT EDIT.

package parser

import (
	"github.com/maxcalandrelli/gocc/example/ctx/ast"
)

import (
	"fmt"
	"github.com/maxcalandrelli/gocc/example/ctx/ctx2.grammar/ctx2/internal/token"
	"github.com/maxcalandrelli/gocc/example/ctx/ctx2.grammar/ctx2/internal/util"
	"strings"
)

func getString(X Attrib) string {
	switch X.(type) {
	case *token.Token:
		return string(X.(*token.Token).Lit)
	case string:
		return X.(string)
	}
	return fmt.Sprintf("%q", X)
}

func unescape(s string) string {
	return util.EscapedString(s).Unescape()
}

func unquote(s string) string {
	r, _, _ := util.EscapedString(s).Unquote()
	return r
}

func lc(s string) string {
	return strings.ToLower(s)
}

func uc(s string) string {
	return strings.ToUpper(s)
}

type (
	//TODO: change type and variable names to be consistent with other tables
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index      int
		NumSymbols int
		ReduceFunc func(interface{}, []Attrib) (Attrib, error)
	}
	Attrib interface {
	}
)

var productionsTable = ProdTab{
	ProdTabEntry{
		String: `S' : Π<StmtList>	<<  >>`,
		Id:         "S'",
		NTType:     0,
		Index:      0,
		NumSymbols: 1,
		ReduceFunc: func(Context interface{}, X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `StmtList : Π<Stmt>	<< ast.NewStmtList($0s) >>`,
		Id:         "StmtList",
		NTType:     1,
		Index:      1,
		NumSymbols: 1,
		ReduceFunc: func(Context interface{}, X []Attrib) (Attrib, error) {
			return ast.NewStmtList(getString(X[0]))
		},
	},
	ProdTabEntry{
		String: `StmtList : Π<StmtList> Π<Stmt>	<< ast.AppendStmt($0.(ast.StmtList), $1s) >>`,
		Id:         "StmtList",
		NTType:     1,
		Index:      2,
		NumSymbols: 2,
		ReduceFunc: func(Context interface{}, X []Attrib) (Attrib, error) {
			return ast.AppendStmt(X[0].(ast.StmtList), getString(X[1]))
		},
	},
	ProdTabEntry{
		String: `Stmt : id	<< ast.NewStmt($0e) >>`,
		Id:         "Stmt",
		NTType:     2,
		Index:      3,
		NumSymbols: 1,
		ReduceFunc: func(Context interface{}, X []Attrib) (Attrib, error) {
			return ast.NewStmt(unescape(getString(X[0])))
		},
	},
	ProdTabEntry{
		String: `Stmt : μ<calc_0>	<< ast.CalcResult($1) >>`,
		Id:         "Stmt",
		NTType:     2,
		Index:      4,
		NumSymbols: 2,
		ReduceFunc: func(Context interface{}, X []Attrib) (Attrib, error) {
			return ast.CalcResult(X[1])
		},
	},
}
