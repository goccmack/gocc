package parser

import (
	"code.google.com/p/gocc/ast"
	"code.google.com/p/gocc/config"
	"code.google.com/p/gocc/parser/lrk"
	"code.google.com/p/gocc/parser/symbols"
)

func Gen(cfg config.Config, prods []*ast.SyntaxBasicProd, symbols *symbols.Symbols) {
	lrk.Gen(cfg, prods, symbols)
}
