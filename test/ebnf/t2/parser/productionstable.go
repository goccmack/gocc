
package parser



type (
	//TODO: change type and variable names to be consistent with other tables
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index int
		NumSymbols int
		ReduceFunc func([]Attrib) (Attrib, error)
	}
	Attrib interface {
	}
)

var productionsTable = ProdTab {
	ProdTabEntry{
		String: `S' : A ;`,
		Id: "S'",
		NTType: 0,
		Index: 0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `A : A~1 ;`,
		Id: "A",
		NTType: 1,
		Index: 1,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `A : ;`,
		Id: "A",
		NTType: 1,
		Index: 2,
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String: `A~1 : A~1_RepTerm << []interface{}{X[0]}, nil >> ;`,
		Id: "A~1",
		NTType: 2,
		Index: 3,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return []interface{}{X[0]}, nil
		},
	},
	ProdTabEntry{
		String: `A~1 : A~1 A~1_RepTerm << append(X[0].([]interface{}), X[1]), nil >> ;`,
		Id: "A~1",
		NTType: 2,
		Index: 4,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return append(X[0].([]interface{}), X[1]), nil
		},
	},
	ProdTabEntry{
		String: `A~1_RepTerm : "a" ;`,
		Id: "A~1_RepTerm",
		NTType: 3,
		Index: 5,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	
}
