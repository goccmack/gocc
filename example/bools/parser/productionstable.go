package parser

import (
	"code.google.com/p/gocc/example/bools/ast"
)

type (
	//TODO: change type and variable names to be consistent with other tables
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index      int
		NumSymbols int
		ReduceFunc func([]Attrib) (Attrib, error)
	}
	Attrib interface {
	}
)

var productionsTable = ProdTab{
	ProdTabEntry{
		String:     `S' : BoolExpr ;`,
		Id:         "S'",
		NTType:     0,
		Index:      0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `BoolExpr : Val << X[0], nil >> ;`,
		Id:         "BoolExpr",
		NTType:     1,
		Index:      1,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `BoolExpr : BoolExpr "&" BoolExpr << ast.NewBoolAndExpr(X[0], X[2]) >> ;`,
		Id:         "BoolExpr",
		NTType:     1,
		Index:      2,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewBoolAndExpr(X[0], X[2])
		},
	},
	ProdTabEntry{
		String:     `BoolExpr : BoolExpr "|" BoolExpr << ast.NewBoolOrExpr(X[0], X[2]) >> ;`,
		Id:         "BoolExpr",
		NTType:     1,
		Index:      3,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewBoolOrExpr(X[0], X[2])
		},
	},
	ProdTabEntry{
		String:     `BoolExpr : "(" BoolExpr ")" << ast.NewBoolGroupExpr(X[1]) >> ;`,
		Id:         "BoolExpr",
		NTType:     1,
		Index:      4,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewBoolGroupExpr(X[1])
		},
	},
	ProdTabEntry{
		String:     `Val : "true" << ast.TRUE, nil >> ;`,
		Id:         "Val",
		NTType:     2,
		Index:      5,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.TRUE, nil
		},
	},
	ProdTabEntry{
		String:     `Val : "false" << ast.FALSE, nil >> ;`,
		Id:         "Val",
		NTType:     2,
		Index:      6,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.FALSE, nil
		},
	},
	ProdTabEntry{
		String:     `Val : CompareExpr << X[0], nil >> ;`,
		Id:         "Val",
		NTType:     2,
		Index:      7,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `Val : SubStringExpr << X[0], nil >> ;`,
		Id:         "Val",
		NTType:     2,
		Index:      8,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `CompareExpr : int_lit "<" int_lit << ast.NewLessThanExpr(X[0], X[2]) >> ;`,
		Id:         "CompareExpr",
		NTType:     3,
		Index:      9,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewLessThanExpr(X[0], X[2])
		},
	},
	ProdTabEntry{
		String:     `CompareExpr : int_lit ">" int_lit << ast.NewLessThanExpr(X[2], X[0]) >> ;`,
		Id:         "CompareExpr",
		NTType:     3,
		Index:      10,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewLessThanExpr(X[2], X[0])
		},
	},
	ProdTabEntry{
		String:     `SubStringExpr : string_lit "in" string_lit << ast.NewSubStringExpr(X[0], X[2]) >> ;`,
		Id:         "SubStringExpr",
		NTType:     4,
		Index:      11,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewSubStringExpr(X[0], X[2])
		},
	},
}
