package ast

import (
	"testing"
)

type actExprTestdataStruct struct {
	src    string
	rawExp string
	exp    string
}

var actExprTestdata = []actExprTestdataStruct{
	{
		src: "<< ast.SomeFunc($...), nil >>",
		exp: "ast.SomeFunc(X...), nil",
	},
	{
		src: "<< ast.SomeFunc($0, $1), nil >>",
		exp: "ast.SomeFunc(X[0], X[1]), nil",
	},
	{
		src: "<< append([]Attrib{}, $...), nil >>",
		exp: "append([]Attrib{}, X...), nil",
	},
}

func TestActExp1(t *testing.T) {
	for _, tst := range actExprTestdata {
		act := ActionExpressionVal([]byte(tst.src))
		if act != tst.exp {
			t.Fatalf("Expected `%s`, got `%s`", tst.exp, act)
		}
	}
}
