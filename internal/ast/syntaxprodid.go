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

// Id or name of a grammar(syntax) production
type SyntaxProdId struct {
	string
	StdSyntaxSymbol
}

func NewSyntaxProdId(tok interface{}) (SyntaxProdId, error) {
	return SyntaxProdId{getString(tok), StdSyntaxSymbol{}}, nil
}

func NewSyntaxProdIdFromString(str string) SyntaxProdId {
	return SyntaxProdId{str, StdSyntaxSymbol{}}
}

func (this SyntaxProdId) SymbolString() string {
	return this.string
}

func (this SyntaxProdId) String() string {
	return this.string
}

func (this SyntaxProdId) SymbolName() string {
	//return this.SymbolString()
	return fmt.Sprintf("%s<%s>", config.INTERNAL_SYMBOL_PROD, this.string)
}

func (this SyntaxProdId) IsTerminal() bool {
	return false
}

func (this SyntaxProdId) IsNonTerminal() bool {
	return true
}
