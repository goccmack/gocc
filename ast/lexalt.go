package ast

import (
	"bytes"
	"fmt"
)

type LexAlt struct {
	Terms []LexTerm
}

func NewLexAlt(lexTerm interface{}) (*LexAlt, error) {
	return &LexAlt{
		Terms: []LexTerm{lexTerm.(LexTerm)},
	}, nil
}

func AppendLexTerm(lexAlt, lexTerm interface{}) (*LexAlt, error) {
	la := lexAlt.(*LexAlt)
	la.Terms = append(la.Terms, lexTerm.(LexTerm))
	return la, nil
}

func (this *LexAlt) Contain(term LexTerm) bool {
	for _, thisTerm := range this.Terms {
		if thisTerm == term {
			return true
		}
	}
	return false
}

func (this *LexAlt) String() string {
	buf := new(bytes.Buffer)
	for i, term := range this.Terms {
		if i > 0 {
			fmt.Fprintf(buf, " ")
		}
		fmt.Fprintf(buf, "%s", term.String())
	}
	return buf.String()
}
