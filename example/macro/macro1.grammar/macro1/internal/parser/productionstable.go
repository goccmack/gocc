// Code generated by gocc; DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/maxcalandrelli/gocc/example/macro/macro1.grammar/macro1/internal/token"
	"github.com/maxcalandrelli/gocc/example/macro/macro1.grammar/macro1/internal/util"
	"strings"
)

func first_one(x string) []string {
	fmt.Printf("start with this this one: %s\n", x)
	return []string{x}
}
func add_one(c []string, x string) []string {
	fmt.Printf("add to %q that %s\n", c, x)
	return append([]string{x}, c...)
}
func summary(v interface{}) string {
	return fmt.Sprintf("ok, so we got <%q> as a final result.\n", v)
}
func definition(deftype, defname string, defval int64) (interface{}, error) {
	fmt.Printf("Definition of variable <%s> with '%s' as %d\n", defname, deftype, defval)
	return fmt.Sprintf("%s <- %d", defname, defval), nil
}

func getString(X Attrib) string {
	switch X.(type) {
	case *token.Token:
		return string(X.(*token.Token).Lit)
	case string:
		return X.(string)
	}
	return fmt.Sprintf("%q", X)
}

func unescape(s string) string {
	return util.EscapedString(s).Unescape()
}

func unquote(s string) string {
	r, _, _ := util.EscapedString(s).Unquote()
	return r
}

func lc(s string) string {
	return strings.ToLower(s)
}

func uc(s string) string {
	return strings.ToUpper(s)
}

type (
	//TODO: change type and variable names to be consistent with other tables
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index      int
		NumSymbols int
		ReduceFunc func(interface{}, []Attrib) (Attrib, error)
	}
	Attrib interface {
	}
)

var productionsTable = ProdTab{
	ProdTabEntry{
		String: `S' : Π<StmtList>	<<  >>`,
		Id:         "S'",
		NTType:     0,
		Index:      0,
		NumSymbols: 1,
		ReduceFunc: func(Context interface{}, X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `StmtList : Π<Stmt_s>	<< func() (interface{}, error) {fmt.Printf("%s", summary($0)); return 42, nil}() >>`,
		Id:         "StmtList",
		NTType:     1,
		Index:      1,
		NumSymbols: 1,
		ReduceFunc: func(Context interface{}, X []Attrib) (Attrib, error) {
			return func() (interface{}, error) { fmt.Printf("%s", summary(X[0])); return 42, nil }()
		},
	},
	ProdTabEntry{
		String: `Stmt_s : Π<Stmt>	<< first_one($0.(string)),nil >>`,
		Id:         "Stmt_s",
		NTType:     2,
		Index:      2,
		NumSymbols: 1,
		ReduceFunc: func(Context interface{}, X []Attrib) (Attrib, error) {
			return first_one(X[0].(string)), nil
		},
	},
	ProdTabEntry{
		String: `Stmt_s : Π<Stmt_s> Λ<;> Π<Stmt>	<< add_one($0.([]string),$2.(string)),nil >>`,
		Id:         "Stmt_s",
		NTType:     2,
		Index:      3,
		NumSymbols: 3,
		ReduceFunc: func(Context interface{}, X []Attrib) (Attrib, error) {
			return add_one(X[0].([]string), X[2].(string)), nil
		},
	},
	ProdTabEntry{
		String: `StmtPref : Λ<let>	<<  >>`,
		Id:         "StmtPref",
		NTType:     4,
		Index:      4,
		NumSymbols: 1,
		ReduceFunc: func(Context interface{}, X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `StmtPref : Λ<define>	<<  >>`,
		Id:         "StmtPref",
		NTType:     4,
		Index:      5,
		NumSymbols: 1,
		ReduceFunc: func(Context interface{}, X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Stmt : Π<StmtPref> id Λ<=> μ<calc_0>	<< definition($0s,$1s,$4.(int64)) >>`,
		Id:         "Stmt",
		NTType:     3,
		Index:      6,
		NumSymbols: 5,
		ReduceFunc: func(Context interface{}, X []Attrib) (Attrib, error) {
			return definition(getString(X[0]), getString(X[1]), X[4].(int64))
		},
	},
}
