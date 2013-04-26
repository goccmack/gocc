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

//Package ast has all the structures for the abstract syntax tree for the parsed gocc bnf.
package ast

import (
	"code.google.com/p/gocc/token"
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

//The list of productions.
type ProdS []*Production

//Creates a new list of productions.
func NewProdS(head *Symbol, bodies BodyS) ProdS {
	ps := make(ProdS, 0, len(bodies))
	for _, bod := range bodies {
		ps = append(ps, NewProduction(head, bod))
	}
	return ps
}

//Appends a list of productions to the list of productions.
func (P ProdS) Append(prods ProdS) ProdS {
	return append(P, prods...)
}

//Returns a string representing the productions.
func (P ProdS) String() string {
	res := ""
	for i, p := range P {
		res += p.String() + "\n"
		if i < len(P)-1 {
			res += "\n"
		}
	}
	return res
}

//The Production.
type Production struct {
	Head *Symbol
	Body *Body
}

//Creates a new Production.
func NewProduction(head *Symbol, body *Body) *Production {
	return &Production{Head: head, Body: body}
}

//Returns whether two productions are equal.
func (this *Production) Equals(that *Production) bool {
	if that == nil {
		return false
	}

	return this.Head == that.Head && this.Body.Equals(that.Body)
}

//Returns a string representing a production.
func (this *Production) String() string {
	return this.Head.TokLit + " : " + this.Body.String() + " ;"
}

//Returns a graphviz DOT representation of the production.
func (this *Production) Dot(i string) string {
	production_node := "node_production_" + i
	res := production_node + ";\n"
	res += production_node + " [label=\"Prod\"] ;\n"
	name_node := "node_production_name_" + i
	res += production_node + " -> " + name_node + ";\n"
	res += name_node + " [label=\"" + this.Head.String() + "\"] ;\n"
	res += production_node + " -> " + this.Body.String()
	return res
}

//The Production Body.
type Body struct {
	Symbols SymbolS
	SDT     string
}

var NULL_BODY = &Body{Symbols: SymbolS{}}

// Creates a null body for an ℇ-production
func NullBody() *Body {
	return &Body{}
}

//Creates a new body.
func NewBody(symbols SymbolS) *Body {
	return &Body{Symbols: symbols}
}

//Creates a new body with a sdt rule.
func NewBodySDT(symbols SymbolS, sdt string) *Body {
	return &Body{Symbols: symbols, SDT: sdt}
}

//Returns whether two Bodies are equal.
func (this *Body) Equals(that *Body) bool {
	if this == nil || that == nil {
		return false
	}

	return this.Symbols.Equals(that.Symbols) &&
		this.SDT == that.SDT
}

//Returns a string representing the Body.
func (B *Body) String() string {
	if len(B.Symbols) == 0 {
		return "ℇ"
	}

	res := B.Symbols.String()
	if B.SDT != "" {
		res += " << " + B.SDT + " >>"
	}

	return res
}

//A list of Bodies.
type BodyS []*Body

//Creates a new list of Bodies.
func NewBodyS(body *Body) BodyS {
	b := make(BodyS, 0, 5)
	return append(b, body)
}

//Appends a Body to the list of Bodies.
func (B BodyS) Append(body *Body) BodyS {
	return append(B, body)
}

//A list of Symbols.
type SymbolS []*Symbol

//Creates a new list of Symbols.
func NewSymbolS(sym *Symbol) SymbolS {
	res := make(SymbolS, 0, 5)
	res = append(res, sym)
	return res
}

//Returns the Union of two lists of Symbols.
func (this SymbolS) AddSet(that SymbolS) SymbolS {
	res := this
	for _, s := range that {
		res = res.AddSetElement(s)
	}
	return res
}

//Adds a Symbol to the list of Symbols if it is not already contained in the list.
func (S SymbolS) AddSetElement(sym *Symbol) SymbolS {
	if S.Contains(sym) {
		return S
	}

	return S.Append(sym)
}

//Appends a list of Symbols to the list Symbols.
func (S SymbolS) Append(s ...*Symbol) SymbolS {
	return append(S, s...)
}

//Returns whether a Symbol is contained in the list of Symbols.
func (S SymbolS) Contains(sym *Symbol) bool {
	for _, s := range S {
		if s.Equals(sym) {
			return true
		}
	}
	return false
}

//Returns whether two lists of Symbols are equal.
func (this SymbolS) Equals(that SymbolS) bool {
	if len(this) != len(that) {
		return false
	}

	for i, s := range this {
		if !s.Equals(that[i]) {
			return false
		}
	}

	return true
}

//Returns the list of Symbols without the Symbols in the given list.
func (S SymbolS) Min(sym *Symbol) SymbolS {
	if len(S) == 0 {
		return SymbolS{}
	}

	res := make(SymbolS, 0, len(S)-1)
	for _, s := range S {
		if !s.Equals(sym) {
			res = append(res, s)
		}
	}
	return res
}

//Returns a string representing the list of Symbols.
func (S SymbolS) String() string {
	res := ""
	for i, s := range S {
		if s == nil {
			fmt.Println("ast.SymbolS.String(): s == nil")
			fmt.Println("len(S)=", len(S), " cap(S)=", cap(S), " i=", i)
		}
		res += s.TokLit
		if i < len(S)-1 {
			res += " "
		}
	}
	return res
}

//Returns a list of strings, representing each symbols respectively.
func (S SymbolS) Strings() []string {
	res := make([]string, len(S))
	for i, s := range S {
		res[i] = s.String()
	}
	return res
}

//The Symbol.
type Symbol struct {
	SymType    SymbolType
	TokType    token.Type
	TokTypeStr string
	TokLit     string
}

//The SymbolType.
type SymbolType int

const (
	ID SymbolType = iota
	KEYWORD
	STRING_LIT
	CHAR_LIT
	EOF
	EPSILON_LIT
	ERROR_LIT
)

var (
	EOF_SYMBOL = &Symbol{EOF, 0, "$", "$"}

	// gocc.bnf does not include ε
	EPSILON_SYM = &Symbol{KEYWORD, 1, "ε", "ε"}
)

//Returns a string representing the SymbolType.
func (this SymbolType) String() string {
	switch this {
	case ID:
		return "ID"
	case KEYWORD:
		return "KEYWORD"
	case STRING_LIT:
		return "STRING_LIT"
	case CHAR_LIT:
		return "CHAR_LIT"
	case EOF:
		return "EOF"
	case EPSILON_LIT:
		return "EPSILON_LIT"
	case ERROR_LIT:
		return "ERROR_LIT"
	}
	return "ILLEGAL"
}

//Creates a new Symbol.
func NewSymbol(symType SymbolType, tok *token.Token, tm *token.TokenMap) (sym *Symbol, err error) {
	sym = new(Symbol)
	switch symType {
	case ID:
		sym.SymType = ID
		sym.TokLit = string(tok.Lit)
		if isTerminal(sym.TokLit) {
			sym.TokTypeStr = sym.TokLit
		} else {
			sym.TokTypeStr = "id"
		}
	case STRING_LIT:
		sym.SymType = KEYWORD
		sym.TokLit = string(tok.Lit[1 : len(tok.Lit)-1])
		sym.TokTypeStr = sym.TokLit
	case CHAR_LIT:
		sym.SymType = CHAR_LIT
		sym.TokLit = string(tok.Lit)
		sym.TokTypeStr = "char_lit"
	case ERROR_LIT:
		sym.SymType = ERROR_LIT
		sym.TokLit = string(tok.Lit)
		sym.TokTypeStr = "error"
	default:
		err = errors.New("expected id/string/char")
		sym.TokLit = string(tok.Lit)
		sym.TokTypeStr = "illegal"
	}
	tm.AddToken(sym.TokTypeStr)
	sym.TokType = tm.Type(sym.TokTypeStr)

	return
}

//Creates a Copy of the Symbol.
func (S *Symbol) Clone() *Symbol {
	return &Symbol{S.SymType, S.TokType, S.TokTypeStr, S.TokLit}
}

//Returns whether two Symbols are equal.
func (this *Symbol) Equals(that *Symbol) bool {
	if this == nil || that == nil {
		return false
	}

	return this.SymType == that.SymType &&
		this.TokType == that.TokType &&
		this.TokLit == that.TokLit
}

//Returns whether the Symbol is a Terminal.
func (S *Symbol) IsTerminal() bool {
	return S.SymType != ID || isTerminal(S.TokLit)
}

func isTerminal(lit string) bool {
	if unicode.IsUpper([]rune(lit)[0]) {
		return false
	}

	return true
}

//Returns a string representing the Symbol.
func (S *Symbol) String() string {
	res := "(" + S.SymType.String() + "(" + strconv.Itoa(int(S.SymType)) + ") ,"
	res += S.TokTypeStr + "(" + strconv.Itoa(int(S.TokType)) + "),"
	res += S.TokLit + ")"
	return res
}
