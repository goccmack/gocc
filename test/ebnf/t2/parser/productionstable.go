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
		String:     `A~1 : "a" << []interface{}{X[0]}, nil >> ;`,
		Id:         "A~1",
		NTType:     2,
		Index:      3,
		NumSymbols: 1,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
			return []interface{}{X[0]}, nil
		},
	},
	ProdTabEntry{
		String:     `A~1 : A~1 "a" << append(X[0].([]interface{}), X[1]), nil >> ;`,
		Id:         "A~1",
		NTType:     2,
		Index:      4,
		NumSymbols: 2,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
			return append(X[0].([]interface{}), X[1]), nil
		},
	},
}
