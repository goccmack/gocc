package ast

import (
	"fmt"
	"code.google.com/p/gocc/frontend/token"
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

func NewErrorBody(symbols, sdtLit interface{}) (*SyntaxBody, error) {
	body, _ := NewSyntaxBody(symbols, sdtLit)
	body.Error = true
	return body, nil
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
