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

	"github.com/goccmack/gocc/internal/io"
	"github.com/goccmack/gocc/internal/token"
)

func GenToken(pkg, outdir string, tokMap *token.TokenMap) {
	tokenPath := path.Join(outdir, "token", "token.go")
	tmpl, err := template.New("token").Parse(TokenMapSrc[1:])
	if err != nil {
		panic(err)
	}
	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, TokenData{TypMap: tokMap.TypeMap, IdMap: typeMap(tokMap)})
	if err != nil {
		println(err.Error())
	}

	// Use go/format to indent the idMap literal correctly.
	source, err := format.Source(buf.Bytes())
	if err != nil {
		println(err.Error())
		// panic(err)
	}
	io.WriteFile(tokenPath, source)
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

//TODO(TokenMap.TokenString): refactor to print pos & token string properly
const TokenMapSrc string = `
// Code generated by gocc; DO NOT EDIT.

package token

import (
	"bytes"
	"fmt"
	"strconv"
	"unicode/utf8"
)

type Token struct {
	Type
	Lit []byte
	Pos
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
	return fmt.Sprintf("%s(%d,%s)", m.Id(tok.Type), tok.Type, tok.Lit)
}

func (m TokenMap) StringType(typ Type) string {
	return fmt.Sprintf("%s(%d)", m.Id(typ), typ)
}

// Equals returns returns true if the token Type and Lit are matches.
func (t *Token) Equals(rhs interface{}) bool {
	switch rhsT := rhs.(type) {
	case *Token:
		return t == rhsT || (t.Type == rhsT.Type && bytes.Equal(t.Lit, rhsT.Lit))
	default:
		return false
	}
}

// CharLiteralValue returns the string value of the char literal.
func (t *Token) CharLiteralValue() string {
	return string(t.Lit[1 : len(t.Lit)-1])
}

// Float32Value returns the float32 value of the token or an error if the token literal does not 
// denote a valid float32.
func (t *Token) Float32Value() (float32, error) {
	if v, err := strconv.ParseFloat(string(t.Lit), 32); err != nil {
		return 0, err
	} else {
		return float32(v), nil
	}
}

// Float64Value returns the float64 value of the token or an error if the token literal does not 
// denote a valid float64.
func (t *Token) Float64Value() (float64, error) {
	return strconv.ParseFloat(string(t.Lit), 64)
}

// IDValue returns the string representation of an identifier token.
func (t *Token) IDValue() string {
	return string(t.Lit)
}

// Int32Value returns the int32 value of the token or an error if the token literal does not
// denote a valid float64.
func (t *Token) Int32Value() (int32, error) {
	if v, err := strconv.ParseInt(string(t.Lit), 10, 64); err != nil {
		return 0, err
	} else {
		return int32(v), nil
	}
}

// Int64Value returns the int64 value of the token or an error if the token literal does not
// denote a valid float64.
func (t *Token) Int64Value() (int64, error) {
	return strconv.ParseInt(string(t.Lit), 10, 64)
}

// UTF8Rune decodes the UTF8 rune in the token literal. It returns utf8.RuneError if
// the token literal contains an invalid rune.
func (t *Token) UTF8Rune() (rune, error) {
	r, _ := utf8.DecodeRune(t.Lit)
	if r == utf8.RuneError {
		err := fmt.Errorf("Invalid rune")
		return r, err
	}
	return r, nil
}

// StringValue returns the string value of the token literal.
func (t *Token) StringValue() string {
	return string(t.Lit[1:len(t.Lit)-1])
}

var TokMap = TokenMap{
	typeMap: []string{
{{- range .TypMap }}
		{{printf "%q" .}},
{{- end }}
	},

	idMap: map[string]Type{
{{- range .IdMap }}
		{{printf "%s" .}},
{{- end }}
	},
}


`
