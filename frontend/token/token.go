package token

import (
	"fmt"
)

type Token struct {
	Type
	Lit []byte
	Pos
}

type Type int

const (
	INVALID Type = iota
	EOF
)

type Pos struct {
	Offset int
	Line   int
	Column int
}

func (this Pos) String() string {
	return fmt.Sprintf("Pos(offset=%d, line=%d, column=%d)", this.Offset, this.Line, this.Column)
}

type TokenMap struct {
	typeMap []string
	idMap   map[string]Type
}

func (this TokenMap) Id(tok Type) string {
	if int(tok) < len(this.typeMap) {
		return this.typeMap[tok]
	}
	return "unknown"
}

func (this TokenMap) Type(tok string) Type {
	if typ, exist := this.idMap[tok]; exist {
		return typ
	}
	return INVALID
}

func (this TokenMap) TokenString(tok *Token) string {
	//TODO: refactor to print pos & token string properly
	return fmt.Sprintf("%s(%d,%s)", this.Id(tok.Type), tok.Type, tok.Lit)
}

func (this TokenMap) StringType(typ Type) string {
	return fmt.Sprintf("%s(%d)", this.Id(typ), typ)
}

var TokMap = TokenMap{
	typeMap: []string{
		"INVALID",
		"$",
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
		"actionLit",
		"prodId",
		"error",
		"stringLit",
	},

	idMap: map[string]Type{
		"INVALID":      0,
		"$":            1,
		"tokId":        2,
		":":            3,
		";":            4,
		"regDefId":     5,
		"ignoredTokId": 6,
		"|":            7,
		".":            8,
		"char_lit":     9,
		"-":            10,
		"[":            11,
		"]":            12,
		"{":            13,
		"}":            14,
		"(":            15,
		")":            16,
		"actionLit":    17,
		"prodId":       18,
		"error":        19,
		"stringLit":    20,
	},
}
