package parser

type (
	ProdTabU      []ProdTabUEntry
	ProdTabUEntry struct {
		String     string
		Head       NT
		HeadIndex	int
		NumSymbols int
		ReduceFunc func([]Attrib) (Attrib, error)
	}
)
