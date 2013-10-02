
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
			reduce(2),		/* $, reduce: A */
			nil,		/* error */
			shift(1),		/* a */
			
		},

	},
	actionRow{ // S1
				canRecover: false,
		actions: [numTerminals]action{
			nil,		/* INVALID */
			reduce(3),		/* $, reduce: A~1 */
			nil,		/* error */
			nil,		/* a */
			
		},

	},
	actionRow{ // S2
				canRecover: false,
		actions: [numTerminals]action{
			nil,		/* INVALID */
			accept(true),		/* $ */
			nil,		/* error */
			nil,		/* a */
			
		},

	},
	actionRow{ // S3
				canRecover: false,
		actions: [numTerminals]action{
			nil,		/* INVALID */
			reduce(1),		/* $, reduce: A */
			nil,		/* error */
			nil,		/* a */
			
		},

	},
	
}

