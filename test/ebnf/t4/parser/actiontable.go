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
			nil,       /* INVALID */
			reduce(1), /* $, reduce: A */
			nil,       /* error */
			shift(1),  /* c1 */
			shift(2),  /* c2 */
			shift(3),  /* c3 */
			shift(4),  /* d1 */
			nil,       /* d2 */
			nil,       /* d3 */
			nil,       /* d4 */
			nil,       /* d5 */

		},
	},
	actionRow{ // S1
		canRecover: false,
		actions: [numTerminals]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* error */
			nil,        /* c1 */
			nil,        /* c2 */
			reduce(10), /* c3, reduce: C~1 */
			nil,        /* d1 */
			nil,        /* d2 */
			nil,        /* d3 */
			nil,        /* d4 */
			nil,        /* d5 */

		},
	},
	actionRow{ // S2
		canRecover: false,
		actions: [numTerminals]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* error */
			nil,        /* c1 */
			nil,        /* c2 */
			reduce(11), /* c3, reduce: C~1 */
			nil,        /* d1 */
			nil,        /* d2 */
			nil,        /* d3 */
			nil,        /* d4 */
			nil,        /* d5 */

		},
	},
	actionRow{ // S3
		canRecover: false,
		actions: [numTerminals]action{
			nil,       /* INVALID */
			reduce(8), /* $, reduce: C */
			nil,       /* error */
			reduce(8), /* c1, reduce: C */
			reduce(8), /* c2, reduce: C */
			reduce(8), /* c3, reduce: C */
			reduce(8), /* d1, reduce: C */
			nil,       /* d2 */
			nil,       /* d3 */
			nil,       /* d4 */
			nil,       /* d5 */

		},
	},
	actionRow{ // S4
		canRecover: false,
		actions: [numTerminals]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* error */
			nil,       /* c1 */
			nil,       /* c2 */
			nil,       /* c3 */
			nil,       /* d1 */
			shift(11), /* d2 */
			shift(12), /* d3 */
			shift(13), /* d4 */
			nil,       /* d5 */

		},
	},
	actionRow{ // S5
		canRecover: false,
		actions: [numTerminals]action{
			nil,          /* INVALID */
			accept(true), /* $ */
			nil,          /* error */
			nil,          /* c1 */
			nil,          /* c2 */
			nil,          /* c3 */
			nil,          /* d1 */
			nil,          /* d2 */
			nil,          /* d3 */
			nil,          /* d4 */
			nil,          /* d5 */

		},
	},
	actionRow{ // S6
		canRecover: false,
		actions: [numTerminals]action{
			nil,       /* INVALID */
			reduce(2), /* $, reduce: A */
			nil,       /* error */
			shift(1),  /* c1 */
			shift(2),  /* c2 */
			shift(3),  /* c3 */
			shift(4),  /* d1 */
			nil,       /* d2 */
			nil,       /* d3 */
			nil,       /* d4 */
			nil,       /* d5 */

		},
	},
	actionRow{ // S7
		canRecover: false,
		actions: [numTerminals]action{
			nil,       /* INVALID */
			reduce(5), /* $, reduce: A */
			nil,       /* error */
			nil,       /* c1 */
			nil,       /* c2 */
			nil,       /* c3 */
			shift(4),  /* d1 */
			nil,       /* d2 */
			nil,       /* d3 */
			nil,       /* d4 */
			nil,       /* d5 */

		},
	},
	actionRow{ // S8
		canRecover: false,
		actions: [numTerminals]action{
			nil,       /* INVALID */
			reduce(3), /* $, reduce: A~1 */
			nil,       /* error */
			reduce(3), /* c1, reduce: A~1 */
			reduce(3), /* c2, reduce: A~1 */
			reduce(3), /* c3, reduce: A~1 */
			reduce(3), /* d1, reduce: A~1 */
			nil,       /* d2 */
			nil,       /* d3 */
			nil,       /* d4 */
			nil,       /* d5 */

		},
	},
	actionRow{ // S9
		canRecover: false,
		actions: [numTerminals]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* error */
			nil,       /* c1 */
			nil,       /* c2 */
			shift(18), /* c3 */
			nil,       /* d1 */
			nil,       /* d2 */
			nil,       /* d3 */
			nil,       /* d4 */
			nil,       /* d5 */

		},
	},
	actionRow{ // S10
		canRecover: false,
		actions: [numTerminals]action{
			nil,        /* INVALID */
			reduce(6),  /* $, reduce: A~2 */
			nil,        /* error */
			nil,        /* c1 */
			nil,        /* c2 */
			reduce(12), /* c3, reduce: C~1 */
			reduce(6),  /* d1, reduce: A~2 */
			nil,        /* d2 */
			nil,        /* d3 */
			nil,        /* d4 */
			nil,        /* d5 */

		},
	},
	actionRow{ // S11
		canRecover: false,
		actions: [numTerminals]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* error */
			nil,        /* c1 */
			nil,        /* c2 */
			nil,        /* c3 */
			nil,        /* d1 */
			nil,        /* d2 */
			nil,        /* d3 */
			nil,        /* d4 */
			reduce(14), /* d5, reduce: D~1 */

		},
	},
	actionRow{ // S12
		canRecover: false,
		actions: [numTerminals]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* error */
			nil,        /* c1 */
			nil,        /* c2 */
			nil,        /* c3 */
			nil,        /* d1 */
			nil,        /* d2 */
			nil,        /* d3 */
			nil,        /* d4 */
			reduce(15), /* d5, reduce: D~1 */

		},
	},
	actionRow{ // S13
		canRecover: false,
		actions: [numTerminals]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* error */
			nil,        /* c1 */
			nil,        /* c2 */
			nil,        /* c3 */
			nil,        /* d1 */
			nil,        /* d2 */
			nil,        /* d3 */
			nil,        /* d4 */
			reduce(16), /* d5, reduce: D~1 */

		},
	},
	actionRow{ // S14
		canRecover: false,
		actions: [numTerminals]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* error */
			nil,       /* c1 */
			nil,       /* c2 */
			nil,       /* c3 */
			nil,       /* d1 */
			nil,       /* d2 */
			nil,       /* d3 */
			nil,       /* d4 */
			shift(19), /* d5 */

		},
	},
	actionRow{ // S15
		canRecover: false,
		actions: [numTerminals]action{
			nil,       /* INVALID */
			reduce(4), /* $, reduce: A~1 */
			nil,       /* error */
			reduce(4), /* c1, reduce: A~1 */
			reduce(4), /* c2, reduce: A~1 */
			reduce(4), /* c3, reduce: A~1 */
			reduce(4), /* d1, reduce: A~1 */
			nil,       /* d2 */
			nil,       /* d3 */
			nil,       /* d4 */
			nil,       /* d5 */

		},
	},
	actionRow{ // S16
		canRecover: false,
		actions: [numTerminals]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* error */
			nil,        /* c1 */
			nil,        /* c2 */
			reduce(12), /* c3, reduce: C~1 */
			nil,        /* d1 */
			nil,        /* d2 */
			nil,        /* d3 */
			nil,        /* d4 */
			nil,        /* d5 */

		},
	},
	actionRow{ // S17
		canRecover: false,
		actions: [numTerminals]action{
			nil,       /* INVALID */
			reduce(7), /* $, reduce: A~2 */
			nil,       /* error */
			nil,       /* c1 */
			nil,       /* c2 */
			nil,       /* c3 */
			reduce(7), /* d1, reduce: A~2 */
			nil,       /* d2 */
			nil,       /* d3 */
			nil,       /* d4 */
			nil,       /* d5 */

		},
	},
	actionRow{ // S18
		canRecover: false,
		actions: [numTerminals]action{
			nil,       /* INVALID */
			reduce(9), /* $, reduce: C */
			nil,       /* error */
			reduce(9), /* c1, reduce: C */
			reduce(9), /* c2, reduce: C */
			reduce(9), /* c3, reduce: C */
			reduce(9), /* d1, reduce: C */
			nil,       /* d2 */
			nil,       /* d3 */
			nil,       /* d4 */
			nil,       /* d5 */

		},
	},
	actionRow{ // S19
		canRecover: false,
		actions: [numTerminals]action{
			nil,        /* INVALID */
			reduce(13), /* $, reduce: D */
			nil,        /* error */
			nil,        /* c1 */
			nil,        /* c2 */
			reduce(13), /* c3, reduce: D */
			reduce(13), /* d1, reduce: D */
			nil,        /* d2 */
			nil,        /* d3 */
			nil,        /* d4 */
			nil,        /* d5 */

		},
	},
}
