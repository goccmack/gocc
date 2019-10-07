// Code generated by gocc; DO NOT EDIT.

package lexer

import (
	"io/ioutil"
	"unicode/utf8"

	"github.com/maxcalandrelli/gocc/internal/frontend/reparsed/internal/token"
)

const (
	NoState    = -1
	NumStates  = 109
	NumSymbols = 75
)

type Position struct {
	pos    int
	line   int
	column int
}

type Lexer struct {
  Position
	src    []byte
}

func NewLexer(src []byte) *Lexer {
	lexer := &Lexer{ src: src }
  lexer.Position.Reset()
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

	return
}

func (l *Lexer) Reset() {
  l.Position.Reset()
}

func (l *Lexer) Reposition(p Position) {
  l.Position = p
}

func (l Lexer) CurrentPosition() Position {
  return l.Position
}

func (l Lexer) Source() []byte {
  return l.src
}

func (l Lexer) Remaining() []byte {
  return l.src[l.pos:]
}


func (p Position) Offset() int {
  return p.pos
}

func (p Position) Line() int {
  return p.line
}

func (p Position) Column() int {
  return p.column
}

func (p *Position) Reset()  {
	p.pos = 0
  p.line = 1
  p.column = 1
}