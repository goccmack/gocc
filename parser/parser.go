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

// Parser for gocc BNF source
package parser

import (
	"code.google.com/p/gocc/ast"
	"code.google.com/p/gocc/scanner"
	"code.google.com/p/gocc/token"
	"errors"
	"fmt"
	"strconv"
)

//Parser for gocc BNF source
type Parser struct {
	lex       *scanner.Scanner
	token     *token.Token
	pos       token.Position
	tokenMap  *token.TokenMap
	gTokenMap *token.TokenMap
}

//Creates a New Parser for gocc BNF Source
func NewParser(lex *scanner.Scanner, tokenMap *token.TokenMap) *Parser {
	return &Parser{lex: lex, tokenMap: tokenMap, gTokenMap: token.NewMap()}
}

//Parses the gocc BNF Source and returns the Productions
func (this *Parser) Parse() (ast.ProdS, *token.TokenMap) {
	this.nextToken()
	return this.Grammar(), this.gTokenMap
}

func (this *Parser) error(msg string) {
	panic(fmt.Sprintf("Error at %v %v", this.pos, msg))
}

func (this *Parser) expect(exp string, errMsg string) {
	if this.tokenMap.Type(exp) != this.token.Type {
		msg := exp + " expected but found "
		msg += this.tokenString()
		if errMsg != "" {
			msg += " - " + errMsg
		}
		this.error(msg)
	}
}

func (this *Parser) tokenString() string {
	res := this.tokenMap.TokenString(this.token.Type) + "(" + strconv.Itoa(int(this.token.Type)) + ")"
	res += ">" + string(this.token.Lit) + "<"
	res += " @ " + this.pos.String()
	return res
}

func (P *Parser) la(exp string) bool {
	return P.token.Type == P.tokenMap.Type(exp)
}

func (this *Parser) nextToken() {
	this.token, this.pos = this.lex.Scan()
}

//Grammar	: Production					<< ast.NewGrammar($0) >>
//		    | Grammar Production			<< $0.Append($1) >>
//		    ;				.
func (P *Parser) Grammar() ast.ProdS {
	res := P.Production()

	for !P.la("$") {
		res = res.Append(P.Production())
	}

	return res
}

//Production	: id ":" Alternatives		<< ast.NewProdS($0, $2) >>
//			;
func (this *Parser) Production() ast.ProdS {
	this.expect("id", "production name expected")
	head, _ := ast.NewSymbol(ast.ID, this.token, this.gTokenMap)
	this.nextToken()
	this.expect(":", "")
	this.nextToken()
	prods := ast.NewProdS(head, this.Alternatives())
	this.expect(";", "")
	this.nextToken()

	return prods
}

//Alternatives	:	Body					<< ast.NewBodyS($0) >>
//				|	Body "|" Alternatives	<< $1.Append($0) >>
//				;
func (P *Parser) Alternatives() ast.BodyS {
	bodies := ast.NewBodyS(P.Body())
	for P.la("|") {
		P.nextToken()
		bodies = bodies.Append(P.Body())
	}
	return bodies
}

//Body	: Symbols						<< ast.NewBody($0) >>
//		| Symbols sdt_lit				<< ast.NewBodyTrans($0, $1) >>
//		;
func (P *Parser) Body() *ast.Body {
	symbols := P.Symbols()
	if P.la("sdt_lit") {
		sdt := P.token.SDTVal()
		P.nextToken()
		return ast.NewBodySDT(symbols, sdt)
	}
	return ast.NewBody(symbols)
}

//Symbols	: Symbol						<< ast.NewSymbols($0) >>
//		: Symbol Symbols				<< $1.Append($0) >>
//		;
func (P *Parser) Symbols() ast.SymbolS {
	res := ast.NewSymbolS(P.Symbol())
	for P.isSymbol() {
		res = res.Append(P.Symbol())
	}
	return res
}

func (P *Parser) isSymbol() bool {
	switch P.token.Type {
	case P.tokenMap.Type("id"),
		P.tokenMap.Type("string"),
		P.tokenMap.Type("char"):
		return true
	}
	return false
}

//Symbol	: id
//		| string
//		| char
//		| ε
//		;
func (P *Parser) Symbol() *ast.Symbol {
	sym, err := (*ast.Symbol)(nil), error(nil)

	switch P.token.Type {
	case P.tokenMap.Type("id"):
		sym, err = ast.NewSymbol(ast.ID, P.token, P.gTokenMap)
	case P.tokenMap.Type("string"):
		sym, err = ast.NewSymbol(ast.STRING_LIT, P.token, P.gTokenMap)
	case P.tokenMap.Type("char"):
		sym, err = ast.NewSymbol(ast.CHAR_LIT, P.token, P.gTokenMap)
	case P.tokenMap.Type("ε"):
		sym = ast.EPSILON_SYM
	default:
		err = errors.New("Cannot parse Symbol:" + P.token.String())
		P.error("Cannot parse Symbol:" + err.Error())
	}

	P.nextToken()
	return sym

}
