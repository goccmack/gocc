package sdt_error

import (
	"testing"

	"github.com/goccmack/gocc/internal/test/sdt_error/errors"
	"github.com/goccmack/gocc/internal/test/sdt_error/lexer"
	"github.com/goccmack/gocc/internal/test/sdt_error/parser"
)

func parseInput(src string) (interface{}, error) {
	l := lexer.NewLexer([]byte(src))
	p := parser.NewParser()
	return p.Parse(l)
}

func TestSimpleActionErrorToken(t *testing.T) {
	src := "fail"
	_, err := parseInput(src)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	parseErr, ok := err.(*errors.Error)
	if !ok {
		t.Fatalf("expected *errors.Error, got %T (%v)", err, err)
	}

	if parseErr.ErrorToken == nil {
		t.Fatalf("expected non-nil ErrorToken")
	}

	if got := string(parseErr.ErrorToken.Lit); got != "fail" {
		t.Errorf("ErrorToken literal: got %q, want %q", got, "fail")
	}

	if got := parseErr.ErrorToken.Pos.Line; got != 1 {
		t.Errorf("ErrorToken line: got %d, want 1", got)
	}

	if got := parseErr.ErrorToken.Pos.Column; got != 1 {
		t.Errorf("ErrorToken column: got %d, want 1", got)
	}

	if got := parseErr.ErrorToken.Pos.Offset; got != 0 {
		t.Errorf("ErrorToken offset: got %d, want 0", got)
	}

	const expectedMsg = "1:1: error: simple action failed"
	if got := parseErr.Error(); got != expectedMsg {
		t.Errorf("formatted error: got %q, want %q", got, expectedMsg)
	}
}

func TestMultiTokenActionErrorTokenInline(t *testing.T) {
	src := "start middle end"
	_, err := parseInput(src)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	parseErr, ok := err.(*errors.Error)
	if !ok {
		t.Fatalf("expected *errors.Error, got %T (%v)", err, err)
	}

	if parseErr.ErrorToken == nil {
		t.Fatalf("expected non-nil ErrorToken")
	}

	// The error token must correspond to the start of the production ("start"),
	// not the end token ("end") or the EOF lookahead token.
	if got := string(parseErr.ErrorToken.Lit); got != "start" {
		t.Errorf("ErrorToken literal: got %q, want %q", got, "start")
	}

	if got := parseErr.ErrorToken.Pos.Line; got != 1 {
		t.Errorf("ErrorToken line: got %d, want 1", got)
	}

	if got := parseErr.ErrorToken.Pos.Column; got != 1 {
		t.Errorf("ErrorToken column: got %d, want 1", got)
	}

	if got := parseErr.ErrorToken.Pos.Offset; got != 0 {
		t.Errorf("ErrorToken offset: got %d, want 0", got)
	}

	const expectedMsg = "1:1: error: multi token action failed"
	if got := parseErr.Error(); got != expectedMsg {
		t.Errorf("formatted error: got %q, want %q", got, expectedMsg)
	}
}

func TestMultiTokenActionErrorTokenMultiline(t *testing.T) {
	src := "  start\n    middle\n      end"
	_, err := parseInput(src)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	parseErr, ok := err.(*errors.Error)
	if !ok {
		t.Fatalf("expected *errors.Error, got %T (%v)", err, err)
	}

	if parseErr.ErrorToken == nil {
		t.Fatalf("expected non-nil ErrorToken")
	}

	if got := string(parseErr.ErrorToken.Lit); got != "start" {
		t.Errorf("ErrorToken literal: got %q, want %q", got, "start")
	}

	if got := parseErr.ErrorToken.Pos.Line; got != 1 {
		t.Errorf("ErrorToken line: got %d, want 1", got)
	}

	if got := parseErr.ErrorToken.Pos.Column; got != 3 {
		t.Errorf("ErrorToken column: got %d, want 3", got)
	}

	if got := parseErr.ErrorToken.Pos.Offset; got != 2 {
		t.Errorf("ErrorToken offset: got %d, want 2", got)
	}

	const expectedMsg = "1:3: error: multi token action failed"
	if got := parseErr.Error(); got != expectedMsg {
		t.Errorf("formatted error: got %q, want %q", got, expectedMsg)
	}
}

func TestNestedActionErrorToken(t *testing.T) {
	// NestedAction reduces NonTerminalPrefix "suffix".
	// The start token must propagate from NonTerminalPrefix's first token ("nested").
	src := "\n  nested prefix\n    suffix"
	_, err := parseInput(src)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	parseErr, ok := err.(*errors.Error)
	if !ok {
		t.Fatalf("expected *errors.Error, got %T (%v)", err, err)
	}

	if parseErr.ErrorToken == nil {
		t.Fatalf("expected non-nil ErrorToken")
	}

	if got := string(parseErr.ErrorToken.Lit); got != "nested" {
		t.Errorf("ErrorToken literal: got %q, want %q", got, "nested")
	}

	if got := parseErr.ErrorToken.Pos.Line; got != 2 {
		t.Errorf("ErrorToken line: got %d, want 2", got)
	}

	if got := parseErr.ErrorToken.Pos.Column; got != 3 {
		t.Errorf("ErrorToken column: got %d, want 3", got)
	}

	if got := parseErr.ErrorToken.Pos.Offset; got != 3 {
		t.Errorf("ErrorToken offset: got %d, want 3", got)
	}

	const expectedMsg = "2:3: error: nested action failed"
	if got := parseErr.Error(); got != expectedMsg {
		t.Errorf("formatted error: got %q, want %q", got, expectedMsg)
	}
}

func TestSyntaxErrorReportsNextToken(t *testing.T) {
	src := "start fail"
	_, err := parseInput(src)
	if err == nil {
		t.Fatalf("expected syntax error, got nil")
	}

	parseErr, ok := err.(*errors.Error)
	if !ok {
		t.Fatalf("expected *errors.Error, got %T (%v)", err, err)
	}

	if parseErr.ErrorToken == nil {
		t.Fatalf("expected non-nil ErrorToken")
	}

	// On syntax error, the unexpected next token must be reported.
	if got := string(parseErr.ErrorToken.Lit); got != "fail" {
		t.Errorf("ErrorToken literal: got %q, want %q", got, "fail")
	}

	if got := parseErr.ErrorToken.Pos.Column; got != 7 {
		t.Errorf("ErrorToken column: got %d, want 7", got)
	}
}
