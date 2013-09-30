package ast

type LexNodeVisitor interface {
	Visit(LexNode) LexNodeVisitor
}
