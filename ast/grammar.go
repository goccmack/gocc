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

type Grammar struct {
	*LexPart
	*SyntaxPart
}

func NewGrammar(lexPart, syntaxPart interface{}) (*Grammar, error) {
	g := &Grammar{}
	if lexPart != nil {
		g.LexPart = lexPart.(*LexPart)
	} else {
		lp, _ := NewLexPart(nil, nil, nil)
		g.LexPart = lp
	}
	if syntaxPart != nil {
		g.SyntaxPart = syntaxPart.(*SyntaxPart).augment()
	}
	return g, nil
}
