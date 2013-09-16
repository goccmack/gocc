<<<<<<< HEAD
package token

import (
=======

package token

import(
>>>>>>> 9d3e28b6b4a375015991b8ecb296891776774570
	"fmt"
)

type Token struct {
	Type
	Lit []byte
	Pos
}

type Type int

<<<<<<< HEAD
const (
=======
const(
>>>>>>> 9d3e28b6b4a375015991b8ecb296891776774570
	INVALID Type = iota
	EOF
)

type Pos struct {
	Offset int
<<<<<<< HEAD
	Line   int
=======
	Line int
>>>>>>> 9d3e28b6b4a375015991b8ecb296891776774570
	Column int
}

func (this Pos) String() string {
	return fmt.Sprintf("Pos(offset=%d, line=%d, column=%d)", this.Offset, this.Line, this.Column)
}

type TokenMap struct {
<<<<<<< HEAD
	typeMap []string
	idMap   map[string]Type
=======
	typeMap  []string
	idMap map[string]Type
>>>>>>> 9d3e28b6b4a375015991b8ecb296891776774570
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
		"c",
		"empty",
		"b",
	},

<<<<<<< HEAD
	idMap: map[string]Type{
		"INVALID": 0,
		"$":       1,
		"c":       2,
		"empty":   3,
		"b":       4,
	},
}
=======
	idMap: map[string]Type {
		"INVALID": 0,
		"$": 1,
		"c": 2,
		"empty": 3,
		"b": 4,
	},
}

>>>>>>> 9d3e28b6b4a375015991b8ecb296891776774570
