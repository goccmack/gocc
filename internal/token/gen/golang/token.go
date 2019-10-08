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

	"github.com/maxcalandrelli/gocc/internal/io"
	"github.com/maxcalandrelli/gocc/internal/token"
)

func GenToken(pkg, outdir string, tokMap *token.TokenMap, subpath string) {
	tokenPath := path.Join(outdir, subpath, "token", "token.go")
	tmpl, err := template.New("token").Parse(TokenMapSrc[1:])
	if err != nil {
		panic(err)
	}
	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, TokenData{TypMap: makeTypeMap(tokMap), IdMap: makeIdMap(tokMap)})
	// Use go/format to indent the idMap literal correctly.
	source, err := format.Source(buf.Bytes())
	if err != nil {
		panic(fmt.Sprintf("%s in \n%s", err.Error(), buf.String()))
	}
	io.WriteFile(tokenPath, source)
}

func makeIdMap(tokMap *token.TokenMap) []string {
	tm := make([]string, len(tokMap.TypeMap))
	for i, sym := range tokMap.TypeMap {
		tm[i] = fmt.Sprintf("\"%s\": %d", sym.SymbolName(), i)
	}
	return tm
}

func makeTypeMap(tokMap *token.TokenMap) []string {
	tm := make([]string, len(tokMap.TypeMap))
	for i, sym := range tokMap.TypeMap {
		tm[i] = fmt.Sprintf("\"%s\"", sym.SymbolName())
	}
	return tm
}

type TokenData struct {
	IdMap  []string
	TypMap []string
}

const TokenMapSrc string = `
// Code generated by gocc; DO NOT EDIT.

package token

import (
	"fmt"
)

type Token struct {
	Type
	Lit []byte
	Pos
	ForeingAstNode  interface{}
	ForeingAstError error
}

type Type int

const (
	INVALID Type = iota
	EOF
)

type Pos struct {
	Offset int
	Line   int
	Column int
}

func (p Pos) String() string {
	return fmt.Sprintf("Pos(offset=%d, line=%d, column=%d)", p.Offset, p.Line, p.Column)
}

func (p Pos) StartingFrom(base Pos) Pos {
	r := base
	r.Offset += p.Offset
	r.Line += p.Line
	r.Column = p.Column
	if p.Line > 0 && base.Line > 0 {
		r.Line--
	}
	if r.Column < 1 {
		r.Column = 1
	}
	return r
}

type TokenMap struct {
	typeMap []string
	idMap   map[string]Type
}

func (m TokenMap) Id(tok Type) string {
	if int(tok) < len(m.typeMap) {
		return m.typeMap[tok]
	}
	return "unknown"
}

func (m TokenMap) Type(tok string) Type {
	if typ, exist := m.idMap[tok]; exist {
		return typ
	}
	return INVALID
}

func (m TokenMap) TokenString(tok *Token) string {
	return fmt.Sprintf("%s(%d,<%s>)", m.Id(tok.Type), tok.Type, tok.Lit)
}

func (m TokenMap) StringType(typ Type) string {
	return fmt.Sprintf("%s(%d)", m.Id(typ), typ)
}

var TokMap = TokenMap{
	typeMap: []string{
{{- range .TypMap }}
		{{printf "%s" .}},
{{- end }}
	},

	idMap: map[string]Type{
{{- range .IdMap }}
		{{printf "%s" .}},
{{- end }}

	},
}
`
