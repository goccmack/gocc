package ast

import (
	"bytes"
	"fmt"
	"code.google.com/p/gocc/frontend/token"
)

type SyntaxPart struct {
	Header   *FileHeader
	ProdList SyntaxProdList
}

func NewSyntaxPart(header, prodList interface{}) (*SyntaxPart, error) {
	sp := &SyntaxPart{}
	if header != nil {
		sp.Header = header.(*FileHeader)
	} else {
		sp.Header = new(FileHeader)
	}
	if prodList != nil {
		sp.ProdList = prodList.(SyntaxProdList)
	}

	return sp, nil
}

func (this *SyntaxPart) augment() *SyntaxPart {
	startProd := &SyntaxProdNonBasic{
		Id: &token.Token{
			Type: token.TokMap.Type("prodId"),
			Lit:  []byte("S'"),
		},
		SyntaxExpression: SyntaxExpression{
			&SyntaxBody{
				Terms: SyntaxTerms{SyntaxProdId(this.ProdList[0].Id.Lit)},
			},
		},
	}
	newProdList := SyntaxProdList{startProd}
	return &SyntaxPart{
		Header:   this.Header,
		ProdList: append(newProdList, this.ProdList...),
	}
}

func (this *SyntaxPart) String() string {
	w := new(bytes.Buffer)
	fmt.Fprintf(w, "/* Syntax Part */\n\n")
	fmt.Fprintf(w, "%s\n\n", this.Header)
	for _, prod := range this.ProdList {
		fmt.Fprintf(w, "%s\n\n", prod)
	}
	return w.String()
}
