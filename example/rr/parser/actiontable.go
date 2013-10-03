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
			shift(1), /* a */
			shift(2), /* c */

		},
	},
	actionRow{ // S1
		canRecover: false,
		actions: [numTerminals]action{
			nil,       /* INVALID */
			reduce(3), /* $, reduce: B */
			nil,       /* error */
			reduce(4), /* a, reduce: A */
			nil,       /* c */

		},
	},
	actionRow{ // S2
		canRecover: false,
		actions: [numTerminals]action{
			nil,       /* INVALID */
			reduce(6), /* $, reduce: A */
			nil,       /* error */
			reduce(6), /* a, reduce: A */
			nil,       /* c */

		},
	},
	actionRow{ // S3
		canRecover: false,
		actions: [numTerminals]action{
			nil,          /* INVALID */
			accept(true), /* $ */
			nil,          /* error */
			nil,          /* a */
			nil,          /* c */

		},
	},
	actionRow{ // S4
		canRecover: false,
		actions: [numTerminals]action{
			nil,       /* INVALID */
			reduce(2), /* $, reduce: RR */
			nil,       /* error */
			nil,       /* a */
			nil,       /* c */

		},
	},
	actionRow{ // S5
		canRecover: false,
		actions: [numTerminals]action{
			nil,       /* INVALID */
			reduce(1), /* $, reduce: RR */
			nil,       /* error */
			shift(6),  /* a */
			nil,       /* c */

		},
	},
	actionRow{ // S6
		canRecover: false,
		actions: [numTerminals]action{
			nil,       /* INVALID */
			reduce(5), /* $, reduce: A */
			nil,       /* error */
			reduce(5), /* a, reduce: A */
			nil,       /* c */

		},
	},
}
