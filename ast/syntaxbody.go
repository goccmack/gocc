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
	"code.google.com/p/gocc/frontend/token"
	"fmt"
)

type SyntaxBody struct {
	Error  bool
	Terms  SyntaxTerms
	Action string
}

func NewSyntaxBody(terms, sdtLit interface{}) (*SyntaxBody, error) {
	syntaxBody := &SyntaxBody{
		Error: false,
	}
	if terms != nil {
		syntaxBody.Terms = terms.(SyntaxTerms)
	}
	if sdtLit != nil {
		syntaxBody.Action = ActionExpressionVal(sdtLit.(*token.Token).Lit)
	}
	return syntaxBody, nil
}

func NewErrorBody(terms, sdtLit interface{}) (*SyntaxBody, error) {
	body, _ := NewSyntaxBody(terms, sdtLit)
	body.Error = true
	if terms != nil {
		body.Terms = append(SyntaxTerms{errorConst}, terms.(SyntaxTerms)...)
	} else {
		body.Terms = SyntaxTerms{errorConst}
	}
	return body, nil
}

func (this SyntaxBody) Equal(that *SyntaxBody) bool {
	if that == nil {
		return false
	}
	return this.Error == that.Error &&
		this.Action == that.Action &&
		this.Terms.Equal(that.Terms)
}

func NewEmptyBody() (*SyntaxBody, error) {
	return NewSyntaxBody(nil, nil)
}

func (this *SyntaxBody) Basic() bool {
	if this.Action != "" {
		return false
	}
	for _, term := range this.Terms {
		if !term.Basic() {
			return false
		}
	}
	return true
}

func (this *SyntaxBody) Empty() bool {
	return len(this.Terms) == 0
}

func (this *SyntaxBody) String() string {
	return fmt.Sprintf("%s\t<< %s >>", this.Terms.String(), this.Action)
}
