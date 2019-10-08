// Code generated by gocc; DO NOT EDIT.

package lexer

import (
	"io"
	"os"

	"github.com/maxcalandrelli/gocc/internal/frontend/reparsed/internal/io/stream"
	token "github.com/maxcalandrelli/gocc/internal/frontend/reparsed/internal/token/next"
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
	stream stream.WindowReader
	eof    bool
}

func NewLexer(src []byte) *Lexer {
	lexer := &Lexer{stream: stream.NewWindowReaderFromBytes(src)}
	return lexer
}

func NewLexerFile(fpath string) (*Lexer, error) {
	s, err := os.Open(fpath)
	if err != nil {
		return nil, err
	}
	return &Lexer{stream: stream.NewWindowReader(s)}, nil
}

func (l *Lexer) Scan() (tok *token.Token) {
	tok = new(token.Token)
	if l.eof {
		tok.Type = token.EOF
		tok.Pos.Offset, tok.Pos.Line, tok.Pos.Column = l.pos, l.line, l.column
		return
	}
	start, startLine, startColumn, end := l.pos, l.line, l.column, 0
	tok.Type = token.INVALID
	tok.Lit = []byte{}
	state, rune1 := 0, rune(-1)
	for state != -1 {
		if l.eof {
			rune1 = -1
		} else {
			rune2, size, err := l.stream.ReadRune()
			if err == io.EOF {
				l.eof = true
				err = nil
			}
			if err == nil && size > 0 {
				rune1 = rune2
				l.pos += size
				tok.Lit = append(tok.Lit, string(rune1)...)
			}
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
				if l.eof {
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

func (p Position) Offset() int {
	return p.pos
}

func (p Position) Line() int {
	return p.line
}

func (p Position) Column() int {
	return p.column
}

func (p *Position) Reset() {
	p.pos = 0
	p.line = 1
	p.column = 1
}
