//Copyright 2012 Vastech SA (PTY) LTD
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

//The token package contains functions for tokens.
package token

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//Represents a Gocc BNF Token.
type Token struct {
	Type Type
	Lit  []byte
}

//Creates a New Gocc BNF Token.
func NewToken(typ Type, lit []byte) *Token {
	return &Token{typ, lit}
}

//Returns whether one token is equal to another.
func (this *Token) Equals(that *Token) bool {
	if this == nil || that == nil {
		return this == that
	}

	if this.Type != that.Type {
		return false
	}

	return bytes.Equal(this.Lit, that.Lit)
}

//Returns a string describing the token.
func (this *Token) String() string {
	str := ""
	if this.Type == EOF {
		str += `"$"`
	} else {
		str += `"` + string(this.Lit) + `"`
	}
	str += "(" + strconv.Itoa(int(this.Type)) + ")"
	return str
}

//Represents the Token Type
type Type int

const (
	ILLEGAL Type = iota - 1
	EOF
)

// Position describes an arbitrary source position
// including the file, line, and column location.
// A Position is valid if the line number is > 0.
//
type Position struct {
	Offset int // offset, starting at 0
	Line   int // line number, starting at 1
	Column int // column number, starting at 1 (character count)
}

// IsValid returns true if the position is valid.
func (pos *Position) IsValid() bool { return pos.Line > 0 }

// String returns a string in one of several forms:
//
//	file:line:column    valid position with file name
//	line:column         valid position without file name
//	file                invalid position with file name
//	-                   invalid position without file name
//
func (pos Position) String() string {
	s := ""
	if pos.IsValid() {
		s += fmt.Sprintf("%d:%d", pos.Line, pos.Column)
	}
	if s == "" {
		s = "-"
	}
	return s
}

//The int value of the token.
func (T *Token) IntValue() (int64, error) {
	return strconv.ParseInt(string(T.Lit), 10, 64)
}

//The uint value of the token.
func (T *Token) UintValue() (uint64, error) {
	return strconv.ParseUint(string(T.Lit), 10, 64)
}

//The SDT value of the token.
func (T *Token) SDTVal() string {
	sdt := string(T.Lit)
	rex, err := regexp.Compile(`\$[0-9]+`)
	if err != nil {
		panic(err)
	}
	idx := rex.FindAllStringIndex(sdt, -1)
	res := ""
	if len(idx) <= 0 {
		res = sdt
	} else {
		for i, loc := range idx {
			if loc[0] > 0 {
				if i > 0 {
					res += sdt[idx[i-1][1]:loc[0]]
				} else {
					res += sdt[0:loc[0]]
				}
			}
			res += "X["
			res += sdt[loc[0]+1 : loc[1]]
			res += "]"
		}
		if idx[len(idx)-1][1] < len(sdt) {
			res += sdt[idx[len(idx)-1][1]:]
		}
	}
	return strings.TrimSpace(res[2 : len(res)-2])
}
