package calc

import (
	"code.google.com/p/gocc/example/calc/parser"
	"code.google.com/p/gocc/example/calc/scanner"
	"code.google.com/p/gocc/example/calc/token"
	"fmt"
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
	s := &scanner.Scanner{}
	p := parser.NewParser(parser.ActionTable, parser.GotoTable, parser.ProductionsTable, token.CALCTokens)
	pass := true
	for _, ts := range testData {
		s.Init([]byte(ts.src), token.CALCTokens)
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
