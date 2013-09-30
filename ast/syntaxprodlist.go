package ast

import (
	"bytes"
	"fmt"
)

type SyntaxProdList []*SyntaxProdNonBasic

func NewSyntaxProdList(prod interface{}) (SyntaxProdList, error) {
	return SyntaxProdList{prod.(*SyntaxProdNonBasic)}, nil
}

func AddSyntaxProds(prodList, prod interface{}) (SyntaxProdList, error) {
	return append(prodList.(SyntaxProdList), prod.(*SyntaxProdNonBasic)), nil
}

func (this SyntaxProdList) String() string {
	w := new(bytes.Buffer)
	for i, prod := range this {
		fmt.Fprintf(w, "%d: %s\n\n", i, prod.String())
	}
	return w.String()
}
