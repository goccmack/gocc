package ast

import (
	"github.com/maxcalandrelli/gocc/internal/frontend/stock/token"
)

var (
	StringGetter func(interface{}) string
)

func getString(v interface{}) string {
	if StringGetter == nil {
		if str, ok := v.(string); ok {
			return str
		}
		if tok, _ := v.(*token.Token); tok != nil {
			return string(tok.Lit)
		}
	}
	return StringGetter(v)
}

func unquoteString(str string) (string, bool, rune) {
	if len(str) > 1 {
		r := str[0]
		if r == '"' || r == '`' || r == '\'' {
			str = str[1 : len(str)-1]
		}
		return str, true, rune(r)
	}
	return str, false, 0
}
