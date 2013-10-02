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

package lexer

import (
	"code.google.com/p/gocc/ast"
	"code.google.com/p/gocc/config"
	"code.google.com/p/gocc/io"
	"code.google.com/p/gocc/lexer/gen/golang"
	"code.google.com/p/gocc/lexer/items"
	"code.google.com/p/gocc/token"
	"os"
	"path"
)

func Gen(cfg config.Config, lexPart *ast.LexPart, tokenMap *token.TokenMap) {
	lexSetsFName := path.Join(cfg.OutDir(), "lexer_sets.txt")
	removeFiles(lexSetsFName)
	lexSets := items.GetItemSets(lexPart)
	if cfg.Verbose2() {
		io.WriteFileString(lexSetsFName, lexSets.String())
	}
	golang.Gen(cfg.Package(), cfg.OutDir(), lexPart.Header.String(), lexSets, tokenMap, cfg)

}

func removeFiles(names ...string) {
	for _, name := range names {
		os.Remove(name)
	}
}
