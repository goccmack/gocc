package scanner

import (
	"github.com/maxcalandrelli/gocc/example/nolexer/nolexer.grammar/nolexer/iface"
)

type Scanner struct {
	src []byte
	pos int
}

func NewString(s string) *Scanner {
	return &Scanner{[]byte(s), 0}
}

func isWhiteSpace(c byte) bool {
	return c == ' ' ||
		c == '\t' ||
		c == '\n' ||
		c == '\r'
}

func (S *Scanner) skipWhiteSpace() {
	for S.pos < len(S.src) && isWhiteSpace(S.src[S.pos]) {
		S.pos++
	}
}

func (S *Scanner) scanId() string {
	pos := S.pos
	for S.pos < len(S.src) && !isWhiteSpace(S.src[S.pos]) {
		S.pos++
	}
	S.pos++
	return string(S.src[pos : S.pos-1])
}

func (S *Scanner) Scan() (tok *iface.Token) {
	S.skipWhiteSpace()

	if S.pos >= len(S.src) {
		return &iface.Token{Type: iface.EOF}
	}

	pos := S.pos

	lit := S.scanId()
	switch lit {
	case "hiya":
		return &iface.Token{Type: iface.GetTokenMap().Type("hiya"),
			Lit: []byte("hiya"),
			Pos: iface.Pos{Offset: pos}}
	case "hello":
		return &iface.Token{Type: iface.GetTokenMap().Type("hello"),
			Lit: []byte("hello"),
			Pos: iface.Pos{Offset: pos}}
	default:
		return &iface.Token{Type: iface.GetTokenMap().Type("name"),
			Lit: []byte(lit),
			Pos: iface.Pos{Offset: pos}}
	}
}
