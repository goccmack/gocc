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
	"fmt"
	"runtime/debug"

	"github.com/maxcalandrelli/gocc/internal/config"
)

type SyntaxBody struct {
	Error   bool
	Symbols SyntaxSymbols
	SDT     string
	_x      debug.GCStats
}

func checkSymbols(symbols interface{}) interface{} {
	if syms, _ := symbols.(SyntaxSymbols); syms != nil {
		for i, s := range syms {
			switch s.(type) {
			case SyntaxEmpty:
				syms[i] = emptySymbol
			case SyntaxError:
				syms[i] = errorSymbol
			case SyntaxTokId:
				switch s.SymbolString() {
				case config.SYMBOL_EMPTY:
					syms[i] = emptySymbol
				case config.SYMBOL_ERROR:
					syms[i] = errorSymbol
				}
			}
		}
		return syms
	}
	return symbols
}

func NewSyntaxBody(symbols, sdtLit interface{}) (*SyntaxBody, error) {
	checkSymbols(symbols)
	if syms, _ := symbols.(SyntaxSymbols); syms != nil {
		for i, s := range syms {
			switch {
			case s.IsEpsilon():
				syms[i] = emptySymbol
			case s.IsError():
				syms[i] = errorSymbol
			}
		}
	}
	syntaxBody := &SyntaxBody{
		Error: false,
	}
	if symbols != nil {
		syntaxBody.Symbols = symbols.(SyntaxSymbols)
	}
	if sdtLit != nil {
		syntaxBody.SDT = SDTVal(getString(sdtLit))
	}
	if symbols == nil {
		return nil, fmt.Errorf("empty production alternative!")
	}
	if s, _ := symbols.(SyntaxSymbols); len(s) > 0 && s[0] == emptySymbol {
		//fmt.Printf("  NewBody(Empty) Symbols=%#v(%s)\n", syntaxBody.Symbols, syntaxBody.Symbols.String())
	} else if s, _ := symbols.(SyntaxSymbols); len(s) > 0 && s[0] == errorSymbol {
		//fmt.Printf("  NewBody(Error) Symbols=%#v(%s)\n", syntaxBody.Symbols, syntaxBody.Symbols.String())
	} else {
		//fmt.Printf("  NewBody()      Symbols=%#v(%s)\n", syntaxBody.Symbols, syntaxBody.Symbols.String())
	}
	return syntaxBody, nil
}

func NewSyntaxBodyGen(symbols, sdtLit interface{}) (*SyntaxBody, error) {
	checkSymbols(symbols)
	syntaxBody := &SyntaxBody{
		Error: false,
	}
	if symbols != nil {
		syntaxBody.Symbols = symbols.(SyntaxSymbols)
	}
	if sdtLit != nil {
		syntaxBody.SDT = SDTVal(getString(sdtLit))
	}
	if symbols == nil {
		return nil, fmt.Errorf("empty production alternative!")
	}
	if s, _ := symbols.(SyntaxSymbols); len(s) > 0 && s[0] == emptySymbol {
		//fmt.Printf("  NewBodyGen(Empty) Symbols=%#v(%s)\n", syntaxBody.Symbols, syntaxBody.Symbols.String())
	} else if s, _ := symbols.(SyntaxSymbols); len(s) > 0 && s[0] == errorSymbol {
		//fmt.Printf("  NewBodyGen(Error) Symbols=%#v(%s)\n", syntaxBody.Symbols, syntaxBody.Symbols.String())
	} else {
		//fmt.Printf("  NewBodyGen()      Symbols=%#v(%s)\n", syntaxBody.Symbols, syntaxBody.Symbols.String())
	}
	return syntaxBody, nil
}

func NewErrorBody(symbols, sdtLit interface{}) (*SyntaxBody, error) {
	if symbols == nil {
		return nil, fmt.Errorf("empty production alternative")
	} else {
		symbols.(SyntaxSymbols)[0] = errorSymbol
	}
	//fmt.Printf("  NewErrorBody()\n")
	if body, err := NewSyntaxBody(symbols, sdtLit); err != nil {
		return nil, err
	} else {
		body.Error = true
		return body, nil
	}
}

func NewErrorBodyGen(symbols, sdtLit interface{}) (*SyntaxBody, error) {
	syms, _ := symbols.(SyntaxSymbols)
	if syms == nil {
		syms = SyntaxSymbols{}
	}
	//fmt.Printf("  NewErrorBodyGen()\n")
	if body, err := NewSyntaxBodyGen(append(SyntaxSymbols{errorSymbol}, syms...), sdtLit); err != nil {
		return nil, err
	} else {
		body.Error = true
		return body, nil
	}
}

func NewEmptyBody() (*SyntaxBody, error) {
	//fmt.Printf("  NewEmptyBody()\n")
	return NewSyntaxBody(SyntaxSymbols{emptySymbol}, nil)
}

func NewEmptyBodyGen() (*SyntaxBody, error) {
	//fmt.Printf("  NewEmptyBodyGen()\n")
	return NewSyntaxBodyGen(SyntaxSymbols{emptySymbol}, nil)
}

func (this *SyntaxBody) Missing() bool {
	return len(this.Symbols) == 0
}

func (this *SyntaxBody) Empty() bool {
	return len(this.Symbols) == 1 && (this.Symbols[0] == emptySymbol)
}

func (this *SyntaxBody) String() string {
	return fmt.Sprintf("%s\t<< %s >>", this.Symbols.String(), this.SDT)
}
