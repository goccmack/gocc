
package parser

type(
	actionTable [numStates]actionRow
	actionRow struct {
		canRecover bool
		actions [numSymbols]action
	}
)

var actionTab = actionTable{
	actionRow{ // S0
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(6),		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			shift(7),		/* regDefId */
			shift(8),		/* ignoredTokId */
			nil,		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			nil,		/* ] */
			nil,		/* { */
			nil,		/* } */
			nil,		/* ( */
			nil,		/* ) */
			shift(11),		/* actionLit */
			shift(13),		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S1
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			accept(true),		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			nil,		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			nil,		/* ] */
			nil,		/* { */
			nil,		/* } */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S2
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(2),		/* $, reduce: Grammar */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			nil,		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			nil,		/* ] */
			nil,		/* { */
			nil,		/* } */
			nil,		/* ( */
			nil,		/* ) */
			shift(11),		/* actionLit */
			shift(13),		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S3
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(3),		/* $, reduce: Grammar */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			nil,		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			nil,		/* ] */
			nil,		/* { */
			nil,		/* } */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S4
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(4),		/* $, reduce: LexicalPart */
			shift(6),		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			shift(7),		/* regDefId */
			shift(8),		/* ignoredTokId */
			nil,		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			nil,		/* ] */
			nil,		/* { */
			nil,		/* } */
			nil,		/* ( */
			nil,		/* ) */
			reduce(4),		/* actionLit, reduce: LexicalPart */
			reduce(4),		/* prodId, reduce: LexicalPart */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S5
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(5),		/* $, reduce: LexProductions */
			reduce(5),		/* tokId, reduce: LexProductions */
			nil,		/* : */
			nil,		/* ; */
			reduce(5),		/* regDefId, reduce: LexProductions */
			reduce(5),		/* ignoredTokId, reduce: LexProductions */
			nil,		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			nil,		/* ] */
			nil,		/* { */
			nil,		/* } */
			nil,		/* ( */
			nil,		/* ) */
			reduce(5),		/* actionLit, reduce: LexProductions */
			reduce(5),		/* prodId, reduce: LexProductions */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S6
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			shift(16),		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			nil,		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			nil,		/* ] */
			nil,		/* { */
			nil,		/* } */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S7
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			shift(17),		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			nil,		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			nil,		/* ] */
			nil,		/* { */
			nil,		/* } */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S8
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			shift(18),		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			nil,		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			nil,		/* ] */
			nil,		/* { */
			nil,		/* } */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S9
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			nil,		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			nil,		/* ] */
			nil,		/* { */
			nil,		/* } */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			shift(13),		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S10
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(22),		/* $, reduce: SyntaxPart */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			nil,		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			nil,		/* ] */
			nil,		/* { */
			nil,		/* } */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			shift(13),		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S11
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			nil,		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			nil,		/* ] */
			nil,		/* { */
			nil,		/* } */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			reduce(23),		/* prodId, reduce: FileHeader */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S12
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(24),		/* $, reduce: SyntaxProdList */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			nil,		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			nil,		/* ] */
			nil,		/* { */
			nil,		/* } */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			reduce(24),		/* prodId, reduce: SyntaxProdList */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S13
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			shift(21),		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			nil,		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			nil,		/* ] */
			nil,		/* { */
			nil,		/* } */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S14
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(1),		/* $, reduce: Grammar */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			nil,		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			nil,		/* ] */
			nil,		/* { */
			nil,		/* } */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S15
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(6),		/* $, reduce: LexProductions */
			reduce(6),		/* tokId, reduce: LexProductions */
			nil,		/* : */
			nil,		/* ; */
			reduce(6),		/* regDefId, reduce: LexProductions */
			reduce(6),		/* ignoredTokId, reduce: LexProductions */
			nil,		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			nil,		/* ] */
			nil,		/* { */
			nil,		/* } */
			nil,		/* ( */
			nil,		/* ) */
			reduce(6),		/* actionLit, reduce: LexProductions */
			reduce(6),		/* prodId, reduce: LexProductions */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S16
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			shift(23),		/* regDefId */
			nil,		/* ignoredTokId */
			nil,		/* | */
			shift(26),		/* . */
			shift(27),		/* char_lit */
			nil,		/* - */
			shift(28),		/* [ */
			nil,		/* ] */
			shift(29),		/* { */
			nil,		/* } */
			shift(30),		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S17
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			shift(23),		/* regDefId */
			nil,		/* ignoredTokId */
			nil,		/* | */
			shift(26),		/* . */
			shift(27),		/* char_lit */
			nil,		/* - */
			shift(28),		/* [ */
			nil,		/* ] */
			shift(29),		/* { */
			nil,		/* } */
			shift(30),		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S18
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			shift(23),		/* regDefId */
			nil,		/* ignoredTokId */
			nil,		/* | */
			shift(26),		/* . */
			shift(27),		/* char_lit */
			nil,		/* - */
			shift(28),		/* [ */
			nil,		/* ] */
			shift(29),		/* { */
			nil,		/* } */
			shift(30),		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S19
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(21),		/* $, reduce: SyntaxPart */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			nil,		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			nil,		/* ] */
			nil,		/* { */
			nil,		/* } */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			shift(13),		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S20
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(25),		/* $, reduce: SyntaxProdList */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			nil,		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			nil,		/* ] */
			nil,		/* { */
			nil,		/* } */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			reduce(25),		/* prodId, reduce: SyntaxProdList */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S21
				canRecover: true,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(33),		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			nil,		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			shift(34),		/* [ */
			nil,		/* ] */
			shift(35),		/* { */
			nil,		/* } */
			shift(36),		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			shift(37),		/* prodId */
			shift(41),		/* error */
			shift(43),		/* stringLit */
			
		},

	},
	actionRow{ // S22
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			shift(44),		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			shift(45),		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			nil,		/* ] */
			nil,		/* { */
			nil,		/* } */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S23
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			reduce(17),		/* ;, reduce: LexTerm */
			reduce(17),		/* regDefId, reduce: LexTerm */
			nil,		/* ignoredTokId */
			reduce(17),		/* |, reduce: LexTerm */
			reduce(17),		/* ., reduce: LexTerm */
			reduce(17),		/* char_lit, reduce: LexTerm */
			nil,		/* - */
			reduce(17),		/* [, reduce: LexTerm */
			nil,		/* ] */
			reduce(17),		/* {, reduce: LexTerm */
			nil,		/* } */
			reduce(17),		/* (, reduce: LexTerm */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S24
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			reduce(10),		/* ;, reduce: LexPattern */
			shift(23),		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(10),		/* |, reduce: LexPattern */
			shift(26),		/* . */
			shift(27),		/* char_lit */
			nil,		/* - */
			shift(28),		/* [ */
			nil,		/* ] */
			shift(29),		/* { */
			nil,		/* } */
			shift(30),		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S25
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			reduce(12),		/* ;, reduce: LexAlt */
			reduce(12),		/* regDefId, reduce: LexAlt */
			nil,		/* ignoredTokId */
			reduce(12),		/* |, reduce: LexAlt */
			reduce(12),		/* ., reduce: LexAlt */
			reduce(12),		/* char_lit, reduce: LexAlt */
			nil,		/* - */
			reduce(12),		/* [, reduce: LexAlt */
			nil,		/* ] */
			reduce(12),		/* {, reduce: LexAlt */
			nil,		/* } */
			reduce(12),		/* (, reduce: LexAlt */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S26
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			reduce(14),		/* ;, reduce: LexTerm */
			reduce(14),		/* regDefId, reduce: LexTerm */
			nil,		/* ignoredTokId */
			reduce(14),		/* |, reduce: LexTerm */
			reduce(14),		/* ., reduce: LexTerm */
			reduce(14),		/* char_lit, reduce: LexTerm */
			nil,		/* - */
			reduce(14),		/* [, reduce: LexTerm */
			nil,		/* ] */
			reduce(14),		/* {, reduce: LexTerm */
			nil,		/* } */
			reduce(14),		/* (, reduce: LexTerm */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S27
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			reduce(15),		/* ;, reduce: LexTerm */
			reduce(15),		/* regDefId, reduce: LexTerm */
			nil,		/* ignoredTokId */
			reduce(15),		/* |, reduce: LexTerm */
			reduce(15),		/* ., reduce: LexTerm */
			reduce(15),		/* char_lit, reduce: LexTerm */
			shift(47),		/* - */
			reduce(15),		/* [, reduce: LexTerm */
			nil,		/* ] */
			reduce(15),		/* {, reduce: LexTerm */
			nil,		/* } */
			reduce(15),		/* (, reduce: LexTerm */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S28
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			shift(49),		/* regDefId */
			nil,		/* ignoredTokId */
			nil,		/* | */
			shift(52),		/* . */
			shift(53),		/* char_lit */
			nil,		/* - */
			shift(54),		/* [ */
			nil,		/* ] */
			shift(55),		/* { */
			nil,		/* } */
			shift(56),		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S29
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			shift(58),		/* regDefId */
			nil,		/* ignoredTokId */
			nil,		/* | */
			shift(61),		/* . */
			shift(62),		/* char_lit */
			nil,		/* - */
			shift(63),		/* [ */
			nil,		/* ] */
			shift(64),		/* { */
			nil,		/* } */
			shift(65),		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S30
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			shift(67),		/* regDefId */
			nil,		/* ignoredTokId */
			nil,		/* | */
			shift(70),		/* . */
			shift(71),		/* char_lit */
			nil,		/* - */
			shift(72),		/* [ */
			nil,		/* ] */
			shift(73),		/* { */
			nil,		/* } */
			shift(74),		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S31
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			shift(75),		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			shift(45),		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			nil,		/* ] */
			nil,		/* { */
			nil,		/* } */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S32
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			shift(76),		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			shift(45),		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			nil,		/* ] */
			nil,		/* { */
			nil,		/* } */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S33
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(37),		/* tokId, reduce: SyntaxTerm */
			nil,		/* : */
			reduce(37),		/* ;, reduce: SyntaxTerm */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(37),		/* |, reduce: SyntaxTerm */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			reduce(37),		/* [, reduce: SyntaxTerm */
			nil,		/* ] */
			reduce(37),		/* {, reduce: SyntaxTerm */
			nil,		/* } */
			reduce(37),		/* (, reduce: SyntaxTerm */
			nil,		/* ) */
			reduce(37),		/* actionLit, reduce: SyntaxTerm */
			reduce(37),		/* prodId, reduce: SyntaxTerm */
			nil,		/* error */
			reduce(37),		/* stringLit, reduce: SyntaxTerm */
			
		},

	},
	actionRow{ // S34
				canRecover: true,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(77),		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			nil,		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			shift(78),		/* [ */
			nil,		/* ] */
			shift(79),		/* { */
			nil,		/* } */
			shift(80),		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			shift(81),		/* prodId */
			shift(85),		/* error */
			shift(87),		/* stringLit */
			
		},

	},
	actionRow{ // S35
				canRecover: true,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(88),		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			nil,		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			shift(89),		/* [ */
			nil,		/* ] */
			shift(90),		/* { */
			nil,		/* } */
			shift(91),		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			shift(92),		/* prodId */
			shift(96),		/* error */
			shift(98),		/* stringLit */
			
		},

	},
	actionRow{ // S36
				canRecover: true,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(99),		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			nil,		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			shift(100),		/* [ */
			nil,		/* ] */
			shift(101),		/* { */
			nil,		/* } */
			shift(102),		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			shift(103),		/* prodId */
			shift(107),		/* error */
			shift(109),		/* stringLit */
			
		},

	},
	actionRow{ // S37
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(36),		/* tokId, reduce: SyntaxTerm */
			nil,		/* : */
			reduce(36),		/* ;, reduce: SyntaxTerm */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(36),		/* |, reduce: SyntaxTerm */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			reduce(36),		/* [, reduce: SyntaxTerm */
			nil,		/* ] */
			reduce(36),		/* {, reduce: SyntaxTerm */
			nil,		/* } */
			reduce(36),		/* (, reduce: SyntaxTerm */
			nil,		/* ) */
			reduce(36),		/* actionLit, reduce: SyntaxTerm */
			reduce(36),		/* prodId, reduce: SyntaxTerm */
			nil,		/* error */
			reduce(36),		/* stringLit, reduce: SyntaxTerm */
			
		},

	},
	actionRow{ // S38
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			shift(110),		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			shift(111),		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			nil,		/* ] */
			nil,		/* { */
			nil,		/* } */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S39
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			reduce(27),		/* ;, reduce: SyntaxExpression */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(27),		/* |, reduce: SyntaxExpression */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			nil,		/* ] */
			nil,		/* { */
			nil,		/* } */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S40
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(33),		/* tokId */
			nil,		/* : */
			reduce(29),		/* ;, reduce: SyntaxBody */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(29),		/* |, reduce: SyntaxBody */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			shift(34),		/* [ */
			nil,		/* ] */
			shift(35),		/* { */
			nil,		/* } */
			shift(36),		/* ( */
			nil,		/* ) */
			shift(112),		/* actionLit */
			shift(37),		/* prodId */
			nil,		/* error */
			shift(43),		/* stringLit */
			
		},

	},
	actionRow{ // S41
				canRecover: true,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(33),		/* tokId */
			nil,		/* : */
			reduce(31),		/* ;, reduce: SyntaxBody */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(31),		/* |, reduce: SyntaxBody */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			shift(34),		/* [ */
			nil,		/* ] */
			shift(35),		/* { */
			nil,		/* } */
			shift(36),		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			shift(37),		/* prodId */
			nil,		/* error */
			shift(43),		/* stringLit */
			
		},

	},
	actionRow{ // S42
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(34),		/* tokId, reduce: SyntaxTerms */
			nil,		/* : */
			reduce(34),		/* ;, reduce: SyntaxTerms */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(34),		/* |, reduce: SyntaxTerms */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			reduce(34),		/* [, reduce: SyntaxTerms */
			nil,		/* ] */
			reduce(34),		/* {, reduce: SyntaxTerms */
			nil,		/* } */
			reduce(34),		/* (, reduce: SyntaxTerms */
			nil,		/* ) */
			reduce(34),		/* actionLit, reduce: SyntaxTerms */
			reduce(34),		/* prodId, reduce: SyntaxTerms */
			nil,		/* error */
			reduce(34),		/* stringLit, reduce: SyntaxTerms */
			
		},

	},
	actionRow{ // S43
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(38),		/* tokId, reduce: SyntaxTerm */
			nil,		/* : */
			reduce(38),		/* ;, reduce: SyntaxTerm */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(38),		/* |, reduce: SyntaxTerm */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			reduce(38),		/* [, reduce: SyntaxTerm */
			nil,		/* ] */
			reduce(38),		/* {, reduce: SyntaxTerm */
			nil,		/* } */
			reduce(38),		/* (, reduce: SyntaxTerm */
			nil,		/* ) */
			reduce(38),		/* actionLit, reduce: SyntaxTerm */
			reduce(38),		/* prodId, reduce: SyntaxTerm */
			nil,		/* error */
			reduce(38),		/* stringLit, reduce: SyntaxTerm */
			
		},

	},
	actionRow{ // S44
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(7),		/* $, reduce: LexProduction */
			reduce(7),		/* tokId, reduce: LexProduction */
			nil,		/* : */
			nil,		/* ; */
			reduce(7),		/* regDefId, reduce: LexProduction */
			reduce(7),		/* ignoredTokId, reduce: LexProduction */
			nil,		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			nil,		/* ] */
			nil,		/* { */
			nil,		/* } */
			nil,		/* ( */
			nil,		/* ) */
			reduce(7),		/* actionLit, reduce: LexProduction */
			reduce(7),		/* prodId, reduce: LexProduction */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S45
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			shift(23),		/* regDefId */
			nil,		/* ignoredTokId */
			nil,		/* | */
			shift(26),		/* . */
			shift(27),		/* char_lit */
			nil,		/* - */
			shift(28),		/* [ */
			nil,		/* ] */
			shift(29),		/* { */
			nil,		/* } */
			shift(30),		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S46
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			reduce(13),		/* ;, reduce: LexAlt */
			reduce(13),		/* regDefId, reduce: LexAlt */
			nil,		/* ignoredTokId */
			reduce(13),		/* |, reduce: LexAlt */
			reduce(13),		/* ., reduce: LexAlt */
			reduce(13),		/* char_lit, reduce: LexAlt */
			nil,		/* - */
			reduce(13),		/* [, reduce: LexAlt */
			nil,		/* ] */
			reduce(13),		/* {, reduce: LexAlt */
			nil,		/* } */
			reduce(13),		/* (, reduce: LexAlt */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S47
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			nil,		/* | */
			nil,		/* . */
			shift(116),		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			nil,		/* ] */
			nil,		/* { */
			nil,		/* } */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S48
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			shift(117),		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			shift(118),		/* ] */
			nil,		/* { */
			nil,		/* } */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S49
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			reduce(17),		/* regDefId, reduce: LexTerm */
			nil,		/* ignoredTokId */
			reduce(17),		/* |, reduce: LexTerm */
			reduce(17),		/* ., reduce: LexTerm */
			reduce(17),		/* char_lit, reduce: LexTerm */
			nil,		/* - */
			reduce(17),		/* [, reduce: LexTerm */
			reduce(17),		/* ], reduce: LexTerm */
			reduce(17),		/* {, reduce: LexTerm */
			nil,		/* } */
			reduce(17),		/* (, reduce: LexTerm */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S50
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			shift(49),		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(10),		/* |, reduce: LexPattern */
			shift(52),		/* . */
			shift(53),		/* char_lit */
			nil,		/* - */
			shift(54),		/* [ */
			reduce(10),		/* ], reduce: LexPattern */
			shift(55),		/* { */
			nil,		/* } */
			shift(56),		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S51
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			reduce(12),		/* regDefId, reduce: LexAlt */
			nil,		/* ignoredTokId */
			reduce(12),		/* |, reduce: LexAlt */
			reduce(12),		/* ., reduce: LexAlt */
			reduce(12),		/* char_lit, reduce: LexAlt */
			nil,		/* - */
			reduce(12),		/* [, reduce: LexAlt */
			reduce(12),		/* ], reduce: LexAlt */
			reduce(12),		/* {, reduce: LexAlt */
			nil,		/* } */
			reduce(12),		/* (, reduce: LexAlt */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S52
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			reduce(14),		/* regDefId, reduce: LexTerm */
			nil,		/* ignoredTokId */
			reduce(14),		/* |, reduce: LexTerm */
			reduce(14),		/* ., reduce: LexTerm */
			reduce(14),		/* char_lit, reduce: LexTerm */
			nil,		/* - */
			reduce(14),		/* [, reduce: LexTerm */
			reduce(14),		/* ], reduce: LexTerm */
			reduce(14),		/* {, reduce: LexTerm */
			nil,		/* } */
			reduce(14),		/* (, reduce: LexTerm */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S53
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			reduce(15),		/* regDefId, reduce: LexTerm */
			nil,		/* ignoredTokId */
			reduce(15),		/* |, reduce: LexTerm */
			reduce(15),		/* ., reduce: LexTerm */
			reduce(15),		/* char_lit, reduce: LexTerm */
			shift(120),		/* - */
			reduce(15),		/* [, reduce: LexTerm */
			reduce(15),		/* ], reduce: LexTerm */
			reduce(15),		/* {, reduce: LexTerm */
			nil,		/* } */
			reduce(15),		/* (, reduce: LexTerm */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S54
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			shift(49),		/* regDefId */
			nil,		/* ignoredTokId */
			nil,		/* | */
			shift(52),		/* . */
			shift(53),		/* char_lit */
			nil,		/* - */
			shift(54),		/* [ */
			nil,		/* ] */
			shift(55),		/* { */
			nil,		/* } */
			shift(56),		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S55
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			shift(58),		/* regDefId */
			nil,		/* ignoredTokId */
			nil,		/* | */
			shift(61),		/* . */
			shift(62),		/* char_lit */
			nil,		/* - */
			shift(63),		/* [ */
			nil,		/* ] */
			shift(64),		/* { */
			nil,		/* } */
			shift(65),		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S56
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			shift(67),		/* regDefId */
			nil,		/* ignoredTokId */
			nil,		/* | */
			shift(70),		/* . */
			shift(71),		/* char_lit */
			nil,		/* - */
			shift(72),		/* [ */
			nil,		/* ] */
			shift(73),		/* { */
			nil,		/* } */
			shift(74),		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S57
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			shift(124),		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			nil,		/* ] */
			nil,		/* { */
			shift(125),		/* } */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S58
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			reduce(17),		/* regDefId, reduce: LexTerm */
			nil,		/* ignoredTokId */
			reduce(17),		/* |, reduce: LexTerm */
			reduce(17),		/* ., reduce: LexTerm */
			reduce(17),		/* char_lit, reduce: LexTerm */
			nil,		/* - */
			reduce(17),		/* [, reduce: LexTerm */
			nil,		/* ] */
			reduce(17),		/* {, reduce: LexTerm */
			reduce(17),		/* }, reduce: LexTerm */
			reduce(17),		/* (, reduce: LexTerm */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S59
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			shift(58),		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(10),		/* |, reduce: LexPattern */
			shift(61),		/* . */
			shift(62),		/* char_lit */
			nil,		/* - */
			shift(63),		/* [ */
			nil,		/* ] */
			shift(64),		/* { */
			reduce(10),		/* }, reduce: LexPattern */
			shift(65),		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S60
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			reduce(12),		/* regDefId, reduce: LexAlt */
			nil,		/* ignoredTokId */
			reduce(12),		/* |, reduce: LexAlt */
			reduce(12),		/* ., reduce: LexAlt */
			reduce(12),		/* char_lit, reduce: LexAlt */
			nil,		/* - */
			reduce(12),		/* [, reduce: LexAlt */
			nil,		/* ] */
			reduce(12),		/* {, reduce: LexAlt */
			reduce(12),		/* }, reduce: LexAlt */
			reduce(12),		/* (, reduce: LexAlt */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S61
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			reduce(14),		/* regDefId, reduce: LexTerm */
			nil,		/* ignoredTokId */
			reduce(14),		/* |, reduce: LexTerm */
			reduce(14),		/* ., reduce: LexTerm */
			reduce(14),		/* char_lit, reduce: LexTerm */
			nil,		/* - */
			reduce(14),		/* [, reduce: LexTerm */
			nil,		/* ] */
			reduce(14),		/* {, reduce: LexTerm */
			reduce(14),		/* }, reduce: LexTerm */
			reduce(14),		/* (, reduce: LexTerm */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S62
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			reduce(15),		/* regDefId, reduce: LexTerm */
			nil,		/* ignoredTokId */
			reduce(15),		/* |, reduce: LexTerm */
			reduce(15),		/* ., reduce: LexTerm */
			reduce(15),		/* char_lit, reduce: LexTerm */
			shift(127),		/* - */
			reduce(15),		/* [, reduce: LexTerm */
			nil,		/* ] */
			reduce(15),		/* {, reduce: LexTerm */
			reduce(15),		/* }, reduce: LexTerm */
			reduce(15),		/* (, reduce: LexTerm */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S63
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			shift(49),		/* regDefId */
			nil,		/* ignoredTokId */
			nil,		/* | */
			shift(52),		/* . */
			shift(53),		/* char_lit */
			nil,		/* - */
			shift(54),		/* [ */
			nil,		/* ] */
			shift(55),		/* { */
			nil,		/* } */
			shift(56),		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S64
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			shift(58),		/* regDefId */
			nil,		/* ignoredTokId */
			nil,		/* | */
			shift(61),		/* . */
			shift(62),		/* char_lit */
			nil,		/* - */
			shift(63),		/* [ */
			nil,		/* ] */
			shift(64),		/* { */
			nil,		/* } */
			shift(65),		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S65
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			shift(67),		/* regDefId */
			nil,		/* ignoredTokId */
			nil,		/* | */
			shift(70),		/* . */
			shift(71),		/* char_lit */
			nil,		/* - */
			shift(72),		/* [ */
			nil,		/* ] */
			shift(73),		/* { */
			nil,		/* } */
			shift(74),		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S66
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			shift(131),		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			nil,		/* ] */
			nil,		/* { */
			nil,		/* } */
			nil,		/* ( */
			shift(132),		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S67
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			reduce(17),		/* regDefId, reduce: LexTerm */
			nil,		/* ignoredTokId */
			reduce(17),		/* |, reduce: LexTerm */
			reduce(17),		/* ., reduce: LexTerm */
			reduce(17),		/* char_lit, reduce: LexTerm */
			nil,		/* - */
			reduce(17),		/* [, reduce: LexTerm */
			nil,		/* ] */
			reduce(17),		/* {, reduce: LexTerm */
			nil,		/* } */
			reduce(17),		/* (, reduce: LexTerm */
			reduce(17),		/* ), reduce: LexTerm */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S68
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			shift(67),		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(10),		/* |, reduce: LexPattern */
			shift(70),		/* . */
			shift(71),		/* char_lit */
			nil,		/* - */
			shift(72),		/* [ */
			nil,		/* ] */
			shift(73),		/* { */
			nil,		/* } */
			shift(74),		/* ( */
			reduce(10),		/* ), reduce: LexPattern */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S69
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			reduce(12),		/* regDefId, reduce: LexAlt */
			nil,		/* ignoredTokId */
			reduce(12),		/* |, reduce: LexAlt */
			reduce(12),		/* ., reduce: LexAlt */
			reduce(12),		/* char_lit, reduce: LexAlt */
			nil,		/* - */
			reduce(12),		/* [, reduce: LexAlt */
			nil,		/* ] */
			reduce(12),		/* {, reduce: LexAlt */
			nil,		/* } */
			reduce(12),		/* (, reduce: LexAlt */
			reduce(12),		/* ), reduce: LexAlt */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S70
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			reduce(14),		/* regDefId, reduce: LexTerm */
			nil,		/* ignoredTokId */
			reduce(14),		/* |, reduce: LexTerm */
			reduce(14),		/* ., reduce: LexTerm */
			reduce(14),		/* char_lit, reduce: LexTerm */
			nil,		/* - */
			reduce(14),		/* [, reduce: LexTerm */
			nil,		/* ] */
			reduce(14),		/* {, reduce: LexTerm */
			nil,		/* } */
			reduce(14),		/* (, reduce: LexTerm */
			reduce(14),		/* ), reduce: LexTerm */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S71
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			reduce(15),		/* regDefId, reduce: LexTerm */
			nil,		/* ignoredTokId */
			reduce(15),		/* |, reduce: LexTerm */
			reduce(15),		/* ., reduce: LexTerm */
			reduce(15),		/* char_lit, reduce: LexTerm */
			shift(134),		/* - */
			reduce(15),		/* [, reduce: LexTerm */
			nil,		/* ] */
			reduce(15),		/* {, reduce: LexTerm */
			nil,		/* } */
			reduce(15),		/* (, reduce: LexTerm */
			reduce(15),		/* ), reduce: LexTerm */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S72
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			shift(49),		/* regDefId */
			nil,		/* ignoredTokId */
			nil,		/* | */
			shift(52),		/* . */
			shift(53),		/* char_lit */
			nil,		/* - */
			shift(54),		/* [ */
			nil,		/* ] */
			shift(55),		/* { */
			nil,		/* } */
			shift(56),		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S73
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			shift(58),		/* regDefId */
			nil,		/* ignoredTokId */
			nil,		/* | */
			shift(61),		/* . */
			shift(62),		/* char_lit */
			nil,		/* - */
			shift(63),		/* [ */
			nil,		/* ] */
			shift(64),		/* { */
			nil,		/* } */
			shift(65),		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S74
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			shift(67),		/* regDefId */
			nil,		/* ignoredTokId */
			nil,		/* | */
			shift(70),		/* . */
			shift(71),		/* char_lit */
			nil,		/* - */
			shift(72),		/* [ */
			nil,		/* ] */
			shift(73),		/* { */
			nil,		/* } */
			shift(74),		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S75
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(8),		/* $, reduce: LexProduction */
			reduce(8),		/* tokId, reduce: LexProduction */
			nil,		/* : */
			nil,		/* ; */
			reduce(8),		/* regDefId, reduce: LexProduction */
			reduce(8),		/* ignoredTokId, reduce: LexProduction */
			nil,		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			nil,		/* ] */
			nil,		/* { */
			nil,		/* } */
			nil,		/* ( */
			nil,		/* ) */
			reduce(8),		/* actionLit, reduce: LexProduction */
			reduce(8),		/* prodId, reduce: LexProduction */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S76
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(9),		/* $, reduce: LexProduction */
			reduce(9),		/* tokId, reduce: LexProduction */
			nil,		/* : */
			nil,		/* ; */
			reduce(9),		/* regDefId, reduce: LexProduction */
			reduce(9),		/* ignoredTokId, reduce: LexProduction */
			nil,		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			nil,		/* ] */
			nil,		/* { */
			nil,		/* } */
			nil,		/* ( */
			nil,		/* ) */
			reduce(9),		/* actionLit, reduce: LexProduction */
			reduce(9),		/* prodId, reduce: LexProduction */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S77
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(37),		/* tokId, reduce: SyntaxTerm */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(37),		/* |, reduce: SyntaxTerm */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			reduce(37),		/* [, reduce: SyntaxTerm */
			reduce(37),		/* ], reduce: SyntaxTerm */
			reduce(37),		/* {, reduce: SyntaxTerm */
			nil,		/* } */
			reduce(37),		/* (, reduce: SyntaxTerm */
			nil,		/* ) */
			reduce(37),		/* actionLit, reduce: SyntaxTerm */
			reduce(37),		/* prodId, reduce: SyntaxTerm */
			nil,		/* error */
			reduce(37),		/* stringLit, reduce: SyntaxTerm */
			
		},

	},
	actionRow{ // S78
				canRecover: true,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(77),		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			nil,		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			shift(78),		/* [ */
			nil,		/* ] */
			shift(79),		/* { */
			nil,		/* } */
			shift(80),		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			shift(81),		/* prodId */
			shift(85),		/* error */
			shift(87),		/* stringLit */
			
		},

	},
	actionRow{ // S79
				canRecover: true,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(88),		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			nil,		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			shift(89),		/* [ */
			nil,		/* ] */
			shift(90),		/* { */
			nil,		/* } */
			shift(91),		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			shift(92),		/* prodId */
			shift(96),		/* error */
			shift(98),		/* stringLit */
			
		},

	},
	actionRow{ // S80
				canRecover: true,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(99),		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			nil,		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			shift(100),		/* [ */
			nil,		/* ] */
			shift(101),		/* { */
			nil,		/* } */
			shift(102),		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			shift(103),		/* prodId */
			shift(107),		/* error */
			shift(109),		/* stringLit */
			
		},

	},
	actionRow{ // S81
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(36),		/* tokId, reduce: SyntaxTerm */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(36),		/* |, reduce: SyntaxTerm */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			reduce(36),		/* [, reduce: SyntaxTerm */
			reduce(36),		/* ], reduce: SyntaxTerm */
			reduce(36),		/* {, reduce: SyntaxTerm */
			nil,		/* } */
			reduce(36),		/* (, reduce: SyntaxTerm */
			nil,		/* ) */
			reduce(36),		/* actionLit, reduce: SyntaxTerm */
			reduce(36),		/* prodId, reduce: SyntaxTerm */
			nil,		/* error */
			reduce(36),		/* stringLit, reduce: SyntaxTerm */
			
		},

	},
	actionRow{ // S82
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			shift(141),		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			shift(142),		/* ] */
			nil,		/* { */
			nil,		/* } */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S83
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(27),		/* |, reduce: SyntaxExpression */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			reduce(27),		/* ], reduce: SyntaxExpression */
			nil,		/* { */
			nil,		/* } */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S84
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(77),		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(29),		/* |, reduce: SyntaxBody */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			shift(78),		/* [ */
			reduce(29),		/* ], reduce: SyntaxBody */
			shift(79),		/* { */
			nil,		/* } */
			shift(80),		/* ( */
			nil,		/* ) */
			shift(143),		/* actionLit */
			shift(81),		/* prodId */
			nil,		/* error */
			shift(87),		/* stringLit */
			
		},

	},
	actionRow{ // S85
				canRecover: true,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(77),		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(31),		/* |, reduce: SyntaxBody */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			shift(78),		/* [ */
			reduce(31),		/* ], reduce: SyntaxBody */
			shift(79),		/* { */
			nil,		/* } */
			shift(80),		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			shift(81),		/* prodId */
			nil,		/* error */
			shift(87),		/* stringLit */
			
		},

	},
	actionRow{ // S86
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(34),		/* tokId, reduce: SyntaxTerms */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(34),		/* |, reduce: SyntaxTerms */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			reduce(34),		/* [, reduce: SyntaxTerms */
			reduce(34),		/* ], reduce: SyntaxTerms */
			reduce(34),		/* {, reduce: SyntaxTerms */
			nil,		/* } */
			reduce(34),		/* (, reduce: SyntaxTerms */
			nil,		/* ) */
			reduce(34),		/* actionLit, reduce: SyntaxTerms */
			reduce(34),		/* prodId, reduce: SyntaxTerms */
			nil,		/* error */
			reduce(34),		/* stringLit, reduce: SyntaxTerms */
			
		},

	},
	actionRow{ // S87
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(38),		/* tokId, reduce: SyntaxTerm */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(38),		/* |, reduce: SyntaxTerm */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			reduce(38),		/* [, reduce: SyntaxTerm */
			reduce(38),		/* ], reduce: SyntaxTerm */
			reduce(38),		/* {, reduce: SyntaxTerm */
			nil,		/* } */
			reduce(38),		/* (, reduce: SyntaxTerm */
			nil,		/* ) */
			reduce(38),		/* actionLit, reduce: SyntaxTerm */
			reduce(38),		/* prodId, reduce: SyntaxTerm */
			nil,		/* error */
			reduce(38),		/* stringLit, reduce: SyntaxTerm */
			
		},

	},
	actionRow{ // S88
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(37),		/* tokId, reduce: SyntaxTerm */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(37),		/* |, reduce: SyntaxTerm */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			reduce(37),		/* [, reduce: SyntaxTerm */
			nil,		/* ] */
			reduce(37),		/* {, reduce: SyntaxTerm */
			reduce(37),		/* }, reduce: SyntaxTerm */
			reduce(37),		/* (, reduce: SyntaxTerm */
			nil,		/* ) */
			reduce(37),		/* actionLit, reduce: SyntaxTerm */
			reduce(37),		/* prodId, reduce: SyntaxTerm */
			nil,		/* error */
			reduce(37),		/* stringLit, reduce: SyntaxTerm */
			
		},

	},
	actionRow{ // S89
				canRecover: true,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(77),		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			nil,		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			shift(78),		/* [ */
			nil,		/* ] */
			shift(79),		/* { */
			nil,		/* } */
			shift(80),		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			shift(81),		/* prodId */
			shift(85),		/* error */
			shift(87),		/* stringLit */
			
		},

	},
	actionRow{ // S90
				canRecover: true,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(88),		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			nil,		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			shift(89),		/* [ */
			nil,		/* ] */
			shift(90),		/* { */
			nil,		/* } */
			shift(91),		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			shift(92),		/* prodId */
			shift(96),		/* error */
			shift(98),		/* stringLit */
			
		},

	},
	actionRow{ // S91
				canRecover: true,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(99),		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			nil,		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			shift(100),		/* [ */
			nil,		/* ] */
			shift(101),		/* { */
			nil,		/* } */
			shift(102),		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			shift(103),		/* prodId */
			shift(107),		/* error */
			shift(109),		/* stringLit */
			
		},

	},
	actionRow{ // S92
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(36),		/* tokId, reduce: SyntaxTerm */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(36),		/* |, reduce: SyntaxTerm */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			reduce(36),		/* [, reduce: SyntaxTerm */
			nil,		/* ] */
			reduce(36),		/* {, reduce: SyntaxTerm */
			reduce(36),		/* }, reduce: SyntaxTerm */
			reduce(36),		/* (, reduce: SyntaxTerm */
			nil,		/* ) */
			reduce(36),		/* actionLit, reduce: SyntaxTerm */
			reduce(36),		/* prodId, reduce: SyntaxTerm */
			nil,		/* error */
			reduce(36),		/* stringLit, reduce: SyntaxTerm */
			
		},

	},
	actionRow{ // S93
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			shift(149),		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			nil,		/* ] */
			nil,		/* { */
			shift(150),		/* } */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S94
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(27),		/* |, reduce: SyntaxExpression */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			nil,		/* ] */
			nil,		/* { */
			reduce(27),		/* }, reduce: SyntaxExpression */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S95
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(88),		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(29),		/* |, reduce: SyntaxBody */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			shift(89),		/* [ */
			nil,		/* ] */
			shift(90),		/* { */
			reduce(29),		/* }, reduce: SyntaxBody */
			shift(91),		/* ( */
			nil,		/* ) */
			shift(151),		/* actionLit */
			shift(92),		/* prodId */
			nil,		/* error */
			shift(98),		/* stringLit */
			
		},

	},
	actionRow{ // S96
				canRecover: true,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(88),		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(31),		/* |, reduce: SyntaxBody */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			shift(89),		/* [ */
			nil,		/* ] */
			shift(90),		/* { */
			reduce(31),		/* }, reduce: SyntaxBody */
			shift(91),		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			shift(92),		/* prodId */
			nil,		/* error */
			shift(98),		/* stringLit */
			
		},

	},
	actionRow{ // S97
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(34),		/* tokId, reduce: SyntaxTerms */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(34),		/* |, reduce: SyntaxTerms */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			reduce(34),		/* [, reduce: SyntaxTerms */
			nil,		/* ] */
			reduce(34),		/* {, reduce: SyntaxTerms */
			reduce(34),		/* }, reduce: SyntaxTerms */
			reduce(34),		/* (, reduce: SyntaxTerms */
			nil,		/* ) */
			reduce(34),		/* actionLit, reduce: SyntaxTerms */
			reduce(34),		/* prodId, reduce: SyntaxTerms */
			nil,		/* error */
			reduce(34),		/* stringLit, reduce: SyntaxTerms */
			
		},

	},
	actionRow{ // S98
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(38),		/* tokId, reduce: SyntaxTerm */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(38),		/* |, reduce: SyntaxTerm */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			reduce(38),		/* [, reduce: SyntaxTerm */
			nil,		/* ] */
			reduce(38),		/* {, reduce: SyntaxTerm */
			reduce(38),		/* }, reduce: SyntaxTerm */
			reduce(38),		/* (, reduce: SyntaxTerm */
			nil,		/* ) */
			reduce(38),		/* actionLit, reduce: SyntaxTerm */
			reduce(38),		/* prodId, reduce: SyntaxTerm */
			nil,		/* error */
			reduce(38),		/* stringLit, reduce: SyntaxTerm */
			
		},

	},
	actionRow{ // S99
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(37),		/* tokId, reduce: SyntaxTerm */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(37),		/* |, reduce: SyntaxTerm */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			reduce(37),		/* [, reduce: SyntaxTerm */
			nil,		/* ] */
			reduce(37),		/* {, reduce: SyntaxTerm */
			nil,		/* } */
			reduce(37),		/* (, reduce: SyntaxTerm */
			reduce(37),		/* ), reduce: SyntaxTerm */
			reduce(37),		/* actionLit, reduce: SyntaxTerm */
			reduce(37),		/* prodId, reduce: SyntaxTerm */
			nil,		/* error */
			reduce(37),		/* stringLit, reduce: SyntaxTerm */
			
		},

	},
	actionRow{ // S100
				canRecover: true,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(77),		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			nil,		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			shift(78),		/* [ */
			nil,		/* ] */
			shift(79),		/* { */
			nil,		/* } */
			shift(80),		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			shift(81),		/* prodId */
			shift(85),		/* error */
			shift(87),		/* stringLit */
			
		},

	},
	actionRow{ // S101
				canRecover: true,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(88),		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			nil,		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			shift(89),		/* [ */
			nil,		/* ] */
			shift(90),		/* { */
			nil,		/* } */
			shift(91),		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			shift(92),		/* prodId */
			shift(96),		/* error */
			shift(98),		/* stringLit */
			
		},

	},
	actionRow{ // S102
				canRecover: true,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(99),		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			nil,		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			shift(100),		/* [ */
			nil,		/* ] */
			shift(101),		/* { */
			nil,		/* } */
			shift(102),		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			shift(103),		/* prodId */
			shift(107),		/* error */
			shift(109),		/* stringLit */
			
		},

	},
	actionRow{ // S103
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(36),		/* tokId, reduce: SyntaxTerm */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(36),		/* |, reduce: SyntaxTerm */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			reduce(36),		/* [, reduce: SyntaxTerm */
			nil,		/* ] */
			reduce(36),		/* {, reduce: SyntaxTerm */
			nil,		/* } */
			reduce(36),		/* (, reduce: SyntaxTerm */
			reduce(36),		/* ), reduce: SyntaxTerm */
			reduce(36),		/* actionLit, reduce: SyntaxTerm */
			reduce(36),		/* prodId, reduce: SyntaxTerm */
			nil,		/* error */
			reduce(36),		/* stringLit, reduce: SyntaxTerm */
			
		},

	},
	actionRow{ // S104
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			shift(157),		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			nil,		/* ] */
			nil,		/* { */
			nil,		/* } */
			nil,		/* ( */
			shift(158),		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S105
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(27),		/* |, reduce: SyntaxExpression */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			nil,		/* ] */
			nil,		/* { */
			nil,		/* } */
			nil,		/* ( */
			reduce(27),		/* ), reduce: SyntaxExpression */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S106
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(99),		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(29),		/* |, reduce: SyntaxBody */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			shift(100),		/* [ */
			nil,		/* ] */
			shift(101),		/* { */
			nil,		/* } */
			shift(102),		/* ( */
			reduce(29),		/* ), reduce: SyntaxBody */
			shift(159),		/* actionLit */
			shift(103),		/* prodId */
			nil,		/* error */
			shift(109),		/* stringLit */
			
		},

	},
	actionRow{ // S107
				canRecover: true,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(99),		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(31),		/* |, reduce: SyntaxBody */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			shift(100),		/* [ */
			nil,		/* ] */
			shift(101),		/* { */
			nil,		/* } */
			shift(102),		/* ( */
			reduce(31),		/* ), reduce: SyntaxBody */
			nil,		/* actionLit */
			shift(103),		/* prodId */
			nil,		/* error */
			shift(109),		/* stringLit */
			
		},

	},
	actionRow{ // S108
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(34),		/* tokId, reduce: SyntaxTerms */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(34),		/* |, reduce: SyntaxTerms */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			reduce(34),		/* [, reduce: SyntaxTerms */
			nil,		/* ] */
			reduce(34),		/* {, reduce: SyntaxTerms */
			nil,		/* } */
			reduce(34),		/* (, reduce: SyntaxTerms */
			reduce(34),		/* ), reduce: SyntaxTerms */
			reduce(34),		/* actionLit, reduce: SyntaxTerms */
			reduce(34),		/* prodId, reduce: SyntaxTerms */
			nil,		/* error */
			reduce(34),		/* stringLit, reduce: SyntaxTerms */
			
		},

	},
	actionRow{ // S109
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(38),		/* tokId, reduce: SyntaxTerm */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(38),		/* |, reduce: SyntaxTerm */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			reduce(38),		/* [, reduce: SyntaxTerm */
			nil,		/* ] */
			reduce(38),		/* {, reduce: SyntaxTerm */
			nil,		/* } */
			reduce(38),		/* (, reduce: SyntaxTerm */
			reduce(38),		/* ), reduce: SyntaxTerm */
			reduce(38),		/* actionLit, reduce: SyntaxTerm */
			reduce(38),		/* prodId, reduce: SyntaxTerm */
			nil,		/* error */
			reduce(38),		/* stringLit, reduce: SyntaxTerm */
			
		},

	},
	actionRow{ // S110
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(26),		/* $, reduce: SyntaxProduction */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			nil,		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			nil,		/* ] */
			nil,		/* { */
			nil,		/* } */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			reduce(26),		/* prodId, reduce: SyntaxProduction */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S111
				canRecover: true,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(33),		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			nil,		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			shift(34),		/* [ */
			nil,		/* ] */
			shift(35),		/* { */
			nil,		/* } */
			shift(36),		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			shift(37),		/* prodId */
			shift(41),		/* error */
			shift(43),		/* stringLit */
			
		},

	},
	actionRow{ // S112
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			reduce(30),		/* ;, reduce: SyntaxBody */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(30),		/* |, reduce: SyntaxBody */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			nil,		/* ] */
			nil,		/* { */
			nil,		/* } */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S113
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(35),		/* tokId, reduce: SyntaxTerms */
			nil,		/* : */
			reduce(35),		/* ;, reduce: SyntaxTerms */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(35),		/* |, reduce: SyntaxTerms */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			reduce(35),		/* [, reduce: SyntaxTerms */
			nil,		/* ] */
			reduce(35),		/* {, reduce: SyntaxTerms */
			nil,		/* } */
			reduce(35),		/* (, reduce: SyntaxTerms */
			nil,		/* ) */
			reduce(35),		/* actionLit, reduce: SyntaxTerms */
			reduce(35),		/* prodId, reduce: SyntaxTerms */
			nil,		/* error */
			reduce(35),		/* stringLit, reduce: SyntaxTerms */
			
		},

	},
	actionRow{ // S114
				canRecover: true,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(33),		/* tokId */
			nil,		/* : */
			reduce(32),		/* ;, reduce: SyntaxBody */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(32),		/* |, reduce: SyntaxBody */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			shift(34),		/* [ */
			nil,		/* ] */
			shift(35),		/* { */
			nil,		/* } */
			shift(36),		/* ( */
			nil,		/* ) */
			shift(163),		/* actionLit */
			shift(37),		/* prodId */
			nil,		/* error */
			shift(43),		/* stringLit */
			
		},

	},
	actionRow{ // S115
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			reduce(11),		/* ;, reduce: LexPattern */
			shift(23),		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(11),		/* |, reduce: LexPattern */
			shift(26),		/* . */
			shift(27),		/* char_lit */
			nil,		/* - */
			shift(28),		/* [ */
			nil,		/* ] */
			shift(29),		/* { */
			nil,		/* } */
			shift(30),		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S116
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			reduce(16),		/* ;, reduce: LexTerm */
			reduce(16),		/* regDefId, reduce: LexTerm */
			nil,		/* ignoredTokId */
			reduce(16),		/* |, reduce: LexTerm */
			reduce(16),		/* ., reduce: LexTerm */
			reduce(16),		/* char_lit, reduce: LexTerm */
			nil,		/* - */
			reduce(16),		/* [, reduce: LexTerm */
			nil,		/* ] */
			reduce(16),		/* {, reduce: LexTerm */
			nil,		/* } */
			reduce(16),		/* (, reduce: LexTerm */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S117
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			shift(49),		/* regDefId */
			nil,		/* ignoredTokId */
			nil,		/* | */
			shift(52),		/* . */
			shift(53),		/* char_lit */
			nil,		/* - */
			shift(54),		/* [ */
			nil,		/* ] */
			shift(55),		/* { */
			nil,		/* } */
			shift(56),		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S118
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			reduce(18),		/* ;, reduce: LexTerm */
			reduce(18),		/* regDefId, reduce: LexTerm */
			nil,		/* ignoredTokId */
			reduce(18),		/* |, reduce: LexTerm */
			reduce(18),		/* ., reduce: LexTerm */
			reduce(18),		/* char_lit, reduce: LexTerm */
			nil,		/* - */
			reduce(18),		/* [, reduce: LexTerm */
			nil,		/* ] */
			reduce(18),		/* {, reduce: LexTerm */
			nil,		/* } */
			reduce(18),		/* (, reduce: LexTerm */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S119
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			reduce(13),		/* regDefId, reduce: LexAlt */
			nil,		/* ignoredTokId */
			reduce(13),		/* |, reduce: LexAlt */
			reduce(13),		/* ., reduce: LexAlt */
			reduce(13),		/* char_lit, reduce: LexAlt */
			nil,		/* - */
			reduce(13),		/* [, reduce: LexAlt */
			reduce(13),		/* ], reduce: LexAlt */
			reduce(13),		/* {, reduce: LexAlt */
			nil,		/* } */
			reduce(13),		/* (, reduce: LexAlt */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S120
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			nil,		/* | */
			nil,		/* . */
			shift(165),		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			nil,		/* ] */
			nil,		/* { */
			nil,		/* } */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S121
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			shift(117),		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			shift(166),		/* ] */
			nil,		/* { */
			nil,		/* } */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S122
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			shift(124),		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			nil,		/* ] */
			nil,		/* { */
			shift(167),		/* } */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S123
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			shift(131),		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			nil,		/* ] */
			nil,		/* { */
			nil,		/* } */
			nil,		/* ( */
			shift(168),		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S124
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			shift(58),		/* regDefId */
			nil,		/* ignoredTokId */
			nil,		/* | */
			shift(61),		/* . */
			shift(62),		/* char_lit */
			nil,		/* - */
			shift(63),		/* [ */
			nil,		/* ] */
			shift(64),		/* { */
			nil,		/* } */
			shift(65),		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S125
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			reduce(19),		/* ;, reduce: LexTerm */
			reduce(19),		/* regDefId, reduce: LexTerm */
			nil,		/* ignoredTokId */
			reduce(19),		/* |, reduce: LexTerm */
			reduce(19),		/* ., reduce: LexTerm */
			reduce(19),		/* char_lit, reduce: LexTerm */
			nil,		/* - */
			reduce(19),		/* [, reduce: LexTerm */
			nil,		/* ] */
			reduce(19),		/* {, reduce: LexTerm */
			nil,		/* } */
			reduce(19),		/* (, reduce: LexTerm */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S126
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			reduce(13),		/* regDefId, reduce: LexAlt */
			nil,		/* ignoredTokId */
			reduce(13),		/* |, reduce: LexAlt */
			reduce(13),		/* ., reduce: LexAlt */
			reduce(13),		/* char_lit, reduce: LexAlt */
			nil,		/* - */
			reduce(13),		/* [, reduce: LexAlt */
			nil,		/* ] */
			reduce(13),		/* {, reduce: LexAlt */
			reduce(13),		/* }, reduce: LexAlt */
			reduce(13),		/* (, reduce: LexAlt */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S127
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			nil,		/* | */
			nil,		/* . */
			shift(170),		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			nil,		/* ] */
			nil,		/* { */
			nil,		/* } */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S128
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			shift(117),		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			shift(171),		/* ] */
			nil,		/* { */
			nil,		/* } */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S129
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			shift(124),		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			nil,		/* ] */
			nil,		/* { */
			shift(172),		/* } */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S130
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			shift(131),		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			nil,		/* ] */
			nil,		/* { */
			nil,		/* } */
			nil,		/* ( */
			shift(173),		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S131
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			shift(67),		/* regDefId */
			nil,		/* ignoredTokId */
			nil,		/* | */
			shift(70),		/* . */
			shift(71),		/* char_lit */
			nil,		/* - */
			shift(72),		/* [ */
			nil,		/* ] */
			shift(73),		/* { */
			nil,		/* } */
			shift(74),		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S132
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			reduce(20),		/* ;, reduce: LexTerm */
			reduce(20),		/* regDefId, reduce: LexTerm */
			nil,		/* ignoredTokId */
			reduce(20),		/* |, reduce: LexTerm */
			reduce(20),		/* ., reduce: LexTerm */
			reduce(20),		/* char_lit, reduce: LexTerm */
			nil,		/* - */
			reduce(20),		/* [, reduce: LexTerm */
			nil,		/* ] */
			reduce(20),		/* {, reduce: LexTerm */
			nil,		/* } */
			reduce(20),		/* (, reduce: LexTerm */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S133
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			reduce(13),		/* regDefId, reduce: LexAlt */
			nil,		/* ignoredTokId */
			reduce(13),		/* |, reduce: LexAlt */
			reduce(13),		/* ., reduce: LexAlt */
			reduce(13),		/* char_lit, reduce: LexAlt */
			nil,		/* - */
			reduce(13),		/* [, reduce: LexAlt */
			nil,		/* ] */
			reduce(13),		/* {, reduce: LexAlt */
			nil,		/* } */
			reduce(13),		/* (, reduce: LexAlt */
			reduce(13),		/* ), reduce: LexAlt */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S134
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			nil,		/* | */
			nil,		/* . */
			shift(175),		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			nil,		/* ] */
			nil,		/* { */
			nil,		/* } */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S135
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			shift(117),		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			shift(176),		/* ] */
			nil,		/* { */
			nil,		/* } */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S136
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			shift(124),		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			nil,		/* ] */
			nil,		/* { */
			shift(177),		/* } */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S137
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			shift(131),		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			nil,		/* ] */
			nil,		/* { */
			nil,		/* } */
			nil,		/* ( */
			shift(178),		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S138
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			shift(141),		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			shift(179),		/* ] */
			nil,		/* { */
			nil,		/* } */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S139
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			shift(149),		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			nil,		/* ] */
			nil,		/* { */
			shift(180),		/* } */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S140
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			shift(157),		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			nil,		/* ] */
			nil,		/* { */
			nil,		/* } */
			nil,		/* ( */
			shift(181),		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S141
				canRecover: true,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(77),		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			nil,		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			shift(78),		/* [ */
			nil,		/* ] */
			shift(79),		/* { */
			nil,		/* } */
			shift(80),		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			shift(81),		/* prodId */
			shift(85),		/* error */
			shift(87),		/* stringLit */
			
		},

	},
	actionRow{ // S142
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(40),		/* tokId, reduce: SyntaxTerm */
			nil,		/* : */
			reduce(40),		/* ;, reduce: SyntaxTerm */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(40),		/* |, reduce: SyntaxTerm */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			reduce(40),		/* [, reduce: SyntaxTerm */
			nil,		/* ] */
			reduce(40),		/* {, reduce: SyntaxTerm */
			nil,		/* } */
			reduce(40),		/* (, reduce: SyntaxTerm */
			nil,		/* ) */
			reduce(40),		/* actionLit, reduce: SyntaxTerm */
			reduce(40),		/* prodId, reduce: SyntaxTerm */
			nil,		/* error */
			reduce(40),		/* stringLit, reduce: SyntaxTerm */
			
		},

	},
	actionRow{ // S143
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(30),		/* |, reduce: SyntaxBody */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			reduce(30),		/* ], reduce: SyntaxBody */
			nil,		/* { */
			nil,		/* } */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S144
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(35),		/* tokId, reduce: SyntaxTerms */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(35),		/* |, reduce: SyntaxTerms */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			reduce(35),		/* [, reduce: SyntaxTerms */
			reduce(35),		/* ], reduce: SyntaxTerms */
			reduce(35),		/* {, reduce: SyntaxTerms */
			nil,		/* } */
			reduce(35),		/* (, reduce: SyntaxTerms */
			nil,		/* ) */
			reduce(35),		/* actionLit, reduce: SyntaxTerms */
			reduce(35),		/* prodId, reduce: SyntaxTerms */
			nil,		/* error */
			reduce(35),		/* stringLit, reduce: SyntaxTerms */
			
		},

	},
	actionRow{ // S145
				canRecover: true,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(77),		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(32),		/* |, reduce: SyntaxBody */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			shift(78),		/* [ */
			reduce(32),		/* ], reduce: SyntaxBody */
			shift(79),		/* { */
			nil,		/* } */
			shift(80),		/* ( */
			nil,		/* ) */
			shift(183),		/* actionLit */
			shift(81),		/* prodId */
			nil,		/* error */
			shift(87),		/* stringLit */
			
		},

	},
	actionRow{ // S146
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			shift(141),		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			shift(184),		/* ] */
			nil,		/* { */
			nil,		/* } */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S147
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			shift(149),		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			nil,		/* ] */
			nil,		/* { */
			shift(185),		/* } */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S148
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			shift(157),		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			nil,		/* ] */
			nil,		/* { */
			nil,		/* } */
			nil,		/* ( */
			shift(186),		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S149
				canRecover: true,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(88),		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			nil,		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			shift(89),		/* [ */
			nil,		/* ] */
			shift(90),		/* { */
			nil,		/* } */
			shift(91),		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			shift(92),		/* prodId */
			shift(96),		/* error */
			shift(98),		/* stringLit */
			
		},

	},
	actionRow{ // S150
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(41),		/* tokId, reduce: SyntaxTerm */
			nil,		/* : */
			reduce(41),		/* ;, reduce: SyntaxTerm */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(41),		/* |, reduce: SyntaxTerm */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			reduce(41),		/* [, reduce: SyntaxTerm */
			nil,		/* ] */
			reduce(41),		/* {, reduce: SyntaxTerm */
			nil,		/* } */
			reduce(41),		/* (, reduce: SyntaxTerm */
			nil,		/* ) */
			reduce(41),		/* actionLit, reduce: SyntaxTerm */
			reduce(41),		/* prodId, reduce: SyntaxTerm */
			nil,		/* error */
			reduce(41),		/* stringLit, reduce: SyntaxTerm */
			
		},

	},
	actionRow{ // S151
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(30),		/* |, reduce: SyntaxBody */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			nil,		/* ] */
			nil,		/* { */
			reduce(30),		/* }, reduce: SyntaxBody */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S152
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(35),		/* tokId, reduce: SyntaxTerms */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(35),		/* |, reduce: SyntaxTerms */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			reduce(35),		/* [, reduce: SyntaxTerms */
			nil,		/* ] */
			reduce(35),		/* {, reduce: SyntaxTerms */
			reduce(35),		/* }, reduce: SyntaxTerms */
			reduce(35),		/* (, reduce: SyntaxTerms */
			nil,		/* ) */
			reduce(35),		/* actionLit, reduce: SyntaxTerms */
			reduce(35),		/* prodId, reduce: SyntaxTerms */
			nil,		/* error */
			reduce(35),		/* stringLit, reduce: SyntaxTerms */
			
		},

	},
	actionRow{ // S153
				canRecover: true,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(88),		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(32),		/* |, reduce: SyntaxBody */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			shift(89),		/* [ */
			nil,		/* ] */
			shift(90),		/* { */
			reduce(32),		/* }, reduce: SyntaxBody */
			shift(91),		/* ( */
			nil,		/* ) */
			shift(188),		/* actionLit */
			shift(92),		/* prodId */
			nil,		/* error */
			shift(98),		/* stringLit */
			
		},

	},
	actionRow{ // S154
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			shift(141),		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			shift(189),		/* ] */
			nil,		/* { */
			nil,		/* } */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S155
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			shift(149),		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			nil,		/* ] */
			nil,		/* { */
			shift(190),		/* } */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S156
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			shift(157),		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			nil,		/* ] */
			nil,		/* { */
			nil,		/* } */
			nil,		/* ( */
			shift(191),		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S157
				canRecover: true,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(99),		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			nil,		/* | */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			shift(100),		/* [ */
			nil,		/* ] */
			shift(101),		/* { */
			nil,		/* } */
			shift(102),		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			shift(103),		/* prodId */
			shift(107),		/* error */
			shift(109),		/* stringLit */
			
		},

	},
	actionRow{ // S158
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(39),		/* tokId, reduce: SyntaxTerm */
			nil,		/* : */
			reduce(39),		/* ;, reduce: SyntaxTerm */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(39),		/* |, reduce: SyntaxTerm */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			reduce(39),		/* [, reduce: SyntaxTerm */
			nil,		/* ] */
			reduce(39),		/* {, reduce: SyntaxTerm */
			nil,		/* } */
			reduce(39),		/* (, reduce: SyntaxTerm */
			nil,		/* ) */
			reduce(39),		/* actionLit, reduce: SyntaxTerm */
			reduce(39),		/* prodId, reduce: SyntaxTerm */
			nil,		/* error */
			reduce(39),		/* stringLit, reduce: SyntaxTerm */
			
		},

	},
	actionRow{ // S159
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(30),		/* |, reduce: SyntaxBody */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			nil,		/* ] */
			nil,		/* { */
			nil,		/* } */
			nil,		/* ( */
			reduce(30),		/* ), reduce: SyntaxBody */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S160
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(35),		/* tokId, reduce: SyntaxTerms */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(35),		/* |, reduce: SyntaxTerms */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			reduce(35),		/* [, reduce: SyntaxTerms */
			nil,		/* ] */
			reduce(35),		/* {, reduce: SyntaxTerms */
			nil,		/* } */
			reduce(35),		/* (, reduce: SyntaxTerms */
			reduce(35),		/* ), reduce: SyntaxTerms */
			reduce(35),		/* actionLit, reduce: SyntaxTerms */
			reduce(35),		/* prodId, reduce: SyntaxTerms */
			nil,		/* error */
			reduce(35),		/* stringLit, reduce: SyntaxTerms */
			
		},

	},
	actionRow{ // S161
				canRecover: true,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(99),		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(32),		/* |, reduce: SyntaxBody */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			shift(100),		/* [ */
			nil,		/* ] */
			shift(101),		/* { */
			nil,		/* } */
			shift(102),		/* ( */
			reduce(32),		/* ), reduce: SyntaxBody */
			shift(193),		/* actionLit */
			shift(103),		/* prodId */
			nil,		/* error */
			shift(109),		/* stringLit */
			
		},

	},
	actionRow{ // S162
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			reduce(28),		/* ;, reduce: SyntaxExpression */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(28),		/* |, reduce: SyntaxExpression */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			nil,		/* ] */
			nil,		/* { */
			nil,		/* } */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S163
				canRecover: true,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			reduce(33),		/* ;, reduce: SyntaxBody */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(33),		/* |, reduce: SyntaxBody */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			nil,		/* ] */
			nil,		/* { */
			nil,		/* } */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S164
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			shift(49),		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(11),		/* |, reduce: LexPattern */
			shift(52),		/* . */
			shift(53),		/* char_lit */
			nil,		/* - */
			shift(54),		/* [ */
			reduce(11),		/* ], reduce: LexPattern */
			shift(55),		/* { */
			nil,		/* } */
			shift(56),		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S165
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			reduce(16),		/* regDefId, reduce: LexTerm */
			nil,		/* ignoredTokId */
			reduce(16),		/* |, reduce: LexTerm */
			reduce(16),		/* ., reduce: LexTerm */
			reduce(16),		/* char_lit, reduce: LexTerm */
			nil,		/* - */
			reduce(16),		/* [, reduce: LexTerm */
			reduce(16),		/* ], reduce: LexTerm */
			reduce(16),		/* {, reduce: LexTerm */
			nil,		/* } */
			reduce(16),		/* (, reduce: LexTerm */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S166
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			reduce(18),		/* regDefId, reduce: LexTerm */
			nil,		/* ignoredTokId */
			reduce(18),		/* |, reduce: LexTerm */
			reduce(18),		/* ., reduce: LexTerm */
			reduce(18),		/* char_lit, reduce: LexTerm */
			nil,		/* - */
			reduce(18),		/* [, reduce: LexTerm */
			reduce(18),		/* ], reduce: LexTerm */
			reduce(18),		/* {, reduce: LexTerm */
			nil,		/* } */
			reduce(18),		/* (, reduce: LexTerm */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S167
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			reduce(19),		/* regDefId, reduce: LexTerm */
			nil,		/* ignoredTokId */
			reduce(19),		/* |, reduce: LexTerm */
			reduce(19),		/* ., reduce: LexTerm */
			reduce(19),		/* char_lit, reduce: LexTerm */
			nil,		/* - */
			reduce(19),		/* [, reduce: LexTerm */
			reduce(19),		/* ], reduce: LexTerm */
			reduce(19),		/* {, reduce: LexTerm */
			nil,		/* } */
			reduce(19),		/* (, reduce: LexTerm */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S168
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			reduce(20),		/* regDefId, reduce: LexTerm */
			nil,		/* ignoredTokId */
			reduce(20),		/* |, reduce: LexTerm */
			reduce(20),		/* ., reduce: LexTerm */
			reduce(20),		/* char_lit, reduce: LexTerm */
			nil,		/* - */
			reduce(20),		/* [, reduce: LexTerm */
			reduce(20),		/* ], reduce: LexTerm */
			reduce(20),		/* {, reduce: LexTerm */
			nil,		/* } */
			reduce(20),		/* (, reduce: LexTerm */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S169
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			shift(58),		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(11),		/* |, reduce: LexPattern */
			shift(61),		/* . */
			shift(62),		/* char_lit */
			nil,		/* - */
			shift(63),		/* [ */
			nil,		/* ] */
			shift(64),		/* { */
			reduce(11),		/* }, reduce: LexPattern */
			shift(65),		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S170
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			reduce(16),		/* regDefId, reduce: LexTerm */
			nil,		/* ignoredTokId */
			reduce(16),		/* |, reduce: LexTerm */
			reduce(16),		/* ., reduce: LexTerm */
			reduce(16),		/* char_lit, reduce: LexTerm */
			nil,		/* - */
			reduce(16),		/* [, reduce: LexTerm */
			nil,		/* ] */
			reduce(16),		/* {, reduce: LexTerm */
			reduce(16),		/* }, reduce: LexTerm */
			reduce(16),		/* (, reduce: LexTerm */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S171
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			reduce(18),		/* regDefId, reduce: LexTerm */
			nil,		/* ignoredTokId */
			reduce(18),		/* |, reduce: LexTerm */
			reduce(18),		/* ., reduce: LexTerm */
			reduce(18),		/* char_lit, reduce: LexTerm */
			nil,		/* - */
			reduce(18),		/* [, reduce: LexTerm */
			nil,		/* ] */
			reduce(18),		/* {, reduce: LexTerm */
			reduce(18),		/* }, reduce: LexTerm */
			reduce(18),		/* (, reduce: LexTerm */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S172
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			reduce(19),		/* regDefId, reduce: LexTerm */
			nil,		/* ignoredTokId */
			reduce(19),		/* |, reduce: LexTerm */
			reduce(19),		/* ., reduce: LexTerm */
			reduce(19),		/* char_lit, reduce: LexTerm */
			nil,		/* - */
			reduce(19),		/* [, reduce: LexTerm */
			nil,		/* ] */
			reduce(19),		/* {, reduce: LexTerm */
			reduce(19),		/* }, reduce: LexTerm */
			reduce(19),		/* (, reduce: LexTerm */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S173
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			reduce(20),		/* regDefId, reduce: LexTerm */
			nil,		/* ignoredTokId */
			reduce(20),		/* |, reduce: LexTerm */
			reduce(20),		/* ., reduce: LexTerm */
			reduce(20),		/* char_lit, reduce: LexTerm */
			nil,		/* - */
			reduce(20),		/* [, reduce: LexTerm */
			nil,		/* ] */
			reduce(20),		/* {, reduce: LexTerm */
			reduce(20),		/* }, reduce: LexTerm */
			reduce(20),		/* (, reduce: LexTerm */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S174
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			shift(67),		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(11),		/* |, reduce: LexPattern */
			shift(70),		/* . */
			shift(71),		/* char_lit */
			nil,		/* - */
			shift(72),		/* [ */
			nil,		/* ] */
			shift(73),		/* { */
			nil,		/* } */
			shift(74),		/* ( */
			reduce(11),		/* ), reduce: LexPattern */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S175
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			reduce(16),		/* regDefId, reduce: LexTerm */
			nil,		/* ignoredTokId */
			reduce(16),		/* |, reduce: LexTerm */
			reduce(16),		/* ., reduce: LexTerm */
			reduce(16),		/* char_lit, reduce: LexTerm */
			nil,		/* - */
			reduce(16),		/* [, reduce: LexTerm */
			nil,		/* ] */
			reduce(16),		/* {, reduce: LexTerm */
			nil,		/* } */
			reduce(16),		/* (, reduce: LexTerm */
			reduce(16),		/* ), reduce: LexTerm */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S176
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			reduce(18),		/* regDefId, reduce: LexTerm */
			nil,		/* ignoredTokId */
			reduce(18),		/* |, reduce: LexTerm */
			reduce(18),		/* ., reduce: LexTerm */
			reduce(18),		/* char_lit, reduce: LexTerm */
			nil,		/* - */
			reduce(18),		/* [, reduce: LexTerm */
			nil,		/* ] */
			reduce(18),		/* {, reduce: LexTerm */
			nil,		/* } */
			reduce(18),		/* (, reduce: LexTerm */
			reduce(18),		/* ), reduce: LexTerm */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S177
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			reduce(19),		/* regDefId, reduce: LexTerm */
			nil,		/* ignoredTokId */
			reduce(19),		/* |, reduce: LexTerm */
			reduce(19),		/* ., reduce: LexTerm */
			reduce(19),		/* char_lit, reduce: LexTerm */
			nil,		/* - */
			reduce(19),		/* [, reduce: LexTerm */
			nil,		/* ] */
			reduce(19),		/* {, reduce: LexTerm */
			nil,		/* } */
			reduce(19),		/* (, reduce: LexTerm */
			reduce(19),		/* ), reduce: LexTerm */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S178
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			reduce(20),		/* regDefId, reduce: LexTerm */
			nil,		/* ignoredTokId */
			reduce(20),		/* |, reduce: LexTerm */
			reduce(20),		/* ., reduce: LexTerm */
			reduce(20),		/* char_lit, reduce: LexTerm */
			nil,		/* - */
			reduce(20),		/* [, reduce: LexTerm */
			nil,		/* ] */
			reduce(20),		/* {, reduce: LexTerm */
			nil,		/* } */
			reduce(20),		/* (, reduce: LexTerm */
			reduce(20),		/* ), reduce: LexTerm */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S179
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(40),		/* tokId, reduce: SyntaxTerm */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(40),		/* |, reduce: SyntaxTerm */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			reduce(40),		/* [, reduce: SyntaxTerm */
			reduce(40),		/* ], reduce: SyntaxTerm */
			reduce(40),		/* {, reduce: SyntaxTerm */
			nil,		/* } */
			reduce(40),		/* (, reduce: SyntaxTerm */
			nil,		/* ) */
			reduce(40),		/* actionLit, reduce: SyntaxTerm */
			reduce(40),		/* prodId, reduce: SyntaxTerm */
			nil,		/* error */
			reduce(40),		/* stringLit, reduce: SyntaxTerm */
			
		},

	},
	actionRow{ // S180
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(41),		/* tokId, reduce: SyntaxTerm */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(41),		/* |, reduce: SyntaxTerm */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			reduce(41),		/* [, reduce: SyntaxTerm */
			reduce(41),		/* ], reduce: SyntaxTerm */
			reduce(41),		/* {, reduce: SyntaxTerm */
			nil,		/* } */
			reduce(41),		/* (, reduce: SyntaxTerm */
			nil,		/* ) */
			reduce(41),		/* actionLit, reduce: SyntaxTerm */
			reduce(41),		/* prodId, reduce: SyntaxTerm */
			nil,		/* error */
			reduce(41),		/* stringLit, reduce: SyntaxTerm */
			
		},

	},
	actionRow{ // S181
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(39),		/* tokId, reduce: SyntaxTerm */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(39),		/* |, reduce: SyntaxTerm */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			reduce(39),		/* [, reduce: SyntaxTerm */
			reduce(39),		/* ], reduce: SyntaxTerm */
			reduce(39),		/* {, reduce: SyntaxTerm */
			nil,		/* } */
			reduce(39),		/* (, reduce: SyntaxTerm */
			nil,		/* ) */
			reduce(39),		/* actionLit, reduce: SyntaxTerm */
			reduce(39),		/* prodId, reduce: SyntaxTerm */
			nil,		/* error */
			reduce(39),		/* stringLit, reduce: SyntaxTerm */
			
		},

	},
	actionRow{ // S182
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(28),		/* |, reduce: SyntaxExpression */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			reduce(28),		/* ], reduce: SyntaxExpression */
			nil,		/* { */
			nil,		/* } */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S183
				canRecover: true,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(33),		/* |, reduce: SyntaxBody */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			reduce(33),		/* ], reduce: SyntaxBody */
			nil,		/* { */
			nil,		/* } */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S184
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(40),		/* tokId, reduce: SyntaxTerm */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(40),		/* |, reduce: SyntaxTerm */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			reduce(40),		/* [, reduce: SyntaxTerm */
			nil,		/* ] */
			reduce(40),		/* {, reduce: SyntaxTerm */
			reduce(40),		/* }, reduce: SyntaxTerm */
			reduce(40),		/* (, reduce: SyntaxTerm */
			nil,		/* ) */
			reduce(40),		/* actionLit, reduce: SyntaxTerm */
			reduce(40),		/* prodId, reduce: SyntaxTerm */
			nil,		/* error */
			reduce(40),		/* stringLit, reduce: SyntaxTerm */
			
		},

	},
	actionRow{ // S185
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(41),		/* tokId, reduce: SyntaxTerm */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(41),		/* |, reduce: SyntaxTerm */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			reduce(41),		/* [, reduce: SyntaxTerm */
			nil,		/* ] */
			reduce(41),		/* {, reduce: SyntaxTerm */
			reduce(41),		/* }, reduce: SyntaxTerm */
			reduce(41),		/* (, reduce: SyntaxTerm */
			nil,		/* ) */
			reduce(41),		/* actionLit, reduce: SyntaxTerm */
			reduce(41),		/* prodId, reduce: SyntaxTerm */
			nil,		/* error */
			reduce(41),		/* stringLit, reduce: SyntaxTerm */
			
		},

	},
	actionRow{ // S186
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(39),		/* tokId, reduce: SyntaxTerm */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(39),		/* |, reduce: SyntaxTerm */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			reduce(39),		/* [, reduce: SyntaxTerm */
			nil,		/* ] */
			reduce(39),		/* {, reduce: SyntaxTerm */
			reduce(39),		/* }, reduce: SyntaxTerm */
			reduce(39),		/* (, reduce: SyntaxTerm */
			nil,		/* ) */
			reduce(39),		/* actionLit, reduce: SyntaxTerm */
			reduce(39),		/* prodId, reduce: SyntaxTerm */
			nil,		/* error */
			reduce(39),		/* stringLit, reduce: SyntaxTerm */
			
		},

	},
	actionRow{ // S187
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(28),		/* |, reduce: SyntaxExpression */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			nil,		/* ] */
			nil,		/* { */
			reduce(28),		/* }, reduce: SyntaxExpression */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S188
				canRecover: true,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(33),		/* |, reduce: SyntaxBody */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			nil,		/* ] */
			nil,		/* { */
			reduce(33),		/* }, reduce: SyntaxBody */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S189
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(40),		/* tokId, reduce: SyntaxTerm */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(40),		/* |, reduce: SyntaxTerm */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			reduce(40),		/* [, reduce: SyntaxTerm */
			nil,		/* ] */
			reduce(40),		/* {, reduce: SyntaxTerm */
			nil,		/* } */
			reduce(40),		/* (, reduce: SyntaxTerm */
			reduce(40),		/* ), reduce: SyntaxTerm */
			reduce(40),		/* actionLit, reduce: SyntaxTerm */
			reduce(40),		/* prodId, reduce: SyntaxTerm */
			nil,		/* error */
			reduce(40),		/* stringLit, reduce: SyntaxTerm */
			
		},

	},
	actionRow{ // S190
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(41),		/* tokId, reduce: SyntaxTerm */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(41),		/* |, reduce: SyntaxTerm */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			reduce(41),		/* [, reduce: SyntaxTerm */
			nil,		/* ] */
			reduce(41),		/* {, reduce: SyntaxTerm */
			nil,		/* } */
			reduce(41),		/* (, reduce: SyntaxTerm */
			reduce(41),		/* ), reduce: SyntaxTerm */
			reduce(41),		/* actionLit, reduce: SyntaxTerm */
			reduce(41),		/* prodId, reduce: SyntaxTerm */
			nil,		/* error */
			reduce(41),		/* stringLit, reduce: SyntaxTerm */
			
		},

	},
	actionRow{ // S191
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(39),		/* tokId, reduce: SyntaxTerm */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(39),		/* |, reduce: SyntaxTerm */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			reduce(39),		/* [, reduce: SyntaxTerm */
			nil,		/* ] */
			reduce(39),		/* {, reduce: SyntaxTerm */
			nil,		/* } */
			reduce(39),		/* (, reduce: SyntaxTerm */
			reduce(39),		/* ), reduce: SyntaxTerm */
			reduce(39),		/* actionLit, reduce: SyntaxTerm */
			reduce(39),		/* prodId, reduce: SyntaxTerm */
			nil,		/* error */
			reduce(39),		/* stringLit, reduce: SyntaxTerm */
			
		},

	},
	actionRow{ // S192
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(28),		/* |, reduce: SyntaxExpression */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			nil,		/* ] */
			nil,		/* { */
			nil,		/* } */
			nil,		/* ( */
			reduce(28),		/* ), reduce: SyntaxExpression */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	actionRow{ // S193
				canRecover: true,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* tokId */
			nil,		/* : */
			nil,		/* ; */
			nil,		/* regDefId */
			nil,		/* ignoredTokId */
			reduce(33),		/* |, reduce: SyntaxBody */
			nil,		/* . */
			nil,		/* char_lit */
			nil,		/* - */
			nil,		/* [ */
			nil,		/* ] */
			nil,		/* { */
			nil,		/* } */
			nil,		/* ( */
			reduce(33),		/* ), reduce: SyntaxBody */
			nil,		/* actionLit */
			nil,		/* prodId */
			nil,		/* error */
			nil,		/* stringLit */
			
		},

	},
	
}

