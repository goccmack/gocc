package errormsg

import (
	"fmt"
	"testing"

	"github.com/johnkerl/gocc/example/errormsg/errors"
	"github.com/johnkerl/gocc/example/errormsg/lexer"
	"github.com/johnkerl/gocc/example/errormsg/parser"
	"github.com/johnkerl/gocc/example/errormsg/token"
)

// assertEqual is a crude implementation of testify's assert.Equal.
func assertEqual(t *testing.T, lhs interface{}, rhs interface{}) {
	t.Helper()

	if lhs != rhs {
		t.Fatalf("mismatch: expected %v, got %v", lhs, rhs)
	}
}

// parse boolean-checks the parseability of a given string against 'expectPass'.
func parse(t *testing.T, code string, expectPass bool) {
	t.Helper()

	l := lexer.NewLexer([]byte(code))
	p := parser.NewParser()
	_, err := p.Parse(l)

	// Match err to expectPass (nil = passed, !nil = failed), and log if not nill
	if err == nil {
		if !expectPass {
			t.Fatal("test should have failed, got nil error instead")
		}
	} else {
		if expectPass && err != nil {
			t.Fatalf("test should have passed, got error instead: %s", err.Error())
		}
		t.Log(err.Error())
	}
}

func TestParsedErrors_passes(t *testing.T) {
	// Validate basic assumptions, that supposed-good strings parse ok.
	candidates := []string{
		"var abcd = 123;",
		"  var  _  :=  abcd  ;",
		"var a  = 1.23 ;",
		"var x;",
	}
	for _, candidate := range candidates {
		candidate := candidate
		t.Run(candidate, func(t *testing.T) {
			parse(t, candidate, true)
		})
	}
}

func TestParsedError_errors(t *testing.T) {
	// Verify that supposed-bad strings generate errors.
	candidates := []string{
		"var a = 1",    // missing ';'.
		"var a = 1\n",  // '\n' not listsed in grammar.
		"let a = 1;",   // first option is var and only var.
		"var = 1;",     // second field must be identifier or '_'.
		"var _ = ;",    // third field must be one of '=', ':=' or ';'.
		"var xyz = {}", // fourth, 'Default' has four candidates.
	}
	for _, candidate := range candidates {
		candidate := candidate
		t.Run(candidate, func(t *testing.T) {
			parse(t, candidate, false)
		})
	}
}

func TestParsedErrors_ExtraTokens(t *testing.T) {
	// Input is supposed to end at ';', so errors.Errors will have an
	// empty expected tokens list.
	parse(t, "var end = 1; oops", false)
}

func mockToken(tokenType token.Type, lit string, line, col int) *token.Token {
	return &token.Token{Type: tokenType, Lit: []byte(lit), Pos: token.Pos{Line: line, Column: col}}
}

func TestErrors_DescribeExpected(t *testing.T) {
	assertEqual(t, "unexpected additional tokens", errors.DescribeExpected([]string{}))
	assertEqual(t, "expected TREE", errors.DescribeExpected([]string{"TREE"}))
	assertEqual(t, "expected either TREE or ENT", errors.DescribeExpected([]string{"TREE", "ENT"}))
	assertEqual(t, "expected one of TREE, ENT or i am groot", errors.DescribeExpected([]string{"TREE", "ENT", "i am groot"}))
	assertEqual(t, "expected one of a, b, c, d, e, f, or g", errors.DescribeExpected([]string{"a", "b", "c", "d", "e", "f", "g"}))
}

func TestErrors_DescribeToken(t *testing.T) {
	t.Run("eof", func(t *testing.T) {
		tok := mockToken(token.EOF, "-not-eof-", 1, 1)
		assertEqual(t, "end-of-file", errors.DescribeToken(tok))
	})

	t.Run("eof", func(t *testing.T) {
		tok := mockToken(token.INVALID, "-not-eof-", 1, 1)
		assertEqual(t, "unknown/invalid token \"-not-eof-\"", errors.DescribeToken(tok))
	})

	t.Run("eof", func(t *testing.T) {
		tok := mockToken(9001, "-not-eof-", 1, 1)
		assertEqual(t, "\"-not-eof-\"", errors.DescribeToken(tok))
	})
}

func TestErrors_Error(t *testing.T) {
	// Direct testing by manually constructing Error objects.
	t.Run("custom error", func(t *testing.T) {
		err := &errors.Error{ErrorToken: mockToken(999, "", 6, 7), Err: fmt.Errorf("source on fire")}
		assertEqual(t, "6:7: error: source on fire", err.Error())
	})

	t.Run("no tokens", func(t *testing.T) {
		err := &errors.Error{ErrorToken: mockToken(888, "biscuit", 10, 12), ExpectedTokens: []string{}}
		assertEqual(t, "10:12: error: unexpected additional tokens; got: \"biscuit\"", err.Error())
	})

	t.Run("unexpected EOF", func(t *testing.T) {
		err := &errors.Error{ErrorToken: mockToken(token.EOF, "", 7, 11), ExpectedTokens: []string{"something-else"}}
		assertEqual(t, "7:11: error: expected something-else; got: end-of-file", err.Error())
	})

	t.Run("nominal error", func(t *testing.T) {
		err := &errors.Error{ErrorToken: mockToken(100, "42", 7, 6), ExpectedTokens: []string{"var", "let", "struct"}}
		assertEqual(t, "7:6: error: expected one of var, let or struct; got: \"42\"", err.Error())
	})
}

func TestErrors_ErrorConext(t *testing.T) {
	// "Before": error without a way to know the source file.
	err := &errors.Error{ErrorToken: mockToken(111, "moyles", 16, 5), ExpectedTokens: []string{"ant", "dec"}}
	assertEqual(t, `16:5: error: expected either ant or dec; got: "moyles"`, err.Error())

	// Now attach a Context that implements `Source() string` and verify we
	// get a proper File-Line-Column error using the same err.
	err.ErrorToken.Context = &lexer.SourceContext{Filepath: "/addicted/to/plaice.lyrics"}
	assertEqual(t, `/addicted/to/plaice.lyrics:16:5: error: expected either ant or dec; got: "moyles"`, err.Error())
}
