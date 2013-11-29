
package parser

type(
	actionTable [numStates]actionRow
	actionRow struct {
		canRecover bool
		actions [numTerminals]action
	}
)

var actionTab = actionTable{
	actionRow{ // S0
				canRecover: false,
		actions: [numTerminals]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* error */
			nil,		/* & */
			shift(1),		/* ( */
			nil,		/* ) */
			nil,		/* < */
			nil,		/* > */
			shift(2),		/* false */
			nil,		/* in */
			shift(3),		/* int_lit */
			shift(4),		/* string_lit */
			shift(5),		/* true */
			nil,		/* | */
			
		},

	},
	actionRow{ // S1
				canRecover: false,
		actions: [numTerminals]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* error */
			nil,		/* & */
			shift(1),		/* ( */
			nil,		/* ) */
			nil,		/* < */
			nil,		/* > */
			shift(2),		/* false */
			nil,		/* in */
			shift(3),		/* int_lit */
			shift(4),		/* string_lit */
			shift(5),		/* true */
			nil,		/* | */
			
		},

	},
	actionRow{ // S2
				canRecover: false,
		actions: [numTerminals]action{
			nil,		/* INVALID */
			reduce(7),		/* $, reduce: Val */
			nil,		/* error */
			reduce(7),		/* &, reduce: Val */
			nil,		/* ( */
			reduce(7),		/* ), reduce: Val */
			nil,		/* < */
			nil,		/* > */
			nil,		/* false */
			nil,		/* in */
			nil,		/* int_lit */
			nil,		/* string_lit */
			nil,		/* true */
			reduce(7),		/* |, reduce: Val */
			
		},

	},
	actionRow{ // S3
				canRecover: false,
		actions: [numTerminals]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* error */
			nil,		/* & */
			nil,		/* ( */
			nil,		/* ) */
			shift(12),		/* < */
			shift(13),		/* > */
			nil,		/* false */
			nil,		/* in */
			nil,		/* int_lit */
			nil,		/* string_lit */
			nil,		/* true */
			nil,		/* | */
			
		},

	},
	actionRow{ // S4
				canRecover: false,
		actions: [numTerminals]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* error */
			nil,		/* & */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* < */
			nil,		/* > */
			nil,		/* false */
			shift(14),		/* in */
			nil,		/* int_lit */
			nil,		/* string_lit */
			nil,		/* true */
			nil,		/* | */
			
		},

	},
	actionRow{ // S5
				canRecover: false,
		actions: [numTerminals]action{
			nil,		/* INVALID */
			reduce(6),		/* $, reduce: Val */
			nil,		/* error */
			reduce(6),		/* &, reduce: Val */
			nil,		/* ( */
			reduce(6),		/* ), reduce: Val */
			nil,		/* < */
			nil,		/* > */
			nil,		/* false */
			nil,		/* in */
			nil,		/* int_lit */
			nil,		/* string_lit */
			nil,		/* true */
			reduce(6),		/* |, reduce: Val */
			
		},

	},
	actionRow{ // S6
				canRecover: false,
		actions: [numTerminals]action{
			nil,		/* INVALID */
			accept(true),		/* $ */
			nil,		/* error */
			shift(15),		/* & */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* < */
			nil,		/* > */
			nil,		/* false */
			nil,		/* in */
			nil,		/* int_lit */
			nil,		/* string_lit */
			nil,		/* true */
			shift(16),		/* | */
			
		},

	},
	actionRow{ // S7
				canRecover: false,
		actions: [numTerminals]action{
			nil,		/* INVALID */
			reduce(1),		/* $, reduce: BoolExpr */
			nil,		/* error */
			reduce(1),		/* &, reduce: BoolExpr */
			nil,		/* ( */
			reduce(1),		/* ), reduce: BoolExpr */
			nil,		/* < */
			nil,		/* > */
			nil,		/* false */
			nil,		/* in */
			nil,		/* int_lit */
			nil,		/* string_lit */
			nil,		/* true */
			reduce(1),		/* |, reduce: BoolExpr */
			
		},

	},
	actionRow{ // S8
				canRecover: false,
		actions: [numTerminals]action{
			nil,		/* INVALID */
			reduce(2),		/* $, reduce: BoolExpr1 */
			nil,		/* error */
			reduce(2),		/* &, reduce: BoolExpr1 */
			nil,		/* ( */
			reduce(2),		/* ), reduce: BoolExpr1 */
			nil,		/* < */
			nil,		/* > */
			nil,		/* false */
			nil,		/* in */
			nil,		/* int_lit */
			nil,		/* string_lit */
			nil,		/* true */
			reduce(2),		/* |, reduce: BoolExpr1 */
			
		},

	},
	actionRow{ // S9
				canRecover: false,
		actions: [numTerminals]action{
			nil,		/* INVALID */
			reduce(8),		/* $, reduce: Val */
			nil,		/* error */
			reduce(8),		/* &, reduce: Val */
			nil,		/* ( */
			reduce(8),		/* ), reduce: Val */
			nil,		/* < */
			nil,		/* > */
			nil,		/* false */
			nil,		/* in */
			nil,		/* int_lit */
			nil,		/* string_lit */
			nil,		/* true */
			reduce(8),		/* |, reduce: Val */
			
		},

	},
	actionRow{ // S10
				canRecover: false,
		actions: [numTerminals]action{
			nil,		/* INVALID */
			reduce(9),		/* $, reduce: Val */
			nil,		/* error */
			reduce(9),		/* &, reduce: Val */
			nil,		/* ( */
			reduce(9),		/* ), reduce: Val */
			nil,		/* < */
			nil,		/* > */
			nil,		/* false */
			nil,		/* in */
			nil,		/* int_lit */
			nil,		/* string_lit */
			nil,		/* true */
			reduce(9),		/* |, reduce: Val */
			
		},

	},
	actionRow{ // S11
				canRecover: false,
		actions: [numTerminals]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* error */
			shift(15),		/* & */
			nil,		/* ( */
			shift(17),		/* ) */
			nil,		/* < */
			nil,		/* > */
			nil,		/* false */
			nil,		/* in */
			nil,		/* int_lit */
			nil,		/* string_lit */
			nil,		/* true */
			shift(16),		/* | */
			
		},

	},
	actionRow{ // S12
				canRecover: false,
		actions: [numTerminals]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* error */
			nil,		/* & */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* < */
			nil,		/* > */
			nil,		/* false */
			nil,		/* in */
			shift(18),		/* int_lit */
			nil,		/* string_lit */
			nil,		/* true */
			nil,		/* | */
			
		},

	},
	actionRow{ // S13
				canRecover: false,
		actions: [numTerminals]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* error */
			nil,		/* & */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* < */
			nil,		/* > */
			nil,		/* false */
			nil,		/* in */
			shift(19),		/* int_lit */
			nil,		/* string_lit */
			nil,		/* true */
			nil,		/* | */
			
		},

	},
	actionRow{ // S14
				canRecover: false,
		actions: [numTerminals]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* error */
			nil,		/* & */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* < */
			nil,		/* > */
			nil,		/* false */
			nil,		/* in */
			nil,		/* int_lit */
			shift(20),		/* string_lit */
			nil,		/* true */
			nil,		/* | */
			
		},

	},
	actionRow{ // S15
				canRecover: false,
		actions: [numTerminals]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* error */
			nil,		/* & */
			shift(1),		/* ( */
			nil,		/* ) */
			nil,		/* < */
			nil,		/* > */
			shift(2),		/* false */
			nil,		/* in */
			shift(3),		/* int_lit */
			shift(4),		/* string_lit */
			shift(5),		/* true */
			nil,		/* | */
			
		},

	},
	actionRow{ // S16
				canRecover: false,
		actions: [numTerminals]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* error */
			nil,		/* & */
			shift(1),		/* ( */
			nil,		/* ) */
			nil,		/* < */
			nil,		/* > */
			shift(2),		/* false */
			nil,		/* in */
			shift(3),		/* int_lit */
			shift(4),		/* string_lit */
			shift(5),		/* true */
			nil,		/* | */
			
		},

	},
	actionRow{ // S17
				canRecover: false,
		actions: [numTerminals]action{
			nil,		/* INVALID */
			reduce(5),		/* $, reduce: BoolExpr1 */
			nil,		/* error */
			reduce(5),		/* &, reduce: BoolExpr1 */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* < */
			nil,		/* > */
			nil,		/* false */
			nil,		/* in */
			nil,		/* int_lit */
			nil,		/* string_lit */
			nil,		/* true */
			reduce(5),		/* |, reduce: BoolExpr1 */
			
		},

	},
	actionRow{ // S18
				canRecover: false,
		actions: [numTerminals]action{
			nil,		/* INVALID */
			reduce(10),		/* $, reduce: CompareExpr */
			nil,		/* error */
			reduce(10),		/* &, reduce: CompareExpr */
			nil,		/* ( */
			reduce(10),		/* ), reduce: CompareExpr */
			nil,		/* < */
			nil,		/* > */
			nil,		/* false */
			nil,		/* in */
			nil,		/* int_lit */
			nil,		/* string_lit */
			nil,		/* true */
			reduce(10),		/* |, reduce: CompareExpr */
			
		},

	},
	actionRow{ // S19
				canRecover: false,
		actions: [numTerminals]action{
			nil,		/* INVALID */
			reduce(11),		/* $, reduce: CompareExpr */
			nil,		/* error */
			reduce(11),		/* &, reduce: CompareExpr */
			nil,		/* ( */
			reduce(11),		/* ), reduce: CompareExpr */
			nil,		/* < */
			nil,		/* > */
			nil,		/* false */
			nil,		/* in */
			nil,		/* int_lit */
			nil,		/* string_lit */
			nil,		/* true */
			reduce(11),		/* |, reduce: CompareExpr */
			
		},

	},
	actionRow{ // S20
				canRecover: false,
		actions: [numTerminals]action{
			nil,		/* INVALID */
			reduce(12),		/* $, reduce: SubStringExpr */
			nil,		/* error */
			reduce(12),		/* &, reduce: SubStringExpr */
			nil,		/* ( */
			reduce(12),		/* ), reduce: SubStringExpr */
			nil,		/* < */
			nil,		/* > */
			nil,		/* false */
			nil,		/* in */
			nil,		/* int_lit */
			nil,		/* string_lit */
			nil,		/* true */
			reduce(12),		/* |, reduce: SubStringExpr */
			
		},

	},
	actionRow{ // S21
				canRecover: false,
		actions: [numTerminals]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* error */
			shift(15),		/* & */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* < */
			nil,		/* > */
			nil,		/* false */
			nil,		/* in */
			nil,		/* int_lit */
			nil,		/* string_lit */
			nil,		/* true */
			shift(16),		/* | */
			
		},

	},
	actionRow{ // S22
				canRecover: false,
		actions: [numTerminals]action{
			nil,		/* INVALID */
			reduce(3),		/* $, reduce: BoolExpr1 */
			nil,		/* error */
			reduce(1),		/* &, reduce: BoolExpr */
			nil,		/* ( */
			reduce(3),		/* ), reduce: BoolExpr1 */
			nil,		/* < */
			nil,		/* > */
			nil,		/* false */
			nil,		/* in */
			nil,		/* int_lit */
			nil,		/* string_lit */
			nil,		/* true */
			reduce(1),		/* |, reduce: BoolExpr */
			
		},

	},
	actionRow{ // S23
				canRecover: false,
		actions: [numTerminals]action{
			nil,		/* INVALID */
			reduce(4),		/* $, reduce: BoolExpr1 */
			nil,		/* error */
			reduce(1),		/* &, reduce: BoolExpr */
			nil,		/* ( */
			reduce(4),		/* ), reduce: BoolExpr1 */
			nil,		/* < */
			nil,		/* > */
			nil,		/* false */
			nil,		/* in */
			nil,		/* int_lit */
			nil,		/* string_lit */
			nil,		/* true */
			reduce(1),		/* |, reduce: BoolExpr */
			
		},

	},
	
}

