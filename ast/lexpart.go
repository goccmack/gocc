package ast

import (
	"bytes"
	"errors"
	"fmt"
	"sort"
)

// All maps are indexed by production id
type LexPart struct {
	Header *FileHeader
	*LexImports
	TokDefs            map[string]*LexTokDef
	tokDefsList        []*LexTokDef
	RegDefs            map[string]*LexRegDef
	regDefsList        []*LexRegDef
	IgnoredTokDefs     map[string]*LexIgnoredTokDef
	ignoredTokDefsList []*LexIgnoredTokDef
	ProdList           *LexProductions
	ProdMap            *LexProdMap
	stringLitTokDefs   map[string]bool
}

func NewLexPart(header, imports, prodList interface{}) (*LexPart, error) {
	lexPart := &LexPart{
		TokDefs:          make(map[string]*LexTokDef, 16),
		RegDefs:          make(map[string]*LexRegDef, 16),
		IgnoredTokDefs:   make(map[string]*LexIgnoredTokDef, 16),
		stringLitTokDefs: make(map[string]bool),
	}
	if header != nil {
		lexPart.Header = header.(*FileHeader)
	} else {
		lexPart.Header = new(FileHeader)
	}
	if imports != nil {
		lexPart.LexImports = imports.(*LexImports)
	} else {
		lexPart.LexImports = newLexImports()
	}
	if prodList != nil {
		lexPart.ProdList = prodList.(*LexProductions)
		lexPart.ProdMap = NewLexProdMap(prodList.(*LexProductions))
		for _, p := range prodList.(*LexProductions).Productions {
			pid := p.Id()

			switch p1 := p.(type) {
			case *LexTokDef:
				//TODO: decide whether to handle in separate symantic check
				if _, exist := lexPart.TokDefs[pid]; exist {
					return nil, errors.New(fmt.Sprintf("Duplicate token def: %s", pid))
				}
				lexPart.TokDefs[pid] = p1
			case *LexRegDef:
				//TODO: decide whether to handle in separate symantic check
				if _, exist := lexPart.RegDefs[pid]; exist {
					return nil, errors.New(fmt.Sprintf("Duplicate token def: %s", pid))
				}
				lexPart.RegDefs[pid] = p1
			case *LexIgnoredTokDef:
				if _, exist := lexPart.IgnoredTokDefs[pid]; exist {
					return nil, errors.New(fmt.Sprintf("Duplicate ignored token def: %s", pid))
				}
				lexPart.IgnoredTokDefs[pid] = p1
			}
		}
	} else {
		lexPart.ProdList = newLexProductions()
		lexPart.ProdMap = newLexProdMap()
	}
	lexPart.makeLists()
	return lexPart, nil
}

func (this *LexPart) IgnoredTokDefsList() []*LexIgnoredTokDef {
	return this.ignoredTokDefsList
}

func (this *LexPart) Production(id string) LexProduction {
	idx, exist := this.ProdMap.idMap[id]
	if !exist {
		panic(fmt.Sprintf("Unknown production: %s", id))
	}
	return this.ProdList.Productions[idx]
}

func (this *LexPart) ProdIndex(id string) LexProdIndex {
	idx, exist := this.ProdMap.idMap[id]
	if !exist {
		panic(fmt.Sprintf("Unknown production %s", id))
	}
	return idx
}

func (this *LexPart) RegDefsList() []*LexRegDef {
	return this.regDefsList
}

/*
Return true iff sym is a string literal token defined in the syntax part
*/
func (this *LexPart) StringLitTokDef(sym string) bool {
	return this.stringLitTokDefs[sym]
}

func (this *LexPart) TokDefsList() []*LexTokDef {
	return this.tokDefsList
}

func (this *LexPart) TokenIds() []string {
	tids := make([]string, 0, len(this.TokDefs))
	for tid, _ := range this.TokDefs {
		tids = append(tids, tid)
	}
	return tids
}

func (this *LexPart) UpdateStringLitTokens(tokens []string) {
	for _, strLit := range tokens {
		tokDef := NewLexStringLitTokDef(strLit)
		this.ProdMap.Add(tokDef)
		this.TokDefs[strLit] = tokDef
		this.ProdList.Productions = append(this.ProdList.Productions, tokDef)
		this.stringLitTokDefs[strLit] = true
	}
}

func (this *LexPart) String() string {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "LexPart.ProdMap:\n")
	if this.ProdList != nil {
		for i, p := range this.ProdList.Productions {
			fmt.Fprintf(buf, "\t%d: %s\n", i, p)
		}
	}
	fmt.Fprintf(buf, "LexPart.TokDefs:\n")
	for t := range this.TokDefs {
		fmt.Fprintf(buf, "\t%s\n", t)
	}
	return buf.String()
	//TODO: add RegDefs and TokDefs
}

/*** Utility functions ***/

func (this *LexPart) makeLists() {
	tokdefs := []string{}
	for sym, _ := range this.TokDefs {
		tokdefs = append(tokdefs, sym)
	}
	sort.Strings(tokdefs)
	for _, sym := range tokdefs {
		this.tokDefsList = append(this.tokDefsList, this.TokDefs[sym])
	}

	regDefs := []string{}
	for sym, _ := range this.RegDefs {
		regDefs = append(regDefs, sym)
	}
	sort.Strings(regDefs)
	for _, sym := range regDefs {
		this.regDefsList = append(this.regDefsList, this.RegDefs[sym])
	}

	ignoredTokDefs := []string{}
	for sym, _ := range this.IgnoredTokDefs {
		ignoredTokDefs = append(ignoredTokDefs, sym)
	}
	sort.Strings(ignoredTokDefs)
	for _, sym := range ignoredTokDefs {
		this.ignoredTokDefsList = append(this.ignoredTokDefsList, this.IgnoredTokDefs[sym])
	}
}
