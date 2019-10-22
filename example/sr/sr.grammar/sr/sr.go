// Code generated by gocc; DO NOT EDIT.

package sr

import (
	"io"

	"github.com/maxcalandrelli/gocc/example/sr/sr.grammar/sr/internal/errors"
	"github.com/maxcalandrelli/gocc/example/sr/sr.grammar/sr/internal/lexer"
	"github.com/maxcalandrelli/gocc/example/sr/sr.grammar/sr/internal/parser"
	"github.com/maxcalandrelli/gocc/example/sr/sr.grammar/sr/internal/token"
)

const (
	INVALID = token.INVALID
	EOF     = token.EOF
)

type (
	Token       = token.Token
	TokenMap    = token.TokenMap
	Pos         = token.Pos
	ErrorSymbol = errors.ErrorSymbol
	Error       = errors.Error
	Lexer       = lexer.Lexer
	Parser      = parser.Parser
)

func NewParser() *parser.Parser {
	return parser.NewParser()
}

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

func GetTokenMap() TokenMap {
	return token.TokMap
}
