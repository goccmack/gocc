package ast

import (
	"bytes"
	"fmt"
	"code.google.com/p/gocc/frontend/token"
)

type LexIgnoredTokDef struct {
	id      string
	pattern *LexPattern
}

func NewLexIgnoredTokDef(tokId, lexPattern interface{}) (*LexIgnoredTokDef, error) {
	tokDef := &LexIgnoredTokDef{
		id:      string(tokId.(*token.Token).Lit),
		pattern: lexPattern.(*LexPattern),
	}
	return tokDef, nil
}

// func (this *LexIgnoredTokDef) Equal(that Node) bool {
// 	that1, ok := that.(*LexIgnoredTokDef)
// 	if !ok {
// 		return false
// 	}
// 	if this.id != that1.id {
// 		return false
// 	}
// 	return this.pattern.Equal(that1.pattern)
// }

func (*LexIgnoredTokDef) RegDef() bool {
	return false
}

func (this *LexIgnoredTokDef) String() string {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "%s : %s", this.id, this.pattern.String())
	return buf.String()
}

func (this *LexIgnoredTokDef) Id() string {
	return this.id
}

func (this *LexIgnoredTokDef) LexPattern() *LexPattern {
	return this.pattern
}
