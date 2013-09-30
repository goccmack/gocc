package ast

import (
	"code.google.com/p/gocc/frontend/token"
)

type LexRegDefId struct {
	Id string
}

func NewLexRegDefId(regDefId interface{}) (*LexRegDefId, error) {
	return &LexRegDefId{
		Id: string(regDefId.(*token.Token).Lit),
	}, nil
}

func (this *LexRegDefId) IsTerminal() bool {
	return false
}

func (this *LexRegDefId) String() string {
	return this.Id
}
