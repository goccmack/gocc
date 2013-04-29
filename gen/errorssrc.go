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

package gen

const errorsSrcHeader = `
package errors

`

const errorsSrcBody = `
type ErrorSymbol interface {
}

type Error struct {
	Err	error
	ErrorToken	*token.Token
	ErrorPos	token.Position
	ErrorSymbols	[]ErrorSymbol
	ExpectedTokens	[]string
}

func (E *Error) String() string {
	errmsg := "Got " + E.ErrorToken.String() + " @ " + E.ErrorPos.String()
	if E.Err != nil {
		errmsg += " " + E.Err.Error()
	} else {
		errmsg += ", expected one of: "
		for _, t := range E.ExpectedTokens {
			errmsg += t + " "
		}
	}
	return errmsg
}

`
