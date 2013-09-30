package check

import (
	"fmt"
	"code.google.com/p/gocc/ast"
	"code.google.com/p/gocc/parser/symbols"
	"code.google.com/p/gocc/semantic/errors"
)

func UnreachableProductions(prods ast.SyntaxProdList, s *symbols.Symbols) (errs []*errors.Error) {
	for i, prod := range prods {
		if i > 0 && !s.IsNT(string(prod.Id.Lit)) {
			errs = append(errs,
				&errors.Error{
					Msg:    fmt.Sprintf("Unreachable production: %s", string(prod.Id.Lit)),
					Offset: prod.Id.Offset,
				})
		}
	}
	return
}
