/*
 */
package parser

const numNTSymbols = 8

type (
	gotoTable [numStates]gotoRow
	gotoRow   [numNTSymbols]int
)

var gotoTab = gotoTable{
	gotoRow{ // S0

		-1, // S'
		3,  // A
		4,  // A~2
		5,  // A~2_RepTerm
		6,  // A~1
		7,  // A~4
		-1, // A~3
		-1, // A~3_RepTerm

	},
	gotoRow{ // S1

		-1, // S'
		-1, // A
		-1, // A~2
		-1, // A~2_RepTerm
		-1, // A~1
		-1, // A~4
		-1, // A~3
		-1, // A~3_RepTerm

	},
	gotoRow{ // S2

		-1, // S'
		-1, // A
		-1, // A~2
		-1, // A~2_RepTerm
		-1, // A~1
		-1, // A~4
		-1, // A~3
		-1, // A~3_RepTerm

	},
	gotoRow{ // S3

		-1, // S'
		-1, // A
		-1, // A~2
		-1, // A~2_RepTerm
		-1, // A~1
		-1, // A~4
		-1, // A~3
		-1, // A~3_RepTerm

	},
	gotoRow{ // S4

		-1, // S'
		-1, // A
		-1, // A~2
		8,  // A~2_RepTerm
		-1, // A~1
		-1, // A~4
		-1, // A~3
		-1, // A~3_RepTerm

	},
	gotoRow{ // S5

		-1, // S'
		-1, // A
		-1, // A~2
		-1, // A~2_RepTerm
		-1, // A~1
		-1, // A~4
		-1, // A~3
		-1, // A~3_RepTerm

	},
	gotoRow{ // S6

		-1, // S'
		-1, // A
		-1, // A~2
		-1, // A~2_RepTerm
		-1, // A~1
		-1, // A~4
		10, // A~3
		11, // A~3_RepTerm

	},
	gotoRow{ // S7

		-1, // S'
		-1, // A
		-1, // A~2
		-1, // A~2_RepTerm
		-1, // A~1
		-1, // A~4
		-1, // A~3
		-1, // A~3_RepTerm

	},
	gotoRow{ // S8

		-1, // S'
		-1, // A
		-1, // A~2
		-1, // A~2_RepTerm
		-1, // A~1
		-1, // A~4
		-1, // A~3
		-1, // A~3_RepTerm

	},
	gotoRow{ // S9

		-1, // S'
		-1, // A
		-1, // A~2
		-1, // A~2_RepTerm
		-1, // A~1
		-1, // A~4
		-1, // A~3
		-1, // A~3_RepTerm

	},
	gotoRow{ // S10

		-1, // S'
		-1, // A
		-1, // A~2
		-1, // A~2_RepTerm
		-1, // A~1
		-1, // A~4
		-1, // A~3
		12, // A~3_RepTerm

	},
	gotoRow{ // S11

		-1, // S'
		-1, // A
		-1, // A~2
		-1, // A~2_RepTerm
		-1, // A~1
		-1, // A~4
		-1, // A~3
		-1, // A~3_RepTerm

	},
	gotoRow{ // S12

		-1, // S'
		-1, // A
		-1, // A~2
		-1, // A~2_RepTerm
		-1, // A~1
		-1, // A~4
		-1, // A~3
		-1, // A~3_RepTerm

	},
}
