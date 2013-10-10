package comment

import (
	"code.google.com/p/gocc/test/lexer/comment/lexer"
	"code.google.com/p/gocc/test/lexer/comment/token"
	"testing"
)

func Test1(t *testing.T) {
	lex := lexer.NewLexerString("token1 \n ; comment \n token2 token3")
	i, tok := 0, lex.Scan()
	for ; tok.Type != token.EOF; i, tok = i+1, lex.Scan() {
		if tok.Type != token.TokMap.Type("token") {
			t.Errorf("Expected token type %d, got %d, lit=%s", token.TokMap.Type("token"), tok.Type, tok.Lit)
		}
	}
	if i != 3 {
		t.Errorf("Expected 3 tokens, got %d", i)
	}
}
