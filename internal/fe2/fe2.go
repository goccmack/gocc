package fe2

import (
	"github.com/maxcalandrelli/gocc/internal/fe2/lexer"
	"github.com/maxcalandrelli/gocc/internal/fe2/parser"
)

func NewParser() *parser.Parser {
	return parser.NewParser()
}

func NewLexer(src []byte) *lexer.Lexer {
	return lexer.NewLexer(src)
}

func NewLexerFile(fpath string) (*lexer.Lexer, error) {
	return lexer.NewLexerFile(fpath)
}
