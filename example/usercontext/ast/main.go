package ast

import "github.com/goccmack/gocc/example/usercontext/token"

type Attrib interface{}

type Context struct {
	Path     string
	Filename string
}

type Import struct {
	Token   *token.Token
	Context *Context
}

type ImportList []*Import

func NewImportList(importAttr Attrib) (ImportList, error) {
	list := make(ImportList, 0)
	return AppendImportToList(list, importAttr)
}

func AppendImportToList(listAttr, importAttr Attrib) (ImportList, error) {
	return append(listAttr.(ImportList), importAttr.(*Import)), nil
}

// Invoked by << ast.NewImport($0, $Context) >>
func NewImport(nameAttr Attrib, context interface{}) (*Import, error) {
	nameTok := nameAttr.(*token.Token)
	return &Import{Token: nameTok, Context: context.(*Context)}, nil
}
