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
		case r == 97: // ['a','a']
			return 1

		}
		return NoState
	},

	// S1
	func(r rune) int {
		switch {

		}
		return NoState
	},
}
