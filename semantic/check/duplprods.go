package check

import (
	"fmt"
	"code.google.com/p/gocc/ast"
	"code.google.com/p/gocc/semantic/errors"
)

func DuplicateProductions(g *ast.Grammar) (errs []*errors.Error) {
	ntMap := make(map[string]int)
	for _, prod := range g.SyntaxPart.ProdList {
		pid := string(prod.Id.Lit)
		if count, dup := ntMap[pid]; dup {
			errs = append(errs, &errors.Error{
				Msg:    fmt.Sprintf("Duplicate production Id: %s", pid),
				Offset: prod.Id.Pos.Offset,
			})
			ntMap[pid] = count + 1
		} else {
			ntMap[pid] = 1
		}
	}
	return
}
