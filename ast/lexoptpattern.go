package ast

import (
	"bytes"
	"fmt"
)

type LexOptPattern struct {
	*LexPattern
}

func NewLexOptPattern(pattern interface{}) (*LexOptPattern, error) {
	return &LexOptPattern{pattern.(*LexPattern)}, nil
}

func (this *LexOptPattern) IsTerminal() bool {
	return false
}

func (this *LexOptPattern) String() string {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "[%s]", this.LexPattern.String())
	return buf.String()
}
