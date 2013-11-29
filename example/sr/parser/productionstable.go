
package parser

import "code.google.com/p/gocc/example/sr/ast"

type (
	//TODO: change type and variable names to be consistent with other tables
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index int
		NumSymbols int
		ReduceFunc func([]interface{}) (interface{}, error)
	}
)

var productionsTable = ProdTab {
	ProdTabEntry{
		String: `S' : Stmt ;`,
		Id: "S'",
		NTType: 0,
		Index: 0,
		NumSymbols: 1,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Stmt : "if" id "then" Stmt << ast.NewIf(X[1], X[3]), nil >> ;`,
		Id: "Stmt",
		NTType: 1,
		Index: 1,
		NumSymbols: 4,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
			return ast.NewIf(X[1], X[3]), nil
		},
	},
	ProdTabEntry{
		String: `Stmt : "if" id "then" Stmt "else" Stmt << ast.NewIfElse(X[1], X[3], X[5]), nil >> ;`,
		Id: "Stmt",
		NTType: 1,
		Index: 2,
		NumSymbols: 6,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
			return ast.NewIfElse(X[1], X[3], X[5]), nil
		},
	},
	ProdTabEntry{
		String: `Stmt : id << ast.NewIdStmt(X[0]), nil >> ;`,
		Id: "Stmt",
		NTType: 1,
		Index: 3,
		NumSymbols: 1,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
			return ast.NewIdStmt(X[0]), nil
		},
	},
	
}
