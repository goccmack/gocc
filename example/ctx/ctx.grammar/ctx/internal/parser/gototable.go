// Code generated by gocc; DO NOT EDIT.

package parser

const numNTSymbols = 3

type (
	gotoTable [numStates]gotoRow
	gotoRow   [numNTSymbols]int
)

var gotoTab = gotoTable{
	gotoRow{ // S0
		-1, // Π<S'>
		1, // Π<StmtList>
		2, // Π<Stmt>
	},
	gotoRow{ // S1
		-1, // Π<S'>
		-1, // Π<StmtList>
		5, // Π<Stmt>
	},
	gotoRow{ // S2
		-1, // Π<S'>
		-1, // Π<StmtList>
		-1, // Π<Stmt>
	},
	gotoRow{ // S3
		-1, // Π<S'>
		-1, // Π<StmtList>
		-1, // Π<Stmt>
	},
	gotoRow{ // S4
		-1, // Π<S'>
		-1, // Π<StmtList>
		-1, // Π<Stmt>
	},
	gotoRow{ // S5
		-1, // Π<S'>
		-1, // Π<StmtList>
		-1, // Π<Stmt>
	},
	gotoRow{ // S6
		-1, // Π<S'>
		-1, // Π<StmtList>
		-1, // Π<Stmt>
	},
}
