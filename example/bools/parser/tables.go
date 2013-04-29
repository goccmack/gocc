
package parser

import ( "code.google.com/p/gocc/example/bools/ast" )

var ProductionsTable = ProdTab{
	// [0]
	ProdTabEntry{
		"S! : BoolExpr ;",
		"S!",
		1,
		func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	// [1]
	ProdTabEntry{
		"BoolExpr : BoolExpr1 << X[0], nil >> ;",
		"BoolExpr",
		1,
		func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	// [2]
	ProdTabEntry{
		"BoolExpr1 : Val << X[0], nil >> ;",
		"BoolExpr1",
		1,
		func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	// [3]
	ProdTabEntry{
		"BoolExpr1 : BoolExpr & BoolExpr1 << ast.NewBoolAndExpr(X[0], X[2]) >> ;",
		"BoolExpr1",
		3,
		func(X []Attrib) (Attrib, error) {
			return ast.NewBoolAndExpr(X[0], X[2])
		},
	},
	// [4]
	ProdTabEntry{
		"BoolExpr1 : BoolExpr | BoolExpr1 << ast.NewBoolOrExpr(X[0], X[2]) >> ;",
		"BoolExpr1",
		3,
		func(X []Attrib) (Attrib, error) {
			return ast.NewBoolOrExpr(X[0], X[2])
		},
	},
	// [5]
	ProdTabEntry{
		"BoolExpr1 : ( BoolExpr ) << ast.NewBoolGroupExpr(X[1]) >> ;",
		"BoolExpr1",
		3,
		func(X []Attrib) (Attrib, error) {
			return ast.NewBoolGroupExpr(X[1])
		},
	},
	// [6]
	ProdTabEntry{
		"Val : true << ast.TRUE, nil >> ;",
		"Val",
		1,
		func(X []Attrib) (Attrib, error) {
			return ast.TRUE, nil
		},
	},
	// [7]
	ProdTabEntry{
		"Val : false << ast.FALSE, nil >> ;",
		"Val",
		1,
		func(X []Attrib) (Attrib, error) {
			return ast.FALSE, nil
		},
	},
	// [8]
	ProdTabEntry{
		"Val : CompareExpr << X[0], nil >> ;",
		"Val",
		1,
		func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	// [9]
	ProdTabEntry{
		"Val : SubStringExpr << X[0], nil >> ;",
		"Val",
		1,
		func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	// [10]
	ProdTabEntry{
		"CompareExpr : int_lit < int_lit << ast.NewLessThanExpr(X[0], X[2]) >> ;",
		"CompareExpr",
		3,
		func(X []Attrib) (Attrib, error) {
			return ast.NewLessThanExpr(X[0], X[2])
		},
	},
	// [11]
	ProdTabEntry{
		"CompareExpr : int_lit > int_lit << ast.NewLessThanExpr(X[2], X[0]) >> ;",
		"CompareExpr",
		3,
		func(X []Attrib) (Attrib, error) {
			return ast.NewLessThanExpr(X[2], X[0])
		},
	},
	// [12]
	ProdTabEntry{
		"SubStringExpr : string_lit in string_lit << ast.NewSubStringExpr(X[0], X[2]) >> ;",
		"SubStringExpr",
		3,
		func(X []Attrib) (Attrib, error) {
			return ast.NewSubStringExpr(X[0], X[2])
		},
	},
}

var ActionTable ActionTab = ActionTab {
	// state 0
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			7: Shift(6), // false
			8: Shift(9), // int_lit
			6: Shift(5), // true
			11: Shift(10), // string_lit
			4: Shift(4), // (
		},
	},

	// state 1
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			3: Shift(12), // |
			2: Shift(11), // &
			0: Accept(0), // $
		},
	},

	// state 2
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			2: Reduce(1), // &
			3: Reduce(1), // |
			0: Reduce(1), // $
		},
	},

	// state 3
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			2: Reduce(2), // &
			3: Reduce(2), // |
			0: Reduce(2), // $
		},
	},

	// state 4
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			4: Shift(16), // (
			7: Shift(18), // false
			8: Shift(21), // int_lit
			11: Shift(22), // string_lit
			6: Shift(17), // true
		},
	},

	// state 5
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			2: Reduce(6), // &
			3: Reduce(6), // |
			0: Reduce(6), // $
		},
	},

	// state 6
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			0: Reduce(7), // $
			2: Reduce(7), // &
			3: Reduce(7), // |
		},
	},

	// state 7
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			3: Reduce(8), // |
			2: Reduce(8), // &
			0: Reduce(8), // $
		},
	},

	// state 8
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			0: Reduce(9), // $
			3: Reduce(9), // |
			2: Reduce(9), // &
		},
	},

	// state 9
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			9: Shift(23), // <
			10: Shift(24), // >
		},
	},

	// state 10
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			12: Shift(25), // in
		},
	},

	// state 11
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			4: Shift(4), // (
			7: Shift(6), // false
			6: Shift(5), // true
			8: Shift(9), // int_lit
			11: Shift(10), // string_lit
		},
	},

	// state 12
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			4: Shift(4), // (
			7: Shift(6), // false
			6: Shift(5), // true
			8: Shift(9), // int_lit
			11: Shift(10), // string_lit
		},
	},

	// state 13
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			3: Shift(30), // |
			2: Shift(29), // &
			5: Shift(31), // )
		},
	},

	// state 14
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			2: Reduce(1), // &
			3: Reduce(1), // |
			5: Reduce(1), // )
		},
	},

	// state 15
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			2: Reduce(2), // &
			3: Reduce(2), // |
			5: Reduce(2), // )
		},
	},

	// state 16
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			7: Shift(18), // false
			11: Shift(22), // string_lit
			6: Shift(17), // true
			4: Shift(16), // (
			8: Shift(21), // int_lit
		},
	},

	// state 17
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			5: Reduce(6), // )
			3: Reduce(6), // |
			2: Reduce(6), // &
		},
	},

	// state 18
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			5: Reduce(7), // )
			3: Reduce(7), // |
			2: Reduce(7), // &
		},
	},

	// state 19
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			2: Reduce(8), // &
			3: Reduce(8), // |
			5: Reduce(8), // )
		},
	},

	// state 20
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			2: Reduce(9), // &
			5: Reduce(9), // )
			3: Reduce(9), // |
		},
	},

	// state 21
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			9: Shift(33), // <
			10: Shift(34), // >
		},
	},

	// state 22
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			12: Shift(35), // in
		},
	},

	// state 23
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			8: Shift(36), // int_lit
		},
	},

	// state 24
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			8: Shift(37), // int_lit
		},
	},

	// state 25
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			11: Shift(38), // string_lit
		},
	},

	// state 26
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			2: Shift(11), // &
			3: Shift(12), // |
		},
	},

	// state 27
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			0: Reduce(3), // $
			2: Reduce(1), // &
			3: Reduce(1), // |
		},
	},

	// state 28
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			0: Reduce(4), // $
			2: Reduce(1), // &
			3: Reduce(1), // |
		},
	},

	// state 29
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			6: Shift(17), // true
			4: Shift(16), // (
			11: Shift(22), // string_lit
			8: Shift(21), // int_lit
			7: Shift(18), // false
		},
	},

	// state 30
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			6: Shift(17), // true
			4: Shift(16), // (
			11: Shift(22), // string_lit
			8: Shift(21), // int_lit
			7: Shift(18), // false
		},
	},

	// state 31
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			2: Reduce(5), // &
			0: Reduce(5), // $
			3: Reduce(5), // |
		},
	},

	// state 32
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			5: Shift(42), // )
			3: Shift(30), // |
			2: Shift(29), // &
		},
	},

	// state 33
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			8: Shift(43), // int_lit
		},
	},

	// state 34
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			8: Shift(44), // int_lit
		},
	},

	// state 35
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			11: Shift(45), // string_lit
		},
	},

	// state 36
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			3: Reduce(10), // |
			2: Reduce(10), // &
			0: Reduce(10), // $
		},
	},

	// state 37
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			2: Reduce(11), // &
			0: Reduce(11), // $
			3: Reduce(11), // |
		},
	},

	// state 38
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			2: Reduce(12), // &
			3: Reduce(12), // |
			0: Reduce(12), // $
		},
	},

	// state 39
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			2: Shift(29), // &
			3: Shift(30), // |
		},
	},

	// state 40
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			2: Reduce(1), // &
			3: Reduce(1), // |
			5: Reduce(3), // )
		},
	},

	// state 41
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			2: Reduce(1), // &
			5: Reduce(4), // )
			3: Reduce(1), // |
		},
	},

	// state 42
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			2: Reduce(5), // &
			5: Reduce(5), // )
			3: Reduce(5), // |
		},
	},

	// state 43
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			2: Reduce(10), // &
			5: Reduce(10), // )
			3: Reduce(10), // |
		},
	},

	// state 44
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			5: Reduce(11), // )
			3: Reduce(11), // |
			2: Reduce(11), // &
		},
	},

	// state 45
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			3: Reduce(12), // |
			5: Reduce(12), // )
			2: Reduce(12), // &
		},
	},

}

var GotoTable GotoTab = GotoTab{
	// state 0
	GotoRow{
		"BoolExpr1": State(2),
		"BoolExpr": State(1),
		"SubStringExpr": State(8),
		"CompareExpr": State(7),
		"Val": State(3),
	},
	// state 1
	GotoRow{
	},
	// state 2
	GotoRow{
	},
	// state 3
	GotoRow{
	},
	// state 4
	GotoRow{
		"SubStringExpr": State(20),
		"CompareExpr": State(19),
		"BoolExpr1": State(14),
		"Val": State(15),
		"BoolExpr": State(13),
	},
	// state 5
	GotoRow{
	},
	// state 6
	GotoRow{
	},
	// state 7
	GotoRow{
	},
	// state 8
	GotoRow{
	},
	// state 9
	GotoRow{
	},
	// state 10
	GotoRow{
	},
	// state 11
	GotoRow{
		"Val": State(3),
		"CompareExpr": State(7),
		"SubStringExpr": State(8),
		"BoolExpr": State(26),
		"BoolExpr1": State(27),
	},
	// state 12
	GotoRow{
		"BoolExpr1": State(28),
		"SubStringExpr": State(8),
		"CompareExpr": State(7),
		"Val": State(3),
		"BoolExpr": State(26),
	},
	// state 13
	GotoRow{
	},
	// state 14
	GotoRow{
	},
	// state 15
	GotoRow{
	},
	// state 16
	GotoRow{
		"Val": State(15),
		"BoolExpr1": State(14),
		"SubStringExpr": State(20),
		"CompareExpr": State(19),
		"BoolExpr": State(32),
	},
	// state 17
	GotoRow{
	},
	// state 18
	GotoRow{
	},
	// state 19
	GotoRow{
	},
	// state 20
	GotoRow{
	},
	// state 21
	GotoRow{
	},
	// state 22
	GotoRow{
	},
	// state 23
	GotoRow{
	},
	// state 24
	GotoRow{
	},
	// state 25
	GotoRow{
	},
	// state 26
	GotoRow{
	},
	// state 27
	GotoRow{
	},
	// state 28
	GotoRow{
	},
	// state 29
	GotoRow{
		"BoolExpr1": State(40),
		"SubStringExpr": State(20),
		"BoolExpr": State(39),
		"CompareExpr": State(19),
		"Val": State(15),
	},
	// state 30
	GotoRow{
		"Val": State(15),
		"CompareExpr": State(19),
		"BoolExpr": State(39),
		"BoolExpr1": State(41),
		"SubStringExpr": State(20),
	},
	// state 31
	GotoRow{
	},
	// state 32
	GotoRow{
	},
	// state 33
	GotoRow{
	},
	// state 34
	GotoRow{
	},
	// state 35
	GotoRow{
	},
	// state 36
	GotoRow{
	},
	// state 37
	GotoRow{
	},
	// state 38
	GotoRow{
	},
	// state 39
	GotoRow{
	},
	// state 40
	GotoRow{
	},
	// state 41
	GotoRow{
	},
	// state 42
	GotoRow{
	},
	// state 43
	GotoRow{
	},
	// state 44
	GotoRow{
	},
	// state 45
	GotoRow{
	},
}

