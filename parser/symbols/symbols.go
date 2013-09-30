package symbols

import (
	"code.google.com/p/gocc/ast"
	"sort"
)

type Symbols struct {
	StartSymbol string

	ntList []string
	ntMap  map[string]ast.SyntaxTerm

	tList []string
	tMap  map[string]ast.SyntaxTerm

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

func NewSymbols(prods []*ast.SyntaxBasicProd) *Symbols {
	symbols := &Symbols{
		StartSymbol: prods[0].Id,
		ntMap:       make(map[string]ast.SyntaxTerm),
		tMap:        make(map[string]ast.SyntaxTerm),
		symMap:      make(map[string]bool),
		prodMap:     make(map[string]int),
	}
	for _, prod := range prods {
		if _, exist := symbols.prodMap[prod.Id]; !exist {
			symbols.ProdList = append(symbols.ProdList, prod.Id)
			symbols.prodMap[prod.Id] = len(symbols.ProdList) - 1
		}
		symbols.addSymbols(prod.Terms)
	}
	symbols.makeLists()
	return symbols
}

func (this *Symbols) makeLists() {
	for sym, _ := range this.ntMap {
		this.ntList = append(this.ntList, sym)
	}
	sort.Strings(this.ntList)

	for sym, _ := range this.tMap {
		this.tList = append(this.tList, sym)
	}
	sort.Strings(this.tList)

	for sym, _ := range this.stringLitMap {
		this.stringLitList = append(this.stringLitList, sym)
	}
	sort.Strings(this.stringLitList)

	this.list = append(this.list, this.tList...)
	this.list = append(this.list, this.ntList...)
}

func (this *Symbols) addSymbols(terms ast.SyntaxTerms) {
	for _, term := range terms {
		this.addSymbol(term)
	}
}

func (this *Symbols) addSymbol(term ast.SyntaxTerm) {
	switch t := term.(type) {
	case ast.SyntaxStringLit:
		this.symMap[string(t)] = true
		this.tMap[string(t)] = term
		this.stringLitMap[string(t)] = t
	case ast.SyntaxTokId:
		this.symMap[string(t)] = true
		this.tMap[string(t)] = term
	case ast.SyntaxProdId:
		this.symMap[string(t)] = true
		this.ntMap[string(t)] = term
	}
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

func (this *Symbols) IsNT(prodId string) bool {
	return this.ntMap[prodId] != nil
}
