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
		String:     `S' : X ;`,
		Id:         "S'",
		NTType:     0,
		Index:      0,
		NumSymbols: 1,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `X : "a" Y "d" ;`,
		Id:         "X",
		NTType:     1,
		Index:      1,
		NumSymbols: 3,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `X : "a" Z "c" ;`,
		Id:         "X",
		NTType:     1,
		Index:      2,
		NumSymbols: 3,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `X : "a" T ;`,
		Id:         "X",
		NTType:     1,
		Index:      3,
		NumSymbols: 2,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `X : "b" Y "e" ;`,
		Id:         "X",
		NTType:     1,
		Index:      4,
		NumSymbols: 3,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `X : "b" Z "d" ;`,
		Id:         "X",
		NTType:     1,
		Index:      5,
		NumSymbols: 3,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `X : "b" T ;`,
		Id:         "X",
		NTType:     1,
		Index:      6,
		NumSymbols: 2,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `Y : "t" W ;`,
		Id:         "Y",
		NTType:     2,
		Index:      7,
		NumSymbols: 2,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `Y : "u" X ;`,
		Id:         "Y",
		NTType:     2,
		Index:      8,
		NumSymbols: 2,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `Z : "t" "u" ;`,
		Id:         "Z",
		NTType:     3,
		Index:      9,
		NumSymbols: 2,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `T : "u" X "a" ;`,
		Id:         "T",
		NTType:     4,
		Index:      10,
		NumSymbols: 3,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String:     `W : "u" ;`,
		Id:         "W",
		NTType:     5,
		Index:      11,
		NumSymbols: 1,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
			return X[0], nil
		},
	},
}
