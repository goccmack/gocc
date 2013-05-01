package parser

import "code.google.com/p/gocc/test/sr/ast"

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
		"Stmt : if id then Stmt << ast.NewIf(X[1], X[3]), nil >> ;",
		"Stmt",
		4,
		func(X []Attrib) (Attrib, error) {
			return ast.NewIf(X[1], X[3]), nil
		},
	},
	// [2]
	ProdTabEntry{
		"Stmt : if id then Stmt else Stmt << ast.NewIfElse(X[1], X[3], X[5]), nil >> ;",
		"Stmt",
		6,
		func(X []Attrib) (Attrib, error) {
			return ast.NewIfElse(X[1], X[3], X[5]), nil
		},
	},
	// [3]
	ProdTabEntry{
		"Stmt : id << ast.NewIdStmt(X[0]), nil >> ;",
		"Stmt",
		1,
		func(X []Attrib) (Attrib, error) {
			return ast.NewIdStmt(X[0]), nil
		},
	},
}

var ActionTable ActionTab = ActionTab{
	// state 0
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			1: Shift(3), // id
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
			1: Shift(4), // id
		},
	},

	// state 3
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			0: Reduce(3), // $
		},
	},

	// state 4
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			3: Shift(5), // then
		},
	},

	// state 5
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			2: Shift(7), // if
			1: Shift(8), // id
		},
	},

	// state 6
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			0: Reduce(1), // $
			4: Shift(9),  // else
		},
	},

	// state 7
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			1: Shift(10), // id
		},
	},

	// state 8
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			4: Reduce(3), // else
			0: Reduce(3), // $
		},
	},

	// state 9
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			1: Shift(3), // id
			2: Shift(2), // if
		},
	},

	// state 10
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			3: Shift(12), // then
		},
	},

	// state 11
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			0: Reduce(2), // $
		},
	},

	// state 12
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			1: Shift(8), // id
			2: Shift(7), // if
		},
	},

	// state 13
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			4: Shift(14), // else
			0: Reduce(1), // $
		},
	},

	// state 14
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			2: Shift(7), // if
			1: Shift(8), // id
		},
	},

	// state 15
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			0: Reduce(2), // $
			4: Reduce(2), // else
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
	GotoRow{},
	// state 5
	GotoRow{
		"Stmt": State(6),
	},
	// state 6
	GotoRow{},
	// state 7
	GotoRow{},
	// state 8
	GotoRow{},
	// state 9
	GotoRow{
		"Stmt": State(11),
	},
	// state 10
	GotoRow{},
	// state 11
	GotoRow{},
	// state 12
	GotoRow{
		"Stmt": State(13),
	},
	// state 13
	GotoRow{},
	// state 14
	GotoRow{
		"Stmt": State(15),
	},
	// state 15
	GotoRow{},
}
