//Copyright 2012 Vastech SA (PTY) LTD
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

package token

import (
	"io/ioutil"
	"strconv"
	"strings"
)

//The map of tokens.
type TokenMap struct {
	tokenMap  []string
	stringMap map[string]Type
}

//Creates a new TokenMap with a $ and a ε.
func NewMap() *TokenMap {
	tm := &TokenMap{make([]string, 0, 10), make(map[string]Type)}
	tm.AddToken("$")
	tm.AddToken("ε")
	return tm
}

//Adds a token to the TokenMap.
func (this *TokenMap) AddToken(str string) {
	if _, exists := this.stringMap[str]; exists {
		return
	}
	this.stringMap[str] = Type(len(this.tokenMap))
	this.tokenMap = append(this.tokenMap, str)
}

//Creates a new TokenMap from the contents of a file.
func NewMapFromFile(file string) (*TokenMap, error) {
	src, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return NewMapFromString(string(src)), nil
}

//Creates a new TokenMap from a list of tokens.
func NewMapFromStrings(input []string) *TokenMap {
	tm := NewMap()
	for _, s := range input {
		tm.AddToken(s)
	}
	return tm
}

//Create a new TokenMap by splitting the input string around spaces into multiple tokens.
func NewMapFromString(input string) *TokenMap {
	tokens := strings.Fields(input)
	return NewMapFromStrings(tokens)
}

//Returns the Type of a Token from the TokenMap.
func (this *TokenMap) Type(key string) Type {
	tok, ok := this.stringMap[key]
	if !ok {
		return ILLEGAL
	}
	return tok
}

//Returns the Token given the Type.
func (this *TokenMap) TokenString(typ Type) string {
	tok := int(typ)
	if tok < 0 || tok >= len(this.tokenMap) {
		return "illegal " + strconv.Itoa(tok)
	}
	return this.tokenMap[tok]
}

//Describes the TokenMap as a string.
func (this *TokenMap) String() string {
	res := ""
	for str, tok := range this.stringMap {
		res += str + " : " + strconv.Itoa(int(tok)) + "\n"
	}
	return res
}

//Returns all the tokens except the first one, which is the $.
func (this *TokenMap) Strings() []string {
	return this.tokenMap[1:]
}

//Tests whether two TokenMaps are equal.
func (this *TokenMap) Equals(that *TokenMap) bool {
	if this == nil || that == nil {
		return false
	}

	if len(this.stringMap) != len(that.stringMap) ||
		len(this.tokenMap) != len(that.tokenMap) {
		return false
	}

	for str, tok := range this.stringMap {
		if tok1, ok := that.stringMap[str]; !ok || tok1 != tok {
			return false
		}
	}

	return true
}

//Returns a list of Tokens from the TokenMap.
func (this *TokenMap) Tokens() []*Token {
	res := make([]*Token, 0, len(this.stringMap))
	for typ, str := range this.tokenMap {
		res = append(res, &Token{Type(typ), []byte(str)})
	}
	return res
}

//Rewrites a file with all the tokens.
func (this *TokenMap) WriteFile(file string) error {
	out := ""
	for i := 1; i < len(this.tokenMap); i++ {
		out += this.TokenString(Type(i)) + "\n"
	}
	return ioutil.WriteFile(file, []byte(out), 0644)
}
