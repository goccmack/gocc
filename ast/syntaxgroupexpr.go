package ast

import (
	"bytes"
	"fmt"
)

type SyntaxGroupExpression SyntaxExpression

func NewSyntaxGroupExpression(expr interface{}) (SyntaxGroupExpression, error) {
	return SyntaxGroupExpression(expr.(SyntaxExpression)), nil
}

func (this SyntaxGroupExpression) ExpressionIsBasic() bool {
	return SyntaxExpression(this).Basic()
}

func (this SyntaxGroupExpression) String() string {
	w := new(bytes.Buffer)
	fmt.Fprintf(w, "(")
	for i, body := range this {
		if i < len(this)-1 {
			fmt.Fprintf(w, "%s | ", body)
		} else {
			fmt.Fprintf(w, "%s", body)
		}
	}
	fmt.Fprintf(w, ")")
	return w.String()
}
