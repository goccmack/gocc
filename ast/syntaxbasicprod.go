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

package ast

import (
	"bytes"
	"fmt"
)

type SyntaxBasicProd struct {
	Prev, Next *SyntaxBasicProd
	Id         string
	Error      bool
	Terms      SyntaxTerms
	Action     string
}

func NewSyntaxBasicProd(id string, body *SyntaxBody) *SyntaxBasicProd {
	return &SyntaxBasicProd{
		Id:     id,
		Error:  body.Error,
		Terms:  body.Terms,
		Action: body.Action,
	}
}

func (this *SyntaxBasicProd) Clone(newTerms SyntaxTerms) (clone *SyntaxBasicProd) {
	clone = &SyntaxBasicProd{
		Id:     this.Id,
		Error:  this.Error,
		Terms:  newTerms,
		Action: this.Action,
	}
	return
}

func (this *SyntaxBasicProd) Equal(that *SyntaxBasicProd) bool {
	if this == nil || that == nil {
		return false
	}
	return this.Id == that.Id &&
		this.Error == that.Error &&
		this.Terms.Equal(that.Terms) &&
		this.Action == that.Action
}

func (this *SyntaxBasicProd) String() string {
	w := new(bytes.Buffer)
	fmt.Fprintf(w, "%s : ", this.Id)
	for _, term := range this.Terms {
		fmt.Fprintf(w, "%s ", term)
	}
	if this.Action != "" {
		fmt.Fprintf(w, "<< %s >> ", this.Action)
	}
	w.WriteString(";")
	return w.String()
}

func (this *SyntaxBasicProd) DotString() string {
	w := new(bytes.Buffer)
	fmt.Fprintf(w, "%s : ", this.Id)
	for _, term := range this.Terms {
		fmt.Fprintf(w, "%s ", term.DotString())
	}
	return w.String()
}

func (this *SyntaxBasicProd) Equivalent(that *SyntaxBasicProd) bool {
	return this.Error == that.Error &&
		this.Terms.Equal(that.Terms) &&
		this.Action == that.Action
}
