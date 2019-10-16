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

package util

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func RuneToString(r rune) string {
	if r >= 0x20 && r < 0x7f {
		return fmt.Sprintf("'%c'", r)
	}
	switch r {
	case 0x07:
		return "'\\a'"
	case 0x08:
		return "'\\b'"
	case 0x0C:
		return "'\\f'"
	case 0x0A:
		return "'\\n'"
	case 0x0D:
		return "'\\r'"
	case 0x09:
		return "'\\t'"
	case 0x0b:
		return "'\\v'"
	case 0x5c:
		return "'\\\\\\'"
	case 0x27:
		return "'\\''"
	case 0x22:
		return "'\\\"'"
	}
	if r < 0x10000 {
		return fmt.Sprintf("\\u%04x", r)
	}
	return fmt.Sprintf("\\U%08x", r)
}

type EscapedString string

func (str EscapedString) Unquote() (unquoted string, wasQuoted bool, quoteType rune) {
	if len(str) > 1 {
		r := str[0]
		if (r == '"' || r == '\x60' || r == '\'') && r == str[len(str)-1] {
			str = str[1 : len(str)-1]
		}
		return string(str), true, rune(r)
	}
	return string(str), false, 0
}

func (str EscapedString) Unescape() string {
	var (
		res  string
		r    rune
		size int
	)
	for s := 0; s < len(str); s += size {
		if str[s] == '\\' {
			r, size = str[s+1:].EscapedFirstCharValue()
			size++
		} else {
			r, size = utf8.DecodeRuneInString(string(str)[s:])
		}
		res += string(r)
	}
	return res
}

func (str EscapedString) EscapedFirstCharValue() (rune, int) {
	var i, base, max uint32
	offset := 0
	switch str[offset] {
	case '\'', '"':
		return rune(str[offset]), 1
	case 'a':
		return '\a', 1
	case 'b':
		return '\b', 1
	case 'f':
		return '\f', 1
	case 'n':
		return '\n', 1
	case 'r':
		return '\r', 1
	case 't':
		return '\t', 1
	case 'v':
		return '\v', 1
	case '\\':
		return '\\', 1
	case '0', '1', '2', '3', '4', '5', '6', '7':
		i, base, max = 3, 8, 255
	case 'x':
		i, base, max = 2, 16, 255
		offset++
	case 'u':
		i, base, max = 4, 16, unicode.MaxRune
		offset++
	case 'U':
		i, base, max = 8, 32, unicode.MaxRune
		offset++
	default:
		panic(fmt.Sprintf("Error decoding character literal: %s\n", str[offset:]))
	}

	var x uint32
	for ; i > 0 && offset < len(str)-1; i-- {
		ch, size := utf8.DecodeRuneInString(string(str)[offset:])
		offset += size
		d := uint32(HexDigitValue(ch))
		if d >= base {
			panic(fmt.Sprintf("charVal(%s): illegal character (%c) in escape sequence. size=%d, offset=%d", str, str[offset], size, offset))
		}
		x = x*base + d
	}
	if x > max || 0xD800 <= x && x < 0xE000 {
		panic(fmt.Sprintf("Error decoding escape char value. Lit:%s, offset:%d, escape sequence is invalid Unicode code point\n", str, offset))
	}

	return rune(x), offset
}

func HexDigitValue(ch rune) int {
	switch {
	case '0' <= ch && ch <= '9':
		return int(ch) - '0'
	case 'a' <= ch && ch <= 'f':
		return int(ch) - 'a' + 10
	case 'A' <= ch && ch <= 'F':
		return int(ch) - 'A' + 10
	}
	return 16 // larger than any legal digit val
}
