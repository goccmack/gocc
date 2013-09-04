
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
			shift(1),		/* RR */
			shift(2),		/* A */
			shift(3),		/* B */
			shift(4),		/* a */
			shift(5),		/* c */
			
		},

	},
	actionRow{ // S1
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			accept(true),		/* $ */
			nil,		/* S' */
			nil,		/* RR */
			nil,		/* A */
			nil,		/* B */
			nil,		/* a */
			nil,		/* c */
			
		},

	},
	actionRow{ // S2
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(1),		/* $, reduce: RR */
			nil,		/* S' */
			nil,		/* RR */
			nil,		/* A */
			nil,		/* B */
			shift(6),		/* a */
			nil,		/* c */
			
		},

	},
	actionRow{ // S3
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(2),		/* $, reduce: RR */
			nil,		/* S' */
			nil,		/* RR */
			nil,		/* A */
			nil,		/* B */
			nil,		/* a */
			nil,		/* c */
			
		},

	},
	actionRow{ // S4
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(3),		/* $, reduce: B */
			nil,		/* S' */
			nil,		/* RR */
			nil,		/* A */
			nil,		/* B */
			reduce(4),		/* a, reduce: A */
			nil,		/* c */
			
		},

	},
	actionRow{ // S5
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(6),		/* $, reduce: A */
			nil,		/* S' */
			nil,		/* RR */
			nil,		/* A */
			nil,		/* B */
			reduce(6),		/* a, reduce: A */
			nil,		/* c */
			
		},

	},
	actionRow{ // S6
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(5),		/* $, reduce: A */
			nil,		/* S' */
			nil,		/* RR */
			nil,		/* A */
			nil,		/* B */
			reduce(5),		/* a, reduce: A */
			nil,		/* c */
			
		},

	},
	
}

