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

	"github.com/maxcalandrelli/gocc/internal/config"
)

type SyntaxContextDependentTokId struct {
	string
	StdSyntaxSymbol
}

func NewContextDependentTokId(tokId interface{}) (SyntaxContextDependentTokId, error) {
	return SyntaxContextDependentTokId{getString(tokId), StdSyntaxSymbol{}}, nil
}

func NewNewContextDependentTokIdFromString(str string) SyntaxContextDependentTokId {
	return SyntaxContextDependentTokId{str, StdSyntaxSymbol{}}
}

func (this SyntaxContextDependentTokId) SymbolString() string {
	return this.string
}

func (this SyntaxContextDependentTokId) String() string {
	return fmt.Sprintf("%s<%s>", config.INTERNAL_SYMBOL_CDTOK, this.string)
}

func (this SyntaxContextDependentTokId) SymbolName() string {
	return fmt.Sprintf("%s<%s>", config.INTERNAL_SYMBOL_CDTOK, this.string)
}

func (this SyntaxContextDependentTokId) IsTerminal() bool {
	return true
}
