package ast

import (
	"bytes"
	"fmt"
	"code.google.com/p/gocc/frontend/token"
)

type LexRegDef struct {
	id      string
	pattern *LexPattern
}

func NewLexRegDef(regDefId, lexPattern interface{}) (*LexRegDef, error) {
	regDef := &LexRegDef{
		id:      string(regDefId.(*token.Token).Lit),
		pattern: lexPattern.(*LexPattern),
	}
	return regDef, nil
}

// func (this *LexRegDef) Equal(that Node) bool {
// 	that1, ok := that.(LexRegDef)
// 	if !ok {
// 		return false
// 	}
// 	if this.id != that1.id {
// 		return false
// 	}
// 	return this.pattern.Equal(that1.pattern)
// }

func (*LexRegDef) RegDef() bool {
	return true
}

func (this *LexRegDef) String() string {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "%s : %s", this.id, this.pattern.String())
	return buf.String()
}

func (this *LexRegDef) Id() string {
	return this.id
}

func (this *LexRegDef) LexPattern() *LexPattern {
	return this.pattern
}
