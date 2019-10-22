//Copyright 2012 Vastech SA (PTY) LTD
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
	"github.com/maxcalandrelli/gocc/internal/config"
	"github.com/maxcalandrelli/gocc/internal/io"
	"github.com/maxcalandrelli/gocc/internal/parser/lr1/items"
	"github.com/maxcalandrelli/gocc/internal/parser/symbols"
)

func GenParser(pkg, outDir string, prods ast.SyntaxProdList, itemSets *items.ItemSets, symbols *symbols.Symbols, cfg config.Config, internal, iface string) {
	tmpl, err := template.New("parser").Parse(parserSrc[1:])
	if err != nil {
		panic(err)
	}
	wr := new(bytes.Buffer)
	if err := tmpl.Execute(wr, getParserData(pkg, internal, iface, prods, itemSets, symbols, cfg)); err != nil {
		panic(err)
	}
	source, err := format.Source(wr.Bytes())
	if err != nil {
		panic(fmt.Sprintf("%s in\n%s", err.Error(), wr.String()))
	}
	io.WriteFile(path.Join(outDir, internal, "parser", "parser.go"), source)
}

type parserData struct {
	Debug          bool
	PkgPath        string
	InternalSubdir string
	IfaceSubdir    string
	Config         config.Config
	NumProductions int
	NumStates      int
	NumSymbols     int
	CdTokList      ast.SyntaxSymbols
	CdSubList      ast.SyntaxSymbols
	CdSubImports   map[string]string
	ErrorTokenName string
	MyName         string
}

func getParserData(pkg, internal, iface string, prods ast.SyntaxProdList, itemSets *items.ItemSets, symbols *symbols.Symbols, cfg config.Config) *parserData {
	ret := &parserData{
		Debug:          cfg.DebugParser(),
		PkgPath:        pkg,
		InternalSubdir: internal,
		IfaceSubdir:    iface,
		Config:         cfg,
		NumProductions: len(prods),
		NumStates:      itemSets.Size(),
		NumSymbols:     symbols.NumSymbols(),
		CdTokList:      symbols.ListContextDependentTokenSymbols(),
		CdSubList:      symbols.ListSubParserSymbols(),
		CdSubImports:   map[string]string{},
		ErrorTokenName: config.INTERNAL_SYMBOL_ERROR,
		MyName:         cfg.ProjectName(),
	}
	for _, sym := range ret.CdSubList {
		sub := sym.(ast.SyntaxSubParser)
		if imp, found := ret.CdSubImports[sub.Alias]; found {
			if imp != sub.Import {
				panic(fmt.Sprintf("alias %s cannot refer to %s, already used by %s", sub.Alias, sub.Import, imp))
			}
		} else {
			ret.CdSubImports[sub.Alias] = sub.Import
		}
	}
	return ret
}

const parserSrc = `
// Code generated by gocc; DO NOT EDIT.

package parser

import (
	"fmt"
	"strings"
  "errors"

	parseError "{{.PkgPath}}/{{.InternalSubdir}}/errors"
	"{{.PkgPath}}/{{.InternalSubdir}}/token"
	"{{.PkgPath}}/{{.IfaceSubdir}}"

  {{ range $alias, $import := .CdSubImports }}
    {{$alias}} "{{$import}}"
  {{end}}

)

const (
	numProductions = {{.NumProductions}}
	numStates      = {{.NumStates}}
	numSymbols     = {{.NumSymbols}}
)

// Stack

type stack struct {
	state  []int
	attrib []Attrib
}

const INITIAL_STACK_SIZE = 100

func newStack() *stack {
	return &stack{
		state:  make([]int, 0, INITIAL_STACK_SIZE),
		attrib: make([]Attrib, 0, INITIAL_STACK_SIZE),
	}
}

func (s *stack) reset() {
	s.state = s.state[:0]
	s.attrib = s.attrib[:0]
}

func (s *stack) push(state int, a Attrib) {
	s.state = append(s.state, state)
	s.attrib = append(s.attrib, a)
}

func (s *stack) top() int {
	return s.state[len(s.state)-1]
}

func (s *stack) peek(pos int) int {
	return s.state[pos]
}

func (s *stack) topIndex() int {
	return len(s.state) - 1
}

func (s *stack) popN(items int) []Attrib {
	lo, hi := len(s.state)-items, len(s.state)

	attrib := s.attrib[lo:hi]

	s.state = s.state[:lo]
	s.attrib = s.attrib[:lo]

	return attrib
}

func (s *stack) String() string {
	w := new(strings.Builder)
	fmt.Fprintf(w, "stack:\n")
	for i, st := range s.state {
		fmt.Fprintf(w, "\t%d: %d , ", i, st)
		if s.attrib[i] == nil {
			fmt.Fprintf(w, "nil")
		} else {
			switch attr := s.attrib[i].(type) {
			case *token.Token:
				fmt.Fprintf(w, "%s", attr.Lit)
			default:
				fmt.Fprintf(w, "%v", attr)
			}
		}
		fmt.Fprintf(w, "\n")
	}
	return w.String()
}

// Parser

type Parser struct {
	stack       *stack
	nextToken   *token.Token
  userContext interface{}

  //
  //  Working data, should be method parameters/locals
  //  but putting them here simplifies a bit the code
  //  Useful only for longest prefix partial parsing
  //
  needsRepositioning bool
  isNonDeterministic bool
  tokens           iface.CheckPointable
  afterPos         iface.CheckPoint
  checkPoint       iface.CheckPoint
  underlyingStream iface.TokenStream
}

type TokenStream = iface.TokenStream

type (
  fakeCp struct{}
  fakeCheckPointable struct{}
)

func (f fakeCheckPointable) GetCheckPoint() iface.CheckPoint {
  return fakeCp{}
}

func (f fakeCp) DistanceFrom(o iface.CheckPoint) int {
  return 0
}

func (f fakeCp) Advance (o int) iface.CheckPoint {
  return fakeCp{}
}

func (f fakeCheckPointable) GotoCheckPoint(iface.CheckPoint) {}

{{ range $c := .CdTokList }}
func {{ printf "cdFunc_%s" $c.SymbolString }} (Stream TokenStream, Context interface{}) (interface{}, error, []byte) {
  return {{ $c.ContexDependentParseFunctionCall }}
}
{{ end }}

{{ range $c := .CdSubList }}
func {{ printf "cdFunc_%s" $c.SymbolString }} (Stream TokenStream, Context interface{}) (interface{}, error, []byte) {
  return {{$c.Alias}}.ParseWithDataPartial(Stream, Context)
}
{{ end }}


func NewParser() *Parser {
	return NewParserWithContext(nil)
}

func NewParserWithContext(u interface{}) *Parser {
	p := &Parser{stack: newStack(), userContext: u }
	p.Reset()
	return p
}

func (p *Parser) Reset() {
	p.stack.reset()
	p.stack.push(0, nil)
}

func (p *Parser) SetContext (ctx interface{}) *Parser {
  p.userContext = ctx
  return p
}

func (p Parser) GetContext () interface{} {
  return p.userContext
}

func (p *Parser) Error(err error, scanner iface.Scanner) (recovered bool, errorAttrib *parseError.Error) {
	errorAttrib = &parseError.Error{
		Err:            err,
		ErrorToken:     p.nextToken,
		ErrorSymbols:   p.popNonRecoveryStates(),
		ExpectedTokens: make([]string, 0, 8),
	}
	for t, action := range parserActions.table[p.stack.top()].actions {
		if action != nil {
			errorAttrib.ExpectedTokens = append(errorAttrib.ExpectedTokens, token.TokMap.Id(token.Type(t)))
		}
	}

	if action := parserActions.table[p.stack.top()].actions[token.TokMap.Type("{{.ErrorTokenName}}")]; action != nil {
		p.stack.push(int(action.(shift)), errorAttrib) // action can only be shift
	} else {
		return
	}

	if action := parserActions.table[p.stack.top()].actions[p.nextToken.Type]; action != nil {
		recovered = true
	}
	for !recovered && p.nextToken.Type != token.EOF {
		p.nextToken = scanner.Scan()
		if action := parserActions.table[p.stack.top()].actions[p.nextToken.Type]; action != nil {
			recovered = true
		}
	}

	return
}

func (p *Parser) popNonRecoveryStates() (removedAttribs []parseError.ErrorSymbol) {
	if rs, ok := p.firstRecoveryState(); ok {
		errorSymbols := p.stack.popN(p.stack.topIndex() - rs)
		removedAttribs = make([]parseError.ErrorSymbol, len(errorSymbols))
		for i, e := range errorSymbols {
			removedAttribs[i] = e
		}
	} else {
		removedAttribs = []parseError.ErrorSymbol{}
	}
	return
}

// recoveryState points to the highest state on the stack, which can recover
func (p *Parser) firstRecoveryState() (recoveryState int, canRecover bool) {
	recoveryState, canRecover = p.stack.topIndex(), parserActions.table[p.stack.top()].canRecover
	for recoveryState > 0 && !canRecover {
		recoveryState--
		canRecover = parserActions.table[p.stack.peek(recoveryState)].canRecover
	}
	return
}

func (p *Parser) newError(err error) error {
	e := &parseError.Error{
		Err:        err,
		StackTop:   p.stack.top(),
		ErrorToken: p.nextToken,
	}
	actRow := parserActions.table[p.stack.top()]
	for i, t := range actRow.actions {
		if t != nil {
			e.ExpectedTokens = append(e.ExpectedTokens, token.TokMap.Id(token.Type(i)))
		}
	}
	return e
}

func (p *Parser) prepareParsing(scanner iface.Scanner, longest bool) error {
	p.Reset()
  {{- if and (eq (len .CdTokList) 0) (eq (len .CdSubList) 0) }}
  p.needsRepositioning = longest
  p.isNonDeterministic = false
  {{- else }}
  p.needsRepositioning = true
  p.isNonDeterministic = true
  {{- end }}
  if p.needsRepositioning {
    if streamScanner, _ := scanner.(iface.StreamScanner) ; streamScanner != nil {
      p.underlyingStream = streamScanner.GetStream()
    }
    if p.tokens, _ = scanner.(iface.CheckPointable) ; p.tokens == nil || p.underlyingStream == nil {
      return errNotRepositionable
    }
  } else {
    if p.tokens, _ = scanner.(iface.CheckPointable) ; p.tokens == nil {
      p.tokens = fakeCheckPointable{}
    }
  }
  return nil
}

func (p *Parser) Parse(scanner iface.Scanner) (res interface{}, err error) {
  if err := p.prepareParsing (scanner, false); err != nil {
    return nil, err
  }
  return p.parse(scanner, false)
}

func (p *Parser) ParseLongestPrefix(scanner iface.Scanner) (res interface{}, err error, parsed []byte) {
  if err := p.prepareParsing (scanner, true); err != nil {
    return nil, err, []byte{}
  }
  startPoint := p.tokens.GetCheckPoint()
  res, err = p.parse(scanner, true)
  if err == nil {
    currPoint := p.tokens.GetCheckPoint()
    p.tokens.GotoCheckPoint(startPoint)
    parsed = make([]byte, currPoint.DistanceFrom(startPoint))
    p.underlyingStream.Read(parsed)
    p.tokens.GotoCheckPoint(currPoint)
  }
  return
}

var errNotRepositionable = errors.New("scanner not repositionable")


func (p *Parser) parse(scanner iface.Scanner, longest bool) (res interface{}, err error) {
  readNextToken := func () {
    p.checkPoint = p.tokens.GetCheckPoint()
 	  p.nextToken = scanner.Scan()
    p.afterPos = p.tokens.GetCheckPoint()
    return
  }
  readNextToken()
	for acc := false; !acc; {
		action := parserActions.table[p.stack.top()].actions[p.nextToken.Type]
		if action == nil && p.isNonDeterministic {
      //
      // If no action, check if we have some context dependent parsing to try
      //
      p.checkPoint = p.checkPoint.Advance(len(p.nextToken.IgnoredPrefix))
  		{{- if .Debug }}
  		fmt.Printf("{{.MyName}}:parser:S%d advanced %d bytes because of prefix <%s> of token <%s>\n", p.stack.top(), len(p.nextToken.IgnoredPrefix), string(p.nextToken.IgnoredPrefix), string(p.nextToken.Lit))
  		{{- end }}
			for _, cdAction := range parserActions.table[p.stack.top()].cdActions {
    		{{- if .Debug }}
    		fmt.Printf("{{.MyName}}:parser:S%d trying=<%#v> \n", p.stack.top(), cdAction)
    		{{- end }}
				p.tokens.GotoCheckPoint(p.checkPoint)
				cd_res, cd_err, cd_parsed := cdAction.tokenScanner(p.underlyingStream, p.userContext)
				if cd_err == nil && len(cd_parsed) > 0 {
					action = parserActions.table[p.stack.top()].actions[cdAction.tokenIndex]
          if action != nil {
            p.nextToken.Foreign = true
            p.nextToken.ForeignAstNode = cd_res
            p.nextToken.Lit = cd_parsed
            p.nextToken.Type = token.Type(cdAction.tokenIndex)
        		{{- if .Debug }}
        		fmt.Printf("{{.MyName}}:parser:S%d got pseudo token=<%s> \n", p.stack.top(), string(p.nextToken.Lit))
        		{{- end }}
					  break
          }
				}
			}
    }
    //
    //  Still no action? If a longest possible parsing is requested in place
    //  of a full text, we should try to check if an EOF here would have led
    //  to a successful parsing instead
    //  Rember that token.EOF is 1, that is the index of SyntaxEof into symbol table
    //  Dirty, dirty, dirty... but I found it as it is, and I prefer not to touch
    //
		if action == nil && longest {
			p.tokens.GotoCheckPoint(p.checkPoint)
      action = parserActions.table[p.stack.top()].actions[token.EOF]
      if action == nil {
        //
        //  ok, let's consume the token then
        //
    		{{- if .Debug }}
    		fmt.Printf("{{.MyName}}:parser:S%d unrecognized=<%#v> restoring checkpoint at %#v => %s\n", p.stack.top(), p.nextToken, p.afterPos, action)
    		{{- end }}
				p.tokens.GotoCheckPoint(p.afterPos)
      }
    }

    //
    //  Well, no action is no action after all...
    //
    if action == nil {
			if recovered, errAttrib := p.Error(nil, scanner); !recovered {
				p.nextToken = errAttrib.ErrorToken
				return nil, p.newError(nil)
			}
			if action = parserActions.table[p.stack.top()].actions[p.nextToken.Type]; action == nil {
				panic("Error recovery led to invalid action")
			}
    }
		{{- if .Debug }}
		fmt.Printf("{{.MyName}}:parser:S%d %s %s\n", p.stack.top(), token.TokMap.TokenString(p.nextToken), action)
		{{- end }}

		switch act := action.(type) {
		case accept:
			res = p.stack.popN(1)[0]
			acc = true
		case shift:
		  p.stack.push(int(act), p.nextToken)
      if p.nextToken.Foreign {
			  p.stack.push(int(act), p.nextToken.ForeignAstNode)
      }
      readNextToken()
		case reduce:
			prod := productionsTable[int(act)]
			attrib, err := prod.ReduceFunc(p.userContext, p.stack.popN(prod.NumSymbols))
			if err != nil {
				return nil, p.newError(err)
			} else {
				p.stack.push(gotoTab[p.stack.top()][prod.NTType], attrib)
			}
		default:
			panic("unknown action: " + action.String())
		}
	}
	return res, nil
}
`
