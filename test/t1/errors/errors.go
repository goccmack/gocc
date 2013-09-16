<<<<<<< HEAD
package errors

import (
	"bytes"
	"code.google.com/p/gocc/test/t1/token"
	"fmt"
=======

package errors

import(
	"bytes"
	"fmt"
	"code.google.com/p/gocc/test/t1/token"
>>>>>>> 9d3e28b6b4a375015991b8ecb296891776774570
)

type ErrorSymbol interface {
}

type Error struct {
	Err            error
	ErrorToken     *token.Token
	ErrorSymbols   []ErrorSymbol
	ExpectedTokens []string
}

func (E *Error) String() string {
	w := new(bytes.Buffer)
	fmt.Fprintf(w, "Error")
	if E.Err != nil {
		fmt.Fprintf(w, " %s\n", E.Err)
	} else {
		fmt.Fprintf(w, "\n")
	}
	fmt.Fprintf(w, "Token: type=%d, lit=%s\n", E.ErrorToken.Type, E.ErrorToken.Lit)
	fmt.Fprintf(w, "Pos: offset=%d, line=%d, column=%d\n", E.ErrorToken.Pos.Offset, E.ErrorToken.Pos.Line, E.ErrorToken.Pos.Column)
	fmt.Fprintf(w, "Expected one of: ")
	for _, sym := range E.ExpectedTokens {
		fmt.Fprintf(w, "%s ", sym)
	}
	fmt.Fprintf(w, "ErrorSymbol:\n")
	for _, sym := range E.ErrorSymbols {
		fmt.Fprintf(w, "%v\n", sym)
	}
	return w.String()
}
