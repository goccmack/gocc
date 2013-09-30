package ast

type LexTerm interface {
	LexNode

	//TODO: remove from all LexTerm classes
	// Equal(LexTerm) bool

	lexTerm()
}

func (*LexDot) lexTerm()          {}
func (*LexCharLit) lexTerm()      {}
func (*LexCharRange) lexTerm()    {}
func (*LexRegDefId) lexTerm()     {}
func (*LexOptPattern) lexTerm()   {}
func (*LexRepPattern) lexTerm()   {}
func (*LexGroupPattern) lexTerm() {}
