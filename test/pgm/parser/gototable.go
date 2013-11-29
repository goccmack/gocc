
/*
*/
package parser

const numNTSymbols = 6
type(
	gotoTable [numStates]gotoRow
	gotoRow	[numNTSymbols] int
)

var gotoTab = gotoTable{
	gotoRow{ // S0
		
		-1, // S'
		3, // X
		-1, // Y
		-1, // Z
		-1, // T
		-1, // W
		

	},
	gotoRow{ // S1
		
		-1, // S'
		-1, // X
		6, // Y
		7, // Z
		8, // T
		-1, // W
		

	},
	gotoRow{ // S2
		
		-1, // S'
		-1, // X
		10, // Y
		11, // Z
		12, // T
		-1, // W
		

	},
	gotoRow{ // S3
		
		-1, // S'
		-1, // X
		-1, // Y
		-1, // Z
		-1, // T
		-1, // W
		

	},
	gotoRow{ // S4
		
		-1, // S'
		-1, // X
		-1, // Y
		-1, // Z
		-1, // T
		14, // W
		

	},
	gotoRow{ // S5
		
		-1, // S'
		15, // X
		-1, // Y
		-1, // Z
		-1, // T
		-1, // W
		

	},
	gotoRow{ // S6
		
		-1, // S'
		-1, // X
		-1, // Y
		-1, // Z
		-1, // T
		-1, // W
		

	},
	gotoRow{ // S7
		
		-1, // S'
		-1, // X
		-1, // Y
		-1, // Z
		-1, // T
		-1, // W
		

	},
	gotoRow{ // S8
		
		-1, // S'
		-1, // X
		-1, // Y
		-1, // Z
		-1, // T
		-1, // W
		

	},
	gotoRow{ // S9
		
		-1, // S'
		-1, // X
		-1, // Y
		-1, // Z
		-1, // T
		14, // W
		

	},
	gotoRow{ // S10
		
		-1, // S'
		-1, // X
		-1, // Y
		-1, // Z
		-1, // T
		-1, // W
		

	},
	gotoRow{ // S11
		
		-1, // S'
		-1, // X
		-1, // Y
		-1, // Z
		-1, // T
		-1, // W
		

	},
	gotoRow{ // S12
		
		-1, // S'
		-1, // X
		-1, // Y
		-1, // Z
		-1, // T
		-1, // W
		

	},
	gotoRow{ // S13
		
		-1, // S'
		-1, // X
		-1, // Y
		-1, // Z
		-1, // T
		-1, // W
		

	},
	gotoRow{ // S14
		
		-1, // S'
		-1, // X
		-1, // Y
		-1, // Z
		-1, // T
		-1, // W
		

	},
	gotoRow{ // S15
		
		-1, // S'
		-1, // X
		-1, // Y
		-1, // Z
		-1, // T
		-1, // W
		

	},
	gotoRow{ // S16
		
		-1, // S'
		-1, // X
		-1, // Y
		-1, // Z
		-1, // T
		-1, // W
		

	},
	gotoRow{ // S17
		
		-1, // S'
		-1, // X
		-1, // Y
		-1, // Z
		-1, // T
		-1, // W
		

	},
	gotoRow{ // S18
		
		-1, // S'
		-1, // X
		-1, // Y
		-1, // Z
		-1, // T
		-1, // W
		

	},
	gotoRow{ // S19
		
		-1, // S'
		-1, // X
		-1, // Y
		-1, // Z
		-1, // T
		-1, // W
		

	},
	gotoRow{ // S20
		
		-1, // S'
		-1, // X
		-1, // Y
		-1, // Z
		-1, // T
		-1, // W
		

	},
	gotoRow{ // S21
		
		-1, // S'
		-1, // X
		-1, // Y
		-1, // Z
		-1, // T
		-1, // W
		

	},
	
}
