package ast

import (
	"fmt"
	"code.google.com/p/gocc/frontend/token"
)

type LexImport struct {
	Id      string
	ExtFunc string
}

func NewLexImport(regDefId, extFunc interface{}) (*LexImport, error) {
	return &LexImport{
		Id:      string(regDefId.(*token.Token).Lit),
		ExtFunc: getExtFunc(extFunc),
	}, nil
}

func (this *LexImport) IsTerminal() bool {
	return true
}

func (this *LexImport) String() string {
	return fmt.Sprintf("%s \"%s\"", this.Id, this.ExtFunc)
}

func getExtFunc(strLit interface{}) string {
	lit := strLit.(*token.Token).Lit
	return string(lit[1 : len(lit)-1])
}
