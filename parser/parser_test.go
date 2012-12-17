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

package parser

import (
	"code.google.com/p/gocc/scanner"
	"code.google.com/p/gocc/token"
	"fmt"
	"strconv"
	"testing"
)

const (
	dragon_book_grammar_4_55_bnf = `S : C C ;
C : c C ;
C : d ;`
)

func Test1(t *testing.T) {
	lex := &scanner.Scanner{}
	src := []byte(dragon_book_grammar_4_55_bnf)
	tokenmap := token.NewMapFromStrings(token.GoccStrings)
	lex.Init([]byte(src), tokenmap)
	// tok := lex.Scan()
	// fmt.Println(tok)
	for tok, pos := lex.Scan(); tok.Type != tokenmap.Type("$"); tok, pos = lex.Scan() {
		msg := tokenmap.TokenString(tok.Type) + "(" + strconv.Itoa(int(tok.Type)) + ")"
		msg += ">" + string(tok.Lit) + "<"
		msg += " @ " + pos.String()
		fmt.Println(msg)
	}
}
