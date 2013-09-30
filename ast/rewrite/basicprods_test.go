package rewrite

import (
	"fmt"
	"code.google.com/p/gocc/ast"
	"code.google.com/p/gocc/frontend/lexer"
	"code.google.com/p/gocc/frontend/parser"
	"testing"
)

type testDataStruct struct {
	src      string
	expected []string
}

var testData = []*testDataStruct{
	&testDataStruct{
		src: `A : "c" ;`,
		expected: []string{
			`S' : A ;`,
			`A : "c" ;`,
		},
	},

	&testDataStruct{
		src: `A : "a" << a >>| "b" << b >> | "c" << c >>;`,
		expected: []string{
			`S' : A ;`,
			`A : "a" << a >> ;`,
			`A : "b" << b >> ;`,
			`A : "c" << c >> ;`,
		},
	},

	&testDataStruct{
		src: `A : "a" ("b" << b >> | "c" << c >> ) "d" ;`,
		expected: []string{
			`S' : A ;`,
			`A : "a" A~1 "d" ;`,
			`A~1 : "b" << b >> ;`,
			`A~1 : "c" << c >> ;`,
		},
	},

	&testDataStruct{
		src: `A : "a" ["b" << b >> | "c" << c >> ] "d" << d >> ;`,
		expected: []string{
			`S' : A ;`,
			`A : "a" A~1 "d" << d >> ;`,
			`A : "a" "d" << d >> ;`,
			`A~1 : "b" << b >> ;`,
			`A~1 : "c" << c >> ;`,
		},
	},

	&testDataStruct{
		src: `A : "a" {"b" << b >> | "c" << c >> } "d" << d >> ;`,
		expected: []string{
			`S' : A ;`,
			`A : "a" A~1 "d" << d >> ;`,
			`A : "a" "d" << d >> ;`,
			`A~1 : A~1_RepTerm << []interface{}{$0}, nil >> ;`,
			`A~1 : A~1 A~1_RepTerm << append($0.([]interface{}), $1), nil >> ;`,
			`A~1_RepTerm : "b" << b >> ;`,
			`A~1_RepTerm : "c" << c >> ;`,
		},
	},

	&testDataStruct{
		src: `A : "a" {["b"] "b" << b >> | "c" <<c>>} "d" <<d>> ;`,
		expected: []string{
			`S' : A ;`,
			`A : "a" A~1 "d" << d >> ;`,
			`A : "a" "d" << d >> ;`,
			`A~1 : A~1_RepTerm << []interface{}{$0}, nil >> ;`,
			`A~1 : A~1 A~1_RepTerm << append($0.([]interface{}), $1), nil >> ;`,
			`A~1_RepTerm : A~1_RepTerm~1 "b" << b >> ;`,
			`A~1_RepTerm : "b" << b >> ;`,
			`A~1_RepTerm~1 : "b" ;`,
			`A~1_RepTerm : "c" << c >> ;`,
		},
	},

	&testDataStruct{
		src: `A : "a" {{"b" << b >> } "c" <<c>>} "d" <<d>> ;`,
		expected: []string{
			`S' : A ;`,
			`A : "a" A~1 "d" << d >> ;`,
			`A : "a" "d" << d >> ;`,
			`A~1 : A~1_RepTerm << []interface{}{$0}, nil >> ;`,
			`A~1 : A~1 A~1_RepTerm << append($0.([]interface{}), $1), nil >> ;`,
			`A~1_RepTerm : A~1_RepTerm~1 "c" << c >> ;`,
			`A~1_RepTerm : "c" << c >> ;`,
			`A~1_RepTerm~1 : A~1_RepTerm~1_RepTerm << []interface{}{$0}, nil >> ;`,
			`A~1_RepTerm~1 : A~1_RepTerm~1 A~1_RepTerm~1_RepTerm << append($0.([]interface{}), $1), nil >> ;`,
			`A~1_RepTerm~1_RepTerm : "b" << b >> ;`,
		},
	},
}

func Test1(t *testing.T) {
	for i, data := range testData {
		fmt.Printf("%s\n", data.src)
		g, err := parser.NewParser().Parse(lexer.NewLexer([]byte(data.src)))
		if err != nil {
			t.Fatalf("%s", err)
		}
		prods := BasicProds(g.(*ast.Grammar).SyntaxPart.ProdList)
		fmt.Printf("Prods:\n")
		for _, prod := range prods {
			fmt.Printf("\t%s\n", prod)
		}
		if len(prods) != len(data.expected) {
			t.Errorf("%d: len(prods)=%d, expected %d", i, len(prods), len(data.expected))
		} else {
			for i, prod := range prods {
				if prod.String() != data.expected[i] {
					t.Errorf("Prod[%d]=`%s`, expected `%s`", i, prod.String(), data.expected[i])
				}
			}
		}
	}
}

func parse(src string) (*ast.Grammar, error) {
	if g, err := parser.NewParser().Parse(lexer.NewLexer([]byte(src))); err != nil {
		return nil, err
	} else {
		return g.(*ast.Grammar), nil
	}
}
