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

/*
This package controls the generation of all parser-related code.
*/
package gen

import (
	"github.com/maxcalandrelli/gocc/internal/ast"
	"github.com/maxcalandrelli/gocc/internal/config"
	"github.com/maxcalandrelli/gocc/internal/parser/gen/golang"
	"github.com/maxcalandrelli/gocc/internal/parser/lr1/items"
	"github.com/maxcalandrelli/gocc/internal/parser/symbols"
	"github.com/maxcalandrelli/gocc/internal/token"
)

func Gen(pkg, outDir, header string, prods ast.SyntaxProdList, symbols *symbols.Symbols,
	itemsets *items.ItemSets, tokMap *token.TokenMap, cfg config.Config, subpath string) (conflicts map[int]items.RowConflicts) {

	golang.GenAction(outDir, subpath)
	conflicts = golang.GenActionTable(outDir, prods, itemsets, tokMap, subpath, header)
	golang.GenErrors(pkg, outDir, subpath)
	golang.GenGotoTable(outDir, itemsets, symbols, subpath)
	golang.GenParser(pkg, outDir, prods, itemsets, symbols, cfg, subpath)
	golang.GenProductionsTable(pkg, outDir, header, prods, symbols, itemsets, tokMap, subpath)
	return
}
