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
	"bytes"
)

type SyntaxTerms []SyntaxTerm

func NewSyntaxTerms(term interface{}) (SyntaxTerms, error) {
	return SyntaxTerms{term.(SyntaxTerm)}, nil
}

func AddSyntaxTerm(terms, term interface{}) (SyntaxTerms, error) {
	return append(terms.(SyntaxTerms), term.(SyntaxTerm)), nil
}

func (this SyntaxTerms) CloneDeleteTerm(tidx int) SyntaxTerms {
	clone := make(SyntaxTerms, len(this)-1)
	copy(clone, this[0:tidx])
	copy(clone[tidx:], this[tidx+1:])
	return clone
}

func (this SyntaxTerms) CloneInsertTerm(tidx int, term SyntaxTerm) SyntaxTerms {
	clone := make(SyntaxTerms, len(this)+1)
	if tidx > 0 {
		copy(clone, this[0:tidx])
		copy(clone[tidx+1:], this[tidx:])
	} else {
		copy(clone[1:], this)
	}
	clone[tidx] = term
	return clone
}

func (this SyntaxTerms) CloneReplace(tidx int, terms SyntaxTerms) SyntaxTerms {
	clone := make(SyntaxTerms, len(this)-1+len(terms))
	copy(clone, this[0:tidx])
	copy(clone[tidx:], terms)
	copy(clone[tidx+len(terms):], this[tidx+1:])
	return clone
}

func (this SyntaxTerms) CloneReplaceTerm(tidx int, term SyntaxTerm) SyntaxTerms {
	clone := make(SyntaxTerms, len(this))
	copy(clone, this)
	clone[tidx] = term
	return clone
}

func (this SyntaxTerms) Equal(that SyntaxTerms) bool {
	if len(this) != len(that) {
		return false
	}
	for i, term := range this {
		if !term.Equal(that[i]) {
			return false
		}
	}
	return true
}

func (this SyntaxTerms) String() string {
	w := new(bytes.Buffer)
	for i, term := range this {
		if i > 0 {
			w.WriteString(" ")
		}
		w.WriteString(term.String())
	}
	return w.String()
}
