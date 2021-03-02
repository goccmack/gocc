// Code generated by gocc; DO NOT EDIT.

package errors

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/goccmack/gocc/example/calc/token"
)

// Severity is the prefix given to error messages in the string representation.
// It's exposed as a var so that it can be changed during testing to avoid the CI
// system thinking it's seeing actual errors.
var Severity = "error"

type ErrorSymbol interface {
}

type Error struct {
	Err            error
	ErrorToken     *token.Token
	ErrorSymbols   []ErrorSymbol
	ExpectedTokens []string
	StackTop       int
}

func (e *Error) String() string {
	w := new(strings.Builder)
	if e.Err != nil {
		fmt.Fprintln(w, "Error ", e.Err)
	} else {
		fmt.Fprintln(w, "Error")
	}
	fmt.Fprintf(w, "Token: type=%d, lit=%s\n", e.ErrorToken.Type, e.ErrorToken.Lit)
	fmt.Fprintf(w, "Pos: offset=%d, line=%d, column=%d\n", e.ErrorToken.Pos.Offset, e.ErrorToken.Pos.Line, e.ErrorToken.Pos.Column)
	fmt.Fprint(w, "Expected one of: ")
	for _, sym := range e.ExpectedTokens {
		fmt.Fprint(w, string(sym), " ")
	}
	fmt.Fprintln(w, "ErrorSymbol:")
	for _, sym := range e.ErrorSymbols {
		fmt.Fprintf(w, "%v\n", sym)
	}

	return w.String()
}

func describeExpected(tokens []string) string {
	switch len(tokens) {
	case 0:
		return "unexpected additional tokens"

	case 1:
		return "expected " + tokens[0]

	case 2:
		return "expected either " + tokens[0] + " or " + tokens[1]

	default: // e.g. expected one of bool, int, float, or string
		tokens = append(tokens[:len(tokens)-1], "or "+tokens[len(tokens)-1])
		return "expected one of " + strings.Join(tokens, ", ")
	}
}

func (e *Error) Error() string {
	// identify the line and column of the error in 'gnu' style so it can be understood
	// by editors and IDEs; user will need to prefix it with a filename.
	text := fmt.Sprintf("%d:%d: %s: ", e.ErrorToken.Pos.Line, e.ErrorToken.Pos.Column, Severity)

	if e.Err != nil {
		// Custom error specified, e.g. by << nil, errors.New("missing newline") >>
		text += e.Err.Error()
	} else {
		tokens := make([]string, len(e.ExpectedTokens))
		for idx, token := range e.ExpectedTokens {
			if !unicode.IsLetter(rune(token[0])) {
				token = strconv.Quote(token)
			}
			tokens[idx] = token
		}
		text += describeExpected(tokens)
		text += fmt.Sprintf("; got: %s", e.ErrorToken.Lit)
	}

	return text
}
