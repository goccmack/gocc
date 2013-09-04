
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
			shift(1),		/* Calc */
			shift(2),		/* Expr */
			nil,		/* + */
			shift(3),		/* Term */
			nil,		/* * */
			shift(4),		/* Factor */
			shift(5),		/* ( */
			nil,		/* ) */
			shift(6),		/* int64 */
			
		},

	},
	actionRow{ // S1
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			accept(true),		/* $ */
			nil,		/* S' */
			nil,		/* Calc */
			nil,		/* Expr */
			nil,		/* + */
			nil,		/* Term */
			nil,		/* * */
			nil,		/* Factor */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int64 */
			
		},

	},
	actionRow{ // S2
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(1),		/* $, reduce: Calc */
			nil,		/* S' */
			nil,		/* Calc */
			nil,		/* Expr */
			shift(7),		/* + */
			nil,		/* Term */
			nil,		/* * */
			nil,		/* Factor */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int64 */
			
		},

	},
	actionRow{ // S3
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(3),		/* $, reduce: Expr */
			nil,		/* S' */
			nil,		/* Calc */
			nil,		/* Expr */
			reduce(3),		/* +, reduce: Expr */
			nil,		/* Term */
			shift(8),		/* * */
			nil,		/* Factor */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int64 */
			
		},

	},
	actionRow{ // S4
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(5),		/* $, reduce: Term */
			nil,		/* S' */
			nil,		/* Calc */
			nil,		/* Expr */
			reduce(5),		/* +, reduce: Term */
			nil,		/* Term */
			reduce(5),		/* *, reduce: Term */
			nil,		/* Factor */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int64 */
			
		},

	},
	actionRow{ // S5
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* S' */
			nil,		/* Calc */
			shift(9),		/* Expr */
			nil,		/* + */
			shift(10),		/* Term */
			nil,		/* * */
			shift(11),		/* Factor */
			shift(12),		/* ( */
			nil,		/* ) */
			shift(13),		/* int64 */
			
		},

	},
	actionRow{ // S6
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(7),		/* $, reduce: Factor */
			nil,		/* S' */
			nil,		/* Calc */
			nil,		/* Expr */
			reduce(7),		/* +, reduce: Factor */
			nil,		/* Term */
			reduce(7),		/* *, reduce: Factor */
			nil,		/* Factor */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int64 */
			
		},

	},
	actionRow{ // S7
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* S' */
			nil,		/* Calc */
			nil,		/* Expr */
			nil,		/* + */
			shift(14),		/* Term */
			nil,		/* * */
			shift(4),		/* Factor */
			shift(5),		/* ( */
			nil,		/* ) */
			shift(6),		/* int64 */
			
		},

	},
	actionRow{ // S8
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* S' */
			nil,		/* Calc */
			nil,		/* Expr */
			nil,		/* + */
			nil,		/* Term */
			nil,		/* * */
			shift(15),		/* Factor */
			shift(5),		/* ( */
			nil,		/* ) */
			shift(6),		/* int64 */
			
		},

	},
	actionRow{ // S9
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* S' */
			nil,		/* Calc */
			nil,		/* Expr */
			shift(16),		/* + */
			nil,		/* Term */
			nil,		/* * */
			nil,		/* Factor */
			nil,		/* ( */
			shift(17),		/* ) */
			nil,		/* int64 */
			
		},

	},
	actionRow{ // S10
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* S' */
			nil,		/* Calc */
			nil,		/* Expr */
			reduce(3),		/* +, reduce: Expr */
			nil,		/* Term */
			shift(18),		/* * */
			nil,		/* Factor */
			nil,		/* ( */
			reduce(3),		/* ), reduce: Expr */
			nil,		/* int64 */
			
		},

	},
	actionRow{ // S11
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* S' */
			nil,		/* Calc */
			nil,		/* Expr */
			reduce(5),		/* +, reduce: Term */
			nil,		/* Term */
			reduce(5),		/* *, reduce: Term */
			nil,		/* Factor */
			nil,		/* ( */
			reduce(5),		/* ), reduce: Term */
			nil,		/* int64 */
			
		},

	},
	actionRow{ // S12
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* S' */
			nil,		/* Calc */
			shift(19),		/* Expr */
			nil,		/* + */
			shift(10),		/* Term */
			nil,		/* * */
			shift(11),		/* Factor */
			shift(12),		/* ( */
			nil,		/* ) */
			shift(13),		/* int64 */
			
		},

	},
	actionRow{ // S13
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* S' */
			nil,		/* Calc */
			nil,		/* Expr */
			reduce(7),		/* +, reduce: Factor */
			nil,		/* Term */
			reduce(7),		/* *, reduce: Factor */
			nil,		/* Factor */
			nil,		/* ( */
			reduce(7),		/* ), reduce: Factor */
			nil,		/* int64 */
			
		},

	},
	actionRow{ // S14
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(2),		/* $, reduce: Expr */
			nil,		/* S' */
			nil,		/* Calc */
			nil,		/* Expr */
			reduce(2),		/* +, reduce: Expr */
			nil,		/* Term */
			shift(8),		/* * */
			nil,		/* Factor */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int64 */
			
		},

	},
	actionRow{ // S15
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(4),		/* $, reduce: Term */
			nil,		/* S' */
			nil,		/* Calc */
			nil,		/* Expr */
			reduce(4),		/* +, reduce: Term */
			nil,		/* Term */
			reduce(4),		/* *, reduce: Term */
			nil,		/* Factor */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int64 */
			
		},

	},
	actionRow{ // S16
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* S' */
			nil,		/* Calc */
			nil,		/* Expr */
			nil,		/* + */
			shift(20),		/* Term */
			nil,		/* * */
			shift(11),		/* Factor */
			shift(12),		/* ( */
			nil,		/* ) */
			shift(13),		/* int64 */
			
		},

	},
	actionRow{ // S17
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(6),		/* $, reduce: Factor */
			nil,		/* S' */
			nil,		/* Calc */
			nil,		/* Expr */
			reduce(6),		/* +, reduce: Factor */
			nil,		/* Term */
			reduce(6),		/* *, reduce: Factor */
			nil,		/* Factor */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* int64 */
			
		},

	},
	actionRow{ // S18
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* S' */
			nil,		/* Calc */
			nil,		/* Expr */
			nil,		/* + */
			nil,		/* Term */
			nil,		/* * */
			shift(21),		/* Factor */
			shift(12),		/* ( */
			nil,		/* ) */
			shift(13),		/* int64 */
			
		},

	},
	actionRow{ // S19
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* S' */
			nil,		/* Calc */
			nil,		/* Expr */
			shift(16),		/* + */
			nil,		/* Term */
			nil,		/* * */
			nil,		/* Factor */
			nil,		/* ( */
			shift(22),		/* ) */
			nil,		/* int64 */
			
		},

	},
	actionRow{ // S20
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* S' */
			nil,		/* Calc */
			nil,		/* Expr */
			reduce(2),		/* +, reduce: Expr */
			nil,		/* Term */
			shift(18),		/* * */
			nil,		/* Factor */
			nil,		/* ( */
			reduce(2),		/* ), reduce: Expr */
			nil,		/* int64 */
			
		},

	},
	actionRow{ // S21
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* S' */
			nil,		/* Calc */
			nil,		/* Expr */
			reduce(4),		/* +, reduce: Term */
			nil,		/* Term */
			reduce(4),		/* *, reduce: Term */
			nil,		/* Factor */
			nil,		/* ( */
			reduce(4),		/* ), reduce: Term */
			nil,		/* int64 */
			
		},

	},
	actionRow{ // S22
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* S' */
			nil,		/* Calc */
			nil,		/* Expr */
			reduce(6),		/* +, reduce: Factor */
			nil,		/* Term */
			reduce(6),		/* *, reduce: Factor */
			nil,		/* Factor */
			nil,		/* ( */
			reduce(6),		/* ), reduce: Factor */
			nil,		/* int64 */
			
		},

	},
	
}

