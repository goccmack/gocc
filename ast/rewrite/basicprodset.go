package rewrite

import (
	"code.google.com/p/gocc/ast"
)

type basicProdSet []*ast.SyntaxBasicProd

func (this basicProdSet) add(prod *ast.SyntaxBasicProd) basicProdSet {
	if this == nil {
		return basicProdSet{prod}
	} else {
		if this.contain(prod) {
			return this
		} else {
			return append(this, prod)
		}
	}
}

func (this basicProdSet) contain(prod *ast.SyntaxBasicProd) bool {
	for _, p := range this {
		if p.Equal(prod) {
			return true
		}
	}
	return false
}
