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
)

type LexCharRange struct {
	From   *LexCharLit
	To     *LexCharLit
	s      string
	Negate bool
}

func NewLexCharRange(from, to interface{}) (*LexCharRange, error) {
	return NewLexCharRangeExt(from, to, false)
}

func NewLexCharRangeExt(from, to interface{}, negate bool) (*LexCharRange, error) {
	cr := &LexCharRange{
		From:   newLexCharLit(from, negate),
		To:     newLexCharLit(to, negate),
		Negate: negate,
	}
	if cr.From.Val > cr.To.Val {
		t := cr.From
		cr.From = cr.To
		cr.To = t
	}
	return cr, nil
}

func (this *LexCharRange) IsTerminal() bool {
	return true
}

func (this *LexCharRange) String() string {
	if this.s == "" {
		this.s = fmt.Sprintf("%s-%s", this.From.String(), this.To.String())
	}
	return this.s
}
