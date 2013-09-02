
package errors

import(
	"bytes"
	"fmt"
	"code.google.com/p/gocc/test/sr/token"
)

type ErrorSymbol interface {
}

type Error struct {
	Err            error
	ErrorToken     *token.Token
	ErrorLit       []byte
	ErrorSymbols   []ErrorSymbol
	ExpectedTokens []string
}

func (E *Error) String() string {
	w := new(bytes.Buffer)
	//TODO: refactor to  print token string & pos properly
	fmt.Fprintf(w, "Got %s ", token.TokMap.TokenString(E.ErrorToken))
	if E.Err != nil {
		fmt.Fprintf(w, "%s", E.Err.Error())
	} else {
		fmt.Fprintf(w, ", expected one of: ")
		for _, t := range E.ExpectedTokens {
			fmt.Fprintf(w, "%s ", t)
		}
	}
	return w.String()
}
