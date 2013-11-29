
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
			reduce(1),		/* $, reduce: A */
			nil,		/* error */
			shift(1),		/* a */
			shift(2),		/* b */
			
		},

	},
	actionRow{ // S1
				canRecover: false,
		actions: [numTerminals]action{
			nil,		/* INVALID */
			reduce(3),		/* $, reduce: A~2 */
			nil,		/* error */
			reduce(3),		/* a, reduce: A~2 */
			nil,		/* b */
			
		},

	},
	actionRow{ // S2
				canRecover: false,
		actions: [numTerminals]action{
			nil,		/* INVALID */
			reduce(7),		/* $, reduce: A~1 */
			nil,		/* error */
			reduce(7),		/* a, reduce: A~1 */
			nil,		/* b */
			
		},

	},
	actionRow{ // S3
				canRecover: false,
		actions: [numTerminals]action{
			nil,		/* INVALID */
			accept(true),		/* $ */
			nil,		/* error */
			nil,		/* a */
			nil,		/* b */
			
		},

	},
	actionRow{ // S4
				canRecover: false,
		actions: [numTerminals]action{
			nil,		/* INVALID */
			reduce(2),		/* $, reduce: A */
			nil,		/* error */
			shift(6),		/* a */
			nil,		/* b */
			
		},

	},
	actionRow{ // S5
				canRecover: false,
		actions: [numTerminals]action{
			nil,		/* INVALID */
			reduce(5),		/* $, reduce: A */
			nil,		/* error */
			shift(1),		/* a */
			nil,		/* b */
			
		},

	},
	actionRow{ // S6
				canRecover: false,
		actions: [numTerminals]action{
			nil,		/* INVALID */
			reduce(4),		/* $, reduce: A~2 */
			nil,		/* error */
			reduce(4),		/* a, reduce: A~2 */
			nil,		/* b */
			
		},

	},
	actionRow{ // S7
				canRecover: false,
		actions: [numTerminals]action{
			nil,		/* INVALID */
			reduce(6),		/* $, reduce: A */
			nil,		/* error */
			shift(6),		/* a */
			nil,		/* b */
			
		},

	},
	
}

