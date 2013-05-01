package parser

import "code.google.com/p/gocc/example/errorrecovery/ast"

var ProductionsTable = ProdTab{
	// [0]
	ProdTabEntry{
		"S! : StmtList ;",
		"S!",
		1,
		func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	// [1]
	ProdTabEntry{
		"StmtList : Stmt << ast.NewStmtList(X[0]) >> ;",
		"StmtList",
		1,
		func(X []Attrib) (Attrib, error) {
			return ast.NewStmtList(X[0])
		},
	},
	// [2]
	ProdTabEntry{
		"StmtList : StmtList Stmt << ast.AppendStmt(X[0], X[1]) >> ;",
		"StmtList",
		2,
		func(X []Attrib) (Attrib, error) {
			return ast.AppendStmt(X[0], X[1])
		},
	},
	// [3]
	ProdTabEntry{
		"Stmt : id << ast.NewStmt(X[0]) >> ;",
		"Stmt",
		1,
		func(X []Attrib) (Attrib, error) {
			return ast.NewStmt(X[0])
		},
	},
	// [4]
	ProdTabEntry{
		"Stmt : error ;",
		"Stmt",
		1,
		func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
}

var ActionTable ActionTab = ActionTab{
	// state 0
	&ActionRow{
		CanRecover: true,
		Actions: Actions{
			2: Shift(4), // error
			1: Shift(3), // id
		},
	},

	// state 1
	&ActionRow{
		CanRecover: true,
		Actions: Actions{
			0: Accept(0), // $
			2: Shift(4),  // error
			1: Shift(3),  // id
		},
	},

	// state 2
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			0: Reduce(1), // $
			1: Reduce(1), // id
			2: Reduce(1), // error
		},
	},

	// state 3
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			2: Reduce(3), // error
			1: Reduce(3), // id
			0: Reduce(3), // $
		},
	},

	// state 4
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			1: Reduce(4), // id
			0: Reduce(4), // $
			2: Reduce(4), // error
		},
	},

	// state 5
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			2: Reduce(2), // error
			0: Reduce(2), // $
			1: Reduce(2), // id
		},
	},
}

var GotoTable GotoTab = GotoTab{
	// state 0
	GotoRow{
		"Stmt":     State(2),
		"StmtList": State(1),
	},
	// state 1
	GotoRow{
		"Stmt": State(5),
	},
	// state 2
	GotoRow{},
	// state 3
	GotoRow{},
	// state 4
	GotoRow{},
	// state 5
	GotoRow{},
}
