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
		canRecover: true,
		actions: [numTerminals]action{
			nil,      /* INVALID */
			nil,      /* $ */
			shift(1), /* error */
			shift(2), /* id */

		},
	},
	actionRow{ // S1
		canRecover: false,
		actions: [numTerminals]action{
			nil,       /* INVALID */
			reduce(4), /* $, reduce: Stmt */
			reduce(4), /* error, reduce: Stmt */
			reduce(4), /* id, reduce: Stmt */

		},
	},
	actionRow{ // S2
		canRecover: false,
		actions: [numTerminals]action{
			nil,       /* INVALID */
			reduce(3), /* $, reduce: Stmt */
			reduce(3), /* error, reduce: Stmt */
			reduce(3), /* id, reduce: Stmt */

		},
	},
	actionRow{ // S3
		canRecover: true,
		actions: [numTerminals]action{
			nil,          /* INVALID */
			accept(true), /* $ */
			shift(1),     /* error */
			shift(2),     /* id */

		},
	},
	actionRow{ // S4
		canRecover: false,
		actions: [numTerminals]action{
			nil,       /* INVALID */
			reduce(1), /* $, reduce: StmtList */
			reduce(1), /* error, reduce: StmtList */
			reduce(1), /* id, reduce: StmtList */

		},
	},
	actionRow{ // S5
		canRecover: false,
		actions: [numTerminals]action{
			nil,       /* INVALID */
			reduce(2), /* $, reduce: StmtList */
			reduce(2), /* error, reduce: StmtList */
			reduce(2), /* id, reduce: StmtList */

		},
	},
}
