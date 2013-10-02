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
	"code.google.com/p/gocc/token"
	"fmt"
	"path"
	"text/template"
)

func GenToken(pkg, outdir string, tokMap *token.TokenMap) {
	tokenPath := path.Join(outdir, "token", "token.go")
	tmpl, err := template.New("token").Parse(TokenMapSrc)
	if err != nil {
		panic(err)
	}
	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, TokenData{TypMap: tokMap.TypeMap, IdMap: typeMap(tokMap)})
	io.WriteFile(tokenPath, buf.Bytes())
}

func typeMap(tokMap *token.TokenMap) []string {
	tm := make([]string, len(tokMap.TypeMap))
	for i, sym := range tokMap.TypeMap {
		tm[i] = fmt.Sprintf("\"%s\": %d", sym, i)
	}
	return tm
}

type TokenData struct {
	IdMap  []string
	TypMap []string
}

const TokenMapSrc string = `
package token

import(
	"fmt"
)

type Token struct {
	Type
	Lit []byte
	Pos
}

type Type int

const(
	INVALID Type = iota
	EOF
)

type Pos struct {
	Offset int
	Line int
	Column int
}

func (this Pos) String() string {
	return fmt.Sprintf("Pos(offset=%d, line=%d, column=%d)", this.Offset, this.Line, this.Column)
}

type TokenMap struct {
	typeMap  []string
	idMap map[string]Type
}

func (this TokenMap) Id(tok Type) string {
	if int(tok) < len(this.typeMap) - 1 {
		return this.typeMap[tok+1]
	}
	return "unknown"
}

func (this TokenMap) Type(tok string) Type {
	if typ, exist := this.idMap[tok]; exist {
		return typ
	}
	return INVALID
}

func (this TokenMap) TokenString(tok *Token) string {
	//TODO: refactor to print pos & token string properly
	return fmt.Sprintf("%s(%d,%s)", this.Id(tok.Type), tok.Type, tok.Lit)
}

func (this TokenMap) StringType(typ Type) string {
	return fmt.Sprintf("%s(%d)", this.Id(typ), typ)
}

var TokMap = TokenMap{
	typeMap: []string{
{{range .TypMap}}{{"\t\t"}}"{{printf "%s" .}}",{{"\n"}}{{end}}{{"\t"}}},

	idMap: map[string]Type {
{{range .IdMap}}{{"\t\t"}}{{printf "%s" .}},{{"\n"}}{{end}}{{"\t"}}},
}

`
