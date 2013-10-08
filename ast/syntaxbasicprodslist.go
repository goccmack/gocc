package ast

import (
	"bytes"
	"fmt"
)

type SyntaxBasicProdsList struct {
	first, last *SyntaxBasicProd
	length      int
}

func NewSyntaxBasicProdsList() *SyntaxBasicProdsList {
	return &SyntaxBasicProdsList{}
}

func (this *SyntaxBasicProdsList) Append(prods ...*SyntaxBasicProd) *SyntaxBasicProdsList {
	for _, prod := range prods {
		prod.Next = nil
		if this.first == nil {
			prod.Prev = nil
			this.first = prod
			this.last = prod
		} else {
			prod.Prev = this.last
			this.last.Next = prod
			this.last = prod
		}
		this.length++
	}
	return this
}

func (this *SyntaxBasicProdsList) First() *SyntaxBasicProd {
	return this.first
}

func (this *SyntaxBasicProdsList) Last() *SyntaxBasicProd {
	return this.last
}

func (this *SyntaxBasicProdsList) Len() int {
	return this.length
}

func (this *SyntaxBasicProdsList) List() []*SyntaxBasicProd {
	prods := make([]*SyntaxBasicProd, this.Len())
	for prod, i := this.First(), 0; prod != nil; prod, i = prod.Next, i+1 {
		prods[i] = prod
	}
	return prods
}

// Used for testing:
// func (this *SyntaxBasicProdsList) checkLength() {
// 	err := false
// 	len := 0
// 	for p := this.First(); p != nil; p = p.Next {
// 		len++
// 	}
// 	if len != this.length {
// 		err = true
// 		fmt.Printf("checkLength: exp=%d, fwd=%d\n", this.length, len)
// 	}
// 	len = 0
// 	for p := this.Last(); p != nil; p = p.Prev {
// 		len++
// 	}
// 	if len != this.length {
// 		err = true
// 		fmt.Printf("checkLength: exp=%d, bwd=%d\n", this.length, len)
// 	}
// 	if err {
// 		panic("checkLenght")
// 	}
// }

/*
Replace prod in this with prods. The SyntaxBasicProdList, prods, is destroyed in the process.
*/
func (this *SyntaxBasicProdsList) Replace(prod *SyntaxBasicProd, prods *SyntaxBasicProdsList) *SyntaxBasicProdsList {
	if this.first == prod {
		this.first = prods.first
	} else {
		prods.first.Prev = prod.Prev
		prod.Prev.Next = prods.first
	}
	prods.last.Next = prod.Next
	if prod.Next != nil {
		prod.Next.Prev = prods.last
	}
	if this.last == prod {
		this.last = prods.last
	}
	this.length += prods.length - 1

	return this
}

func (this *SyntaxBasicProdsList) String() string {
	w := new(bytes.Buffer)
	for prod := this.First(); prod != nil; prod = prod.Next {
		fmt.Fprintf(w, "%s\n", prod)
	}
	return w.String()
}
