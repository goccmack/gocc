package semantic

import (
	"code.google.com/p/gocc/ast"
	"code.google.com/p/gocc/parser/symbols"
	"code.google.com/p/gocc/semantic/check"
	"code.google.com/p/gocc/semantic/errors"
)

func Check(g *ast.Grammar, basicProds []*ast.SyntaxBasicProd, s *symbols.Symbols) (errs []*errors.Error) {
	errs = append(errs, check.DuplicateProductions(g)...)
	errs = append(errs, check.UnreachableProductions(g.SyntaxPart.ProdList, s)...)
	return
}
