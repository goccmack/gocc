package ast

import (
	"bytes"
	"fmt"
)

type LexGroupPattern struct {
	*LexPattern
}

func NewLexGroupPattern(pattern interface{}) (*LexGroupPattern, error) {
	return &LexGroupPattern{pattern.(*LexPattern)}, nil
}

func (this *LexGroupPattern) IsTerminal() bool {
	return false
}

func (this *LexGroupPattern) String() string {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "(%s)", this.LexPattern.String())
	return buf.String()
}
