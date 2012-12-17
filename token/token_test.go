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

package token

import (
	// "fmt"
	"testing"
)

// func Test1(t *testing.T) {
// 	testEqual(
// 		"<< $1.Type1.Append($2)   >>",
// 		"X[1].Type1.Append(X[2])",
// 		t)
// }
//
// func testEqual(src, exp string, t *testing.T) {
// 	tok := &Token{Lit:[]byte(src)}
// 	out, err := tok.SDTVal()
// 	if err != nil {
// 		t.Fail()
// 		fmt.Println(err)
// 	} else if out != exp {
// 		t.Fail()
// 		fmt.Println("Expected: >" + exp + "< but got >" + out + "<" )
// 	}
// }

func TestTokenEquals1(t *testing.T) {
	t1, t2 := (*Token)(nil), (*Token)(nil)
	if !t1.Equals(t2) {
		t.Fail()
	}
}

func TestTokenEquals2(t *testing.T) {
	t1, t2 := (*Token)(nil), NewToken(EOF, []byte("$"))
	if t1.Equals(t2) {
		t.Fail()
	}
	if t2.Equals(t1) {
		t.Fail()
	}
}

func TestTokenEquals3(t *testing.T) {
	t1, t2 := NewToken(Type(1), []byte("id0")), NewToken(Type(1), []byte("id0"))
	if !t1.Equals(t2) {
		t.Fail()
	}
}

func TestTokenEquals4(t *testing.T) {
	t1, t2 := NewToken(Type(1), []byte("id0")), NewToken(Type(2), []byte("id0"))
	if t1.Equals(t2) {
		t.Fail()
	}
	if t2.Equals(t1) {
		t.Fail()
	}
	if !t1.Equals(t1) {
		t.Fail()
	}
}

func TestTokenEquals5(t *testing.T) {
	t1, t2 := NewToken(Type(1), []byte("id0")), NewToken(Type(1), []byte("id1"))
	if t1.Equals(t2) {
		t.Fail()
	}
	if t2.Equals(t1) {
		t.Fail()
	}
	if !t1.Equals(t1) {
		t.Fail()
	}
}
