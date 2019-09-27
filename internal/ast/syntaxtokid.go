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
	"github.com/maxcalandrelli/gocc/internal/frontend/token"
)

type SyntaxTokId string

func NewTokId(tokId interface{}) (SyntaxTokId, error) {
	return SyntaxTokId(string(tokId.(*token.Token).Lit)), nil
}

func (this SyntaxTokId) SymbolString() string {
	return string(this)
}

func (this SyntaxTokId) String() string {
	return string(this)
}
