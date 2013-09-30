package ast

import (
	"bytes"
	"fmt"
)

type SyntaxOptionalExpression SyntaxExpression

func NewSyntaxOptionalExpression(expr interface{}) (SyntaxOptionalExpression, error) {
	return SyntaxOptionalExpression(expr.(SyntaxExpression)), nil
}

func (this SyntaxOptionalExpression) ExpressionIsBasic() bool {
	return SyntaxExpression(this).Basic()
}

func (this SyntaxOptionalExpression) String() string {
	w := new(bytes.Buffer)
	fmt.Fprintf(w, "[")
	for i, body := range this {
		if i < len(this)-1 {
			fmt.Fprintf(w, "%s | ", body)
		} else {
			fmt.Fprintf(w, "%s", body)
		}
	}
	fmt.Fprintf(w, "]")
	return w.String()
}
