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
	genIface(d)
}

func genBase(d data) {
	basePath := path.Join(d.Outdir, "grammar.go")
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

func genMain(d data) {
	basePath := path.Join(d.Outdir, "main", "main.go")
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

func ParseFile(fpath string) (interface{}, error, int) {
  if lexer, err := NewLexerFile(fpath); err == nil {
    return NewParser().Parse(lexer)
  } else {
    return nil, err, 0
  }
}

func ParseText(text string) (interface{}, error, int) {
  return NewParser().Parse(NewLexerBytes([]byte(text)))
}

func Parse(stream io.Reader) (interface{}, error, int) {
  lex, err := NewLexer(stream)
  if lex == nil {
    return nil, err, 0
  }
  return NewParser().Parse(lex)
}

func ParseFileWithData(fpath string, userData interface{}) (interface{}, error, int) {
  if lexer, err := NewLexerFile(fpath); err == nil {
    return NewParser().SetContext(userData).Parse(lexer)
  } else {
    return nil, err, 0
  }
}

func ParseTextWithData(text string, userData interface{}) (interface{}, error, int) {
  return NewParser().SetContext(userData).Parse(NewLexerBytes([]byte(text)))
}

func ParseWithData(stream io.Reader, userData interface{}) (interface{}, error, int) {
  lex, err := NewLexer(stream)
  if lex == nil {
    return nil, err, 0
  }
  return NewParser().SetContext(userData).Parse(lex)
}


func ParseFilePartial(fpath string) (interface{}, error, int) {
  if lexer, err := NewLexerFile(fpath); err == nil {
    return NewParser().ParseLongestPrefix(lexer)
  } else {
    return nil, err, 0
  }
}

func ParseTextPartial(text string) (interface{}, error, int) {
  return NewParser().ParseLongestPrefix(NewLexerBytes([]byte(text)))
}

func ParsePartial(stream io.Reader) (interface{}, error, int) {
  lex, err := NewLexer(stream)
  if lex == nil {
    return nil, err, 0
  }
  return NewParser().ParseLongestPrefix(lex)
}

func ParseFileWithDataPartial(fpath string, userData interface{}) (interface{}, error, int) {
  if lexer, err := NewLexerFile(fpath); err == nil {
    return NewParser().SetContext(userData).ParseLongestPrefix(lexer)
  } else {
    return nil, err, 0
  }
}

func ParseTextWithDataPartial(text string, userData interface{}) (interface{}, error, int) {
  return NewParser().SetContext(userData).ParseLongestPrefix(NewLexerBytes([]byte(text)))
}

func ParseWithDataPartial(stream io.Reader, userData interface{}) (interface{}, error, int) {
  lex, err := NewLexer(stream)
  if lex == nil {
    return nil, err, 0
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

func showResult (r interface{}, e error, l int) {
  if e != nil {
    fmt.Fprintf(os.Stderr, "parsing returned the following error: %s\n", e.Error())
  } else {
    fmt.Printf("r=%#v, %d bytes\n", r, l)
  }
}

var (
  File string
  Text string
	Longest bool
)

func parse (longest bool, lex *{{.MyName}}.Lexer) (res interface{}, err error, ptl int) {
  if longest {
    return {{.MyName}}.NewParser().ParseLongestPrefix(lex)
  } else {
    return {{.MyName}}.NewParser().Parse(lex)
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


`
