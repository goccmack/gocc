//Copyright 2013 Vastech SA (PTY) LTD
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

package mail

import (
	"code.google.com/p/gocc/example/mail/lexer"
	"code.google.com/p/gocc/example/mail/token"
	"testing"
)

var testData1 = map[string]bool{
	"mymail@google.com":          true,
	"@google.com":                false,
	`"quoted string"@mymail.com`: true,
	`"unclosed quote@mymail.com`: false,
}

func Test1(t *testing.T) {
	for input, ok := range testData1 {
		l := lexer.NewLexer([]byte(input))
		tok := l.Scan()
		switch {
		case tok.Type == token.INVALID:
			if ok {
				t.Errorf("%s", input)
			}
		case tok.Type == token.TokMap.Type("addrspec"):
			if !ok {
				t.Errorf("%s", input)
			}
		default:
			t.Fatalf("This must not happen")
		}
	}
}

var checkData2 = []string{
	"addr1@gmail.com",
	"addr2@gmail.com",
	"addr3@gmail.com",
}

var testData2 = `
	addr1@gmail.com
	addr2@gmail.com
	addr3@gmail.com
`

func Test2(t *testing.T) {
	l := lexer.NewLexer([]byte(testData2))
	num := 0
	for tok := l.Scan(); tok.Type == token.TokMap.Type("addrspec"); tok = l.Scan() {
		if string(tok.Lit) != checkData2[num] {
			t.Errorf("%s != %s", string(tok.Lit), checkData2[num])
		}
		num++
	}
	if num != len(checkData2) {
		t.Fatalf("%d addresses parsed", num)
	}
}
