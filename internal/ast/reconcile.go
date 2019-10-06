package ast

import (
	"github.com/maxcalandrelli/gocc/internal/frontend/stock/token"
)

var (
	StringGetter func(interface{}) string
)

func getString(v interface{}) string {
	if StringGetter == nil {
		return string(v.(*token.Token).Lit)
	}
	return StringGetter(v)
}
