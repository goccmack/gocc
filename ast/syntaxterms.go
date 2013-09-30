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
