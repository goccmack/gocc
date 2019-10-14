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
	"fmt"
	"go/format"
	"path"
	"text/template"

	"github.com/maxcalandrelli/gocc/internal/ast"
	"github.com/maxcalandrelli/gocc/internal/io"
	"github.com/maxcalandrelli/gocc/internal/parser/lr1/action"
	"github.com/maxcalandrelli/gocc/internal/parser/lr1/items"
	"github.com/maxcalandrelli/gocc/internal/token"
)

type actionTableData struct {
	Rows             []*actRow
	CdTokenFunctions []string
}

type tmplData struct {
	Header string
	Table  *actionTableData
}

func GenActionTable(outDir string, prods ast.SyntaxProdList, itemSets *items.ItemSets, tokMap *token.TokenMap, subpath, header string) map[int]items.RowConflicts {
	tmpl, err := template.New("parser action table").Parse(actionTableSrc[1:])
	if err != nil {
		panic(err)
	}
	wr := new(bytes.Buffer)
	data, conflicts := getActionTableData(prods, itemSets, tokMap)
	if err := tmpl.Execute(wr, tmplData{Table: data, Header: header}); err != nil {
		panic(err)
	}
	source, err := format.Source(wr.Bytes())
	if err != nil {
		panic(fmt.Sprintf("%s in\n%s", err.Error(), wr.String()))
	}
	io.WriteFile(path.Join(outDir, subpath, "parser", "actiontable.go"), source)
	return conflicts
}

func getActionTableData(prods ast.SyntaxProdList, itemSets *items.ItemSets,
	tokMap *token.TokenMap) (actTab *actionTableData, conflicts map[int]items.RowConflicts) {

	actTab = &actionTableData{
		Rows:             make([]*actRow, itemSets.Size()),
		CdTokenFunctions: make([]string, len(tokMap.TypeMap)),
	}
	conflicts = make(map[int]items.RowConflicts)
	var row *actRow
	cnflcts := items.RowConflicts{}
	for i := range actTab.Rows {
		if row, cnflcts = getActionRowData(prods, itemSets.Set(i), tokMap); len(cnflcts) > 0 {
			conflicts[i] = cnflcts
		}
		actTab.Rows[i] = row
	}
	return
}

type cdActionFunc struct {
	TokenType int
	TokenFunc string
}

type actRow struct {
	CanRecover bool
	Actions    []string
	CdActions  []cdActionFunc
}

func getActionRowData(prods ast.SyntaxProdList, set *items.ItemSet, tokMap *token.TokenMap) (data *actRow, conflicts items.RowConflicts) {
	data = &actRow{
		CanRecover: set.CanRecover(),
		Actions:    make([]string, len(tokMap.TypeMap)),
		CdActions:  []cdActionFunc{},
	}
	conflicts = make(items.RowConflicts)
	// calculate padding.
	for i, sym := range tokMap.TypeMap {
		act, symConflicts := set.Action(sym)
		if len(symConflicts) > 0 {
			conflicts[sym] = symConflicts
		}
		switch act1 := act.(type) {
		case action.Accept:
			data.Actions[i] = fmt.Sprintf("accept(true), // %s", sym)
		case action.Error:
			data.Actions[i] = fmt.Sprintf("nil, // %s", sym)
		case action.Reduce:
			data.Actions[i] = fmt.Sprintf("reduce(%d), // %s, reduce: %s", int(act1), sym, prods[int(act1)].Id)
		case action.Shift:
			data.Actions[i] = fmt.Sprintf("shift(%d), // %s", int(act1), sym)
			switch s := sym.(type) {
			case ast.SyntaxContextDependentTokId, ast.SyntaxSubParser:
				data.CdActions = append(data.CdActions, cdActionFunc{
					i,
					fmt.Sprintf("cdFunc_%s", s.SymbolString()),
				})
			}
		default:
			panic(fmt.Sprintf("Unknown action type: %T", act1))
		}
	}
	return
}

const actionTableSrc = `
// Code generated by gocc; DO NOT EDIT.

package parser


type (
	actionTable [numStates]actionRow
  cdFunc func(TokenStream, interface{}) (interface{}, error, []byte)
  cdAction struct {
    tokenIndex   int
    tokenScanner cdFunc
  }
	actionRow   struct {
		canRecover bool
		actions    [numSymbols]action
    cdActions  []cdAction
	}
  actions struct {
    table    actionTable
  }
)
var parserActions = actions {
  table: actionTable{
  	{{- range $i, $r := .Table.Rows }}
  	actionRow{ // S{{$i}}
  		canRecover: {{printf "%t" .CanRecover}},
  		actions: [numSymbols]action{
  			{{- range $a := .Actions }}
  			{{$a}}
  			{{- end }}
  		},
      cdActions: []cdAction{
  			{{- range $c := .CdActions }}
  			cdAction{tokenIndex: {{$c.TokenType}}, tokenScanner: {{$c.TokenFunc}} },
  			{{- end }}
      },
  	},
  	{{- end }}
  },
}

`
