package ast

import (
	"code.google.com/p/gocc/test/ebnf/t3/token"
)

type Result struct {
	Opt string
	Rep []string
}

func Test(attrib ...interface{}) (*Result, error) {
	switch len(attrib) {
	case 0:
		return &Result{}, nil
	case 2:
		return &Result{
			Opt: getOpt(attrib[0]),
			Rep: getRep(attrib[1]),
		}, nil
	case 1:
		switch attrib[0].(type) {
		case *token.Token:
			return &Result{
				Opt: getOpt(attrib[0]),
			}, nil
		case []interface{}:
			return &Result{
				Rep: getRep(attrib[0]),
			}, nil
		}
	}
	panic("Should not get here")
}

func getOpt(attrib interface{}) string {
	return string(attrib.(*token.Token).Lit)
}

func getRep(attrib interface{}) []string {
	var rep []string
	for _, tok := range attrib.([]interface{}) {
		rep = append(rep, string(tok.(*token.Token).Lit))
	}
	return rep
}
