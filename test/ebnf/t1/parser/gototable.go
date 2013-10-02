
/*
*/
package parser

const numNTSymbols = 3
type(
	gotoTable [numStates]gotoRow
	gotoRow	[numNTSymbols] int
)

var gotoTab = gotoTable{
	gotoRow{ // S0
		
		-1, // S'
		2, // A
		3, // A~1
		

	},
	gotoRow{ // S1
		
		-1, // S'
		-1, // A
		-1, // A~1
		

	},
	gotoRow{ // S2
		
		-1, // S'
		-1, // A
		-1, // A~1
		

	},
	gotoRow{ // S3
		
		-1, // S'
		-1, // A
		-1, // A~1
		

	},
	
}
