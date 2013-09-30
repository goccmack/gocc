package ast

import (
	"fmt"
)

type LexProdMap struct {
	//key: production id
	idMap map[string]LexProdIndex

	//value: lex production id
	idxMap map[LexProdIndex]string
}

type LexProdIndex int

func NewLexProdMap(prodList *LexProductions) *LexProdMap {
	lpm := &LexProdMap{
		idMap:  make(map[string]LexProdIndex),
		idxMap: make(map[LexProdIndex]string),
	}
	lpm.Add(prodList.Productions...)

	return lpm
}

func newLexProdMap() *LexProdMap {
	return &LexProdMap{
		idMap:  make(map[string]LexProdIndex),
		idxMap: make(map[LexProdIndex]string),
	}
}

func (this *LexProdMap) Index(id string) LexProdIndex {
	idx, exist := this.idMap[id]
	if exist {
		return idx
	}
	return -1
}

func (this *LexProdMap) Id(index LexProdIndex) string {
	id, exist := this.idxMap[index]
	if exist {
		return id
	}
	return ""
}

func (this *LexProdMap) Add(prods ...LexProduction) {
	for _, prod := range prods {
		if _, exist := this.idMap[prod.Id()]; exist {
			panic(fmt.Sprintf("Production %s already exists", prod.Id()))
		}
		idx := len(this.idxMap)
		this.idMap[prod.Id()] = LexProdIndex(idx)
		this.idxMap[LexProdIndex(idx)] = prod.Id()
	}
}
