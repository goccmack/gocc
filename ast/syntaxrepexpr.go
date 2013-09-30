package ast

import (
	"bytes"
	"fmt"
)

type SyntaxRepeatedExpression SyntaxExpression

func NewSyntaxRepeatedExpression(expr interface{}) (SyntaxRepeatedExpression, error) {
	return SyntaxRepeatedExpression(expr.(SyntaxExpression)), nil
}

func (this SyntaxRepeatedExpression) ExpressionIsBasic() bool {
	return SyntaxExpression(this).Basic()
}

func (this SyntaxRepeatedExpression) String() string {
	w := new(bytes.Buffer)
	fmt.Fprintf(w, "{")
	for i, body := range this {
		if i < len(this)-1 {
			fmt.Fprintf(w, "%s | ", body)
		} else {
			fmt.Fprintf(w, "%s", body)
		}
	}
	fmt.Fprintf(w, "}")
	return w.String()
}
