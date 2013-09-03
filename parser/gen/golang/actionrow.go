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
	"code.google.com/p/gocc/parser/lr1/action"
	"code.google.com/p/gocc/parser/lr1/items"
	"code.google.com/p/gocc/parser/symbols"
	"fmt"
	"text/template"
)

func genActionRow(prods ast.SyntaxProdList, set *items.ItemSet, symbols *symbols.Symbols) (string, items.RowConflicts) {
	wr := new(bytes.Buffer)
	tmpl, err := template.New("parser action row").Parse(actionRowSrc)
	if err != nil {
		panic(err)
	}
	data, conflicts := getActionRowData(prods, set, symbols)
	tmpl.Execute(wr, data)
	return wr.String(), conflicts
}

type actRow struct {
	CanRecover bool
	Actions    []string
}

func getActionRowData(prods ast.SyntaxProdList, set *items.ItemSet, symbols *symbols.Symbols) (data *actRow, conflicts items.RowConflicts) {
	data = &actRow{
		CanRecover: set.CanRecover(),
		Actions:    make([]string, symbols.NumSymbols()),
	}
	conflicts = make(items.RowConflicts)
	for i, sym := range symbols.List() {
		act, symConflicts := set.Action(sym)
		if len(symConflicts) > 0 {
			conflicts[sym] = symConflicts
		}

		switch act1 := act.(type) {
		case action.Accept:
			data.Actions[i] = fmt.Sprintf("accept(true),\t\t/* %s */", sym)
		case action.Error:
			data.Actions[i] = fmt.Sprintf("nil,\t\t/* %s */", sym)
		case action.Reduce:
			data.Actions[i] = fmt.Sprintf("reduce(%d),\t\t/* %s, reduce: %s */", int(act1), sym, prods[int(act1)].Id)
		case action.Shift:
			data.Actions[i] = fmt.Sprintf("shift(%d),\t\t/* %s */", int(act1), sym)
		default:
			panic(fmt.Sprintf("Unknown action type: %T", act1))
		}
	}
	return
}

const actionRowSrc = `		canRecover: {{printf "%t" .CanRecover}},
		actions: [numSymbols]action{
			{{range $a := .Actions}}{{$a}}
			{{end}}
		},
`
