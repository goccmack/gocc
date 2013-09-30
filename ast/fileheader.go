package ast

import (
	"fmt"
	"code.google.com/p/gocc/frontend/token"
)

type FileHeader struct {
	Lit string
	str string
}

func NewFileHeader(actExpr interface{}) (*FileHeader, error) {
	sh := &FileHeader{
		Lit: ActionExpressionVal(actExpr.(*token.Token).Lit),
	}
	sh.str = fmt.Sprintf("<< %s >>", sh.Lit)
	return sh, nil
}

func (this *FileHeader) String() string {
	return this.str
}
