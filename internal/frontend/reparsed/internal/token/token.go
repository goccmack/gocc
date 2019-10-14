// Code generated by gocc; DO NOT EDIT.

package token

import (
	"fmt"
)

type Token struct {
	Type
	Lit           []byte
	IgnoredPrefix []byte
	Pos
	ForeignAstNode interface{}
	Foreign        bool
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

func (p Pos) String() string {
	return fmt.Sprintf("Pos(offset=%d, line=%d, column=%d)", p.Offset, p.Line, p.Column)
}

func (p Pos) StartingFrom(base Pos) Pos {
	r := base
	r.Offset += p.Offset
	r.Line += p.Line
	r.Column = p.Column
	if p.Line > 0 && base.Line > 0 {
		r.Line--
	}
	if r.Column < 1 {
		r.Column = 1
	}
	return r
}

type TokenMap struct {
	typeMap []string
	idMap   map[string]Type
	litMap  map[string]Type
}

func (m TokenMap) Id(tok Type) string {
	if int(tok) < len(m.typeMap) {
		return m.typeMap[tok]
	}
	return "unknown"
}

func (m TokenMap) Type(tok string) Type {
	if typ, exist := m.idMap[tok]; exist {
		return typ
	}
	return INVALID
}

func (m TokenMap) TokenString(tok *Token) string {
	return fmt.Sprintf("%s(%d,<%s>)", m.Id(tok.Type), tok.Type, tok.Lit)
}

func (m TokenMap) StringType(typ Type) string {
	return fmt.Sprintf("%s(%d)", m.Id(typ), typ)
}

var TokMap = TokenMap{
	typeMap: []string{
		"INVALID",
		"Ω",
		"tokId",
		"Λ<:>",
		"Λ<;>",
		"regDefId",
		"ignoredTokId",
		"Λ<|>",
		"Λ<.>",
		"char_lit",
		"Λ<->",
		"Λ<~>",
		"Λ<(>",
		"Λ<)>",
		"Λ<[>",
		"Λ<]>",
		"Λ<{>",
		"Λ<}>",
		"g_sdt_lit",
		"prodId",
		"string_lit",
		"g_ctxdep_lit",
		"Λ<@>",
		"Λ<error>",
		"Λ<λ>",
		"Λ<empty>",
		"Λ<ε>",
	},

	idMap: map[string]Type{
		"INVALID":      0,
		"Ω":            1,
		"tokId":        2,
		"Λ<:>":         3,
		"Λ<;>":         4,
		"regDefId":     5,
		"ignoredTokId": 6,
		"Λ<|>":         7,
		"Λ<.>":         8,
		"char_lit":     9,
		"Λ<->":         10,
		"Λ<~>":         11,
		"Λ<(>":         12,
		"Λ<)>":         13,
		"Λ<[>":         14,
		"Λ<]>":         15,
		"Λ<{>":         16,
		"Λ<}>":         17,
		"g_sdt_lit":    18,
		"prodId":       19,
		"string_lit":   20,
		"g_ctxdep_lit": 21,
		"Λ<@>":         22,
		"Λ<error>":     23,
		"Λ<λ>":         24,
		"Λ<empty>":     25,
		"Λ<ε>":         26,
	},

	litMap: map[string]Type{
		".":     8,
		"~":     11,
		"(":     12,
		"empty": 25,
		"|":     7,
		";":     4,
		"{":     16,
		":":     3,
		"-":     10,
		")":     13,
		"[":     14,
		"]":     15,
		"}":     17,
		"@":     22,
		"error": 23,
		"λ":     24,
		"ε":     26,
	},
}
