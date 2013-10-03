package parser

import "code.google.com/p/gocc/test/ebnf/t3/ast"

type (
	//TODO: change type and variable names to be consistent with other tables
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index      int
		NumSymbols int
		ReduceFunc func([]interface{}) (interface{}, error)
	}
)

var productionsTable = ProdTab{
	ProdTabEntry{
		String:     `S' : A ;`,
		Id:         "S'",
		NTType:     0,
		Index:      0,
		NumSymbols: 1,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `A : A~1 A~3 << ast.Test(X...) >> ;`,
		Id:         "A",
		NTType:     1,
		Index:      1,
		NumSymbols: 2,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
			return ast.Test(X...)
		},
	},
	ProdTabEntry{
		String:     `A : A~2 << ast.Test(X...) >> ;`,
		Id:         "A",
		NTType:     1,
		Index:      2,
		NumSymbols: 1,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
			return ast.Test(X...)
		},
	},
	ProdTabEntry{
		String:     `A : << ast.Test(X...) >> ;`,
		Id:         "A",
		NTType:     1,
		Index:      3,
		NumSymbols: 0,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String:     `A~2 : A~2_RepTerm << []interface{}{X[0]}, nil >> ;`,
		Id:         "A~2",
		NTType:     2,
		Index:      4,
		NumSymbols: 1,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
			return []interface{}{X[0]}, nil
		},
	},
	ProdTabEntry{
		String:     `A~2 : A~2 A~2_RepTerm << append(X[0].([]interface{}), X[1]), nil >> ;`,
		Id:         "A~2",
		NTType:     2,
		Index:      5,
		NumSymbols: 2,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
			return append(X[0].([]interface{}), X[1]), nil
		},
	},
	ProdTabEntry{
		String:     `A~2_RepTerm : "a" ;`,
		Id:         "A~2_RepTerm",
		NTType:     3,
		Index:      6,
		NumSymbols: 1,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `A~1 : "b" ;`,
		Id:         "A~1",
		NTType:     4,
		Index:      7,
		NumSymbols: 1,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `A : A~4 << ast.Test(X...) >> ;`,
		Id:         "A",
		NTType:     1,
		Index:      8,
		NumSymbols: 1,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
			return ast.Test(X...)
		},
	},
	ProdTabEntry{
		String:     `A : << ast.Test(X...) >> ;`,
		Id:         "A",
		NTType:     1,
		Index:      9,
		NumSymbols: 0,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String:     `A~4 : "b" ;`,
		Id:         "A~4",
		NTType:     5,
		Index:      10,
		NumSymbols: 1,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `A~3 : A~3_RepTerm << []interface{}{X[0]}, nil >> ;`,
		Id:         "A~3",
		NTType:     6,
		Index:      11,
		NumSymbols: 1,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
			return []interface{}{X[0]}, nil
		},
	},
	ProdTabEntry{
		String:     `A~3 : A~3 A~3_RepTerm << append(X[0].([]interface{}), X[1]), nil >> ;`,
		Id:         "A~3",
		NTType:     6,
		Index:      12,
		NumSymbols: 2,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
			return append(X[0].([]interface{}), X[1]), nil
		},
	},
	ProdTabEntry{
		String:     `A~3_RepTerm : "a" ;`,
		Id:         "A~3_RepTerm",
		NTType:     7,
		Index:      13,
		NumSymbols: 1,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
			return X[0], nil
		},
	},
}
