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

package golang

import (
	"github.com/maxcalandrelli/gocc/internal/lexer/symbols"
)

type importType struct {
	Id      string
	ExtFunc string
	Type    int
}

func getImports(symbols *symbols.Symbols) (imports []*importType) {
	imports = make([]*importType, 0, len(symbols.ImportIdList))
	for _, id := range symbols.ImportIdList {
		impType := &importType{
			Id:      id,
			ExtFunc: symbols.ExternalFunction(id),
			Type:    symbols.Type(id),
		}
		imports = append(imports, impType)
	}
	return
}
