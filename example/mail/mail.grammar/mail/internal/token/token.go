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
		"addrspec",
	},

	idMap: map[string]Type{
		"INVALID":  0,
		"Ω":        1,
		"addrspec": 2,
	},

	litMap: map[string]Type{},
}
