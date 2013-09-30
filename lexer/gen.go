package lexer

import (
	"code.google.com/p/gocc/ast"
	"code.google.com/p/gocc/config"
	"code.google.com/p/gocc/io"
	"code.google.com/p/gocc/lexer/gen/golang"
	"code.google.com/p/gocc/lexer/items"
	"code.google.com/p/gocc/token"
	"path"
)

func Gen(cfg config.Config, lexPart *ast.LexPart, tokenMap *token.TokenMap) {
	lexSets := items.GetItemSets(lexPart)
	if cfg.Verbose() {
		io.WriteFileString(path.Join(cfg.OutDir(), "lexer_sets.txt"), lexSets.String())
	}
	golang.Gen(cfg.Package(), cfg.OutDir(), lexPart.Header.String(), lexSets, tokenMap, cfg)

}
