
package parser

type(
	actionTable [numStates]actionRow
	actionRow struct {
		canRecover bool
		actions [numTerminals]action
	}
)

var actionTab = actionTable{
	actionRow{ // S0
				canRecover: false,
		actions: [numTerminals]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* error */
			shift(1),		/* id */
			
		},

	},
	actionRow{ // S1
				canRecover: false,
		actions: [numTerminals]action{
			nil,		/* INVALID */
			reduce(3),		/* $, reduce: Stmt */
			nil,		/* error */
			reduce(3),		/* id, reduce: Stmt */
			
		},

	},
	actionRow{ // S2
				canRecover: false,
		actions: [numTerminals]action{
			nil,		/* INVALID */
			accept(true),		/* $ */
			nil,		/* error */
			shift(1),		/* id */
			
		},

	},
	actionRow{ // S3
				canRecover: false,
		actions: [numTerminals]action{
			nil,		/* INVALID */
			reduce(1),		/* $, reduce: StmtList */
			nil,		/* error */
			reduce(1),		/* id, reduce: StmtList */
			
		},

	},
	actionRow{ // S4
				canRecover: false,
		actions: [numTerminals]action{
			nil,		/* INVALID */
			reduce(2),		/* $, reduce: StmtList */
			nil,		/* error */
			reduce(2),		/* id, reduce: StmtList */
			
		},

	},
	
}

