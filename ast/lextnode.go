package ast

type LexTNode interface {
	LexNode
	lexSymbol()
}

func (*LexCharLit) lexSymbol()   {}
func (*LexCharRange) lexSymbol() {}
func (*LexDot) lexSymbol()       {}
func (*LexRegDefId) lexSymbol()  {}
