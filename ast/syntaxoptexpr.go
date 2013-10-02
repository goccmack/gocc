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

type SyntaxOptionalExpression SyntaxExpression

func NewSyntaxOptionalExpression(expr interface{}) (SyntaxOptionalExpression, error) {
	return SyntaxOptionalExpression(expr.(SyntaxExpression)), nil
}

func (this SyntaxOptionalExpression) ExpressionIsBasic() bool {
	return SyntaxExpression(this).Basic()
}

func (this SyntaxOptionalExpression) String() string {
	w := new(bytes.Buffer)
	fmt.Fprintf(w, "[")
	for i, body := range this {
		if i < len(this)-1 {
			fmt.Fprintf(w, "%s | ", body)
		} else {
			fmt.Fprintf(w, "%s", body)
		}
	}
	fmt.Fprintf(w, "]")
	return w.String()
}
