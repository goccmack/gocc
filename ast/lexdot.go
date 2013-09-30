package ast

type LexDot struct{}

var LexDOT = &LexDot{}

func (this *LexDot) String() string {
	return "."
}

func (this *LexDot) IsTerminal() bool {
	return true
}
