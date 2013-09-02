
package parser

type(
	actionTable [numStates]actionRow
	actionRow struct {
		canRecover bool
		actions [numSymbols]action
	}
)

var actionTab = actionTable{
	actionRow{ // S0
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* S' */
			shift(1),		/* StmtList */
			shift(2),		/* Stmt */
			shift(3),		/* id */
			
		},

	},
	actionRow{ // S1
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			accept(true),		/* $ */
			nil,		/* S' */
			nil,		/* StmtList */
			shift(4),		/* Stmt */
			shift(3),		/* id */
			
		},

	},
	actionRow{ // S2
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(1),		/* $, reduce: StmtList */
			nil,		/* S' */
			nil,		/* StmtList */
			nil,		/* Stmt */
			reduce(1),		/* id, reduce: StmtList */
			
		},

	},
	actionRow{ // S3
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(3),		/* $, reduce: Stmt */
			nil,		/* S' */
			nil,		/* StmtList */
			nil,		/* Stmt */
			reduce(3),		/* id, reduce: Stmt */
			
		},

	},
	actionRow{ // S4
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(2),		/* $, reduce: StmtList */
			nil,		/* S' */
			nil,		/* StmtList */
			nil,		/* Stmt */
			reduce(2),		/* id, reduce: StmtList */
			
		},

	},
	
}

