package parser

var ProductionsTable = ProdTab{
	// [0]
	ProdTabEntry{
		"S! : RR ;",
		"S!",
		1,
		func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	// [1]
	ProdTabEntry{
		"RR : A ;",
		"RR",
		1,
		func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	// [2]
	ProdTabEntry{
		"RR : B ;",
		"RR",
		1,
		func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	// [3]
	ProdTabEntry{
		"B : a << `B `, nil >> ;",
		"B",
		1,
		func(X []Attrib) (Attrib, error) {
			return `B `, nil
		},
	},
	// [4]
	ProdTabEntry{
		"A : a << `A0 `, nil >> ;",
		"A",
		1,
		func(X []Attrib) (Attrib, error) {
			return `A0 `, nil
		},
	},
	// [5]
	ProdTabEntry{
		"A : A a << `A1 `, nil >> ;",
		"A",
		2,
		func(X []Attrib) (Attrib, error) {
			return `A1 `, nil
		},
	},
	// [6]
	ProdTabEntry{
		"A : c << `A2 `, nil >> ;",
		"A",
		1,
		func(X []Attrib) (Attrib, error) {
			return `A2 `, nil
		},
	},
}

var ActionTable ActionTab = ActionTab{
	// state 0
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			3: Shift(5), // c
			2: Shift(4), // a
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
			2: Shift(6),  // a
			0: Reduce(1), // $
		},
	},

	// state 3
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			0: Reduce(2), // $
		},
	},

	// state 4
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			2: Reduce(4), // a
			0: Reduce(3), // $
		},
	},

	// state 5
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			0: Reduce(6), // $
			2: Reduce(6), // a
		},
	},

	// state 6
	&ActionRow{
		CanRecover: false,
		Actions: Actions{
			0: Reduce(5), // $
			2: Reduce(5), // a
		},
	},
}

var GotoTable GotoTab = GotoTab{
	// state 0
	GotoRow{
		"A":  State(2),
		"RR": State(1),
		"B":  State(3),
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
	GotoRow{},
	// state 6
	GotoRow{},
}
