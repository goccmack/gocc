package ast

import (
	"code.google.com/p/gocc/frontend/token"
)

// Id or name of a grammar(syntax) production
type SyntaxProdId string

func NewSyntaxProdId(tok interface{}) (SyntaxProdId, error) {
	return SyntaxProdId(string(tok.(*token.Token).Lit)), nil
}

func (this SyntaxProdId) String() string {
	return string(this)
}
