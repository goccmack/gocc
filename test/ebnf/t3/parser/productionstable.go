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
		String:     `A : << ast.Test(X...) >> ;`,
		Id:         "A",
		NTType:     1,
		Index:      1,
		NumSymbols: 0,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
			return nil, nil
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
		String:     `A~2 : "a" << []interface{}{X[0]}, nil >> ;`,
		Id:         "A~2",
		NTType:     2,
		Index:      3,
		NumSymbols: 1,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
			return []interface{}{X[0]}, nil
		},
	},
	ProdTabEntry{
		String:     `A~2 : A~2 "a" << append(X[0].([]interface{}), X[1]), nil >> ;`,
		Id:         "A~2",
		NTType:     2,
		Index:      4,
		NumSymbols: 2,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
			return append(X[0].([]interface{}), X[1]), nil
		},
	},
	ProdTabEntry{
		String:     `A : A~1 << ast.Test(X...) >> ;`,
		Id:         "A",
		NTType:     1,
		Index:      5,
		NumSymbols: 1,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
			return ast.Test(X...)
		},
	},
	ProdTabEntry{
		String:     `A : A~1 A~2 << ast.Test(X...) >> ;`,
		Id:         "A",
		NTType:     1,
		Index:      6,
		NumSymbols: 2,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
			return ast.Test(X...)
		},
	},
	ProdTabEntry{
		String:     `A~1 : "b" ;`,
		Id:         "A~1",
		NTType:     3,
		Index:      7,
		NumSymbols: 1,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
			return X[0], nil
		},
	},
}
