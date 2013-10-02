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

/*
All syntax symbols are types of string
*/
type SyntaxTerm interface {
	String() string
	Basic() bool
}

func (SyntaxEmpty) Basic() bool              { return true }
func (SyntaxEof) Basic() bool                { return true }
func (SyntaxError) Basic() bool              { return true }
func (SyntaxProdId) Basic() bool             { return true }
func (SyntaxTokId) Basic() bool              { return true }
func (SyntaxStringLit) Basic() bool          { return true }
func (SyntaxOptionalExpression) Basic() bool { return false }
func (SyntaxGroupExpression) Basic() bool    { return false }
func (SyntaxRepeatedExpression) Basic() bool { return false }
