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
	"errors"
	"fmt"
	"os"
)

type Grammar struct {
	*LexPart
	*SyntaxPart
}

func NewGrammar(lexPart, syntaxPart interface{}) (*Grammar, error) {
	g := &Grammar{}
	if lexPart != nil {
		g.LexPart = lexPart.(*LexPart)
	} else {
		lp, _ := NewLexPart(nil, nil, nil)
		g.LexPart = lp
	}
	if syntaxPart != nil {
		g.SyntaxPart = syntaxPart.(*SyntaxPart).augment()
	}

	return g, consistent(g)
}

var (
	errUndefined = errors.New("undefined symbol used in production")
)

func consistent(g *Grammar) (err error) {
	if g.SyntaxPart == nil {
		return
	}

	defs := make(map[SyntaxSymbol]struct{})
	used := make(map[SyntaxSymbol]SyntaxSymbols)

	for _, tok := range g.LexPart.TokDefsList {
		defs[SyntaxTokId{tok.id, StdSyntaxSymbol{}}] = struct{}{}
	}
	for _, prod := range g.SyntaxPart.ProdList {
		if prod.Body.Missing() {
			return fmt.Errorf("empty production alternative: Maybe you are missing the \"empty\" keyword in %q", prod)
		}
		defs[prod.Id] = struct{}{}
		for _, s := range prod.Body.Symbols {
			if s.String()[0] == '"' {
				continue
			}
			used[s] = append(used[s], prod.Id)
		}
	}
	for s := range defs {
		if s.SymbolString() == "S'" {
			continue
		}
		if _, ok := used[s]; !ok {
			fmt.Fprintf(os.Stderr, "warning: symbol %q defined but never used\n", s)
		}
	}
	for s, in := range used {
		if _, ok := defs[s]; !ok {
			switch s.(type) {
			case SyntaxEmpty, SyntaxError, SyntaxContextDependentTokId, SyntaxSubParser:
				continue
			}
			if !s.IsTerminal() {
				fmt.Fprintf(os.Stderr, "error: undefined symbol %T{%q} used in productions %q\n", s, s, in)
				err = errUndefined
			} else {
				fmt.Fprintf(os.Stderr, "warning: undefined symbol %T{%q} used in productions %q\n", s, s, in)
			}
		}
	}
	return
}
