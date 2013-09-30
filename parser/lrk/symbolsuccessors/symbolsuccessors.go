package symbolsuccessors

// key: symbol, value: []state no
type SymbolSuccessors map[string][]int

func NewSymbolSuccessors() SymbolSuccessors {
	return make(SymbolSuccessors)
}
