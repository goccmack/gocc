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
	"path"

	"github.com/maxcalandrelli/gocc/internal/config"
	"github.com/maxcalandrelli/gocc/internal/util"
)

type SyntaxSubParser struct {
	string
	StdSyntaxSymbol
	Alias  string
	Import string
}

func NewAliasedSubParser(_alias, _import interface{}) (SyntaxSubParser, error) {
	return newSubParser(getString(_alias), getString(_import))
}

func NewSubParser(_import interface{}) (SyntaxSubParser, error) {
	return newSubParser("", getString(_import))
}

func newSubParser(alias, imp string) (SyntaxSubParser, error) {
	imp, _, _ = util.EscapedString(imp).Unquote()
	if alias == "" {
		alias = path.Base(imp)
	} else {
		alias, _, _ = util.EscapedString(alias).Unquote()
	}
	return SyntaxSubParser{
		fmt.Sprintf("%s_%d", alias, tokenIdCount),
		StdSyntaxSymbol{},
		alias,
		imp,
	}, nil
}

func (this SyntaxSubParser) SymbolString() string {
	return this.string
}

func (this SyntaxSubParser) String() string {
	return fmt.Sprintf("%s<%s>", config.INTERNAL_SYMBOL_CDTOK, this.string)
}

func (this SyntaxSubParser) SymbolName() string {
	return fmt.Sprintf("%s<%s>", config.INTERNAL_SYMBOL_CDTOK, this.string)
}

func (this SyntaxSubParser) IsTerminal() bool {
	return true
}
