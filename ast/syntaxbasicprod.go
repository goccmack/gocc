package ast

import (
	"bytes"
	"fmt"
)

type SyntaxBasicProd struct {
	Id     string
	Error  bool
	Terms  SyntaxTerms
	Action string
}

func (this *SyntaxBasicProd) String() string {
	w := new(bytes.Buffer)
	fmt.Fprintf(w, "%s : ", this.Id)
	for _, term := range this.Terms {
		fmt.Fprintf(w, "%s ", term)
	}
	if this.Action != "" {
		fmt.Fprintf(w, "<< %s >> ", this.Action)
	}
	w.WriteString(";")
	return w.String()
}
