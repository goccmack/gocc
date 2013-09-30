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
