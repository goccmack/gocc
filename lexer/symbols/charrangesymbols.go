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

// key: string of range, e.g.: 'a'-'z'
type CharRangeSymbols map[string]*ast.LexCharRange

func NewCharRangeSymbols() CharRangeSymbols {
	return make(CharRangeSymbols)
}

func (this CharRangeSymbols) Add(cr *ast.LexCharRange) {
	this[cr.String()] = cr
}

func (this CharRangeSymbols) Len() int {
	return len(this)
}

func (this CharRangeSymbols) List() []*ast.LexCharRange {
	list := make([]*ast.LexCharRange, 0, len(this))
	for _, sym := range this {
		list = append(list, sym)
	}
	return list
}

func (this CharRangeSymbols) StringList() []string {
	symbols := make([]string, 0, len(this))
	for key := range this {
		symbols = append(symbols, key)
	}
	return symbols
}
