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
	"code.google.com/p/gocc/frontend/token"
)

// Id or name of a grammar(syntax) production
type SyntaxProdId string

func NewSyntaxProdId(tok interface{}) (SyntaxProdId, error) {
	return SyntaxProdId(string(tok.(*token.Token).Lit)), nil
}

func (this SyntaxProdId) Equal(that SyntaxTerm) bool {
	if thatPId, ok := that.(SyntaxProdId); ok {
		return this == thatPId
	} else {
		return false
	}
}

func (this SyntaxProdId) String() string {
	return string(this)
}

func (this SyntaxProdId) DotString() string {
	return string(this)
}
