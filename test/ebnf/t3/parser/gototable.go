
/*
*/
package parser

const numNTSymbols = 4
type(
	gotoTable [numStates]gotoRow
	gotoRow	[numNTSymbols] int
)

var gotoTab = gotoTable{
	gotoRow{ // S0
		
		-1, // S'
		3, // A
		4, // A~2
		5, // A~1
		

	},
	gotoRow{ // S1
		
		-1, // S'
		-1, // A
		-1, // A~2
		-1, // A~1
		

	},
	gotoRow{ // S2
		
		-1, // S'
		-1, // A
		-1, // A~2
		-1, // A~1
		

	},
	gotoRow{ // S3
		
		-1, // S'
		-1, // A
		-1, // A~2
		-1, // A~1
		

	},
	gotoRow{ // S4
		
		-1, // S'
		-1, // A
		-1, // A~2
		-1, // A~1
		

	},
	gotoRow{ // S5
		
		-1, // S'
		-1, // A
		7, // A~2
		-1, // A~1
		

	},
	gotoRow{ // S6
		
		-1, // S'
		-1, // A
		-1, // A~2
		-1, // A~1
		

	},
	gotoRow{ // S7
		
		-1, // S'
		-1, // A
		-1, // A~2
		-1, // A~1
		

	},
	
}
