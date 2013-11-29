/*
 */
package parser

const numNTSymbols = 2

type (
	gotoTable [numStates]gotoRow
	gotoRow   [numNTSymbols]int
)

var gotoTab = gotoTable{
	gotoRow{ // S0

		-1, // S'
		3,  // Stmt

	},
	gotoRow{ // S1

		-1, // S'
		-1, // Stmt

	},
	gotoRow{ // S2

		-1, // S'
		-1, // Stmt

	},
	gotoRow{ // S3

		-1, // S'
		-1, // Stmt

	},
	gotoRow{ // S4

		-1, // S'
		-1, // Stmt

	},
	gotoRow{ // S5

		-1, // S'
		6,  // Stmt

	},
	gotoRow{ // S6

		-1, // S'
		-1, // Stmt

	},
	gotoRow{ // S7

		-1, // S'
		8,  // Stmt

	},
	gotoRow{ // S8

		-1, // S'
		-1, // Stmt

	},
}
