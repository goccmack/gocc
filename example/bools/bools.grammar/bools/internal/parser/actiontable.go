// Code generated by gocc; DO NOT EDIT.

package parser

type (
	actionTable [numStates]actionRow
	cdFunc      func(TokenStream, interface{}) (interface{}, error, int)
	cdAction    struct {
		tokenIndex   int
		tokenScanner cdFunc
	}
	actionRow struct {
		canRecover bool
		actions    [numSymbols]action
		cdActions  []cdAction
	}
	actions struct {
		table actionTable
	}
)

var parserActions = actions{
	table: actionTable{
		actionRow{ // S0
			canRecover: false,
			actions: [numSymbols]action{
				nil,       // ά<INVALID>
				nil,       // Ω<EOF>
				nil,       // "&"
				nil,       // "|"
				shift(6),  // "("
				nil,       // ")"
				shift(7),  // "true"
				shift(8),  // "false"
				shift(9),  // int_lit
				nil,       // "<"
				nil,       // ">"
				shift(10), // string_lit
				nil,       // "in"
			},
			cdActions: []cdAction{},
		},
		actionRow{ // S1
			canRecover: false,
			actions: [numSymbols]action{
				nil,          // ά<INVALID>
				accept(true), // Ω<EOF>
				shift(11),    // "&"
				shift(12),    // "|"
				nil,          // "("
				nil,          // ")"
				nil,          // "true"
				nil,          // "false"
				nil,          // int_lit
				nil,          // "<"
				nil,          // ">"
				nil,          // string_lit
				nil,          // "in"
			},
			cdActions: []cdAction{},
		},
		actionRow{ // S2
			canRecover: false,
			actions: [numSymbols]action{
				nil,       // ά<INVALID>
				reduce(1), // Ω<EOF>, reduce: BoolExpr
				reduce(1), // "&", reduce: BoolExpr
				reduce(1), // "|", reduce: BoolExpr
				nil,       // "("
				nil,       // ")"
				nil,       // "true"
				nil,       // "false"
				nil,       // int_lit
				nil,       // "<"
				nil,       // ">"
				nil,       // string_lit
				nil,       // "in"
			},
			cdActions: []cdAction{},
		},
		actionRow{ // S3
			canRecover: false,
			actions: [numSymbols]action{
				nil,       // ά<INVALID>
				reduce(2), // Ω<EOF>, reduce: BoolExpr1
				reduce(2), // "&", reduce: BoolExpr1
				reduce(2), // "|", reduce: BoolExpr1
				nil,       // "("
				nil,       // ")"
				nil,       // "true"
				nil,       // "false"
				nil,       // int_lit
				nil,       // "<"
				nil,       // ">"
				nil,       // string_lit
				nil,       // "in"
			},
			cdActions: []cdAction{},
		},
		actionRow{ // S4
			canRecover: false,
			actions: [numSymbols]action{
				nil,       // ά<INVALID>
				reduce(8), // Ω<EOF>, reduce: Val
				reduce(8), // "&", reduce: Val
				reduce(8), // "|", reduce: Val
				nil,       // "("
				nil,       // ")"
				nil,       // "true"
				nil,       // "false"
				nil,       // int_lit
				nil,       // "<"
				nil,       // ">"
				nil,       // string_lit
				nil,       // "in"
			},
			cdActions: []cdAction{},
		},
		actionRow{ // S5
			canRecover: false,
			actions: [numSymbols]action{
				nil,       // ά<INVALID>
				reduce(9), // Ω<EOF>, reduce: Val
				reduce(9), // "&", reduce: Val
				reduce(9), // "|", reduce: Val
				nil,       // "("
				nil,       // ")"
				nil,       // "true"
				nil,       // "false"
				nil,       // int_lit
				nil,       // "<"
				nil,       // ">"
				nil,       // string_lit
				nil,       // "in"
			},
			cdActions: []cdAction{},
		},
		actionRow{ // S6
			canRecover: false,
			actions: [numSymbols]action{
				nil,       // ά<INVALID>
				nil,       // Ω<EOF>
				nil,       // "&"
				nil,       // "|"
				shift(18), // "("
				nil,       // ")"
				shift(19), // "true"
				shift(20), // "false"
				shift(21), // int_lit
				nil,       // "<"
				nil,       // ">"
				shift(22), // string_lit
				nil,       // "in"
			},
			cdActions: []cdAction{},
		},
		actionRow{ // S7
			canRecover: false,
			actions: [numSymbols]action{
				nil,       // ά<INVALID>
				reduce(6), // Ω<EOF>, reduce: Val
				reduce(6), // "&", reduce: Val
				reduce(6), // "|", reduce: Val
				nil,       // "("
				nil,       // ")"
				nil,       // "true"
				nil,       // "false"
				nil,       // int_lit
				nil,       // "<"
				nil,       // ">"
				nil,       // string_lit
				nil,       // "in"
			},
			cdActions: []cdAction{},
		},
		actionRow{ // S8
			canRecover: false,
			actions: [numSymbols]action{
				nil,       // ά<INVALID>
				reduce(7), // Ω<EOF>, reduce: Val
				reduce(7), // "&", reduce: Val
				reduce(7), // "|", reduce: Val
				nil,       // "("
				nil,       // ")"
				nil,       // "true"
				nil,       // "false"
				nil,       // int_lit
				nil,       // "<"
				nil,       // ">"
				nil,       // string_lit
				nil,       // "in"
			},
			cdActions: []cdAction{},
		},
		actionRow{ // S9
			canRecover: false,
			actions: [numSymbols]action{
				nil,       // ά<INVALID>
				nil,       // Ω<EOF>
				nil,       // "&"
				nil,       // "|"
				nil,       // "("
				nil,       // ")"
				nil,       // "true"
				nil,       // "false"
				nil,       // int_lit
				shift(23), // "<"
				shift(24), // ">"
				nil,       // string_lit
				nil,       // "in"
			},
			cdActions: []cdAction{},
		},
		actionRow{ // S10
			canRecover: false,
			actions: [numSymbols]action{
				nil,       // ά<INVALID>
				nil,       // Ω<EOF>
				nil,       // "&"
				nil,       // "|"
				nil,       // "("
				nil,       // ")"
				nil,       // "true"
				nil,       // "false"
				nil,       // int_lit
				nil,       // "<"
				nil,       // ">"
				nil,       // string_lit
				shift(25), // "in"
			},
			cdActions: []cdAction{},
		},
		actionRow{ // S11
			canRecover: false,
			actions: [numSymbols]action{
				nil,       // ά<INVALID>
				nil,       // Ω<EOF>
				nil,       // "&"
				nil,       // "|"
				shift(6),  // "("
				nil,       // ")"
				shift(7),  // "true"
				shift(8),  // "false"
				shift(9),  // int_lit
				nil,       // "<"
				nil,       // ">"
				shift(10), // string_lit
				nil,       // "in"
			},
			cdActions: []cdAction{},
		},
		actionRow{ // S12
			canRecover: false,
			actions: [numSymbols]action{
				nil,       // ά<INVALID>
				nil,       // Ω<EOF>
				nil,       // "&"
				nil,       // "|"
				shift(6),  // "("
				nil,       // ")"
				shift(7),  // "true"
				shift(8),  // "false"
				shift(9),  // int_lit
				nil,       // "<"
				nil,       // ">"
				shift(10), // string_lit
				nil,       // "in"
			},
			cdActions: []cdAction{},
		},
		actionRow{ // S13
			canRecover: false,
			actions: [numSymbols]action{
				nil,       // ά<INVALID>
				nil,       // Ω<EOF>
				shift(29), // "&"
				shift(30), // "|"
				nil,       // "("
				shift(31), // ")"
				nil,       // "true"
				nil,       // "false"
				nil,       // int_lit
				nil,       // "<"
				nil,       // ">"
				nil,       // string_lit
				nil,       // "in"
			},
			cdActions: []cdAction{},
		},
		actionRow{ // S14
			canRecover: false,
			actions: [numSymbols]action{
				nil,       // ά<INVALID>
				nil,       // Ω<EOF>
				reduce(1), // "&", reduce: BoolExpr
				reduce(1), // "|", reduce: BoolExpr
				nil,       // "("
				reduce(1), // ")", reduce: BoolExpr
				nil,       // "true"
				nil,       // "false"
				nil,       // int_lit
				nil,       // "<"
				nil,       // ">"
				nil,       // string_lit
				nil,       // "in"
			},
			cdActions: []cdAction{},
		},
		actionRow{ // S15
			canRecover: false,
			actions: [numSymbols]action{
				nil,       // ά<INVALID>
				nil,       // Ω<EOF>
				reduce(2), // "&", reduce: BoolExpr1
				reduce(2), // "|", reduce: BoolExpr1
				nil,       // "("
				reduce(2), // ")", reduce: BoolExpr1
				nil,       // "true"
				nil,       // "false"
				nil,       // int_lit
				nil,       // "<"
				nil,       // ">"
				nil,       // string_lit
				nil,       // "in"
			},
			cdActions: []cdAction{},
		},
		actionRow{ // S16
			canRecover: false,
			actions: [numSymbols]action{
				nil,       // ά<INVALID>
				nil,       // Ω<EOF>
				reduce(8), // "&", reduce: Val
				reduce(8), // "|", reduce: Val
				nil,       // "("
				reduce(8), // ")", reduce: Val
				nil,       // "true"
				nil,       // "false"
				nil,       // int_lit
				nil,       // "<"
				nil,       // ">"
				nil,       // string_lit
				nil,       // "in"
			},
			cdActions: []cdAction{},
		},
		actionRow{ // S17
			canRecover: false,
			actions: [numSymbols]action{
				nil,       // ά<INVALID>
				nil,       // Ω<EOF>
				reduce(9), // "&", reduce: Val
				reduce(9), // "|", reduce: Val
				nil,       // "("
				reduce(9), // ")", reduce: Val
				nil,       // "true"
				nil,       // "false"
				nil,       // int_lit
				nil,       // "<"
				nil,       // ">"
				nil,       // string_lit
				nil,       // "in"
			},
			cdActions: []cdAction{},
		},
		actionRow{ // S18
			canRecover: false,
			actions: [numSymbols]action{
				nil,       // ά<INVALID>
				nil,       // Ω<EOF>
				nil,       // "&"
				nil,       // "|"
				shift(18), // "("
				nil,       // ")"
				shift(19), // "true"
				shift(20), // "false"
				shift(21), // int_lit
				nil,       // "<"
				nil,       // ">"
				shift(22), // string_lit
				nil,       // "in"
			},
			cdActions: []cdAction{},
		},
		actionRow{ // S19
			canRecover: false,
			actions: [numSymbols]action{
				nil,       // ά<INVALID>
				nil,       // Ω<EOF>
				reduce(6), // "&", reduce: Val
				reduce(6), // "|", reduce: Val
				nil,       // "("
				reduce(6), // ")", reduce: Val
				nil,       // "true"
				nil,       // "false"
				nil,       // int_lit
				nil,       // "<"
				nil,       // ">"
				nil,       // string_lit
				nil,       // "in"
			},
			cdActions: []cdAction{},
		},
		actionRow{ // S20
			canRecover: false,
			actions: [numSymbols]action{
				nil,       // ά<INVALID>
				nil,       // Ω<EOF>
				reduce(7), // "&", reduce: Val
				reduce(7), // "|", reduce: Val
				nil,       // "("
				reduce(7), // ")", reduce: Val
				nil,       // "true"
				nil,       // "false"
				nil,       // int_lit
				nil,       // "<"
				nil,       // ">"
				nil,       // string_lit
				nil,       // "in"
			},
			cdActions: []cdAction{},
		},
		actionRow{ // S21
			canRecover: false,
			actions: [numSymbols]action{
				nil,       // ά<INVALID>
				nil,       // Ω<EOF>
				nil,       // "&"
				nil,       // "|"
				nil,       // "("
				nil,       // ")"
				nil,       // "true"
				nil,       // "false"
				nil,       // int_lit
				shift(33), // "<"
				shift(34), // ">"
				nil,       // string_lit
				nil,       // "in"
			},
			cdActions: []cdAction{},
		},
		actionRow{ // S22
			canRecover: false,
			actions: [numSymbols]action{
				nil,       // ά<INVALID>
				nil,       // Ω<EOF>
				nil,       // "&"
				nil,       // "|"
				nil,       // "("
				nil,       // ")"
				nil,       // "true"
				nil,       // "false"
				nil,       // int_lit
				nil,       // "<"
				nil,       // ">"
				nil,       // string_lit
				shift(35), // "in"
			},
			cdActions: []cdAction{},
		},
		actionRow{ // S23
			canRecover: false,
			actions: [numSymbols]action{
				nil,       // ά<INVALID>
				nil,       // Ω<EOF>
				nil,       // "&"
				nil,       // "|"
				nil,       // "("
				nil,       // ")"
				nil,       // "true"
				nil,       // "false"
				shift(36), // int_lit
				nil,       // "<"
				nil,       // ">"
				nil,       // string_lit
				nil,       // "in"
			},
			cdActions: []cdAction{},
		},
		actionRow{ // S24
			canRecover: false,
			actions: [numSymbols]action{
				nil,       // ά<INVALID>
				nil,       // Ω<EOF>
				nil,       // "&"
				nil,       // "|"
				nil,       // "("
				nil,       // ")"
				nil,       // "true"
				nil,       // "false"
				shift(37), // int_lit
				nil,       // "<"
				nil,       // ">"
				nil,       // string_lit
				nil,       // "in"
			},
			cdActions: []cdAction{},
		},
		actionRow{ // S25
			canRecover: false,
			actions: [numSymbols]action{
				nil,       // ά<INVALID>
				nil,       // Ω<EOF>
				nil,       // "&"
				nil,       // "|"
				nil,       // "("
				nil,       // ")"
				nil,       // "true"
				nil,       // "false"
				nil,       // int_lit
				nil,       // "<"
				nil,       // ">"
				shift(38), // string_lit
				nil,       // "in"
			},
			cdActions: []cdAction{},
		},
		actionRow{ // S26
			canRecover: false,
			actions: [numSymbols]action{
				nil,       // ά<INVALID>
				nil,       // Ω<EOF>
				shift(11), // "&"
				shift(12), // "|"
				nil,       // "("
				nil,       // ")"
				nil,       // "true"
				nil,       // "false"
				nil,       // int_lit
				nil,       // "<"
				nil,       // ">"
				nil,       // string_lit
				nil,       // "in"
			},
			cdActions: []cdAction{},
		},
		actionRow{ // S27
			canRecover: false,
			actions: [numSymbols]action{
				nil,       // ά<INVALID>
				reduce(3), // Ω<EOF>, reduce: BoolExpr1
				reduce(1), // "&", reduce: BoolExpr
				reduce(1), // "|", reduce: BoolExpr
				nil,       // "("
				nil,       // ")"
				nil,       // "true"
				nil,       // "false"
				nil,       // int_lit
				nil,       // "<"
				nil,       // ">"
				nil,       // string_lit
				nil,       // "in"
			},
			cdActions: []cdAction{},
		},
		actionRow{ // S28
			canRecover: false,
			actions: [numSymbols]action{
				nil,       // ά<INVALID>
				reduce(4), // Ω<EOF>, reduce: BoolExpr1
				reduce(1), // "&", reduce: BoolExpr
				reduce(1), // "|", reduce: BoolExpr
				nil,       // "("
				nil,       // ")"
				nil,       // "true"
				nil,       // "false"
				nil,       // int_lit
				nil,       // "<"
				nil,       // ">"
				nil,       // string_lit
				nil,       // "in"
			},
			cdActions: []cdAction{},
		},
		actionRow{ // S29
			canRecover: false,
			actions: [numSymbols]action{
				nil,       // ά<INVALID>
				nil,       // Ω<EOF>
				nil,       // "&"
				nil,       // "|"
				shift(18), // "("
				nil,       // ")"
				shift(19), // "true"
				shift(20), // "false"
				shift(21), // int_lit
				nil,       // "<"
				nil,       // ">"
				shift(22), // string_lit
				nil,       // "in"
			},
			cdActions: []cdAction{},
		},
		actionRow{ // S30
			canRecover: false,
			actions: [numSymbols]action{
				nil,       // ά<INVALID>
				nil,       // Ω<EOF>
				nil,       // "&"
				nil,       // "|"
				shift(18), // "("
				nil,       // ")"
				shift(19), // "true"
				shift(20), // "false"
				shift(21), // int_lit
				nil,       // "<"
				nil,       // ">"
				shift(22), // string_lit
				nil,       // "in"
			},
			cdActions: []cdAction{},
		},
		actionRow{ // S31
			canRecover: false,
			actions: [numSymbols]action{
				nil,       // ά<INVALID>
				reduce(5), // Ω<EOF>, reduce: BoolExpr1
				reduce(5), // "&", reduce: BoolExpr1
				reduce(5), // "|", reduce: BoolExpr1
				nil,       // "("
				nil,       // ")"
				nil,       // "true"
				nil,       // "false"
				nil,       // int_lit
				nil,       // "<"
				nil,       // ">"
				nil,       // string_lit
				nil,       // "in"
			},
			cdActions: []cdAction{},
		},
		actionRow{ // S32
			canRecover: false,
			actions: [numSymbols]action{
				nil,       // ά<INVALID>
				nil,       // Ω<EOF>
				shift(29), // "&"
				shift(30), // "|"
				nil,       // "("
				shift(42), // ")"
				nil,       // "true"
				nil,       // "false"
				nil,       // int_lit
				nil,       // "<"
				nil,       // ">"
				nil,       // string_lit
				nil,       // "in"
			},
			cdActions: []cdAction{},
		},
		actionRow{ // S33
			canRecover: false,
			actions: [numSymbols]action{
				nil,       // ά<INVALID>
				nil,       // Ω<EOF>
				nil,       // "&"
				nil,       // "|"
				nil,       // "("
				nil,       // ")"
				nil,       // "true"
				nil,       // "false"
				shift(43), // int_lit
				nil,       // "<"
				nil,       // ">"
				nil,       // string_lit
				nil,       // "in"
			},
			cdActions: []cdAction{},
		},
		actionRow{ // S34
			canRecover: false,
			actions: [numSymbols]action{
				nil,       // ά<INVALID>
				nil,       // Ω<EOF>
				nil,       // "&"
				nil,       // "|"
				nil,       // "("
				nil,       // ")"
				nil,       // "true"
				nil,       // "false"
				shift(44), // int_lit
				nil,       // "<"
				nil,       // ">"
				nil,       // string_lit
				nil,       // "in"
			},
			cdActions: []cdAction{},
		},
		actionRow{ // S35
			canRecover: false,
			actions: [numSymbols]action{
				nil,       // ά<INVALID>
				nil,       // Ω<EOF>
				nil,       // "&"
				nil,       // "|"
				nil,       // "("
				nil,       // ")"
				nil,       // "true"
				nil,       // "false"
				nil,       // int_lit
				nil,       // "<"
				nil,       // ">"
				shift(45), // string_lit
				nil,       // "in"
			},
			cdActions: []cdAction{},
		},
		actionRow{ // S36
			canRecover: false,
			actions: [numSymbols]action{
				nil,        // ά<INVALID>
				reduce(10), // Ω<EOF>, reduce: CompareExpr
				reduce(10), // "&", reduce: CompareExpr
				reduce(10), // "|", reduce: CompareExpr
				nil,        // "("
				nil,        // ")"
				nil,        // "true"
				nil,        // "false"
				nil,        // int_lit
				nil,        // "<"
				nil,        // ">"
				nil,        // string_lit
				nil,        // "in"
			},
			cdActions: []cdAction{},
		},
		actionRow{ // S37
			canRecover: false,
			actions: [numSymbols]action{
				nil,        // ά<INVALID>
				reduce(11), // Ω<EOF>, reduce: CompareExpr
				reduce(11), // "&", reduce: CompareExpr
				reduce(11), // "|", reduce: CompareExpr
				nil,        // "("
				nil,        // ")"
				nil,        // "true"
				nil,        // "false"
				nil,        // int_lit
				nil,        // "<"
				nil,        // ">"
				nil,        // string_lit
				nil,        // "in"
			},
			cdActions: []cdAction{},
		},
		actionRow{ // S38
			canRecover: false,
			actions: [numSymbols]action{
				nil,        // ά<INVALID>
				reduce(12), // Ω<EOF>, reduce: SubStringExpr
				reduce(12), // "&", reduce: SubStringExpr
				reduce(12), // "|", reduce: SubStringExpr
				nil,        // "("
				nil,        // ")"
				nil,        // "true"
				nil,        // "false"
				nil,        // int_lit
				nil,        // "<"
				nil,        // ">"
				nil,        // string_lit
				nil,        // "in"
			},
			cdActions: []cdAction{},
		},
		actionRow{ // S39
			canRecover: false,
			actions: [numSymbols]action{
				nil,       // ά<INVALID>
				nil,       // Ω<EOF>
				shift(29), // "&"
				shift(30), // "|"
				nil,       // "("
				nil,       // ")"
				nil,       // "true"
				nil,       // "false"
				nil,       // int_lit
				nil,       // "<"
				nil,       // ">"
				nil,       // string_lit
				nil,       // "in"
			},
			cdActions: []cdAction{},
		},
		actionRow{ // S40
			canRecover: false,
			actions: [numSymbols]action{
				nil,       // ά<INVALID>
				nil,       // Ω<EOF>
				reduce(1), // "&", reduce: BoolExpr
				reduce(1), // "|", reduce: BoolExpr
				nil,       // "("
				reduce(3), // ")", reduce: BoolExpr1
				nil,       // "true"
				nil,       // "false"
				nil,       // int_lit
				nil,       // "<"
				nil,       // ">"
				nil,       // string_lit
				nil,       // "in"
			},
			cdActions: []cdAction{},
		},
		actionRow{ // S41
			canRecover: false,
			actions: [numSymbols]action{
				nil,       // ά<INVALID>
				nil,       // Ω<EOF>
				reduce(1), // "&", reduce: BoolExpr
				reduce(1), // "|", reduce: BoolExpr
				nil,       // "("
				reduce(4), // ")", reduce: BoolExpr1
				nil,       // "true"
				nil,       // "false"
				nil,       // int_lit
				nil,       // "<"
				nil,       // ">"
				nil,       // string_lit
				nil,       // "in"
			},
			cdActions: []cdAction{},
		},
		actionRow{ // S42
			canRecover: false,
			actions: [numSymbols]action{
				nil,       // ά<INVALID>
				nil,       // Ω<EOF>
				reduce(5), // "&", reduce: BoolExpr1
				reduce(5), // "|", reduce: BoolExpr1
				nil,       // "("
				reduce(5), // ")", reduce: BoolExpr1
				nil,       // "true"
				nil,       // "false"
				nil,       // int_lit
				nil,       // "<"
				nil,       // ">"
				nil,       // string_lit
				nil,       // "in"
			},
			cdActions: []cdAction{},
		},
		actionRow{ // S43
			canRecover: false,
			actions: [numSymbols]action{
				nil,        // ά<INVALID>
				nil,        // Ω<EOF>
				reduce(10), // "&", reduce: CompareExpr
				reduce(10), // "|", reduce: CompareExpr
				nil,        // "("
				reduce(10), // ")", reduce: CompareExpr
				nil,        // "true"
				nil,        // "false"
				nil,        // int_lit
				nil,        // "<"
				nil,        // ">"
				nil,        // string_lit
				nil,        // "in"
			},
			cdActions: []cdAction{},
		},
		actionRow{ // S44
			canRecover: false,
			actions: [numSymbols]action{
				nil,        // ά<INVALID>
				nil,        // Ω<EOF>
				reduce(11), // "&", reduce: CompareExpr
				reduce(11), // "|", reduce: CompareExpr
				nil,        // "("
				reduce(11), // ")", reduce: CompareExpr
				nil,        // "true"
				nil,        // "false"
				nil,        // int_lit
				nil,        // "<"
				nil,        // ">"
				nil,        // string_lit
				nil,        // "in"
			},
			cdActions: []cdAction{},
		},
		actionRow{ // S45
			canRecover: false,
			actions: [numSymbols]action{
				nil,        // ά<INVALID>
				nil,        // Ω<EOF>
				reduce(12), // "&", reduce: SubStringExpr
				reduce(12), // "|", reduce: SubStringExpr
				nil,        // "("
				reduce(12), // ")", reduce: SubStringExpr
				nil,        // "true"
				nil,        // "false"
				nil,        // int_lit
				nil,        // "<"
				nil,        // ">"
				nil,        // string_lit
				nil,        // "in"
			},
			cdActions: []cdAction{},
		},
	},
}
