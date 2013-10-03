/*
 */
package parser

const numNTSymbols = 4

type (
	gotoTable [numStates]gotoRow
	gotoRow   [numNTSymbols]int
)

var gotoTab = gotoTable{
	gotoRow{ // S0

		-1, // S'
		2,  // A
		3,  // A~1
		4,  // A~1_RepTerm

	},
	gotoRow{ // S1

		-1, // S'
		-1, // A
		-1, // A~1
		-1, // A~1_RepTerm

	},
	gotoRow{ // S2

		-1, // S'
		-1, // A
		-1, // A~1
		-1, // A~1_RepTerm

	},
	gotoRow{ // S3

		-1, // S'
		-1, // A
		-1, // A~1
		5,  // A~1_RepTerm

	},
	gotoRow{ // S4

		-1, // S'
		-1, // A
		-1, // A~1
		-1, // A~1_RepTerm

	},
	gotoRow{ // S5

		-1, // S'
		-1, // A
		-1, // A~1
		-1, // A~1_RepTerm

	},
}
