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

package rewrite

import (
	"bytes"
	"code.google.com/p/gocc/ast"
	"fmt"
	"text/template"
)

type augId map[string]int

func (this augId) Next(prodId string) string {
	if num, exist := this[prodId]; exist {
		this[prodId] = num + 1
	} else {
		this[prodId] = 1
	}
	return fmt.Sprintf("%s~%d", prodId, this[prodId])
}

type rewrite struct {
	augId
	repProds map[string]*ast.SyntaxBody
	prodMap  map[string]basicProdSet
	bprods   *ast.SyntaxBasicProdsList
}

func BasicProds(prods ast.SyntaxProdList) *ast.SyntaxBasicProdsList {
	rw := &rewrite{
		augId:    make(augId),
		repProds: make(map[string]*ast.SyntaxBody),
		prodMap:  make(map[string]basicProdSet),
		bprods:   ast.NewSyntaxBasicProdsList(),
	}
	for _, prod := range prods {
		for _, body := range prod.SyntaxExpression {
			newProd := ast.NewSyntaxBasicProd(string(prod.Id.Lit), body)
			rw.bprods.Append(newProd)
			rw.prodMap[newProd.Id].add(newProd)
		}
	}
	for prod := rw.bprods.First(); prod != nil; {
		if newProd, changed := rw.rewriteBasicProd(prod); !changed {
			prod = prod.Next
		} else {
			prod = newProd
		}
	}
	return rw.bprods
}

func (this *rewrite) getRepProdId(body *ast.SyntaxBody) (podId string, exist bool) {
	for pid, bod := range this.repProds {
		if bod.Equal(body) {
			return pid, true
		}
	}
	return "", false
}

func (this *rewrite) rewriteBasicProd(prod *ast.SyntaxBasicProd) (newProd *ast.SyntaxBasicProd, changed bool) {
	for ti, term := range prod.Terms {
		switch t := term.(type) {
		case ast.SyntaxError, ast.SyntaxProdId, ast.SyntaxTokId, ast.SyntaxStringLit:
			// do nothing
		case ast.SyntaxGroupExpression:
			return this.rewriteGroupExpr2(prod, ti, t), true
		case ast.SyntaxOptionalExpression:
			return this.rewriteOptExpr2(prod, ti, t), true
		case ast.SyntaxRepeatedExpression:
			return this.rewriteRepExpr1(prod, ti, t), true
		default:
			panic(fmt.Sprint("Unexpected type of term: %T", term))
		}
	}
	return prod, false
}

func (this *rewrite) rewriteGroupExpr1(prod *ast.SyntaxBasicProd, termIdx int, term ast.SyntaxGroupExpression) (firstNewProd *ast.SyntaxBasicProd) {
	newProds := ast.NewSyntaxBasicProdsList()
	for _, body := range term {
		if body.Error || body.Action != "" {
			newProd := ast.NewSyntaxBasicProd(this.augId.Next(prod.Id), body)
			prod1 := prod.Clone(prod.Terms.CloneReplaceTerm(termIdx, ast.SyntaxProdId(newProd.Id)))
			newProds.Append(prod1, newProd)
		} else {
			prod1 := prod.Clone(prod.Terms.CloneReplace(termIdx, body.Terms))
			newProds.Append(prod1)
		}
	}

	this.bprods.Replace(prod, newProds)

	return newProds.First()
}

func (this *rewrite) rewriteGroupExpr2(prod *ast.SyntaxBasicProd, termIdx int, term ast.SyntaxGroupExpression) (firstNewProd *ast.SyntaxBasicProd) {
	newProds := ast.NewSyntaxBasicProdsList()
	newProdId := this.augId.Next(prod.Id)
	newProds.Append(prod.Clone(prod.Terms.CloneReplaceTerm(termIdx, ast.SyntaxProdId(newProdId))))
	for _, body := range term {
		newProds.Append(ast.NewSyntaxBasicProd(newProdId, body))
	}

	this.bprods.Replace(prod, newProds)

	return newProds.First()
}

func (this *rewrite) rewriteOptExpr1(prod *ast.SyntaxBasicProd, termIdx int, term ast.SyntaxOptionalExpression) (newProd *ast.SyntaxBasicProd) {
	newProds := ast.NewSyntaxBasicProdsList()
	newProds.Append(prod.Clone(prod.Terms.CloneDeleteTerm(termIdx)))
	for _, body := range term {
		if body.Error || body.Action != "" {
			newProd := ast.NewSyntaxBasicProd(this.augId.Next(prod.Id), body)
			prod1 := prod.Clone(prod.Terms.CloneReplaceTerm(termIdx, ast.SyntaxProdId(newProd.Id)))
			newProds.Append(prod1, newProd)
		} else {
			prod1 := prod.Clone(prod.Terms.CloneReplace(termIdx, body.Terms))
			newProds.Append(prod1)
		}
	}

	this.bprods.Replace(prod, newProds)

	return newProds.First()
}

func (this *rewrite) rewriteOptExpr2(prod *ast.SyntaxBasicProd, termIdx int, term ast.SyntaxOptionalExpression) (newProd *ast.SyntaxBasicProd) {
	newProds := ast.NewSyntaxBasicProdsList()
	newProds.Append(prod.Clone(prod.Terms.CloneDeleteTerm(termIdx)))
	newProdId := this.augId.Next(prod.Id)
	newProds.Append(prod.Clone(prod.Terms.CloneReplaceTerm(termIdx, ast.SyntaxProdId(newProdId))))
	for _, body := range term {
		newProds.Append(ast.NewSyntaxBasicProd(newProdId, body))
	}

	this.bprods.Replace(prod, newProds)

	return newProds.First()
}

func (this *rewrite) rewriteRepExpr1(prod *ast.SyntaxBasicProd, termIdx int, term ast.SyntaxRepeatedExpression) (newProd *ast.SyntaxBasicProd) {
	newProds := ast.NewSyntaxBasicProdsList()
	newProds.Append(prod.Clone(prod.Terms.CloneDeleteTerm(termIdx)))
	for _, body := range term {
		repProdId, repProds := this.getRepProds(prod, body)
		newProds.Append(prod.Clone(prod.Terms.CloneReplaceTerm(termIdx, ast.SyntaxProdId(repProdId))))
		if len(repProds) > 0 {
			newProds.Append(repProds...)
		}
	}

	this.bprods.Replace(prod, newProds)

	return newProds.First()
}

func (this *rewrite) getRepProds(parent *ast.SyntaxBasicProd, body *ast.SyntaxBody) (newProdId string, newProds []*ast.SyntaxBasicProd) {
	newProdId, exist := this.getRepProdId(body)
	if !exist {
		newProdId = this.augId.Next(parent.Id)
		this.repProds[newProdId] = body
		newProd1 := &ast.SyntaxBasicProd{Id: newProdId, Error: body.Error, Terms: body.Terms, Action: repAction1(body.Action)}
		newProd2 := &ast.SyntaxBasicProd{Id: newProdId, Error: body.Error, Terms: body.Terms.CloneInsertTerm(0, ast.SyntaxProdId(newProdId)),
			Action: repAction2(body.Action)}
		newProds = append(newProds, newProd1, newProd2)
	}
	return
}

func repAction1(action string) string {
	if action == "" {
		return "[]interface{}{X[0]}, nil"
	} else {
		w := new(bytes.Buffer)
		tmpl, err := template.New("repAction1").Parse(repAction1Src)
		if err != nil {
			panic(err)
		}
		tmpl.Execute(w, action)
		return w.String()
	}
}

const repAction1Src = `func(X []interface{}) (interface{}, error) {
	var res interface{}
	var err error
	if res, err = {{.}}; err != nil {
		return nil, err
	} else {
		return []interface{}{res}, nil
	}
}(X)`

func repAction2(action string) string {
	if action == "" {
		return "append(X[0].([]interface{}), X[1]), nil"
	} else {
		w := new(bytes.Buffer)
		tmpl, err := template.New("repAction2").Parse(repAction2Src)
		if err != nil {
			panic(err)
		}
		tmpl.Execute(w, action)
		return w.String()
	}
}

const repAction2Src = `func(X []interface{}) (interface{}, error) {
	var res interface{}
	var err error
	if res, err = {{.}}; err != nil {
		return nil, err
	} else {
		return append(X[0].([]interface{}), res), nil
	}
}(X)`
