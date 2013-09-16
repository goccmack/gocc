<<<<<<< HEAD
package parser

=======

package parser



>>>>>>> 9d3e28b6b4a375015991b8ecb296891776774570
type (
	//TODO: change type and variable names to be consistent with other tables
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
<<<<<<< HEAD
		Index      int
=======
		Index int
>>>>>>> 9d3e28b6b4a375015991b8ecb296891776774570
		NumSymbols int
		ReduceFunc func([]Attrib) (Attrib, error)
	}
	Attrib interface {
	}
)

<<<<<<< HEAD
var productionsTable = ProdTab{
	ProdTabEntry{
		String: `S' : A	<<  >>`,
		Id:         "S'",
		NTType:     0,
		Index:      0,
=======
var productionsTable = ProdTab {
	ProdTabEntry{
		String: `S' : A	<<  >>`,
		Id: "S'",
		NTType: 0,
		Index: 0,
>>>>>>> 9d3e28b6b4a375015991b8ecb296891776774570
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `A : B "c"	<< []interface{}{X[0], "c"}, nil >>`,
<<<<<<< HEAD
		Id:         "A",
		NTType:     1,
		Index:      1,
=======
		Id: "A",
		NTType: 1,
		Index: 1,
>>>>>>> 9d3e28b6b4a375015991b8ecb296891776774570
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return []interface{}{X[0], "c"}, nil
		},
	},
	ProdTabEntry{
		String: `B : empty	<<  >>`,
<<<<<<< HEAD
		Id:         "B",
		NTType:     2,
		Index:      2,
=======
		Id: "B",
		NTType: 2,
		Index: 2,
>>>>>>> 9d3e28b6b4a375015991b8ecb296891776774570
		NumSymbols: 0,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return nil, nil
		},
	},
	ProdTabEntry{
		String: `B : "b"	<< "b", nil >>`,
<<<<<<< HEAD
		Id:         "B",
		NTType:     2,
		Index:      3,
=======
		Id: "B",
		NTType: 2,
		Index: 3,
>>>>>>> 9d3e28b6b4a375015991b8ecb296891776774570
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return "b", nil
		},
	},
<<<<<<< HEAD
=======
	
>>>>>>> 9d3e28b6b4a375015991b8ecb296891776774570
}
