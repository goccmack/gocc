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
			reduce(2), /* $, reduce: A */
			nil,       /* error */
			shift(1),  /* a */

		},
	},
	actionRow{ // S1
		canRecover: false,
		actions: [numTerminals]action{
			nil,       /* INVALID */
			reduce(5), /* $, reduce: A~1_RepTerm */
			nil,       /* error */
			reduce(5), /* a, reduce: A~1_RepTerm */

		},
	},
	actionRow{ // S2
		canRecover: false,
		actions: [numTerminals]action{
			nil,          /* INVALID */
			accept(true), /* $ */
			nil,          /* error */
			nil,          /* a */

		},
	},
	actionRow{ // S3
		canRecover: false,
		actions: [numTerminals]action{
			nil,       /* INVALID */
			reduce(1), /* $, reduce: A */
			nil,       /* error */
			shift(1),  /* a */

		},
	},
	actionRow{ // S4
		canRecover: false,
		actions: [numTerminals]action{
			nil,       /* INVALID */
			reduce(3), /* $, reduce: A~1 */
			nil,       /* error */
			reduce(3), /* a, reduce: A~1 */

		},
	},
	actionRow{ // S5
		canRecover: false,
		actions: [numTerminals]action{
			nil,       /* INVALID */
			reduce(4), /* $, reduce: A~1 */
			nil,       /* error */
			reduce(4), /* a, reduce: A~1 */

		},
	},
}
