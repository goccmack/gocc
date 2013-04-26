//Copyright 2012 Vastech SA (PTY) LTD
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

package lr1

import (
	"code.google.com/p/gocc/ast"
	"code.google.com/p/gocc/parser"
	"code.google.com/p/gocc/scanner"
	"code.google.com/p/gocc/token"
	"fmt"
	"testing"
)

const (
	dragon_book_grammar_4_55_bnf = `S : C C ;
C : c C ;
C : d ;`
)

var (
	aug        *ast.Grammar
	item0      *Item
	I0, I1, I2 ItemSet
	C, C_Full  ItemSets
)

func TestSetup(t *testing.T) {
	tokenmap := token.NewMapFromStrings(token.GoccStrings)
	scanner := new(scanner.Scanner)
	srcBuffer := []byte(dragon_book_grammar_4_55_bnf)
	scanner.Init(srcBuffer, tokenmap)
	parser := parser.NewParser(scanner, tokenmap)
	aug, err := ast.NewAugmentedGrammar(parser.Parse())
	if err != nil {
		fmt.Println("Error creating aug:", err)
	}

	item0 = InitialItem(aug)
	fmt.Println("Item0:", item0)

	I0 = Closure(ItemSet{item0}, aug.FirstSets())
	fmt.Println("I0:", I0)

	// I1 = Goto(I0, aug.GTokens.GetToken("S"), aug.FirstSets())
	// fmt.Println("I1", I1)
	//
	// I2 = Goto(I0, aug.GTokens.GetToken("C"), aug.FirstSets())
	// fmt.Println("I2", I2)
	//
	// C = ItemSets{I0, I1}
	//
	// C_Full, _ = GetItemSets(aug)
	// fmt.Println("C_Full:\n" + C_Full.String())
}

// func TestItem0(t *testing.T) {
// 	tmp := &Item{0, 0, aug.GTokens.GetToken(gtoken.EOF_STR), aug}
// 	if !item0.Equals(tmp) {
// 		t.Fail()
// 		fmt.Println("item0: ", item0)
// 		fmt.Println(" temp: ", tmp)
// 	}
// }
//
// func TestFirstS1(t *testing.T) {
// 	gStr := `
// 		A : B C ;
// 		B : b ;
// 		B : ε ;
// 		C : c ;
// 	`
// 	lex := new(scanner.Scanner)
// 	lex.Init("", []byte(gStr), nil, 0)
// 	g := bnfparser.NewParser(lex).Parse()
// 	gfs := g.FirstS(g.Prod[0].Handle)
// 	fsCtl := gtoken.GTokenList{
// 		g.GTokens.GetToken("b"), g.GTokens.GetToken("c") }
// 	if !gfs.EqualsSet(fsCtl) { t.Error("fsCtl:" + fsCtl.String() + " : " + gfs.String())}
// }
//
// func TestFirstS2(t *testing.T) {
// 	gStr := `
// 		A : B C ;
// 		B : b ;
// 		B : ε ;
// 		C : c ;
// 		C : ε ;
// 	`
// 	lex := new(scanner.Scanner)
// 	lex.Init("", []byte(gStr), nil, 0)
// 	g := bnfparser.NewParser(lex).Parse()
// 	gfs := g.FirstS(g.Prod[0].Handle)
// 	fsCtl := gtoken.GTokenList{
// 		g.GTokens.GetToken("b"), g.GTokens.GetToken("c"), g.GTokens.GetToken(gtoken.EPSILON_STR) }
// 	if !gfs.EqualsSet(fsCtl) { t.Error("fsCtl:" + fsCtl.String() + " : " + gfs.String())}
// }
//
// func TestI0(t *testing.T) {
// 	i0Control := ItemSet {
// 		&Item{0, 0, aug.GTokens.GetToken(gtoken.EOF_STR), aug},
// 		&Item{1, 0, aug.GTokens.GetToken(gtoken.EOF_STR), aug},
// 		&Item{2, 0, aug.GTokens.GetToken("c"), aug},
// 		&Item{2, 0, aug.GTokens.GetToken("d"), aug},
// 		&Item{3, 0, aug.GTokens.GetToken("c"), aug},
// 		&Item{3, 0, aug.GTokens.GetToken("d"), aug},
// 	}
// 	if !C_Full[0].Equals(i0Control) {
// 		t.Fail()
// 	}
// }
//
// func TestGetIndex(t *testing.T) {
// 	if C.GetIndex(I0) != 0 {
// 		fmt.Println("C.GetIndex(I0) != 0")
// 		t.Fail()
// 	}
// 	if C.GetIndex(I1) != 1 {
// 		fmt.Println("C.GetIndex(I1) != 1")
// 		t.Fail()
// 	}
// 	if C.GetIndex(I2) != -1 {
// 		fmt.Println("C.GetIndex(I2) != -1")
// 		t.Fail()
// 	}
// }
//
// func TestContains(t *testing.T) {
// 	if !C.Contains(I0) {
// 		fmt.Println("!C.Contains(I0)")
// 		t.Fail()
// 	}
// 	if !C.Contains(I1) {
// 		fmt.Println("!C.Contains(I1)")
// 		t.Fail()
// 	}
// 	if C.Contains(I2) {
// 		fmt.Println("C.Contains(I2)")
// 		t.Fail()
// 	}
// }
//
// func TestCLen(t *testing.T) {
// 	if len(C_Full) != 10 {
// 		fmt.Println("len(C_Full) =", len(C_Full))
// 		t.Fail()
// 	}
// }
