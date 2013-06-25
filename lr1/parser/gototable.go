package parser

type (
	GotoTab []GotoRow
	GotoRow map[NT]State
	State   int
	NT      string
)
