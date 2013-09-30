package ast

type LexNode interface {
	LexTerminal() bool
	String() string
}

func (LexAlt) LexTerminal() bool {
	return false
}

func (LexGroupPattern) LexTerminal() bool {
	return false
}

func (*LexIgnoredTokDef) LexTerminal() bool {
	return false
}

func (LexImports) LexTerminal() bool {
	return false
}

func (LexOptPattern) LexTerminal() bool {
	return false
}

func (LexPattern) LexTerminal() bool {
	return false
}

func (LexProductions) LexTerminal() bool {
	return false
}

func (*LexRegDef) LexTerminal() bool {
	return false
}

func (LexRepPattern) LexTerminal() bool {
	return false
}

func (*LexTokDef) LexTerminal() bool {
	return false
}

func (*LexCharLit) LexTerminal() bool {
	return true
}

func (*LexCharRange) LexTerminal() bool {
	return true
}

func (LexDot) LexTerminal() bool {
	return true
}

func (LexRegDefId) LexTerminal() bool {
	return true
}
