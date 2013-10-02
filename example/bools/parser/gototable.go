/*
 */
package parser

const numNTSymbols = 6

type (
	gotoTable [numStates]gotoRow
	gotoRow   [numNTSymbols]int
)

var gotoTab = gotoTable{
	gotoRow{ // S0

		-1, // S'
		6,  // BoolExpr
		7,  // BoolExpr1
		8,  // Val
		9,  // CompareExpr
		10, // SubStringExpr

	},
	gotoRow{ // S1

		-1, // S'
		11, // BoolExpr
		7,  // BoolExpr1
		8,  // Val
		9,  // CompareExpr
		10, // SubStringExpr

	},
	gotoRow{ // S2

		-1, // S'
		-1, // BoolExpr
		-1, // BoolExpr1
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S3

		-1, // S'
		-1, // BoolExpr
		-1, // BoolExpr1
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S4

		-1, // S'
		-1, // BoolExpr
		-1, // BoolExpr1
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S5

		-1, // S'
		-1, // BoolExpr
		-1, // BoolExpr1
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S6

		-1, // S'
		-1, // BoolExpr
		-1, // BoolExpr1
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S7

		-1, // S'
		-1, // BoolExpr
		-1, // BoolExpr1
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S8

		-1, // S'
		-1, // BoolExpr
		-1, // BoolExpr1
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S9

		-1, // S'
		-1, // BoolExpr
		-1, // BoolExpr1
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S10

		-1, // S'
		-1, // BoolExpr
		-1, // BoolExpr1
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S11

		-1, // S'
		-1, // BoolExpr
		-1, // BoolExpr1
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S12

		-1, // S'
		-1, // BoolExpr
		-1, // BoolExpr1
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S13

		-1, // S'
		-1, // BoolExpr
		-1, // BoolExpr1
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S14

		-1, // S'
		-1, // BoolExpr
		-1, // BoolExpr1
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S15

		-1, // S'
		21, // BoolExpr
		22, // BoolExpr1
		8,  // Val
		9,  // CompareExpr
		10, // SubStringExpr

	},
	gotoRow{ // S16

		-1, // S'
		21, // BoolExpr
		23, // BoolExpr1
		8,  // Val
		9,  // CompareExpr
		10, // SubStringExpr

	},
	gotoRow{ // S17

		-1, // S'
		-1, // BoolExpr
		-1, // BoolExpr1
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S18

		-1, // S'
		-1, // BoolExpr
		-1, // BoolExpr1
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S19

		-1, // S'
		-1, // BoolExpr
		-1, // BoolExpr1
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S20

		-1, // S'
		-1, // BoolExpr
		-1, // BoolExpr1
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S21

		-1, // S'
		-1, // BoolExpr
		-1, // BoolExpr1
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S22

		-1, // S'
		-1, // BoolExpr
		-1, // BoolExpr1
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
	gotoRow{ // S23

		-1, // S'
		-1, // BoolExpr
		-1, // BoolExpr1
		-1, // Val
		-1, // CompareExpr
		-1, // SubStringExpr

	},
}
