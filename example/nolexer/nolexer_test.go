package nolexer

import (
	"testing"

	"github.com/maxcalandrelli/gocc/example/nolexer/nolexer.grammar/nolexer"
	"github.com/maxcalandrelli/gocc/example/nolexer/scanner"
)

func Test1(t *testing.T) {
	S := scanner.NewString("hiya world")
	P := nolexer.NewParser()
	if _, e := P.Parse(S); e != nil {
		t.Error(e.Error())
	}
}

func Test2(t *testing.T) {
	S := scanner.NewString("hello world")
	P := nolexer.NewParser()
	if _, e := P.Parse(S); e != nil {
		t.Error(e.Error())
	}
}
