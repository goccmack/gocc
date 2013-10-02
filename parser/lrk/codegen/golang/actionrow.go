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
	"code.google.com/p/gocc/parser/lrk/action"
	"code.google.com/p/gocc/parser/lrk/states"
	"code.google.com/p/gocc/parser/symbols"
	"fmt"
	"text/template"
)

func genActionRow(prods []*ast.SyntaxBasicProd, symbols *symbols.Symbols, state *states.State, actions map[string]action.Action) string {
	wr := new(bytes.Buffer)
	tmpl, err := template.New("parser action row").Parse(actionRowSrc)
	if err != nil {
		panic(err)
	}
	tmpl.Execute(wr, getActionRowData(prods, state, symbols, actions))
	return wr.String()
}

type actRow struct {
	CanRecover bool
	Actions    []string
}

func getActionRowData(prods []*ast.SyntaxBasicProd, state *states.State, symbols *symbols.Symbols, actions map[string]action.Action) (data *actRow) {
	data = &actRow{
		CanRecover: state.CanRecover(),
		Actions:    make([]string, len(symbols.ListTerminals())),
	}
	for i, sym := range symbols.ListTerminals() {
		if actions[sym] == nil {
			data.Actions[i] = fmt.Sprintf("nil,\t\t/* %s */", sym)
		} else {
			switch act := actions[sym].(type) {
			case action.Accept:
				data.Actions[i] = fmt.Sprintf("accept(true),\t\t/* %s */", sym)
			case action.Reduce:
				data.Actions[i] = fmt.Sprintf("reduce(%d),\t\t/* %s, reduce: %s */", int(act), sym, prods[int(act)].Id)
			case action.Shift:
				data.Actions[i] = fmt.Sprintf("shift(%d),\t\t/* %s */", int(act), sym)
			default:
				panic(fmt.Sprintf("Unknown action type: %T", act))
			}
		}
	}
	return
}

const actionRowSrc = `		canRecover: {{printf "%t" .CanRecover}},
		actions: [numTerminals]action{
			{{range $a := .Actions}}{{$a}}
			{{end}}
		},
`
