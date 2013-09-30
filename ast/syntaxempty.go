package ast

type SyntaxEmpty int

const EMPTY SyntaxEmpty = 0

func (SyntaxEmpty) String() string {
	return "empty"
}
