package first

import (
	"bytes"
	"code.google.com/p/gocc/ast"
	"code.google.com/p/gocc/parser/symbols"
	"fmt"
)

type First struct {
	// key: symbol
	symbol map[string]FirstSet

	// key: prod id
	ntMap  map[string][]*ast.SyntaxBasicProd
	ntList []string

	tList []string
}

func New(symbols *symbols.Symbols, prods []*ast.SyntaxBasicProd) *First {
	first := &First{
		symbol: make(map[string]FirstSet),
		ntMap:  make(map[string][]*ast.SyntaxBasicProd),
		ntList: symbols.ListNonTerminals(),
		tList:  symbols.ListTerminals(),
	}
	first.AddTerminals()
	first.AddNonTerminals(prods)
	return first
}

func (this *First) AddNonTerminals(prods []*ast.SyntaxBasicProd) {
	for again := true; again; {
		again = false
		for _, prod := range prods {
			changed := false
			switch t := prod.Terms[0].(type) {
			case ast.SyntaxTokId:
				this.symbol[prod.Id], changed = this.symbol[prod.Id].Add(string(t))
				break
			case ast.SyntaxProdId:
				this.symbol[prod.Id], changed = this.symbol[prod.Id].Add(this.symbol[string(t)]...)
			case ast.SyntaxStringLit:
				this.symbol[prod.Id], changed = this.symbol[prod.Id].Add(string(t))
			}
			if changed {
				again = true
			}
		}
	}
}

func (this *First) AddTerminals() {
	for _, s := range this.tList {
		if _, exist := this.symbol[s]; exist {
			panic(fmt.Sprintf("Duplicate terminal: %s", s))
		}
		this.symbol[s] = FirstSet{s}
	}
}

func (this *First) FirstSymbol(sym string) FirstSet {
	return this.symbol[sym]
}

func (this *First) FirstString(s []string, context ...string) FirstSet {
	// If the AST is changed to allow e-productions this must iterate over the string and context
	if len(s) == 0 {
		return append(FirstSet{}, context...)
	}
	return this.symbol[s[0]]
}

func (this *First) String() string {
	w := new(bytes.Buffer)
	fmt.Fprintf(w, "Terminals:\n")
	for _, t := range this.tList {
		fmt.Fprintf(w, "\t%s\t%s\n", t, this.symbol[t])
	}
	fmt.Fprintf(w, "\nNon-terminals:\n")
	for _, nt := range this.ntList {
		fmt.Fprintf(w, "\t%s\t%s\n", nt, this.symbol[nt])
	}
	return w.String()
}
