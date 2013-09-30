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
	"code.google.com/p/gocc/io"
	"code.google.com/p/gocc/lexer/symbols"
	"path"
	"text/template"
)

func genAsciiTable(outDir string, symbols *symbols.Symbols) {
	tmpl, err := template.New("ascii table").Parse(asciiTabSrc)
	if err != nil {
		panic(err)
	}
	w := new(bytes.Buffer)
	if err = tmpl.Execute(w, getAsciiTab(symbols)); err != nil {
		panic(err)
	}
	io.WriteFile(path.Join(outDir, "lexer", "asciitable.go"), w.Bytes())
}

func getAsciiTab(symbols *symbols.Symbols) map[rune]asciiType {
	asciiMap := make(map[rune]asciiType)
	for i, sym := range symbols.List() {
		if cl, exist := symbols.CharLitSymbols.GetSymbolId(sym); exist {
			if cl.Val < 0x100 {
				asciiMap[cl.Val] = asciiType{Type: i, Comment: sym}
			}
		}
	}
	return asciiMap
}

type asciiType struct {
	Type    int
	Comment string
}

const asciiTabSrc = `
package lexer

const asciiTabLen = 255

// asciiTable[rune] returns the lexer symbol type of rune
var asciiTable = func() *[asciiTabLen]int {
	atab := new([asciiTabLen] int)
	for i, _ := range atab {
		atab[i] = -1
	}
	{{range $run, $typ := .}}atab[{{$run}}] = {{$typ.Type}}		// {{$typ.Comment}}
	{{end}}
	return atab
}()

`