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
		6,  // BoolExpr
		7,  // Val
		8,  // CompareExpr
		9,  // SubStringExpr

	},
	gotoRow{ // S1

		-1, // S'
		10, // BoolExpr
		7,  // Val
		8,  // CompareExpr
		9,  // SubStringExpr

	},
	gotoRow{ // S2

		-1, // S'
		-1, // BoolExpr
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S3

		-1, // S'
		-1, // BoolExpr
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S4

		-1, // S'
		-1, // BoolExpr
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S5

		-1, // S'
		-1, // BoolExpr
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S6

		-1, // S'
		-1, // BoolExpr
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S7

		-1, // S'
		-1, // BoolExpr
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S8

		-1, // S'
		-1, // BoolExpr
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S9

		-1, // S'
		-1, // BoolExpr
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S10

		-1, // S'
		-1, // BoolExpr
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S11

		-1, // S'
		-1, // BoolExpr
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S12

		-1, // S'
		-1, // BoolExpr
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S13

		-1, // S'
		-1, // BoolExpr
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S14

		-1, // S'
		20, // BoolExpr
		7,  // Val
		8,  // CompareExpr
		9,  // SubStringExpr

	},
	gotoRow{ // S15

		-1, // S'
		21, // BoolExpr
		7,  // Val
		8,  // CompareExpr
		9,  // SubStringExpr

	},
	gotoRow{ // S16

		-1, // S'
		-1, // BoolExpr
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S17

		-1, // S'
		-1, // BoolExpr
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S18

		-1, // S'
		-1, // BoolExpr
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S19

		-1, // S'
		-1, // BoolExpr
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S20

		-1, // S'
		-1, // BoolExpr
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S21

		-1, // S'
		-1, // BoolExpr
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
}
