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
	"code.google.com/p/gocc/parser/lrk/action"
	"code.google.com/p/gocc/parser/lrk/states"
	"code.google.com/p/gocc/parser/symbols"
	"path"
	"text/template"
)

func genActionTable(outDir string, prods []*ast.SyntaxBasicProd, symbols *symbols.Symbols, states *states.States, actions action.Actions) {
	tmpl, err := template.New("parser action table").Parse(actionTableSrc)
	if err != nil {
		panic(err)
	}
	wr := new(bytes.Buffer)
	tmpl.Execute(wr, getActionTableData(prods, symbols, states, actions))
	io.WriteFile(path.Join(outDir, "parser", "actiontable.go"), wr.Bytes())
}

type actionTableData struct {
	Rows []string
}

func getActionTableData(prods []*ast.SyntaxBasicProd, symbols *symbols.Symbols, states *states.States, actions action.Actions) (actTab *actionTableData) {
	actTab = &actionTableData{
		Rows: make([]string, states.Size()),
	}
	for i := range actTab.Rows {
		actTab.Rows[i] = genActionRow(prods, symbols, states.List[i], actions[i])
	}
	return
}

const actionTableSrc = `
package parser

type(
	actionTable [numStates]actionRow
	actionRow struct {
		canRecover bool
		actions [numTerminals]action
	}
)

var actionTab = actionTable{
	{{range $i, $r := .Rows}}actionRow{ // S{{$i}}
		{{$r}}
	},
	{{end}}
}

`
