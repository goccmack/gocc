package nolexer

import (
	"code.google.com/p/gocc/example/nolexer/parser"
	"code.google.com/p/gocc/example/nolexer/scanner"
	"testing"
)

func Test1(t *testing.T) {
	S := scanner.NewString("hiya world")
	P := parser.NewParser()
	if _, e := P.Parse(S); e != nil {
		t.Error(e.Error())
	}
}

func Test2(t *testing.T) {
	S := scanner.NewString("hello world")
	P := parser.NewParser()
	if _, e := P.Parse(S); e != nil {
		t.Error(e.Error())
	}
}
