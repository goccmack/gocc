package ast

import (
	"bytes"
	"fmt"
)

type SyntaxExpression []*SyntaxBody

func NewSyntaxExpression(body interface{}) (SyntaxExpression, error) {
	return SyntaxExpression{(body.(*SyntaxBody))}, nil
}

func AddSyntaxExprBody(expr, body interface{}) (SyntaxExpression, error) {
	return append(expr.(SyntaxExpression), body.(*SyntaxBody)), nil
}

func (this SyntaxExpression) Basic() bool {
	if len(this) > 1 {
		return false
	}
	return this[0].Basic()
}

func (this SyntaxExpression) String() string {
	w := new(bytes.Buffer)
	for i, body := range this {
		if i < len(this)-1 {
			fmt.Fprintf(w, "%s | ", body)
		} else {
			fmt.Fprintf(w, "%s", body)
		}
	}
	return w.String()
}
