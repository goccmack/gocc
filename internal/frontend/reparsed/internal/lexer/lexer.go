// Code generated by gocc; DO NOT EDIT.

package lexer

import (
  
  "io"
  "bytes"
  "os"

  
  "github.com/maxcalandrelli/gocc/internal/frontend/reparsed/iface"
  "github.com/maxcalandrelli/gocc/internal/frontend/reparsed/internal/token"
  "github.com/maxcalandrelli/gocc/internal/frontend/reparsed/internal/io/stream"
)

const (
	NoState    = -1
	NumStates  = 114
	NumSymbols = 84
  INVALID_RUNE = rune(-1)
)

type position struct {
	token.Pos
}

type Lexer struct {
	position
	stream   iface.TokenStream
	eof      bool
}

func NewLexerBytes(src []byte) *Lexer {
	lexer := &Lexer{stream: bytes.NewReader(src)}
	lexer.reset()
	return lexer
}

func NewLexerString(src string) *Lexer {
	return NewLexerBytes([]byte(src))
}

func NewLexerFile(fpath string) (*Lexer, error) {
	s, err := os.Open(fpath)
	if err != nil {
		return nil, err
	}
	lexer := &Lexer{stream: stream.NewWindowReader(s)}
	lexer.reset()
	return lexer, nil
}

func NewLexer(reader io.Reader) (*Lexer, error) {
	lexer := &Lexer{}
	lexer.reset()
	if lexer.stream, _ = reader.(iface.TokenStream); lexer.stream == nil {
		lexer.stream = stream.NewWindowReader(reader)
	}
	return lexer, nil
}

func (l *Lexer) reset () {
  l.position.Reset()
}

func (l Lexer) GetStream() iface.TokenStream {
  return l.stream
}

type checkPoint int64

func (c checkPoint) value () int64 {
  return int64(c)
}

func (c checkPoint) DistanceFrom (o iface.CheckPoint) int {
  return int (c.value() - o.(checkPoint).value())
}

func (l *Lexer) GetCheckPoint() iface.CheckPoint {
  if l == nil {
    return checkPoint(0)
  }
  pos, _ := l.stream.Seek(0, io.SeekCurrent)
  return checkPoint(pos)
}

func (l Lexer) GotoCheckPoint(cp iface.CheckPoint) {
  l.stream.Seek(int64(cp.(checkPoint)), io.SeekStart)
}

func (l *Lexer) Scan() (tok *token.Token) {
	tok = new(token.Token)
	tok.Type = token.INVALID
	tok.Lit = []byte{}
  start := l.position
  state := 0
	for state != -1  {
	  curr, size, err := l.stream.ReadRune()
    if size < 1 || err != nil {
      curr = INVALID_RUNE
    }
		if size > 0 {
  		l.position.Pos.Offset += size
    }
		nextState := -1
		if curr != INVALID_RUNE {
			nextState = TransTab[state](curr)
		}
		state = nextState
		if state != -1 {
      	switch curr {
      	case '\n':
      		l.position.Pos.Line++
      		l.position.Pos.Column = 1
      	case '\r':
      		l.position.Pos.Column = 1
      	case '\t':
      		l.position.Pos.Column += 4
      	default:
      		l.position.Pos.Column++
      	}
			switch {
			case ActTab[state].Accept != -1:
				tok.Type = ActTab[state].Accept
				tok.Lit = append(tok.Lit, string(curr)...)
			case ActTab[state].Ignore != "":
				start = l.position
				state = 0
				tok.Lit = []byte{}
			}
		} else if curr != INVALID_RUNE{
      if len(tok.Lit) == 0 {
			  tok.Lit = append(tok.Lit, string(curr)...)
      } else {
        l.stream.UnreadRune()
      }
    }
  	if err == io.EOF && len(tok.Lit)==0 {
  		tok.Type = token.EOF
  		tok.Pos = start.Pos
  		return
  	}
	}
	tok.Pos = start.Pos
	return
}

func (l *Lexer) Reset() {
	l.position.Reset()
}

func (l Lexer) CurrentPosition() position {
	return l.position
}

func (p *position) Reset() {
	p.Offset = 0
	p.Line = 1
	p.Column = 1
}

func (p position) StartingFrom(base position) position {
	r := p
	r.Pos = p.Pos.StartingFrom(base.Pos)
	return r
}
