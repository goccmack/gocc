<<<<<<< HEAD
/*
 */
package parser

const numNTSymbols = 3

type (
	gotoTable [numStates]gotoRow
	gotoRow   [numNTSymbols]int
=======

/*
*/
package parser

const numNTSymbols = 3
type(
	gotoTable [numStates]gotoRow
	gotoRow	[numNTSymbols] int
>>>>>>> 9d3e28b6b4a375015991b8ecb296891776774570
)

var gotoTab = gotoTable{
	gotoRow{ // S0
<<<<<<< HEAD

		-1, // S'
		1,  // A
		2,  // B

	},
	gotoRow{ // S1

		-1, // S'
		-1, // A
		-1, // B

	},
	gotoRow{ // S2

		-1, // S'
		-1, // A
		-1, // B

	},
	gotoRow{ // S3

		-1, // S'
		-1, // A
		-1, // B

	},
	gotoRow{ // S4

		-1, // S'
		-1, // A
		-1, // B

	},
=======
		
		-1, // S'
		1, // A
		2, // B
		

	},
	gotoRow{ // S1
		
		-1, // S'
		-1, // A
		-1, // B
		

	},
	gotoRow{ // S2
		
		-1, // S'
		-1, // A
		-1, // B
		

	},
	gotoRow{ // S3
		
		-1, // S'
		-1, // A
		-1, // B
		

	},
	gotoRow{ // S4
		
		-1, // S'
		-1, // A
		-1, // B
		

	},
	
>>>>>>> 9d3e28b6b4a375015991b8ecb296891776774570
}
