package ast

type LexProduction interface {
	LexNode
	Id() string
	LexPattern() *LexPattern
	RegDef() bool
}
