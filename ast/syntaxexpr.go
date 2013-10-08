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
	"bytes"
	"fmt"
)

type SyntaxExpression []*SyntaxBody

func NewSyntaxExpression(body interface{}) (SyntaxExpression, error) {
	return SyntaxExpression{(body.(*SyntaxBody))}, nil
}

func AddSyntaxExprBody(expr, body interface{}) (SyntaxExpression, error) {
	return append(expr.(SyntaxExpression), body.(*SyntaxBody)), nil
}

func (this SyntaxExpression) Basic() bool {
	if len(this) > 1 {
		return false
	}
	return this[0].Basic()
}

func (this SyntaxExpression) Equal(that SyntaxExpression) bool {
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if !this[i].Equal(that[i]) {
			return false
		}
	}
	return true
}

func (this SyntaxExpression) String() string {
	w := new(bytes.Buffer)
	for i, body := range this {
		if i < len(this)-1 {
			fmt.Fprintf(w, "%s | ", body)
		} else {
			fmt.Fprintf(w, "%s", body)
		}
	}
	return w.String()
}
