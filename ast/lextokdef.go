package ast

import (
	"bytes"
	"fmt"
	"code.google.com/p/gocc/frontend/token"
)

type LexTokDef struct {
	id      string
	pattern *LexPattern
}

func NewLexTokDef(tokId, lexPattern interface{}) (*LexTokDef, error) {
	tokDef := &LexTokDef{
		id:      string(tokId.(*token.Token).Lit),
		pattern: lexPattern.(*LexPattern),
	}
	return tokDef, nil
}

func NewLexStringLitTokDef(tokId string) *LexTokDef {
	runes := bytes.Runes([]byte(tokId))
	alt, _ := NewLexAlt(newLexCharLitFromRune(runes[0]))
	for i := 1; i < len(runes); i++ {
		alt, _ = AppendLexTerm(alt, newLexCharLitFromRune(runes[i]))
	}
	ptrn, _ := NewLexPattern(alt)
	return &LexTokDef{
		id:      tokId,
		pattern: ptrn,
	}
}

// TODO:
// func (this *LexTokDef) Equal(that Node) bool {
// 	that1, ok := that.(*LexTokDef)
// 	if !ok {
// 		return false
// 	}
// 	if this.id != that1.id {
// 		return false
// 	}
// 	return this.pattern.Equal(that1.pattern)
// }

// func (*LexTokDef) RegDef() bool {
// 	return false
// }

func (*LexTokDef) RegDef() bool {
	return false
}

func (this *LexTokDef) String() string {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "%s : %s", this.id, this.pattern.String())
	return buf.String()
}

func (this *LexTokDef) Id() string {
	return this.id
}

func (this *LexTokDef) LexPattern() *LexPattern {
	return this.pattern
}
