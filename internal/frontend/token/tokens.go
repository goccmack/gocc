package token

import (
	"github.com/maxcalandrelli/gocc/internal/config"
)

var FRONTENDTokens = NewMapFromStrings([]string{
	"id",
	"tokId",
	":",
	";",
	"regDefId",
	"ignoredTokId",
	"|",
	".",
	"char_lit",
	"-",
	"[",
	"]",
	"{",
	"}",
	"(",
	")",
	"prodId",
	"g_sdt_lit",
	config.SYMBOL_ERROR,
	config.SYMBOL_EMPTY,
	"string_lit",
})
