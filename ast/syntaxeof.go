package ast

type SyntaxEof int

var EOF SyntaxEof = 0

func (SyntaxEof) SymbolsString() string {
	return "$"
}

func (SyntaxEof) String() string {
	return "$"
}
