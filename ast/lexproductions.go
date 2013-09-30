package ast

import (
	"bytes"
	"fmt"
)

type LexProductions struct {
	Productions []LexProduction
}

func NewLexProductions(lexProd interface{}) (*LexProductions, error) {
	return &LexProductions{
		Productions: []LexProduction{lexProd.(LexProduction)},
	}, nil
}

func newLexProductions() *LexProductions {
	return &LexProductions{
		Productions: []LexProduction{},
	}
}

func AppendLexProduction(lexProds, prod interface{}) (*LexProductions, error) {
	lp := lexProds.(*LexProductions)
	lp.Productions = append(lp.Productions, prod.(LexProduction))
	return lp, nil
}

func (this *LexProductions) String() string {
	w := new(bytes.Buffer)
	for _, prod := range this.Productions {
		fmt.Fprintf(w, "%s ;", prod.String())
	}
	return w.String()
}
