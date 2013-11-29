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
			shift(2), /* b */
			nil,      /* c */
			nil,      /* d */
			nil,      /* e */
			nil,      /* t */
			nil,      /* u */

		},
	},
	actionRow{ // S1
		canRecover: false,
		actions: [numTerminals]action{
			nil,      /* INVALID */
			nil,      /* $ */
			nil,      /* error */
			nil,      /* a */
			nil,      /* b */
			nil,      /* c */
			nil,      /* d */
			nil,      /* e */
			shift(4), /* t */
			shift(5), /* u */

		},
	},
	actionRow{ // S2
		canRecover: false,
		actions: [numTerminals]action{
			nil,      /* INVALID */
			nil,      /* $ */
			nil,      /* error */
			nil,      /* a */
			nil,      /* b */
			nil,      /* c */
			nil,      /* d */
			nil,      /* e */
			shift(9), /* t */
			shift(5), /* u */

		},
	},
	actionRow{ // S3
		canRecover: false,
		actions: [numTerminals]action{
			nil,          /* INVALID */
			accept(true), /* $ */
			nil,          /* error */
			nil,          /* a */
			nil,          /* b */
			nil,          /* c */
			nil,          /* d */
			nil,          /* e */
			nil,          /* t */
			nil,          /* u */

		},
	},
	actionRow{ // S4
		canRecover: false,
		actions: [numTerminals]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* error */
			nil,       /* a */
			nil,       /* b */
			nil,       /* c */
			nil,       /* d */
			nil,       /* e */
			nil,       /* t */
			shift(13), /* u */

		},
	},
	actionRow{ // S5
		canRecover: false,
		actions: [numTerminals]action{
			nil,      /* INVALID */
			nil,      /* $ */
			nil,      /* error */
			shift(1), /* a */
			shift(2), /* b */
			nil,      /* c */
			nil,      /* d */
			nil,      /* e */
			nil,      /* t */
			nil,      /* u */

		},
	},
	actionRow{ // S6
		canRecover: false,
		actions: [numTerminals]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* error */
			nil,       /* a */
			nil,       /* b */
			nil,       /* c */
			shift(16), /* d */
			nil,       /* e */
			nil,       /* t */
			nil,       /* u */

		},
	},
	actionRow{ // S7
		canRecover: false,
		actions: [numTerminals]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* error */
			nil,       /* a */
			nil,       /* b */
			shift(17), /* c */
			nil,       /* d */
			nil,       /* e */
			nil,       /* t */
			nil,       /* u */

		},
	},
	actionRow{ // S8
		canRecover: false,
		actions: [numTerminals]action{
			nil,       /* INVALID */
			reduce(3), /* $, reduce: X */
			nil,       /* error */
			reduce(3), /* a, reduce: X */
			nil,       /* b */
			nil,       /* c */
			reduce(3), /* d, reduce: X */
			reduce(3), /* e, reduce: X */
			nil,       /* t */
			nil,       /* u */

		},
	},
	actionRow{ // S9
		canRecover: false,
		actions: [numTerminals]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* error */
			nil,       /* a */
			nil,       /* b */
			nil,       /* c */
			nil,       /* d */
			nil,       /* e */
			nil,       /* t */
			shift(18), /* u */

		},
	},
	actionRow{ // S10
		canRecover: false,
		actions: [numTerminals]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* error */
			nil,       /* a */
			nil,       /* b */
			nil,       /* c */
			nil,       /* d */
			shift(19), /* e */
			nil,       /* t */
			nil,       /* u */

		},
	},
	actionRow{ // S11
		canRecover: false,
		actions: [numTerminals]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* error */
			nil,       /* a */
			nil,       /* b */
			nil,       /* c */
			shift(20), /* d */
			nil,       /* e */
			nil,       /* t */
			nil,       /* u */

		},
	},
	actionRow{ // S12
		canRecover: false,
		actions: [numTerminals]action{
			nil,       /* INVALID */
			reduce(6), /* $, reduce: X */
			nil,       /* error */
			reduce(6), /* a, reduce: X */
			nil,       /* b */
			nil,       /* c */
			reduce(6), /* d, reduce: X */
			reduce(6), /* e, reduce: X */
			nil,       /* t */
			nil,       /* u */

		},
	},
	actionRow{ // S13
		canRecover: false,
		actions: [numTerminals]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* error */
			nil,        /* a */
			nil,        /* b */
			reduce(9),  /* c, reduce: Z */
			reduce(11), /* d, reduce: W */
			nil,        /* e */
			nil,        /* t */
			nil,        /* u */

		},
	},
	actionRow{ // S14
		canRecover: false,
		actions: [numTerminals]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* error */
			nil,       /* a */
			nil,       /* b */
			nil,       /* c */
			reduce(7), /* d, reduce: Y */
			reduce(7), /* e, reduce: Y */
			nil,       /* t */
			nil,       /* u */

		},
	},
	actionRow{ // S15
		canRecover: false,
		actions: [numTerminals]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* error */
			shift(21), /* a */
			nil,       /* b */
			nil,       /* c */
			reduce(8), /* d, reduce: Y */
			reduce(8), /* e, reduce: Y */
			nil,       /* t */
			nil,       /* u */

		},
	},
	actionRow{ // S16
		canRecover: false,
		actions: [numTerminals]action{
			nil,       /* INVALID */
			reduce(1), /* $, reduce: X */
			nil,       /* error */
			reduce(1), /* a, reduce: X */
			nil,       /* b */
			nil,       /* c */
			reduce(1), /* d, reduce: X */
			reduce(1), /* e, reduce: X */
			nil,       /* t */
			nil,       /* u */

		},
	},
	actionRow{ // S17
		canRecover: false,
		actions: [numTerminals]action{
			nil,       /* INVALID */
			reduce(2), /* $, reduce: X */
			nil,       /* error */
			reduce(2), /* a, reduce: X */
			nil,       /* b */
			nil,       /* c */
			reduce(2), /* d, reduce: X */
			reduce(2), /* e, reduce: X */
			nil,       /* t */
			nil,       /* u */

		},
	},
	actionRow{ // S18
		canRecover: false,
		actions: [numTerminals]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* error */
			nil,        /* a */
			nil,        /* b */
			nil,        /* c */
			reduce(9),  /* d, reduce: Z */
			reduce(11), /* e, reduce: W */
			nil,        /* t */
			nil,        /* u */

		},
	},
	actionRow{ // S19
		canRecover: false,
		actions: [numTerminals]action{
			nil,       /* INVALID */
			reduce(4), /* $, reduce: X */
			nil,       /* error */
			reduce(4), /* a, reduce: X */
			nil,       /* b */
			nil,       /* c */
			reduce(4), /* d, reduce: X */
			reduce(4), /* e, reduce: X */
			nil,       /* t */
			nil,       /* u */

		},
	},
	actionRow{ // S20
		canRecover: false,
		actions: [numTerminals]action{
			nil,       /* INVALID */
			reduce(5), /* $, reduce: X */
			nil,       /* error */
			reduce(5), /* a, reduce: X */
			nil,       /* b */
			nil,       /* c */
			reduce(5), /* d, reduce: X */
			reduce(5), /* e, reduce: X */
			nil,       /* t */
			nil,       /* u */

		},
	},
	actionRow{ // S21
		canRecover: false,
		actions: [numTerminals]action{
			nil,        /* INVALID */
			reduce(10), /* $, reduce: T */
			nil,        /* error */
			nil,        /* a */
			nil,        /* b */
			nil,        /* c */
			nil,        /* d */
			nil,        /* e */
			nil,        /* t */
			nil,        /* u */

		},
	},
}
