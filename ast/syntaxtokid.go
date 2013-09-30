package ast

import (
	"code.google.com/p/gocc/frontend/token"
)

type SyntaxTokId string

func NewTokId(tokId interface{}) (SyntaxTokId, error) {
	return SyntaxTokId(string(tokId.(*token.Token).Lit)), nil
}

func (this SyntaxTokId) String() string {
	return string(this)
}
