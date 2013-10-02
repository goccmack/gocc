package parser

import "code.google.com/p/gocc/ast"

type (
	//TODO: change type and variable names to be consistent with other tables
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index      int
		NumSymbols int
		ReduceFunc func([]Attrib) (Attrib, error)
	}
	Attrib interface {
	}
)

var productionsTable = ProdTab{
	ProdTabEntry{
		String: `S' : Grammar	<<  >>`,
		Id:         "S'",
		NTType:     0,
		Index:      0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Grammar : LexicalPart SyntaxPart	<< ast.NewGrammar(X[0], X[1]) >>`,
		Id:         "Grammar",
		NTType:     1,
		Index:      1,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewGrammar(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `Grammar : LexicalPart	<< ast.NewGrammar(X[0], nil) >>`,
		Id:         "Grammar",
		NTType:     1,
		Index:      2,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewGrammar(X[0], nil)
		},
	},
	ProdTabEntry{
		String: `Grammar : SyntaxPart	<< ast.NewGrammar(nil, X[0]) >>`,
		Id:         "Grammar",
		NTType:     1,
		Index:      3,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewGrammar(nil, X[0])
		},
	},
	ProdTabEntry{
		String: `LexicalPart : LexProductions	<< ast.NewLexPart(nil, nil, X[0]) >>`,
		Id:         "LexicalPart",
		NTType:     2,
		Index:      4,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewLexPart(nil, nil, X[0])
		},
	},
	ProdTabEntry{
		String: `LexProductions : LexProduction	<< ast.NewLexProductions(X[0]) >>`,
		Id:         "LexProductions",
		NTType:     3,
		Index:      5,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewLexProductions(X[0])
		},
	},
	ProdTabEntry{
		String: `LexProductions : LexProductions LexProduction	<< ast.AppendLexProduction(X[0], X[1]) >>`,
		Id:         "LexProductions",
		NTType:     3,
		Index:      6,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AppendLexProduction(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `LexProduction : tokId ":" LexPattern ";"	<< ast.NewLexTokDef(X[0], X[2]) >>`,
		Id:         "LexProduction",
		NTType:     4,
		Index:      7,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewLexTokDef(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `LexProduction : regDefId ":" LexPattern ";"	<< ast.NewLexRegDef(X[0], X[2]) >>`,
		Id:         "LexProduction",
		NTType:     4,
		Index:      8,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewLexRegDef(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `LexProduction : ignoredTokId ":" LexPattern ";"	<< ast.NewLexIgnoredTokDef(X[0], X[2]) >>`,
		Id:         "LexProduction",
		NTType:     4,
		Index:      9,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewLexIgnoredTokDef(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `LexPattern : LexAlt	<< ast.NewLexPattern(X[0]) >>`,
		Id:         "LexPattern",
		NTType:     5,
		Index:      10,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewLexPattern(X[0])
		},
	},
	ProdTabEntry{
		String: `LexPattern : LexPattern "|" LexAlt	<< ast.AppendLexAlt(X[0], X[2]) >>`,
		Id:         "LexPattern",
		NTType:     5,
		Index:      11,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AppendLexAlt(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `LexAlt : LexTerm	<< ast.NewLexAlt(X[0]) >>`,
		Id:         "LexAlt",
		NTType:     6,
		Index:      12,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewLexAlt(X[0])
		},
	},
	ProdTabEntry{
		String: `LexAlt : LexAlt LexTerm	<< ast.AppendLexTerm(X[0], X[1]) >>`,
		Id:         "LexAlt",
		NTType:     6,
		Index:      13,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AppendLexTerm(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `LexTerm : "."	<< ast.LexDOT, nil >>`,
		Id:         "LexTerm",
		NTType:     7,
		Index:      14,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.LexDOT, nil
		},
	},
	ProdTabEntry{
		String: `LexTerm : char_lit	<< ast.NewLexCharLit(X[0]) >>`,
		Id:         "LexTerm",
		NTType:     7,
		Index:      15,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewLexCharLit(X[0])
		},
	},
	ProdTabEntry{
		String: `LexTerm : char_lit "-" char_lit	<< ast.NewLexCharRange(X[0], X[2]) >>`,
		Id:         "LexTerm",
		NTType:     7,
		Index:      16,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewLexCharRange(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `LexTerm : regDefId	<< ast.NewLexRegDefId(X[0]) >>`,
		Id:         "LexTerm",
		NTType:     7,
		Index:      17,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewLexRegDefId(X[0])
		},
	},
	ProdTabEntry{
		String: `LexTerm : "[" LexPattern "]"	<< ast.NewLexOptPattern(X[1]) >>`,
		Id:         "LexTerm",
		NTType:     7,
		Index:      18,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewLexOptPattern(X[1])
		},
	},
	ProdTabEntry{
		String: `LexTerm : "{" LexPattern "}"	<< ast.NewLexRepPattern(X[1]) >>`,
		Id:         "LexTerm",
		NTType:     7,
		Index:      19,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewLexRepPattern(X[1])
		},
	},
	ProdTabEntry{
		String: `LexTerm : "(" LexPattern ")"	<< ast.NewLexGroupPattern(X[1]) >>`,
		Id:         "LexTerm",
		NTType:     7,
		Index:      20,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewLexGroupPattern(X[1])
		},
	},
	ProdTabEntry{
		String: `SyntaxPart : FileHeader SyntaxProdList	<< ast.NewSyntaxPart(X[0], X[1]) >>`,
		Id:         "SyntaxPart",
		NTType:     8,
		Index:      21,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewSyntaxPart(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `SyntaxPart : SyntaxProdList	<< ast.NewSyntaxPart(nil, X[0]) >>`,
		Id:         "SyntaxPart",
		NTType:     8,
		Index:      22,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewSyntaxPart(nil, X[0])
		},
	},
	ProdTabEntry{
		String: `FileHeader : actionLit	<< ast.NewFileHeader(X[0]) >>`,
		Id:         "FileHeader",
		NTType:     9,
		Index:      23,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewFileHeader(X[0])
		},
	},
	ProdTabEntry{
		String: `SyntaxProdList : SyntaxProduction	<< ast.NewSyntaxProdList(X[0]) >>`,
		Id:         "SyntaxProdList",
		NTType:     10,
		Index:      24,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewSyntaxProdList(X[0])
		},
	},
	ProdTabEntry{
		String: `SyntaxProdList : SyntaxProdList SyntaxProduction	<< ast.AddSyntaxProds(X[0], X[1]) >>`,
		Id:         "SyntaxProdList",
		NTType:     10,
		Index:      25,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AddSyntaxProds(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `SyntaxProduction : prodId ":" SyntaxExpression ";"	<< ast.NewSyntaxProdNonBasic(X[0], X[2]) >>`,
		Id:         "SyntaxProduction",
		NTType:     11,
		Index:      26,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewSyntaxProdNonBasic(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `SyntaxExpression : SyntaxBody	<< ast.NewSyntaxExpression(X[0]) >>`,
		Id:         "SyntaxExpression",
		NTType:     12,
		Index:      27,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewSyntaxExpression(X[0])
		},
	},
	ProdTabEntry{
		String: `SyntaxExpression : SyntaxExpression "|" SyntaxBody	<< ast.AddSyntaxExprBody(X[0], X[2]) >>`,
		Id:         "SyntaxExpression",
		NTType:     12,
		Index:      28,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AddSyntaxExprBody(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `SyntaxBody : SyntaxTerms	<< ast.NewSyntaxBody(X[0], nil) >>`,
		Id:         "SyntaxBody",
		NTType:     13,
		Index:      29,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewSyntaxBody(X[0], nil)
		},
	},
	ProdTabEntry{
		String: `SyntaxBody : SyntaxTerms actionLit	<< ast.NewSyntaxBody(X[0], X[1]) >>`,
		Id:         "SyntaxBody",
		NTType:     13,
		Index:      30,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewSyntaxBody(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `SyntaxBody : "error"	<< ast.NewErrorBody(nil, nil) >>`,
		Id:         "SyntaxBody",
		NTType:     13,
		Index:      31,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewErrorBody(nil, nil)
		},
	},
	ProdTabEntry{
		String: `SyntaxBody : "error" SyntaxTerms	<< ast.NewErrorBody(X[1], nil) >>`,
		Id:         "SyntaxBody",
		NTType:     13,
		Index:      32,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewErrorBody(X[1], nil)
		},
	},
	ProdTabEntry{
		String: `SyntaxBody : "error" SyntaxTerms actionLit	<< ast.NewErrorBody(X[1], X[2]) >>`,
		Id:         "SyntaxBody",
		NTType:     13,
		Index:      33,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewErrorBody(X[1], X[2])
		},
	},
	ProdTabEntry{
		String: `SyntaxTerms : SyntaxTerm	<< ast.NewSyntaxTerms(X[0]) >>`,
		Id:         "SyntaxTerms",
		NTType:     14,
		Index:      34,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewSyntaxTerms(X[0])
		},
	},
	ProdTabEntry{
		String: `SyntaxTerms : SyntaxTerms SyntaxTerm	<< ast.AddSyntaxTerm(X[0], X[1]) >>`,
		Id:         "SyntaxTerms",
		NTType:     14,
		Index:      35,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AddSyntaxTerm(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `SyntaxTerm : prodId	<< ast.NewSyntaxProdId(X[0]) >>`,
		Id:         "SyntaxTerm",
		NTType:     15,
		Index:      36,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewSyntaxProdId(X[0])
		},
	},
	ProdTabEntry{
		String: `SyntaxTerm : tokId	<< ast.NewTokId(X[0]) >>`,
		Id:         "SyntaxTerm",
		NTType:     15,
		Index:      37,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewTokId(X[0])
		},
	},
	ProdTabEntry{
		String: `SyntaxTerm : stringLit	<< ast.NewStringLit(X[0]) >>`,
		Id:         "SyntaxTerm",
		NTType:     15,
		Index:      38,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewStringLit(X[0])
		},
	},
	ProdTabEntry{
		String: `SyntaxTerm : "(" SyntaxExpression ")"	<< ast.NewSyntaxGroupExpression(X[1]) >>`,
		Id:         "SyntaxTerm",
		NTType:     15,
		Index:      39,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewSyntaxGroupExpression(X[1])
		},
	},
	ProdTabEntry{
		String: `SyntaxTerm : "[" SyntaxExpression "]"	<< ast.NewSyntaxOptionalExpression(X[1]) >>`,
		Id:         "SyntaxTerm",
		NTType:     15,
		Index:      40,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewSyntaxOptionalExpression(X[1])
		},
	},
	ProdTabEntry{
		String: `SyntaxTerm : "{" SyntaxExpression "}"	<< ast.NewSyntaxRepeatedExpression(X[1]) >>`,
		Id:         "SyntaxTerm",
		NTType:     15,
		Index:      41,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewSyntaxRepeatedExpression(X[1])
		},
	},
}
