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

	"github.com/maxcalandrelli/gocc/internal/config"
	"github.com/maxcalandrelli/gocc/internal/io"
)

type data struct {
	MyName             string
	Pkg                string
	Outdir             string
	InternalSubdir     string
	Config             config.Config
	HasSyntax          bool
	IfaceDir           string
	ExportWindowReader bool
}

func Gen(pkg, outdir, internal, iface string, cfg config.Config, hasSyntax bool) {
	d := data{
		MyName:             cfg.ProjectName(),
		Pkg:                pkg,
		Outdir:             outdir,
		InternalSubdir:     internal,
		Config:             cfg,
		HasSyntax:          hasSyntax,
		IfaceDir:           iface,
		ExportWindowReader: false,
	}
	if !cfg.NoLexer() && hasSyntax {
		genMain(d)
	}
	genBase(d)
	genUtil(d)
	genIface(d)
}

func genBase(d data) {
	basePath := path.Join(d.Outdir, d.MyName+".go")
	tmpl, err := template.New(d.MyName).Parse(baseSrc[1:])
	if err != nil {
		panic(err)
	}
	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, d)
	source, err := format.Source(buf.Bytes())
	if err != nil {
		panic(err)
	}
	io.WriteFile(basePath, source)
}

func genUtil(d data) {
	basePath := path.Join(d.Outdir, "util.go")
	tmpl, err := template.New(d.MyName).Parse(stringUtilSrc[1:])
	if err != nil {
		panic(err)
	}
	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, d)
	source, err := format.Source(buf.Bytes())
	if err != nil {
		panic(err)
	}
	io.WriteFile(basePath, source)
}

func genMain(d data) {
	basePath := path.Join(d.Outdir, "main", d.MyName+".go")
	tmpl, err := template.New(d.MyName).Parse(mainSrc[1:])
	if err != nil {
		panic(err)
	}
	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, d)
	// Use go/format to indent the idMap literal correctly.
	source, err := format.Source(buf.Bytes())
	if err != nil {
		panic(err)
	}
	io.WriteFile(basePath, source)
}

func genIface(d data) {
	basePath := path.Join(d.Outdir, d.IfaceDir, fmt.Sprintf("%s.go", d.MyName))
	tmpl, err := template.New(d.MyName).Parse(ifaceSrc[1:])
	if err != nil {
		panic(err)
	}
	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, d)
	// Use go/format to indent the idMap literal correctly.
	source, err := format.Source(buf.Bytes())
	if err != nil {
		panic(err)
	}
	io.WriteFile(basePath, source)
}

const baseSrc string = `
// Code generated by gocc; DO NOT EDIT.

package {{.MyName}}

import (
  {{- if not .Config.NoLexer }}
  "io"
  {{- end}}

  "{{.Pkg}}/{{.InternalSubdir}}/token"
  {{- if not .Config.NoLexer }}
  "{{.Pkg}}/{{.InternalSubdir}}/lexer"
  {{- end}}
  {{- if .HasSyntax}}
  "{{.Pkg}}/{{.InternalSubdir}}/errors"
  "{{.Pkg}}/{{.InternalSubdir}}/parser"
  {{- end}}
  {{- if .ExportWindowReader }}
  "{{.Pkg}}/{{.InternalSubdir}}/io/stream"
  {{- end}}
)

const (
  INVALID = token.INVALID
  EOF = token.EOF
)

type (
  Token = token.Token
  TokenMap = token.TokenMap
  Pos = token.Pos
  {{- if .HasSyntax}}
  ErrorSymbol = errors.ErrorSymbol
  Error = errors.Error
  {{- end}}
  {{- if not .Config.NoLexer }}
  Lexer = lexer.Lexer
  {{- end}}
  {{- if .ExportWindowReader }}
  WindowReader = stream.WindowReader
  {{- end}}
  {{- if .HasSyntax}}
  Parser = parser.Parser
  {{- end}}
)

{{- if .HasSyntax}}
func NewParser() *parser.Parser {
  return parser.NewParser()
}

{{- if not .Config.NoLexer }}

func ParseFile(fpath string) (interface{}, error) {
  if lexer, err := NewLexerFile(fpath); err == nil {
    return NewParser().Parse(lexer)
  } else {
    return nil, err
  }
}

func ParseText(text string) (interface{}, error) {
  return NewParser().Parse(NewLexerBytes([]byte(text)))
}

func Parse(stream io.Reader) (interface{}, error) {
  lex, err := NewLexer(stream)
  if lex == nil {
    return nil, err
  }
  return NewParser().Parse(lex)
}

func ParseFileWithData(fpath string, userData interface{}) (interface{}, error) {
  if lexer, err := NewLexerFile(fpath); err == nil {
    return NewParser().SetContext(userData).Parse(lexer)
  } else {
    return nil, err
  }
}

func ParseTextWithData(text string, userData interface{}) (interface{}, error) {
  return NewParser().SetContext(userData).Parse(NewLexerBytes([]byte(text)))
}

func ParseWithData(stream io.Reader, userData interface{}) (interface{}, error) {
  lex, err := NewLexer(stream)
  if lex == nil {
    return nil, err
  }
  return NewParser().SetContext(userData).Parse(lex)
}


func ParseFilePartial(fpath string) (interface{}, error, []byte) {
  if lexer, err := NewLexerFile(fpath); err == nil {
    return NewParser().ParseLongestPrefix(lexer)
  } else {
    return nil, err, []byte{}
  }
}

func ParseTextPartial(text string) (interface{}, error, []byte) {
  return NewParser().ParseLongestPrefix(NewLexerBytes([]byte(text)))
}

func ParsePartial(stream io.Reader) (interface{}, error, []byte) {
  lex, err := NewLexer(stream)
  if lex == nil {
    return nil, err, []byte{}
  }
  return NewParser().ParseLongestPrefix(lex)
}

func ParseFileWithDataPartial(fpath string, userData interface{}) (interface{}, error, []byte) {
  if lexer, err := NewLexerFile(fpath); err == nil {
    return NewParser().SetContext(userData).ParseLongestPrefix(lexer)
  } else {
    return nil, err, []byte{}
  }
}

func ParseTextWithDataPartial(text string, userData interface{}) (interface{}, error, []byte) {
  return NewParser().SetContext(userData).ParseLongestPrefix(NewLexerBytes([]byte(text)))
}

func ParseWithDataPartial(stream io.Reader, userData interface{}) (interface{}, error, []byte) {
  lex, err := NewLexer(stream)
  if lex == nil {
    return nil, err, []byte{}
  }
  return NewParser().SetContext(userData).ParseLongestPrefix(lex)
}


{{- end}}
{{- end}}
{{- if not .Config.NoLexer }}

func NewLexerBytes(src []byte) *lexer.Lexer {
  return lexer.NewLexerBytes(src)
}

func NewLexerString(src string) *lexer.Lexer {
  return lexer.NewLexerBytes([]byte(src))
}

func NewLexerFile(fpath string) (*lexer.Lexer, error) {
  return lexer.NewLexerFile(fpath)
}

func NewLexer(reader io.Reader) (*lexer.Lexer, error) {
  return lexer.NewLexer(reader)
}
{{- end}}

{{- if .ExportWindowReader }}
func NewWindowReaderFromBytes(src []byte) WindowReader {
	return stream.NewWindowReaderFromBytes(src)
}

func NewWindowReader(rdr io.Reader) WindowReader {
	return stream.NewWindowReader(rdr)
}

func NewLimitedWindowReader(rdr io.Reader, sizeMin, sizeMax int) WindowReader {
	return stream.NewLimitedWindowReader(rdr, sizeMin, sizeMax)
}
{{- end}}


func GetTokenMap() TokenMap {
  return token.TokMap
}

`

const mainSrc string = `
package main

import (
  "flag"
  "fmt"
  "os"
  "strings"

  {{.MyName}} "{{.Pkg}}"
)

func showResult (r interface{}, e error, p []byte) {
  if e != nil {
    fmt.Fprintf(os.Stderr, "parsing returned the following error: %s\n", e.Error())
  } else {
    if len(p) > 0 {
      fmt.Printf("r=%#v, (%s)\n", r, string(p))
    } else {
      fmt.Printf("r=%#v\n", r)
    }
  }
}

var (
  File string
  Text string
	Longest bool
)

func parse (longest bool, lex *{{.MyName}}.Lexer) (res interface{}, err error, parsed []byte) {
  if longest {
    return {{.MyName}}.NewParser().ParseLongestPrefix(lex)
  } else {
    res, err = {{.MyName}}.NewParser().Parse(lex)
    parsed = []byte{}
  }
  return
}

func main () {
  flag.StringVar(&File, "file", "", "parse also text in file")
  flag.StringVar(&Text, "text", "", "parse also text given with flag")
	flag.BoolVar(&Longest, "longest", false, "parse longest possible part")
  flag.Parse()
  if Text > "" {
    showResult(parse(Longest, {{.MyName}}.NewLexerString(Text)))
  }
  if File > "" {
    l, e := {{.MyName}}.NewLexerFile(File)
    if e != nil { panic(e) }
    showResult(parse(Longest, l))
  }
  if str := strings.Join(flag.Args(), " "); str > "" {
    showResult(parse(Longest, {{.MyName}}.NewLexerString(str)))
  }
}

`

const ifaceSrc string = `
// Code generated by gocc; DO NOT EDIT.

package iface

import (
  "io"
  "{{.Pkg}}/{{.InternalSubdir}}/token"
  {{- if .HasSyntax}}
  "{{.Pkg}}/{{.InternalSubdir}}/errors"
  {{- end}}
  "{{.Pkg}}/{{.InternalSubdir}}/io/stream"
)

type (
  Token = token.Token
  TokenMap = token.TokenMap
  Pos = token.Pos
  {{- if .HasSyntax}}
  ErrorSymbol = errors.ErrorSymbol
  Error = errors.Error
  {{- end}}

  Scanner interface {
  	  Scan() (tok *Token)
  }

  StreamScanner interface {
    GetStream() TokenStream
  }

  CheckPoint interface {
    DistanceFrom(CheckPoint) int
    Advance(int)CheckPoint
  }

  CheckPointable interface {
  	GetCheckPoint() CheckPoint
    GotoCheckPoint(CheckPoint)
  }

  TokenStream interface {
    io.Reader
    	io.RuneScanner
    	io.Seeker
  }

)


const (
  INVALID = token.INVALID
  EOF = token.EOF
)

func GetTokenMap() TokenMap {
  return token.TokMap
}

func NewWindowReaderFromBytes(src []byte) stream.WindowReader {
	return stream.NewWindowReaderFromBytes(src)
}

func NewWindowReader(rdr io.Reader) stream.WindowReader {
	return stream.NewWindowReader(rdr)
}

func NewLimitedWindowReader(rdr io.Reader, sizeMin, sizeMax int) stream.WindowReader {
	return stream.NewLimitedWindowReader(rdr, sizeMin, sizeMax)
}


`

const stringUtilSrc string = `
// Code generated by gocc; DO NOT EDIT.

package {{.MyName}}

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func RuneToString(r rune) string {
	if r >= 0x20 && r < 0x7f {
		return fmt.Sprintf("'%c'", r)
	}
	switch r {
	case 0x07:
		return "'\\a'"
	case 0x08:
		return "'\\b'"
	case 0x0C:
		return "'\\f'"
	case 0x0A:
		return "'\\n'"
	case 0x0D:
		return "'\\r'"
	case 0x09:
		return "'\\t'"
	case 0x0b:
		return "'\\v'"
	case 0x5c:
		return "'\\\\\\'"
	case 0x27:
		return "'\\''"
	case 0x22:
		return "'\\\"'"
	}
	if r < 0x10000 {
		return fmt.Sprintf("\\u%04x", r)
	}
	return fmt.Sprintf("\\U%08x", r)
}

type EscapedString string

func (str EscapedString) Unquote() (unquoted string, wasQuoted bool, quoteType rune) {
	if len(str) > 1 {
		r := str[0]
		if (r == '"' || r == '\x60' || r == '\'') && r == str[len(str)-1] {
			str = str[1 : len(str)-1]
		}
		return string(str), true, rune(r)
	}
	return string(str), false, 0
}

func (str EscapedString) Unescape() string {
	var (
		res  string
		r    rune
		size int
	)
	for s := 0; s < len(str); s += size {
		if str[s] == '\\' {
			r, size = str[s+1:].EscapedFirstCharValue()
			size++
		} else {
			r, size = utf8.DecodeRuneInString(string(str)[s:])
		}
		res += string(r)
	}
	return res
}

func (str EscapedString) EscapedFirstCharValue() (rune, int) {
	var i, base, max uint32
	offset := 0
	switch str[offset] {
	case 'a':
		return '\a', 1
	case 'b':
		return '\b', 1
	case 'f':
		return '\f', 1
	case 'n':
		return '\n', 1
	case 'r':
		return '\r', 1
	case 't':
		return '\t', 1
	case 'v':
		return '\v', 1
	case '\\':
		return '\\', 1
	case '\'':
		return '\'', 1
	case '0', '1', '2', '3', '4', '5', '6', '7':
		i, base, max = 3, 8, 255
	case 'x':
		i, base, max = 2, 16, 255
		offset++
	case 'u':
		i, base, max = 4, 16, unicode.MaxRune
		offset++
	case 'U':
		i, base, max = 8, 32, unicode.MaxRune
		offset++
	default:
		panic(fmt.Sprintf("Error decoding character literal: %s\n", str))
	}

	var x uint32
	for ; i > 0 && offset < len(str)-1; i-- {
		ch, size := utf8.DecodeRuneInString(string(str)[offset:])
		offset += size
		d := uint32(EscapedString(ch).HexDigitValue())
		if d >= base {
			panic(fmt.Sprintf("charVal(%s): illegal character (%c) in escape sequence. size=%d, offset=%d", str, str[offset], size, offset))
		}
		x = x*base + d
	}
	if x > max || 0xD800 <= x && x < 0xE000 {
		panic(fmt.Sprintf("Error decoding escape char value. Lit:%s, offset:%d, escape sequence is invalid Unicode code point\n", str, offset))
	}

	return rune(x), offset
}

func (str EscapedString) HexDigitValue() int {
	ch := str[0]
	switch {
	case '0' <= ch && ch <= '9':
		return int(ch) - '0'
	case 'a' <= ch && ch <= 'f':
		return int(ch) - 'a' + 10
	case 'A' <= ch && ch <= 'F':
		return int(ch) - 'A' + 10
	}
	return 16 // larger than any legal digit val
}

`
