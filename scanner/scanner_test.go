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

package scanner

import (
	"code.google.com/p/gocc/token"
	"fmt"
	"strconv"
	"testing"
)

func Test1(t *testing.T) {
	src := `id0 'c' "string"  : ; { } ... <<xyz>> | `
	strs := []string{"id", "char", "string", ":", ";", "{", "}", "...", "sdt_lit", "|"}

	tokenmap := token.NewMapFromStrings(token.GoccStrings)
	lex := new(Scanner)
	lex.Init([]byte(src), tokenmap)

	tok, _ := lex.Scan()
	for i := 0; tok.Type != token.EOF; i = i + 1 {
		if tok.Type != tokenmap.Type(strs[i]) {
			t.Fail()
			fmt.Println(i, ": expected", strs[i], "but scanned ", tokenmap.TokenString(tok.Type))
		}
		tok, _ = lex.Scan()
	}
}

func Test2(t *testing.T) {
	src := `<< $1 >>`

	tokenmap := token.NewMapFromStrings(token.GoccStrings)
	lex := new(Scanner)
	lex.Init([]byte(src), tokenmap)

	tok, _ := lex.Scan()
	testLit(tok.Lit, src, t)
	// testSDTVal(tok, "X[1]", t)
}

func Test3(t *testing.T) {
	src := `<< NewType($123.attrib) >>`
	// exp := `NewType(X[123].attrib)`

	tokenmap := token.NewMapFromStrings(token.GoccStrings)
	lex := new(Scanner)
	lex.Init([]byte(src), tokenmap)

	tok, _ := lex.Scan()
	testLit(tok.Lit, src, t)
	// testSDTVal(tok, exp, t)
}

func Test4(t *testing.T) {
	src := `<< $0.(Type1).Append($1.(Type2), $2.(Type3)) >>`
	// exp := `X[0].(Type1).Append(X[1].(Type2), X[2].(Type3))`

	tokenmap := token.NewMapFromStrings(token.GoccStrings)
	lex := new(Scanner)
	lex.Init([]byte(src), tokenmap)

	tok, _ := lex.Scan()
	testLit(tok.Lit, src, t)
	// testSDTVal(tok, exp, t)
}

func testLit(src []byte, exp string, t *testing.T) {
	testString(string(src), exp, t)
}

func testString(src, exp string, t *testing.T) {
	if string(src) != exp {
		t.Fail()
		fmt.Println("Expected >" + exp + "< but got >" + string(src) + "<")
	}
}

// func testSDTVal(tok *token.Token, exp string, t *testing.T) {
// 	sdt, err := tok.SDTVal()
// 	if err != nil {
// 		t.Fail()
// 		fmt.Println(err)
// 	}
// 	testString(sdt, exp, t)
// }

func Test5(t *testing.T) {
	src := []byte(". 123 / 456 .")
	tm := token.NewMapFromStrings([]string{"int_lit", ".", "/"})
	lex := new(Scanner)
	lex.Init(src, tm)
	for tok, pos := lex.Scan(); tok.Type != tm.Type("$"); tok, pos = lex.Scan() {
		fmt.Println(strconv.Itoa(int(tok.Type)), tm.TokenString(tok.Type), string(tok.Lit), pos)
	}
}
