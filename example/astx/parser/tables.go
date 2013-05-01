package parser

import "code.google.com/p/gocc/example/astx/ast"

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
}

var ActionTable ActionTab = ActionTab{
	// state 0
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			1: Shift(3), // id
		},
	},

	// state 1
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			1: Shift(3),  // id
			0: Accept(0), // $
		},
	},

	// state 2
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			1: Reduce(1), // id
			0: Reduce(1), // $
		},
	},

	// state 3
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			1: Reduce(3), // id
			0: Reduce(3), // $
		},
	},

	// state 4
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			1: Reduce(2), // id
			0: Reduce(2), // $
		},
	},
}

var GotoTable GotoTab = GotoTab{
	// state 0
	GotoRow{
		"StmtList": State(1),
		"Stmt":     State(2),
	},
	// state 1
	GotoRow{
		"Stmt": State(4),
	},
	// state 2
	GotoRow{},
	// state 3
	GotoRow{},
	// state 4
	GotoRow{},
}
