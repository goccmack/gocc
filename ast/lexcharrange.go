package ast

import (
	"fmt"
)

type LexCharRange struct {
	From *LexCharLit
	To   *LexCharLit
	s    string
}

func NewLexCharRange(from, to interface{}) (*LexCharRange, error) {
	cr := &LexCharRange{
		From: newLexCharLit(from),
		To:   newLexCharLit(to),
	}

	return cr, nil
}

func (this *LexCharRange) IsTerminal() bool {
	return true
}

func (this *LexCharRange) String() string {
	if this.s == "" {
		this.s = fmt.Sprintf("%s-%s", this.From.String(), this.To.String())
	}
	return this.s
}
