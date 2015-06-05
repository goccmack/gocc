package calc

import (
	"fmt"
	"github.com/goccmack/gocc/example/calc/lexer"
	"github.com/goccmack/gocc/example/calc/parser"
	"testing"
)

type TI struct {
	src    string
	expect int64
}

var testData = []*TI{
	{"1 + 1", 2},
	{"1 * 1", 1},
	{"1 + 2 * 3", 7},
}

func Test1(t *testing.T) {
	p := parser.NewParser()
	pass := true
	for _, ts := range testData {
		s := lexer.NewLexer([]byte(ts.src))
		sum, err := p.Parse(s)
		if err != nil {
			pass = false
			t.Log(err.Error())
		}
		if sum != ts.expect {
			pass = false
			t.Log(fmt.Sprintf("Error: %s = %d. Got %d\n", ts.src, sum, ts.expect))
		}
	}
	if !pass {
		t.Fail()
	}
}
