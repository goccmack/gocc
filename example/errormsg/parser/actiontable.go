// Code generated by gocc; DO NOT EDIT.

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
		canRecover: false,
		actions: [numSymbols]action{
			nil,      // INVALID
			nil,      // ␚
			shift(2), // var
			nil,      // ;
			nil,      // identifier
			nil,      // _
			nil,      // =
			nil,      // :=
			nil,      // float
			nil,      // integer
		},
	},
	actionRow{ // S1
		canRecover: false,
		actions: [numSymbols]action{
			nil,          // INVALID
			accept(true), // ␚
			nil,          // var
			nil,          // ;
			nil,          // identifier
			nil,          // _
			nil,          // =
			nil,          // :=
			nil,          // float
			nil,          // integer
		},
	},
	actionRow{ // S2
		canRecover: false,
		actions: [numSymbols]action{
			nil,      // INVALID
			nil,      // ␚
			nil,      // var
			nil,      // ;
			shift(4), // identifier
			shift(5), // _
			nil,      // =
			nil,      // :=
			nil,      // float
			nil,      // integer
		},
	},
	actionRow{ // S3
		canRecover: false,
		actions: [numSymbols]action{
			nil,      // INVALID
			nil,      // ␚
			nil,      // var
			shift(7), // ;
			nil,      // identifier
			nil,      // _
			shift(8), // =
			shift(9), // :=
			nil,      // float
			nil,      // integer
		},
	},
	actionRow{ // S4
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // ␚
			nil,       // var
			reduce(3), // ;, reduce: Name
			nil,       // identifier
			nil,       // _
			reduce(3), // =, reduce: Name
			reduce(3), // :=, reduce: Name
			nil,       // float
			nil,       // integer
		},
	},
	actionRow{ // S5
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // ␚
			nil,       // var
			reduce(4), // ;, reduce: Name
			nil,       // identifier
			nil,       // _
			reduce(4), // =, reduce: Name
			reduce(4), // :=, reduce: Name
			nil,       // float
			nil,       // integer
		},
	},
	actionRow{ // S6
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // ␚
			nil,       // var
			nil,       // ;
			shift(11), // identifier
			shift(12), // _
			nil,       // =
			nil,       // :=
			shift(13), // float
			shift(14), // integer
		},
	},
	actionRow{ // S7
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			reduce(2), // ␚, reduce: Declaration
			nil,       // var
			nil,       // ;
			nil,       // identifier
			nil,       // _
			nil,       // =
			nil,       // :=
			nil,       // float
			nil,       // integer
		},
	},
	actionRow{ // S8
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // ␚
			nil,       // var
			nil,       // ;
			reduce(5), // identifier, reduce: Equal
			reduce(5), // _, reduce: Equal
			nil,       // =
			nil,       // :=
			reduce(5), // float, reduce: Equal
			reduce(5), // integer, reduce: Equal
		},
	},
	actionRow{ // S9
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // ␚
			nil,       // var
			nil,       // ;
			reduce(6), // identifier, reduce: Equal
			reduce(6), // _, reduce: Equal
			nil,       // =
			nil,       // :=
			reduce(6), // float, reduce: Equal
			reduce(6), // integer, reduce: Equal
		},
	},
	actionRow{ // S10
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // ␚
			nil,       // var
			shift(15), // ;
			nil,       // identifier
			nil,       // _
			nil,       // =
			nil,       // :=
			nil,       // float
			nil,       // integer
		},
	},
	actionRow{ // S11
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // ␚
			nil,       // var
			reduce(7), // ;, reduce: Default
			nil,       // identifier
			nil,       // _
			nil,       // =
			nil,       // :=
			nil,       // float
			nil,       // integer
		},
	},
	actionRow{ // S12
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // ␚
			nil,       // var
			reduce(8), // ;, reduce: Default
			nil,       // identifier
			nil,       // _
			nil,       // =
			nil,       // :=
			nil,       // float
			nil,       // integer
		},
	},
	actionRow{ // S13
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // ␚
			nil,       // var
			reduce(9), // ;, reduce: Default
			nil,       // identifier
			nil,       // _
			nil,       // =
			nil,       // :=
			nil,       // float
			nil,       // integer
		},
	},
	actionRow{ // S14
		canRecover: false,
		actions: [numSymbols]action{
			nil,        // INVALID
			nil,        // ␚
			nil,        // var
			reduce(10), // ;, reduce: Default
			nil,        // identifier
			nil,        // _
			nil,        // =
			nil,        // :=
			nil,        // float
			nil,        // integer
		},
	},
	actionRow{ // S15
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			reduce(1), // ␚, reduce: Declaration
			nil,       // var
			nil,       // ;
			nil,       // identifier
			nil,       // _
			nil,       // =
			nil,       // :=
			nil,       // float
			nil,       // integer
		},
	},
}
