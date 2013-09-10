package parser

type (
	actionTable [numStates]actionRow
	actionRow   struct {
		canRecover bool
		actions    [numSymbols]action
	}
)

var actionTab = actionTable{
	actionRow{ // S0
		canRecover: true,
		actions: [numSymbols]action{
			nil,      /* INVALID */
			nil,      /* $ */
			nil,      /* S' */
			shift(1), /* StmtList */
			shift(2), /* Stmt */
			shift(3), /* id */
			shift(4), /* error */

		},
	},
	actionRow{ // S1
		canRecover: true,
		actions: [numSymbols]action{
			nil,          /* INVALID */
			accept(true), /* $ */
			nil,          /* S' */
			nil,          /* StmtList */
			shift(5),     /* Stmt */
			shift(3),     /* id */
			shift(4),     /* error */

		},
	},
	actionRow{ // S2
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(1), /* $, reduce: StmtList */
			nil,       /* S' */
			nil,       /* StmtList */
			nil,       /* Stmt */
			reduce(1), /* id, reduce: StmtList */
			reduce(1), /* error, reduce: StmtList */

		},
	},
	actionRow{ // S3
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(3), /* $, reduce: Stmt */
			nil,       /* S' */
			nil,       /* StmtList */
			nil,       /* Stmt */
			reduce(3), /* id, reduce: Stmt */
			reduce(3), /* error, reduce: Stmt */

		},
	},
	actionRow{ // S4
		canRecover: true,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(4), /* $, reduce: Stmt */
			nil,       /* S' */
			nil,       /* StmtList */
			nil,       /* Stmt */
			reduce(4), /* id, reduce: Stmt */
			reduce(4), /* error, reduce: Stmt */

		},
	},
	actionRow{ // S5
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(2), /* $, reduce: StmtList */
			nil,       /* S' */
			nil,       /* StmtList */
			nil,       /* Stmt */
			reduce(2), /* id, reduce: StmtList */
			reduce(2), /* error, reduce: StmtList */

		},
	},
}
