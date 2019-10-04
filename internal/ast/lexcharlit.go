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
	"github.com/maxcalandrelli/gocc/internal/util"
)

type LexCharLit struct {
	Val    rune
	Lit    []byte
	s      string
	Negate bool
}

func NewLexCharLit(tok interface{}) (*LexCharLit, error) {
	return newLexCharLit(tok, false), nil
}

func NewLexCharLitExt(tok interface{}, negate bool) (*LexCharLit, error) {
	return newLexCharLit(tok, negate), nil
}

func newLexCharLit(tok interface{}, negate bool) *LexCharLit {
	c := new(LexCharLit)
	lit := []byte(getString(tok))

	c.Val = util.LitToRune(lit)
	c.Lit = lit
	c.s = util.RuneToString(c.Val)
	c.Negate = negate

	return c
}

func newLexCharLitFromRune(c rune) *LexCharLit {
	cl := &LexCharLit{
		Val: c,
		s:   util.RuneToString(c),
	}
	cl.Lit = []byte(cl.s)
	return cl
}

func (this *LexCharLit) IsTerminal() bool {
	return true
}
func (this *LexCharLit) String() string {
	return this.s
}
