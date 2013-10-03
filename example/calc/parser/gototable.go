/*
 */
package parser

const numNTSymbols = 5

type (
	gotoTable [numStates]gotoRow
	gotoRow   [numNTSymbols]int
)

var gotoTab = gotoTable{
	gotoRow{ // S0

		-1, // S'
		3,  // Calc
		4,  // Expr
		5,  // Term
		6,  // Factor

	},
	gotoRow{ // S1

		-1, // S'
		-1, // Calc
		7,  // Expr
		5,  // Term
		6,  // Factor

	},
	gotoRow{ // S2

		-1, // S'
		-1, // Calc
		-1, // Expr
		-1, // Term
		-1, // Factor

	},
	gotoRow{ // S3

		-1, // S'
		-1, // Calc
		-1, // Expr
		-1, // Term
		-1, // Factor

	},
	gotoRow{ // S4

		-1, // S'
		-1, // Calc
		-1, // Expr
		-1, // Term
		-1, // Factor

	},
	gotoRow{ // S5

		-1, // S'
		-1, // Calc
		-1, // Expr
		-1, // Term
		-1, // Factor

	},
	gotoRow{ // S6

		-1, // S'
		-1, // Calc
		-1, // Expr
		-1, // Term
		-1, // Factor

	},
	gotoRow{ // S7

		-1, // S'
		-1, // Calc
		-1, // Expr
		-1, // Term
		-1, // Factor

	},
	gotoRow{ // S8

		-1, // S'
		-1, // Calc
		-1, // Expr
		11, // Term
		6,  // Factor

	},
	gotoRow{ // S9

		-1, // S'
		-1, // Calc
		-1, // Expr
		-1, // Term
		12, // Factor

	},
	gotoRow{ // S10

		-1, // S'
		-1, // Calc
		-1, // Expr
		-1, // Term
		-1, // Factor

	},
	gotoRow{ // S11

		-1, // S'
		-1, // Calc
		-1, // Expr
		-1, // Term
		-1, // Factor

	},
	gotoRow{ // S12

		-1, // S'
		-1, // Calc
		-1, // Expr
		-1, // Term
		-1, // Factor

	},
}
