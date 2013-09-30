package ast

type SyntaxError int

const errorConst = SyntaxError(-1)

func (SyntaxError) String() string {
	return "error"
}
