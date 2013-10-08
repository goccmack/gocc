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

package symbols

import (
	"bytes"
	"code.google.com/p/gocc/ast"
	"errors"
	"fmt"
	"sort"
)

type Symbols struct {
	StartSymbol string

	ntList []string
	ntMap  map[string]int

	tList []string
	tMap  map[string]bool

	list   []string
	symMap map[string]bool

	ProdList []string
	prodMap  map[string]int

	stringLitList []string
	stringLitMap  map[string]ast.SyntaxStringLit
}

type Type int

const (
	T_Type Type = iota
	NT_Type
)

func NewSymbols(lexpart *ast.LexPart, prods *ast.SyntaxBasicProdsList) (symbols *Symbols, errs []error) {
	declared_prods := make(map[string]bool)
	symbols = &Symbols{
		ntMap:        make(map[string]int),
		tMap:         make(map[string]bool),
		symMap:       make(map[string]bool),
		prodMap:      make(map[string]int),
		stringLitMap: make(map[string]ast.SyntaxStringLit),
	}
	if prods != nil && prods.Len() > 0 {
		symbols.StartSymbol = prods.First().Id
		for prod := prods.First(); prod != nil; prod = prod.Next {
			declared_prods[prod.Id] = false
		}
		for prod := prods.First(); prod != nil; prod = prod.Next {
			if _, exist := symbols.prodMap[prod.Id]; !exist {
				symbols.ntList = append(symbols.ntList, prod.Id)
				symbols.ntMap[prod.Id] = len(symbols.ntList) - 1
				symbols.ProdList = append(symbols.ProdList, prod.Id)
				symbols.prodMap[prod.Id] = len(symbols.ProdList) - 1
			}
			if err := symbols.addSymbols(prod.Terms, declared_prods); err != nil {
				errs = append(errs, newError(prod, err))
			}
		}
		errs = append(errs, checkUnreachableProds(declared_prods)...)
	} else {
		for _, sym := range lexpart.TokenIds() {
			symbols.tMap[sym] = true
		}
	}
	if lexpart != nil {
		errs = append(errs, symbols.checkTokens(lexpart.TokenIds())...)
	}

	symbols.makeLists()
	return
}

func (this *Symbols) IsNT(prodId string) bool {
	if _, exist := this.ntMap[prodId]; exist {
		return true
	}
	return false
}

func (this *Symbols) List() []string {
	return this.list
}

func (this *Symbols) ListNonTerminals() []string {
	return this.ntList
}

func (this *Symbols) ListStringLitSymbols() []string {
	return this.stringLitList
}

func (this *Symbols) ListTerminals() []string {
	return this.tList
}

func (this *Symbols) NTType(prodId string) int {
	for i, id := range this.ntList {
		if prodId == id {
			return i
		}
	}
	panic(fmt.Sprintf("Unknown production id: %s", prodId))
}

/*** utility functions ***/

func (this *Symbols) addSymbols(terms ast.SyntaxTerms, declared_prods map[string]bool) (errs []error) {
	for _, term := range terms {
		if err := this.addSymbol(term, declared_prods); err != nil {
			errs = append(errs, err)
		}
	}
	return
}

func (this *Symbols) addSymbol(term ast.SyntaxTerm, declared_prods map[string]bool) (err error) {
	switch t := term.(type) {
	case ast.SyntaxStringLit:
		this.symMap[string(t)] = true
		this.tMap[string(t)] = true
		this.stringLitMap[string(t)] = t
	case ast.SyntaxTokId:
		this.symMap[string(t)] = true
		this.tMap[string(t)] = true
	case ast.SyntaxProdId:
		this.symMap[string(t)] = true
		if _, exist := declared_prods[string(t)]; !exist {
			err = errors.New(fmt.Sprintf("reference to unknown production: %s", string(t)))
		} else {
			declared_prods[string(t)] = true
		}
	}
	return
}

func (this Symbols) checkTokens(tokIds []string) (errs []error) {
	for _, tokId := range tokIds {
		if _, exist := this.tMap[tokId]; !exist {
			errs = append(errs, errors.New(fmt.Sprintf("token %s not used in syntax part", tokId)))
		}
	}
	for tokId := range this.tMap {
		if _, exist := this.stringLitMap[tokId]; !exist && !tokIdsContain(tokIds, tokId) {
			errs = append(errs, errors.New(fmt.Sprintf("token %s referenced in syntax part is not defined in lex part", tokId)))
		}
	}
	return
}

func tokIdsContain(tokIds []string, id string) bool {
	for _, tokId := range tokIds {
		if tokId == id {
			return true
		}
	}
	return false
}

func checkUnreachableProds(declared_prods map[string]bool) (errs []error) {
	for prod, referenced := range declared_prods {
		if prod != "S'" && !referenced {
			errs = append(errs, errors.New(fmt.Sprintf("Unreachable production: %s", prod)))
		}
	}
	return
}

func (this *Symbols) makeLists() {
	for sym := range this.tMap {
		this.tList = append(this.tList, sym)
	}
	sort.Strings(this.tList)
	this.tMap["$"] = true
	this.tMap["error"] = true
	this.tMap["INVALID"] = true
	this.tList = append([]string{"INVALID", "$", "error"}, this.tList...)

	for sym := range this.stringLitMap {
		this.stringLitList = append(this.stringLitList, sym)
	}
	sort.Strings(this.stringLitList)

	this.list = append(this.list, this.tList...)
	this.list = append(this.list, this.ntList...)
}

func newError(prod *ast.SyntaxBasicProd, errs []error) error {
	w := new(bytes.Buffer)
	fmt.Fprintf(w, "%s:\n", prod.String())
	for _, err := range errs {
		fmt.Fprintf(w, "\t%s\n", err)
	}
	return errors.New(w.String())
}
