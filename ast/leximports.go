package ast

import (
	"bytes"
	"errors"
	"fmt"
)

type LexImports struct {
	Imports map[string]*LexImport
}

func NewLexImports(lexImport interface{}) (*LexImports, error) {
	imports, err := newLexImports().Add(lexImport.(*LexImport))
	return imports, err
}

func newLexImports() *LexImports {
	return &LexImports{
		Imports: make(map[string]*LexImport),
	}
}

func AddLexImport(imports, lexImport interface{}) (*LexImports, error) {
	return imports.(*LexImports).Add(lexImport.(*LexImport))
}

/*
Return true if a new lex import has been added.
Return false if lexImport is a duplicate.
*/
func (this *LexImports) Add(lexImport *LexImport) (*LexImports, error) {
	if _, exist := this.Imports[lexImport.Id]; exist {
		return nil, errors.New(fmt.Sprintf("Duplicate builtin declaration: %s", lexImport.String()))
	}
	this.Imports[lexImport.Id] = lexImport
	return this, nil
}

func (this *LexImports) String() string {
	w := new(bytes.Buffer)
	w.WriteString("import(\n")
	for _, imp := range this.Imports {
		fmt.Fprintf(w, "\t%s\n", imp.String())
	}
	w.WriteString(")")
	return w.String()
}
