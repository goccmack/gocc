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
	"code.google.com/p/gocc/parser/lr1/items"
	"code.google.com/p/gocc/parser/symbols"
	"text/template"
)

func genGotoRow(itemSet *items.ItemSet, sym *symbols.Symbols) string {
	tmpl, err := template.New("parser goto table row").Parse(gotoRowSrc)
	if err != nil {
		panic(err)
	}
	wr := new(bytes.Buffer)
	tmpl.Execute(wr, getGotoRowData(itemSet, sym))
	return wr.String()
}

type gotoRowElement struct {
	NT    string
	State int
}

func getGotoRowData(itemSet *items.ItemSet, sym *symbols.Symbols) []gotoRowElement {
	row := make([]gotoRowElement, sym.NumNTSymbols())
	for i, nt := range sym.NTList() {
		row[i].NT = nt
		row[i].State = itemSet.NextSetIndex(nt)
	}
	return row
}

const gotoRowSrc = `
		{{range $i, $gto := .}}{{$gto.State}}, // {{$gto.NT}}
		{{end}}
`
