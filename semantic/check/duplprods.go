//Copyright 2013 Vastech SA (PTY) LTD
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package check

import (
	"code.google.com/p/gocc/ast"
	"code.google.com/p/gocc/semantic/errors"
	"fmt"
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
