package parser

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
		String:     `A : ;`,
		Id:         "A",
		NTType:     1,
		Index:      1,
		NumSymbols: 0,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String:     `A : A~1 ;`,
		Id:         "A",
		NTType:     1,
		Index:      2,
		NumSymbols: 1,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `A~1 : C << []interface{}{X[0]}, nil >> ;`,
		Id:         "A~1",
		NTType:     2,
		Index:      3,
		NumSymbols: 1,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
			return []interface{}{X[0]}, nil
		},
	},
	ProdTabEntry{
		String:     `A~1 : A~1 C << append(X[0].([]interface{}), X[1]), nil >> ;`,
		Id:         "A~1",
		NTType:     2,
		Index:      4,
		NumSymbols: 2,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
			return append(X[0].([]interface{}), X[1]), nil
		},
	},
	ProdTabEntry{
		String:     `A : A~2 ;`,
		Id:         "A",
		NTType:     1,
		Index:      5,
		NumSymbols: 1,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `A~2 : D << []interface{}{X[0]}, nil >> ;`,
		Id:         "A~2",
		NTType:     3,
		Index:      6,
		NumSymbols: 1,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
			return []interface{}{X[0]}, nil
		},
	},
	ProdTabEntry{
		String:     `A~2 : A~2 D << append(X[0].([]interface{}), X[1]), nil >> ;`,
		Id:         "A~2",
		NTType:     3,
		Index:      7,
		NumSymbols: 2,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
			return append(X[0].([]interface{}), X[1]), nil
		},
	},
	ProdTabEntry{
		String:     `C : "c3" ;`,
		Id:         "C",
		NTType:     4,
		Index:      8,
		NumSymbols: 1,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `C : C~1 "c3" ;`,
		Id:         "C",
		NTType:     4,
		Index:      9,
		NumSymbols: 2,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `C~1 : "c1" ;`,
		Id:         "C~1",
		NTType:     5,
		Index:      10,
		NumSymbols: 1,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `C~1 : "c2" ;`,
		Id:         "C~1",
		NTType:     5,
		Index:      11,
		NumSymbols: 1,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `C~1 : D ;`,
		Id:         "C~1",
		NTType:     5,
		Index:      12,
		NumSymbols: 1,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `D : "d1" D~1 "d5" ;`,
		Id:         "D",
		NTType:     6,
		Index:      13,
		NumSymbols: 3,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `D~1 : "d2" ;`,
		Id:         "D~1",
		NTType:     7,
		Index:      14,
		NumSymbols: 1,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `D~1 : "d3" ;`,
		Id:         "D~1",
		NTType:     7,
		Index:      15,
		NumSymbols: 1,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `D~1 : "d4" ;`,
		Id:         "D~1",
		NTType:     7,
		Index:      16,
		NumSymbols: 1,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
			return X[0], nil
		},
	},
}
