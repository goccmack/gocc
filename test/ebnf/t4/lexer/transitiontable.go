package lexer

/*
Let s be the current state
Let r be the current input rune
transitionTable[s](r) returns the next state.
*/
type TransitionTable [NumStates]func(rune) int

var TransTab = TransitionTable{

	// S0
	func(r rune) int {
		switch {
		case r == 99: // ['c','c']
			return 1
		case r == 100: // ['d','d']
			return 2

		}
		return NoState
	},

	// S1
	func(r rune) int {
		switch {
		case r == 49: // ['1','1']
			return 3
		case r == 50: // ['2','2']
			return 4
		case r == 51: // ['3','3']
			return 5

		}
		return NoState
	},

	// S2
	func(r rune) int {
		switch {
		case r == 49: // ['1','1']
			return 6
		case r == 50: // ['2','2']
			return 7
		case r == 51: // ['3','3']
			return 8
		case r == 52: // ['4','4']
			return 9
		case r == 53: // ['5','5']
			return 10

		}
		return NoState
	},

	// S3
	func(r rune) int {
		switch {

		}
		return NoState
	},

	// S4
	func(r rune) int {
		switch {

		}
		return NoState
	},

	// S5
	func(r rune) int {
		switch {

		}
		return NoState
	},

	// S6
	func(r rune) int {
		switch {

		}
		return NoState
	},

	// S7
	func(r rune) int {
		switch {

		}
		return NoState
	},

	// S8
	func(r rune) int {
		switch {

		}
		return NoState
	},

	// S9
	func(r rune) int {
		switch {

		}
		return NoState
	},

	// S10
	func(r rune) int {
		switch {

		}
		return NoState
	},
}
