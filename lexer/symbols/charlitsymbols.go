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

package symbols

import (
	"code.google.com/p/gocc/ast"
)

// key: string of symbols - string(ast.CharLit.Lit). E.g.: "'a'"
type CharLitSymbols map[string]*ast.LexCharLit

func NewCharLitSymbols() CharLitSymbols {
	return make(CharLitSymbols)
}

func (this CharLitSymbols) Add(cl *ast.LexCharLit) {
	this[cl.String()] = cl
}

func (this CharLitSymbols) Len() int {
	return len(this)
}

func (this CharLitSymbols) List() []*ast.LexCharLit {
	list := make([]*ast.LexCharLit, 0, len(this))
	for _, sym := range this {
		list = append(list, sym)
	}
	return list
}

func (this CharLitSymbols) StringList() []string {
	symbols := make([]string, 0, len(this))
	for key, _ := range this {
		symbols = append(symbols, key)
	}
	return symbols
}
