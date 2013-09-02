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
	"code.google.com/p/gocc/config"
	"code.google.com/p/gocc/io"
	"code.google.com/p/gocc/lexer/items"
	"path"
	"text/template"
)

func genLexer(pkg, outDir string, itemsets *items.ItemSets, cfg config.Config) {
	tmpl, err := template.New("lexer").Parse(lexerSrc)
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
package lexer

import (
	{{if .Debug}}
	"fmt"
	"{{.UtilImport}}"
	{{else}}
	// "fmt"
	// "{{.UtilImport}}"
	{{end}}
	"io/ioutil"
	"unicode/utf8"
	"{{.TokenImport}}"
)

const(
	NoState = -1
	NumStates = {{.NumStates}}
	NumSymbols = {{.NumSymbols}}
) 

type Lexer struct {
	src             []byte
	pos             int
	line            int
	column          int
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

func (this *Lexer) Scan() (tok *token.Token) {
	{{if .Debug}}
	fmt.Printf("Lexer.Scan() pos=%d\n", this.pos)
	{{else}}
	// fmt.Printf("Lexer.Scan() pos=%d\n", this.pos)
	{{end}}
	tok = new(token.Token)
	if this.pos >= len(this.src) {
		tok.Type = token.EOF
		tok.Pos.Offset, tok.Pos.Line, tok.Pos.Column = this.pos, this.line, this.column
		return
	}
	start, end := this.pos, 0
	tok.Type = token.INVALID
	state, rune1, size := 0, rune(-1), 0
	for state != -1 {
	{{if .Debug}}
		fmt.Printf("\tpos=%d, line=%d, col=%d, state=%d\n", this.pos, this.line, this.column, state)
	{{else}}
		// fmt.Printf("\tpos=%d, line=%d, col=%d, state=%d\n", this.pos, this.line, this.column, state)
	{{end}}
		if this.pos >= len(this.src) {
			rune1 = -1
		} else {
			rune1, size = utf8.DecodeRune(this.src[this.pos:])
			this.pos += size
		}
		switch rune1 {
		case '\n':
			this.line++
			this.column = 1
		case '\r':
			this.column = 1
		case '\t':
			this.column += 4
		default:
			this.column++
		}

		nextState := -1
		if rune1 != -1 {
			nextState = TransTab[state](rune1)
		}

	{{if .Debug}}
		fmt.Printf("\tS%d, : tok=%s, rune == %s(%x), next state == %d\n", state, token.TokMap.Id(tok.Type), util.RuneToString(rune1), rune1, nextState)
		fmt.Printf("\t\tpos=%d, size=%d, start=%d, end=%d\n", this.pos, size, start, end)
		if nextState != -1 {
			fmt.Printf("\t\taction:%s\n", ActTab[nextState].String())
		}
	{{else}}
		// fmt.Printf("\tS%d, : tok=%s, rune == %s(%x), next state == %d\n", state, token.TokMap.Id(tok.Type), util.RuneToString(rune1), rune1, nextState)
		// fmt.Printf("\t\tpos=%d, size=%d, start=%d, end=%d\n", this.pos, size, start, end)
		// if nextState != -1 {
		// 	fmt.Printf("\t\taction:%s\n", ActTab[nextState].String())
		// }
	{{end}}

		state = nextState

		if state != -1 {
			switch {
			case ActTab[state].Accept != -1:
				tok.Type = ActTab[state].Accept
				// fmt.Printf("\t Accept(%s), %s(%d)\n", string(act), token.TokMap.Id(tok), tok)
				end = this.pos
			case ActTab[state].Ignore != "":
				// fmt.Printf("\t Ignore(%s)\n", string(act))
				start = this.pos
				state = 0
				if start >= len(this.src) {
					tok.Type = token.EOF
				}

			}
		} else {
			if tok.Type == token.INVALID {
				end = this.pos
			}
		}
	}
	if end > start {
		this.pos = end
		tok.Lit = this.src[start:end]
	} else {
		tok.Lit = []byte{}
	}
	tok.Pos.Offset = start
	tok.Pos.Column = this.column
	tok.Pos.Line = this.line
	return
}

func (this *Lexer) Reset() {
	this.pos = 0
}

/*
Lexer symbols:
{{range $i, $sym := .Symbols}}{{printf "%d: %s" $i $sym}}{{"\n"}}{{end}}
*/
`
