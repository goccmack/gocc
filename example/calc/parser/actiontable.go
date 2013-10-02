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
			shift(1), /* ( */
			nil,      /* ) */
			nil,      /* * */
			nil,      /* + */
			shift(2), /* int64 */

		},
	},
	actionRow{ // S1
		canRecover: false,
		actions: [numTerminals]action{
			nil,      /* INVALID */
			nil,      /* $ */
			nil,      /* error */
			shift(1), /* ( */
			nil,      /* ) */
			nil,      /* * */
			nil,      /* + */
			shift(2), /* int64 */

		},
	},
	actionRow{ // S2
		canRecover: false,
		actions: [numTerminals]action{
			nil,       /* INVALID */
			reduce(7), /* $, reduce: Factor */
			nil,       /* error */
			nil,       /* ( */
			reduce(7), /* ), reduce: Factor */
			reduce(7), /* *, reduce: Factor */
			reduce(7), /* +, reduce: Factor */
			nil,       /* int64 */

		},
	},
	actionRow{ // S3
		canRecover: false,
		actions: [numTerminals]action{
			nil,          /* INVALID */
			accept(true), /* $ */
			nil,          /* error */
			nil,          /* ( */
			nil,          /* ) */
			nil,          /* * */
			nil,          /* + */
			nil,          /* int64 */

		},
	},
	actionRow{ // S4
		canRecover: false,
		actions: [numTerminals]action{
			nil,       /* INVALID */
			reduce(1), /* $, reduce: Calc */
			nil,       /* error */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* * */
			shift(8),  /* + */
			nil,       /* int64 */

		},
	},
	actionRow{ // S5
		canRecover: false,
		actions: [numTerminals]action{
			nil,       /* INVALID */
			reduce(3), /* $, reduce: Expr */
			nil,       /* error */
			nil,       /* ( */
			reduce(3), /* ), reduce: Expr */
			shift(9),  /* * */
			reduce(3), /* +, reduce: Expr */
			nil,       /* int64 */

		},
	},
	actionRow{ // S6
		canRecover: false,
		actions: [numTerminals]action{
			nil,       /* INVALID */
			reduce(5), /* $, reduce: Term */
			nil,       /* error */
			nil,       /* ( */
			reduce(5), /* ), reduce: Term */
			reduce(5), /* *, reduce: Term */
			reduce(5), /* +, reduce: Term */
			nil,       /* int64 */

		},
	},
	actionRow{ // S7
		canRecover: false,
		actions: [numTerminals]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* error */
			nil,       /* ( */
			shift(10), /* ) */
			nil,       /* * */
			shift(8),  /* + */
			nil,       /* int64 */

		},
	},
	actionRow{ // S8
		canRecover: false,
		actions: [numTerminals]action{
			nil,      /* INVALID */
			nil,      /* $ */
			nil,      /* error */
			shift(1), /* ( */
			nil,      /* ) */
			nil,      /* * */
			nil,      /* + */
			shift(2), /* int64 */

		},
	},
	actionRow{ // S9
		canRecover: false,
		actions: [numTerminals]action{
			nil,      /* INVALID */
			nil,      /* $ */
			nil,      /* error */
			shift(1), /* ( */
			nil,      /* ) */
			nil,      /* * */
			nil,      /* + */
			shift(2), /* int64 */

		},
	},
	actionRow{ // S10
		canRecover: false,
		actions: [numTerminals]action{
			nil,       /* INVALID */
			reduce(6), /* $, reduce: Factor */
			nil,       /* error */
			nil,       /* ( */
			nil,       /* ) */
			reduce(6), /* *, reduce: Factor */
			reduce(6), /* +, reduce: Factor */
			nil,       /* int64 */

		},
	},
	actionRow{ // S11
		canRecover: false,
		actions: [numTerminals]action{
			nil,       /* INVALID */
			reduce(2), /* $, reduce: Expr */
			nil,       /* error */
			nil,       /* ( */
			reduce(2), /* ), reduce: Expr */
			shift(9),  /* * */
			reduce(2), /* +, reduce: Expr */
			nil,       /* int64 */

		},
	},
	actionRow{ // S12
		canRecover: false,
		actions: [numTerminals]action{
			nil,       /* INVALID */
			reduce(4), /* $, reduce: Term */
			nil,       /* error */
			nil,       /* ( */
			reduce(4), /* ), reduce: Term */
			reduce(4), /* *, reduce: Term */
			reduce(4), /* +, reduce: Term */
			nil,       /* int64 */

		},
	},
}
