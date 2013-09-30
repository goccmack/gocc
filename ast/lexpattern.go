package ast

import (
	"bytes"
	"fmt"
)

type LexPattern struct {
	Alternatives []*LexAlt
}

func NewLexPattern(lexAlt interface{}) (*LexPattern, error) {
	return &LexPattern{
		Alternatives: []*LexAlt{lexAlt.(*LexAlt)},
	}, nil
}

func AppendLexAlt(lexPattern, lexAlt interface{}) (*LexPattern, error) {
	lp := lexPattern.(*LexPattern)
	lp.Alternatives = append(lp.Alternatives, lexAlt.(*LexAlt))
	return lp, nil
}

func (this *LexPattern) String() string {
	buf := new(bytes.Buffer)
	for i, alt := range this.Alternatives {
		if i > 0 {
			fmt.Fprintf(buf, " | ")
		}
		fmt.Fprintf(buf, "%s", alt.String())
	}
	return buf.String()
}
