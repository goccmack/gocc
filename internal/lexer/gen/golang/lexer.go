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
	"path"
	"text/template"

	"github.com/goccmack/gocc/internal/config"
	"github.com/goccmack/gocc/internal/io"
	"github.com/goccmack/gocc/internal/lexer/items"
)

func genLexer(pkg, outDir string, itemsets *items.ItemSets, cfg config.Config) {
	tmpl, err := template.New("lexer").Parse(lexerSrc[1:])
	if err != nil {
		panic(err)
	}
	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, getLexerData(pkg, outDir, itemsets, cfg))
	if err != nil {
		panic(err)
	}
	io.WriteFile(path.Join(outDir, "lexer", "lexer.go"), buf.Bytes())
}

func getLexerData(pkg, outDir string, itemsets *items.ItemSets, cfg config.Config) *lexerData {
	lexSymbols := itemsets.Symbols().List()
	return &lexerData{
		Debug:       cfg.DebugLexer(),
		TokenImport: path.Join(pkg, "token"),
		UtilImport:  path.Join(pkg, "util"),
		NumStates:   itemsets.Size(),
		NumSymbols:  len(lexSymbols),
		Symbols:     lexSymbols,
	}
}

type lexerData struct {
	Debug       bool
	TokenImport string
	UtilImport  string
	NumStates   int
	NumSymbols  int
	NextState   []byte
	Symbols     []string
}

const lexerSrc string = `
// Code generated by gocc; DO NOT EDIT.

package lexer

import (
{{if .Debug}}	"fmt"
{{end}}	"io/ioutil"
	"unicode/utf8"

{{if .Debug}}	"{{.UtilImport}}"
{{end}}	"{{.TokenImport}}"
)

const (
	NoState    = -1
	NumStates  = {{.NumStates}}
	NumSymbols = {{.NumSymbols}}
)

type Lexer struct {
	src    []byte
	pos    int
	line   int
	column int
}

func NewLexer(src []byte) *Lexer {
	lexer := &Lexer{
		src:    src,
		pos:    0,
		line:   1,
		column: 1,
	}
	return lexer
}

func NewLexerFile(fpath string) (*Lexer, error) {
	src, err := ioutil.ReadFile(fpath)
	if err != nil {
		return nil, err
	}
	return NewLexer(src), nil
}

func (l *Lexer) Scan() (tok *token.Token) {
	{{- if .Debug}}
	fmt.Printf("Lexer.Scan() pos=%d\n", l.pos)
	{{- end}}
	tok = new(token.Token)
	if l.pos >= len(l.src) {
		tok.Type = token.EOF
		tok.Pos.Offset, tok.Pos.Line, tok.Pos.Column = l.pos, l.line, l.column
		return
	}
	start, startLine, startColumn, end := l.pos, l.line, l.column, 0
	tok.Type = token.INVALID
	state, rune1, size := 0, rune(-1), 0
	for state != -1 {
		{{- if .Debug}}
		fmt.Printf("\tpos=%d, line=%d, col=%d, state=%d\n", l.pos, l.line, l.column, state)
		{{- end}}
		if l.pos >= len(l.src) {
			rune1 = -1
		} else {
			rune1, size = utf8.DecodeRune(l.src[l.pos:])
			l.pos += size
		}

		nextState := -1
		if rune1 != -1 {
			nextState = TransTab[state](rune1)
		}
		{{- if .Debug}}
		fmt.Printf("\tS%d, : tok=%s, rune == %s(%x), next state == %d\n", state, token.TokMap.Id(tok.Type), util.RuneToString(rune1), rune1, nextState)
		fmt.Printf("\t\tpos=%d, size=%d, start=%d, end=%d\n", l.pos, size, start, end)
		if nextState != -1 {
			fmt.Printf("\t\taction:%s\n", ActTab[nextState].String())
		}
		{{- end}}
		state = nextState

		if state != -1 {

			switch rune1 {
			case '\n':
				l.line++
				l.column = 1
			case '\r':
				l.column = 1
			case '\t':
				l.column += 4
			default:
				l.column++
			}

			switch {
			case ActTab[state].Accept != -1:
				tok.Type = ActTab[state].Accept
				end = l.pos
			case ActTab[state].Ignore != "":
				start, startLine, startColumn = l.pos, l.line, l.column
				state = 0
				if start >= len(l.src) {
					tok.Type = token.EOF
				}

			}
		} else {
			if tok.Type == token.INVALID {
				end = l.pos
			}
		}
	}
	if end > start {
		l.pos = end
		tok.Lit = l.src[start:end]
	} else {
		tok.Lit = []byte{}
	}
	tok.Pos.Offset, tok.Pos.Line, tok.Pos.Column = start, startLine, startColumn

	{{- if .Debug}}
	fmt.Printf("Token at %s: %s \"%s\"\n", tok.String(), token.TokMap.Id(tok.Type), tok.Lit)
	{{- end}}

	return
}

func (l *Lexer) Reset() {
	l.pos = 0
}

/*
Lexer symbols:
{{- range $i, $sym := .Symbols}}
{{- printf "\n%d: %s" $i $sym}}
{{- end}}
*/
`
