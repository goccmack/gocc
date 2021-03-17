package ast

import (
	"fmt"

	"github.com/goccmack/gocc/example/usercontext/token"
)

type Attrib interface{}

// Demonstrates how Lexical and Parser context might differ.
type LexicalContext struct {
	SourceFile     string
	ForbiddenWords []string
}

// Source implements the lexer.Sourcer interface.
func (lc *LexicalContext) Source() string {
	return lc.SourceFile
}

type ParserContext struct {
	ImportedFiles        []*token.Token // files and where we imported them from
	ExtensionToForbidden map[string]string
	Visitors             []string
	CallbackFn           func()
}

// Simple type that will convey both kinds of context.
type Identifier struct {
	Name          *token.Token
	ParserContext *ParserContext
}

func NewIdentifier(name Attrib, context interface{}) (*Identifier, error) {
	switch pc := context.(type) {
	case *ParserContext:
		pc.Visitors = append(pc.Visitors, string(name.(*token.Token).Lit))
		if pc.CallbackFn != nil {
			pc.CallbackFn()
		}
		return &Identifier{Name: name.(*token.Token), ParserContext: pc}, nil

	default:
		return nil, fmt.Errorf("%s: NewIdentifier expected ParserContext, got %+v",
			string(name.(*token.Token).Lit), context)
	}
}
