<<<<<<< HEAD
package parser

type (
	actionTable [numStates]actionRow
	actionRow   struct {
		canRecover bool
		actions    [numSymbols]action
=======

package parser

type(
	actionTable [numStates]actionRow
	actionRow struct {
		canRecover bool
		actions [numSymbols]action
>>>>>>> 9d3e28b6b4a375015991b8ecb296891776774570
	}
)

var actionTab = actionTable{
	actionRow{ // S0
<<<<<<< HEAD
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			reduce(2), /* c, reduce: B */
			nil,       /* empty */
			shift(3),  /* b */

		},
	},
	actionRow{ // S1
		canRecover: false,
		actions: [numSymbols]action{
			nil,          /* INVALID */
			accept(true), /* $ */
			nil,          /* c */
			nil,          /* empty */
			nil,          /* b */

		},
	},
	actionRow{ // S2
		canRecover: false,
		actions: [numSymbols]action{
			nil,      /* INVALID */
			nil,      /* $ */
			shift(4), /* c */
			nil,      /* empty */
			nil,      /* b */

		},
	},
	actionRow{ // S3
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			reduce(3), /* c, reduce: B */
			nil,       /* empty */
			nil,       /* b */

		},
	},
	actionRow{ // S4
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(1), /* $, reduce: A */
			nil,       /* c */
			nil,       /* empty */
			nil,       /* b */

		},
	},
}
=======
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(2),		/* c, reduce: B */
			nil,		/* empty */
			shift(3),		/* b */
			
		},

	},
	actionRow{ // S1
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			accept(true),		/* $ */
			nil,		/* c */
			nil,		/* empty */
			nil,		/* b */
			
		},

	},
	actionRow{ // S2
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(4),		/* c */
			nil,		/* empty */
			nil,		/* b */
			
		},

	},
	actionRow{ // S3
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(3),		/* c, reduce: B */
			nil,		/* empty */
			nil,		/* b */
			
		},

	},
	actionRow{ // S4
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(1),		/* $, reduce: A */
			nil,		/* c */
			nil,		/* empty */
			nil,		/* b */
			
		},

	},
	
}

>>>>>>> 9d3e28b6b4a375015991b8ecb296891776774570
