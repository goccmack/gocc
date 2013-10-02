//Copyright 2013 Vastech SA (PTY) LTD
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package rewrite

import (
	"code.google.com/p/gocc/ast"
	"code.google.com/p/gocc/frontend/lexer"
	"code.google.com/p/gocc/frontend/parser"
	"fmt"
	"testing"
)

type testDataStruct struct {
	src      string
	expected []string
}

var testData = []*testDataStruct{
	{
		src: `A : "c" ;`,
		expected: []string{
			`S' : A ;`,
			`A : "c" ;`,
		},
	},

	{
		src: `A : "a" << a >>| "b" << b >> | "c" << c >>;`,
		expected: []string{
			`S' : A ;`,
			`A : "a" << a >> ;`,
			`A : "b" << b >> ;`,
			`A : "c" << c >> ;`,
		},
	},

	{
		src: `A : "a" ("b" << b >> | "c" << c >> ) "d" ;`,
		expected: []string{
			`S' : A ;`,
			`A : "a" A~1 "d" ;`,
			`A~1 : "b" << b >> ;`,
			`A~1 : "c" << c >> ;`,
		},
	},

	{
		src: `A : "a" ["b" << b >> | "c" << c >> ] "d" << d >> ;`,
		expected: []string{
			`S' : A ;`,
			`A : "a" A~1 "d" << d >> ;`,
			`A : "a" "d" << d >> ;`,
			`A~1 : "b" << b >> ;`,
			`A~1 : "c" << c >> ;`,
		},
	},

	{
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

	{
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

	{
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
