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

package items

import (
	"bytes"
	"code.google.com/p/gocc/ast"
	"fmt"
)

type Items struct {
	List []*Item

	// key: item.haskHey
	idxMap map[string]int

	// key: prod id
	startItems map[string][]*Item
}

func NewItems(prods []*ast.SyntaxBasicProd) *Items {
	items := &Items{
		List:       make([]*Item, 0, 128),
		idxMap:     make(map[string]int),
		startItems: make(map[string][]*Item),
	}
	for i, prod := range prods {
		items.newItems(prod, i)
	}
	return items
}

func (this *Items) StartItems(prodId string) []*Item {
	return this.startItems[prodId]
}

func (this *Items) newItems(prod *ast.SyntaxBasicProd, prodIdx int) {
	items := NewItem(prod, prodIdx)
	this.startItems[prod.Id] = append(this.startItems[prod.Id], items[0])
	for _, item := range items {
		if _, exist := this.idxMap[item.hashKey]; exist {
			panic(fmt.Sprintf("Duplicate item: %s", item))
		}
		this.List = append(this.List, item)
		this.idxMap[item.hashKey] = len(this.List) - 1
	}
}

func (this *Items) String() string {
	w := new(bytes.Buffer)
	for i, item := range this.List {
		fmt.Fprintf(w, "%d: %s\n", i, item)
		if item.NextItem != nil {
			fmt.Fprintf(w, "\tnext: %d\n\n", this.idxMap[item.NextItem.hashKey])
		} else {
			fmt.Fprintf(w, "\n")
		}
	}
	return w.String()
}
