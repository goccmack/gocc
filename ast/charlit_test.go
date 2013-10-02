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

package ast

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	testCharRange(t, 'a', 'z')
	testCharRange(t, '0', '9')
}

func testCharRange(t *testing.T, from, to rune) {
	for ch := from; ch <= to; ch++ {
		src := []byte(fmt.Sprintf("'%c'", ch))
		if ch1 := charVal(src); ch1 != ch {
			t.Error(fmt.Sprintf("charVal(%s) returned %d, expected %d\n", src, ch1, ch))
		}
	}
}

type test2T struct {
	lit string
	ch  rune
}

var test2D = []test2T{
	{`'\a'`, '\a'},
	{`'\b'`, '\b'},
	{`'\f'`, '\f'},
	{`'\n'`, '\n'},
	{`'\r'`, '\r'},
	{`'\t'`, '\t'},
	{`'\v'`, '\v'},
	{`'\\'`, '\\'},
	{`'\''`, '\''},
}

func Test2(t *testing.T) {
	for _, src := range test2D {
		if ch1 := charVal([]byte(src.lit)); ch1 != src.ch {
			t.Error(fmt.Sprintf("charVal(%s) returned %d, expected %d\n", src.lit, src.ch, ch1))
		}
	}
}

func Test3(t *testing.T) {
	for i := 1; i < 0100; i++ {
		src := fmt.Sprintf("'\\0%o'", i)
		if ch := charVal([]byte(src)); ch != rune(i) {
			t.Error(fmt.Printf("charVal(%s) returned %o, expected %o\n", src, i, ch))
		}
	}
}

func Test4(t *testing.T) {
	for i := 0; i < 0x100; i++ {
		src := fmt.Sprintf("'\\x%x'", i)
		if ch := charVal([]byte(src)); ch != rune(i) {
			t.Error(fmt.Printf("charVal(%s) returned %x, expected %x\n", src, i, ch))
		}
	}
}

func Test5(t *testing.T) {
	for i := 1; i <= 0x2000; i++ {
		testUnicode(t, rune(i), 'u')
	}
}

func Test6(t *testing.T) {
	for i := 1; i <= 0xd7ff; i++ {
		testUnicode(t, rune(i), 'U')
	}
}

func testUnicode(t *testing.T, ch rune, u rune) {
	var src string
	switch ch {
	case 'u':
		src = fmt.Sprintf("'\\u%04x'", ch)
	default:
		src = fmt.Sprintf("'\\U%08x'", ch)

	}
	if ch1 := charVal([]byte(src)); ch != ch {
		t.Error(fmt.Printf("charVal(%s) returned %x, expected %x\n", src, ch1, ch))
	}
}
