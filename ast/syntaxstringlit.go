package ast

import (
	"fmt"
	"code.google.com/p/gocc/frontend/token"
)

type SyntaxStringLit string

func NewStringLit(tok interface{}) (SyntaxStringLit, error) {
	lit := tok.(*token.Token).Lit
	return SyntaxStringLit(lit[1 : len(lit)-1]), nil
}

func (this SyntaxStringLit) String() string {
	return fmt.Sprintf("\"%s\"", string(this))
}

func (this SyntaxStringLit) Bytes() []byte {
	return []byte(this)
}
