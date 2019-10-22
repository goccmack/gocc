// Code generated by gocc; DO NOT EDIT.

package parser

type (
	actionTable [numStates]actionRow
	cdFunc      func(TokenStream, interface{}) (interface{}, error, []byte)
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
				nil,      // ά<INVALID>
				nil,      // Ω<EOF>
				nil,      // ";"
				shift(5), // "let"
				shift(6), // "define"
				nil,      // id
				nil,      // "="
				nil,      // μ<calc_0>
			},
			cdActions: []cdAction{},
		},
		actionRow{ // S1
			canRecover: false,
			actions: [numSymbols]action{
				nil,          // ά<INVALID>
				accept(true), // Ω<EOF>
				nil,          // ";"
				nil,          // "let"
				nil,          // "define"
				nil,          // id
				nil,          // "="
				nil,          // μ<calc_0>
			},
			cdActions: []cdAction{},
		},
		actionRow{ // S2
			canRecover: false,
			actions: [numSymbols]action{
				nil,       // ά<INVALID>
				reduce(1), // Ω<EOF>, reduce: StmtList
				shift(7),  // ";"
				nil,       // "let"
				nil,       // "define"
				nil,       // id
				nil,       // "="
				nil,       // μ<calc_0>
			},
			cdActions: []cdAction{},
		},
		actionRow{ // S3
			canRecover: false,
			actions: [numSymbols]action{
				nil,       // ά<INVALID>
				reduce(2), // Ω<EOF>, reduce: Stmt_s
				reduce(2), // ";", reduce: Stmt_s
				nil,       // "let"
				nil,       // "define"
				nil,       // id
				nil,       // "="
				nil,       // μ<calc_0>
			},
			cdActions: []cdAction{},
		},
		actionRow{ // S4
			canRecover: false,
			actions: [numSymbols]action{
				nil,      // ά<INVALID>
				nil,      // Ω<EOF>
				nil,      // ";"
				nil,      // "let"
				nil,      // "define"
				shift(8), // id
				nil,      // "="
				nil,      // μ<calc_0>
			},
			cdActions: []cdAction{},
		},
		actionRow{ // S5
			canRecover: false,
			actions: [numSymbols]action{
				nil,       // ά<INVALID>
				nil,       // Ω<EOF>
				nil,       // ";"
				nil,       // "let"
				nil,       // "define"
				reduce(4), // id, reduce: StmtPref
				nil,       // "="
				nil,       // μ<calc_0>
			},
			cdActions: []cdAction{},
		},
		actionRow{ // S6
			canRecover: false,
			actions: [numSymbols]action{
				nil,       // ά<INVALID>
				nil,       // Ω<EOF>
				nil,       // ";"
				nil,       // "let"
				nil,       // "define"
				reduce(5), // id, reduce: StmtPref
				nil,       // "="
				nil,       // μ<calc_0>
			},
			cdActions: []cdAction{},
		},
		actionRow{ // S7
			canRecover: false,
			actions: [numSymbols]action{
				nil,      // ά<INVALID>
				nil,      // Ω<EOF>
				nil,      // ";"
				shift(5), // "let"
				shift(6), // "define"
				nil,      // id
				nil,      // "="
				nil,      // μ<calc_0>
			},
			cdActions: []cdAction{},
		},
		actionRow{ // S8
			canRecover: false,
			actions: [numSymbols]action{
				nil,       // ά<INVALID>
				nil,       // Ω<EOF>
				nil,       // ";"
				nil,       // "let"
				nil,       // "define"
				nil,       // id
				shift(10), // "="
				nil,       // μ<calc_0>
			},
			cdActions: []cdAction{},
		},
		actionRow{ // S9
			canRecover: false,
			actions: [numSymbols]action{
				nil,       // ά<INVALID>
				reduce(3), // Ω<EOF>, reduce: Stmt_s
				reduce(3), // ";", reduce: Stmt_s
				nil,       // "let"
				nil,       // "define"
				nil,       // id
				nil,       // "="
				nil,       // μ<calc_0>
			},
			cdActions: []cdAction{},
		},
		actionRow{ // S10
			canRecover: false,
			actions: [numSymbols]action{
				nil,       // ά<INVALID>
				nil,       // Ω<EOF>
				nil,       // ";"
				nil,       // "let"
				nil,       // "define"
				nil,       // id
				nil,       // "="
				shift(11), // μ<calc_0>
			},
			cdActions: []cdAction{
				cdAction{tokenIndex: 7, tokenScanner: cdFunc_calc_0},
			},
		},
		actionRow{ // S11
			canRecover: false,
			actions: [numSymbols]action{
				nil,       // ά<INVALID>
				reduce(6), // Ω<EOF>, reduce: Stmt
				reduce(6), // ";", reduce: Stmt
				nil,       // "let"
				nil,       // "define"
				nil,       // id
				nil,       // "="
				nil,       // μ<calc_0>
			},
			cdActions: []cdAction{},
		},
	},
}
