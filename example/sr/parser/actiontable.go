package parser

type (
	actionTable [numStates]actionRow
	actionRow   struct {
		canRecover bool
		actions    [numTerminals]action
	}
)

var actionTab = actionTable{
	actionRow{ // S0
		canRecover: false,
		actions: [numTerminals]action{
			nil,      /* INVALID */
			nil,      /* $ */
			nil,      /* error */
			nil,      /* else */
			shift(1), /* id */
			shift(2), /* if */
			nil,      /* then */

		},
	},
	actionRow{ // S1
		canRecover: false,
		actions: [numTerminals]action{
			nil,       /* INVALID */
			reduce(3), /* $, reduce: Stmt */
			nil,       /* error */
			reduce(3), /* else, reduce: Stmt */
			nil,       /* id */
			nil,       /* if */
			nil,       /* then */

		},
	},
	actionRow{ // S2
		canRecover: false,
		actions: [numTerminals]action{
			nil,      /* INVALID */
			nil,      /* $ */
			nil,      /* error */
			nil,      /* else */
			shift(4), /* id */
			nil,      /* if */
			nil,      /* then */

		},
	},
	actionRow{ // S3
		canRecover: false,
		actions: [numTerminals]action{
			nil,          /* INVALID */
			accept(true), /* $ */
			nil,          /* error */
			nil,          /* else */
			nil,          /* id */
			nil,          /* if */
			nil,          /* then */

		},
	},
	actionRow{ // S4
		canRecover: false,
		actions: [numTerminals]action{
			nil,      /* INVALID */
			nil,      /* $ */
			nil,      /* error */
			nil,      /* else */
			nil,      /* id */
			nil,      /* if */
			shift(5), /* then */

		},
	},
	actionRow{ // S5
		canRecover: false,
		actions: [numTerminals]action{
			nil,      /* INVALID */
			nil,      /* $ */
			nil,      /* error */
			nil,      /* else */
			shift(1), /* id */
			shift(2), /* if */
			nil,      /* then */

		},
	},
	actionRow{ // S6
		canRecover: false,
		actions: [numTerminals]action{
			nil,       /* INVALID */
			reduce(1), /* $, reduce: Stmt */
			nil,       /* error */
			shift(7),  /* else */
			nil,       /* id */
			nil,       /* if */
			nil,       /* then */

		},
	},
	actionRow{ // S7
		canRecover: false,
		actions: [numTerminals]action{
			nil,      /* INVALID */
			nil,      /* $ */
			nil,      /* error */
			nil,      /* else */
			shift(1), /* id */
			shift(2), /* if */
			nil,      /* then */

		},
	},
	actionRow{ // S8
		canRecover: false,
		actions: [numTerminals]action{
			nil,       /* INVALID */
			reduce(2), /* $, reduce: Stmt */
			nil,       /* error */
			nil,       /* else */
			nil,       /* id */
			nil,       /* if */
			nil,       /* then */

		},
	},
}
