package scanner

import (
	"code.google.com/p/gocc/frontend/token"
	"fmt"
	"testing"
)

type testRecord struct {
	src    string
	typ    token.Type
	tokLit string
}

var testData = []testRecord{
	testRecord{"tokId", token.FRONTENDTokens.Type("tokId"), "tokId"},
	testRecord{"!whitespace", token.FRONTENDTokens.Type("ignoredTokId"), "!whitespace"},
	testRecord{":", token.FRONTENDTokens.Type(":"), ":"},
	testRecord{";", token.FRONTENDTokens.Type(";"), ";"},
	testRecord{"_regDefId", token.FRONTENDTokens.Type("regDefId"), "_regDefId"},
	testRecord{"|", token.FRONTENDTokens.Type("|"), "|"},
	testRecord{`'\u0011'`, token.FRONTENDTokens.Type("char_lit"), `'\u0011'`},
	testRecord{"-", token.FRONTENDTokens.Type("-"), "-"},
	testRecord{"(", token.FRONTENDTokens.Type("("), "("},
	testRecord{")", token.FRONTENDTokens.Type(")"), ")"},
	testRecord{"[", token.FRONTENDTokens.Type("["), "["},
	testRecord{"]", token.FRONTENDTokens.Type("]"), "]"},
	testRecord{"{", token.FRONTENDTokens.Type("{"), "{"},
	testRecord{"}", token.FRONTENDTokens.Type("}"), "}"},
	testRecord{"<< sdt lit >>", token.FRONTENDTokens.Type("g_sdt_lit"), "<< sdt lit >>"},
	testRecord{"ProdId", token.FRONTENDTokens.Type("prodId"), "ProdId"},
	testRecord{`"string lit"`, token.FRONTENDTokens.Type("string_lit"), `"string lit"`},
}

func Test1(tst *testing.T) {
	s := &Scanner{}
	for _, t := range testData {
		s.Init([]byte(t.src), token.FRONTENDTokens)
		tok, _ := s.Scan()
		if tok.Type != t.typ {
			tst.Error(fmt.Sprintf("src: %s, type: %d -- got type: %d\n", t.src, t.typ, tok.Type))
		}
		if string(tok.Lit) != t.tokLit {
			tst.Error(fmt.Sprintf("src: %s, expected lit: %s, got: %s\n", t.src, t.tokLit, string(tok.Lit)))
		}
	}
}

func Test2(t *testing.T) {
	s := &Scanner{}
	lit := "The SDT Lit"
	s.Init([]byte(fmt.Sprintf("<< %s >>", lit)), token.FRONTENDTokens)
	tok, _ := s.Scan()
	if tok.Type != token.FRONTENDTokens.Type("g_sdt_lit") {
		t.Error(fmt.Sprintf("Expected tok type: g_sdt_lit, got: %s"), token.FRONTENDTokens.TokenString(tok.Type))
	}
	if tok.SDTVal() != lit {
		t.Error(fmt.Sprintf("Expected SDTVal: %s, got: %s\n", lit, tok.SDTVal()))
	}
}
