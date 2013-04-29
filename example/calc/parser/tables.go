package parser

import "code.google.com/p/gocc/example/calc/token"

var ProductionsTable = ProdTab{
	// [0]
	ProdTabEntry{
		"S! : Calc ;",
		"S!",
		1,
		func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	// [1]
	ProdTabEntry{
		"Calc : Expr ;",
		"Calc",
		1,
		func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	// [2]
	ProdTabEntry{
		"Expr : Expr + Term << X[0].(int64) + X[2].(int64), nil >> ;",
		"Expr",
		3,
		func(X []Attrib) (Attrib, error) {
			return X[0].(int64) + X[2].(int64), nil
		},
	},
	// [3]
	ProdTabEntry{
		"Expr : Term ;",
		"Expr",
		1,
		func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	// [4]
	ProdTabEntry{
		"Term : Term * Factor << X[0].(int64) * X[2].(int64), nil >> ;",
		"Term",
		3,
		func(X []Attrib) (Attrib, error) {
			return X[0].(int64) * X[2].(int64), nil
		},
	},
	// [5]
	ProdTabEntry{
		"Term : Factor ;",
		"Term",
		1,
		func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	// [6]
	ProdTabEntry{
		"Factor : ( Expr ) << X[1], nil >> ;",
		"Factor",
		3,
		func(X []Attrib) (Attrib, error) {
			return X[1], nil
		},
	},
	// [7]
	ProdTabEntry{
		"Factor : int_lit << X[0].(*token.Token).IntValue() >> ;",
		"Factor",
		1,
		func(X []Attrib) (Attrib, error) {
			return X[0].(*token.Token).IntValue()
		},
	},
}

var ActionTable ActionTab = ActionTab{
	// state 0
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			6: Shift(6), // int_lit
			4: Shift(5), // (
		},
	},

	// state 1
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			0: Accept(0), // $
		},
	},

	// state 2
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			0: Reduce(1), // $
			2: Shift(7),  // +
		},
	},

	// state 3
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			0: Reduce(3), // $
			3: Shift(8),  // *
			2: Reduce(3), // +
		},
	},

	// state 4
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			2: Reduce(5), // +
			0: Reduce(5), // $
			3: Reduce(5), // *
		},
	},

	// state 5
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			6: Shift(13), // int_lit
			4: Shift(12), // (
		},
	},

	// state 6
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			2: Reduce(7), // +
			3: Reduce(7), // *
			0: Reduce(7), // $
		},
	},

	// state 7
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			4: Shift(5), // (
			6: Shift(6), // int_lit
		},
	},

	// state 8
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			4: Shift(5), // (
			6: Shift(6), // int_lit
		},
	},

	// state 9
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			2: Shift(16), // +
			5: Shift(17), // )
		},
	},

	// state 10
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			3: Shift(18), // *
			2: Reduce(3), // +
			5: Reduce(3), // )
		},
	},

	// state 11
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			5: Reduce(5), // )
			3: Reduce(5), // *
			2: Reduce(5), // +
		},
	},

	// state 12
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			4: Shift(12), // (
			6: Shift(13), // int_lit
		},
	},

	// state 13
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			5: Reduce(7), // )
			2: Reduce(7), // +
			3: Reduce(7), // *
		},
	},

	// state 14
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			0: Reduce(2), // $
			3: Shift(8),  // *
			2: Reduce(2), // +
		},
	},

	// state 15
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			3: Reduce(4), // *
			0: Reduce(4), // $
			2: Reduce(4), // +
		},
	},

	// state 16
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			6: Shift(13), // int_lit
			4: Shift(12), // (
		},
	},

	// state 17
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			2: Reduce(6), // +
			3: Reduce(6), // *
			0: Reduce(6), // $
		},
	},

	// state 18
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			4: Shift(12), // (
			6: Shift(13), // int_lit
		},
	},

	// state 19
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			5: Shift(22), // )
			2: Shift(16), // +
		},
	},

	// state 20
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			2: Reduce(2), // +
			5: Reduce(2), // )
			3: Shift(18), // *
		},
	},

	// state 21
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			3: Reduce(4), // *
			2: Reduce(4), // +
			5: Reduce(4), // )
		},
	},

	// state 22
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			2: Reduce(6), // +
			3: Reduce(6), // *
			5: Reduce(6), // )
		},
	},
}

var GotoTable GotoTab = GotoTab{
	// state 0
	GotoRow{
		"Factor": State(4),
		"Calc":   State(1),
		"Term":   State(3),
		"Expr":   State(2),
	},
	// state 1
	GotoRow{},
	// state 2
	GotoRow{},
	// state 3
	GotoRow{},
	// state 4
	GotoRow{},
	// state 5
	GotoRow{
		"Term":   State(10),
		"Expr":   State(9),
		"Factor": State(11),
	},
	// state 6
	GotoRow{},
	// state 7
	GotoRow{
		"Term":   State(14),
		"Factor": State(4),
	},
	// state 8
	GotoRow{
		"Factor": State(15),
	},
	// state 9
	GotoRow{},
	// state 10
	GotoRow{},
	// state 11
	GotoRow{},
	// state 12
	GotoRow{
		"Term":   State(10),
		"Factor": State(11),
		"Expr":   State(19),
	},
	// state 13
	GotoRow{},
	// state 14
	GotoRow{},
	// state 15
	GotoRow{},
	// state 16
	GotoRow{
		"Term":   State(20),
		"Factor": State(11),
	},
	// state 17
	GotoRow{},
	// state 18
	GotoRow{
		"Factor": State(21),
	},
	// state 19
	GotoRow{},
	// state 20
	GotoRow{},
	// state 21
	GotoRow{},
	// state 22
	GotoRow{},
}
