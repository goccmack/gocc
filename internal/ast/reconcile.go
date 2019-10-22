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
