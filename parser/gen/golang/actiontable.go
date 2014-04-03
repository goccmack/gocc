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

package golang

import (
	"bytes"
	"code.google.com/p/gocc/ast"
	"code.google.com/p/gocc/io"
	"code.google.com/p/gocc/parser/lr1/items"
	"code.google.com/p/gocc/token"
	"path"
	"text/template"
)

func GenActionTable(outDir string, prods ast.SyntaxProdList, itemSets *items.ItemSets, tokMap *token.TokenMap) map[int]items.RowConflicts {
	tmpl, err := template.New("parser action table").Parse(actionTableSrc)
	if err != nil {
		panic(err)
	}
	wr := new(bytes.Buffer)
	data, conflicts := getActionTableData(prods, itemSets, tokMap)
	tmpl.Execute(wr, data)
	io.WriteFile(path.Join(outDir, "parser", "actiontable.go"), wr.Bytes())
	return conflicts
}

type actionTableData struct {
	Rows []string
}

func getActionTableData(prods ast.SyntaxProdList, itemSets *items.ItemSets,
	tokMap *token.TokenMap) (actTab *actionTableData, conflicts map[int]items.RowConflicts) {

	actTab = &actionTableData{
		Rows: make([]string, itemSets.Size()),
	}
	conflicts = make(map[int]items.RowConflicts)
	row, cnflcts := "", items.RowConflicts{}
	for i := range actTab.Rows {
		if row, cnflcts = genActionRow(prods, itemSets.Set(i), tokMap); len(cnflcts) > 0 {
			conflicts[i] = cnflcts
		}
		actTab.Rows[i] = row
	}
	return
}

const actionTableSrc = `
package parser

type(
	actionTable [numStates]actionRow
	actionRow struct {
		canRecover bool
		actions [numSymbols]action
	}
)

var actionTab = actionTable{
	{{range $i, $r := .Rows}}actionRow{ // S{{$i}}
		{{$r}}
	},
	{{end}}
}

`
