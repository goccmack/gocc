package ast

import (
	"bytes"
	"fmt"
)

type LexRepPattern struct {
	*LexPattern
}

func NewLexRepPattern(pattern interface{}) (*LexRepPattern, error) {
	return &LexRepPattern{
		LexPattern: pattern.(*LexPattern),
	}, nil
}

func (this *LexRepPattern) IsTerminal() bool {
	return false
}

func (this *LexRepPattern) String() string {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "{%s}", this.LexPattern.String())
	return buf.String()
}
