package parser

var ProductionsTable = ProdTab{
	// [0]
	ProdTabEntry{
		"S! : Stmt ;",
		"S!",
		1,
		func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	// [1]
	ProdTabEntry{
		"Stmt : if expr then Stmt ;",
		"Stmt",
		4,
		func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	// [2]
	ProdTabEntry{
		"Stmt : if expr then Stmt else Stmt ;",
		"Stmt",
		6,
		func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
}

var ActionTable ActionTab = ActionTab{
	// state 0
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			2: Shift(2), // if
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
			3: Shift(3), // expr
		},
	},

	// state 3
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			4: Shift(4), // then
		},
	},

	// state 4
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			2: Shift(6), // if
		},
	},

	// state 5
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			5: Shift(7),  // else
			0: Reduce(1), // $
		},
	},

	// state 6
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			3: Shift(8), // expr
		},
	},

	// state 7
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			2: Shift(2), // if
		},
	},

	// state 8
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			4: Shift(10), // then
		},
	},

	// state 9
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			0: Reduce(2), // $
		},
	},

	// state 10
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			2: Shift(6), // if
		},
	},

	// state 11
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			0: Reduce(1), // $
			5: Shift(12), // else
		},
	},

	// state 12
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			2: Shift(6), // if
		},
	},

	// state 13
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			0: Reduce(2), // $
			5: Reduce(2), // else
		},
	},
}

var GotoTable GotoTab = GotoTab{
	// state 0
	GotoRow{
		"Stmt": State(1),
	},
	// state 1
	GotoRow{},
	// state 2
	GotoRow{},
	// state 3
	GotoRow{},
	// state 4
	GotoRow{
		"Stmt": State(5),
	},
	// state 5
	GotoRow{},
	// state 6
	GotoRow{},
	// state 7
	GotoRow{
		"Stmt": State(9),
	},
	// state 8
	GotoRow{},
	// state 9
	GotoRow{},
	// state 10
	GotoRow{
		"Stmt": State(11),
	},
	// state 11
	GotoRow{},
	// state 12
	GotoRow{
		"Stmt": State(13),
	},
	// state 13
	GotoRow{},
}
