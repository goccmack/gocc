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

type SyntaxStringLit struct {
	string
	StdSyntaxSymbol
}

func NewStringLit(tok interface{}) (SyntaxStringLit, error) {
	lit := getString(tok)
	return SyntaxStringLit{lit[1 : len(lit)-1], StdSyntaxSymbol{}}, nil
}

func (this SyntaxStringLit) SymbolString() string {
	return this.string
}

func (this SyntaxStringLit) String() string {
	return fmt.Sprintf("\"%s\"", this.string)
}

func (this SyntaxStringLit) Bytes() []byte {
	return []byte(this.string)
}

func (this SyntaxStringLit) SymbolName() string {
	return config.INTERNAL_SYMBOL_LIT + this.string
}
