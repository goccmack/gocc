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
	"github.com/maxcalandrelli/gocc/internal/config"
)

/*
All syntax symbols are types of string
*/
type SyntaxSymbol interface {
	SymbolName() string
	SymbolString() string
	String() string
	IsError() bool
	IsEpsilon() bool
	IsTerminal() bool
	IsNonTerminal() bool
	gSymbol()
}

type StdSyntaxSymbol struct{}
type InvalidSyntaxSymbol struct{ StdSyntaxSymbol }

func (StdSyntaxSymbol) IsError() bool       { return false }
func (StdSyntaxSymbol) IsEpsilon() bool     { return false }
func (StdSyntaxSymbol) IsTerminal() bool    { return false }
func (StdSyntaxSymbol) IsNonTerminal() bool { return false }

func (SyntaxEmpty) gSymbol()                 {}
func (SyntaxEof) gSymbol()                   {}
func (SyntaxError) gSymbol()                 {}
func (SyntaxProdId) gSymbol()                {}
func (SyntaxTokId) gSymbol()                 {}
func (SyntaxStringLit) gSymbol()             {}
func (InvalidSyntaxSymbol) gSymbol()         {}
func (SyntaxContextDependentTokId) gSymbol() {}
func (SyntaxSubParser) gSymbol()             {}

func (InvalidSyntaxSymbol) SymbolName() string {
	return config.INTERNAL_SYMBOL_INVALID
}

func (InvalidSyntaxSymbol) SymbolString() string {
	return config.INTERNAL_SYMBOL_INVALID
}

func (InvalidSyntaxSymbol) String() string {
	return config.SYMBOL_INVALID
}

func (InvalidSyntaxSymbol) IsTerminal() bool {
	return true
}
