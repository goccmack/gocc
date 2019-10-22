// Code generated by gocc; DO NOT EDIT.

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
		case r == 9: // ['\t','\t']
			return 1
		case r == 10: // ['\n','\n']
			return 1
		case r == 13: // ['\r','\r']
			return 1
		case r == 32: // [' ',' ']
			return 1
		case r == 945: // [\u03b1,\u03b1]
			return 2
		case r == 946: // [\u03b2,\u03b2]
			return 3
		}
		return NoState
	},
	// S1
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S2
	func(r rune) int {
		switch {
		case r == 945: // [\u03b1,\u03b1]
			return 4
		case r == 946: // [\u03b2,\u03b2]
			return 5
		case r == 947: // [\u03b3,\u03b3]
			return 6
		}
		return NoState
	},
	// S3
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 7
		}
		return NoState
	},
	// S4
	func(r rune) int {
		switch {
		case r == 945: // [\u03b1,\u03b1]
			return 4
		case r == 946: // [\u03b2,\u03b2]
			return 5
		case r == 947: // [\u03b3,\u03b3]
			return 6
		}
		return NoState
	},
	// S5
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 8
		}
		return NoState
	},
	// S6
	func(r rune) int {
		switch {
		case r == 945: // [\u03b1,\u03b1]
			return 4
		case r == 946: // [\u03b2,\u03b2]
			return 5
		case r == 947: // [\u03b3,\u03b3]
			return 6
		}
		return NoState
	},
	// S7
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 7
		case r == 945: // [\u03b1,\u03b1]
			return 9
		case r == 946: // [\u03b2,\u03b2]
			return 10
		case r == 947: // [\u03b3,\u03b3]
			return 6
		}
		return NoState
	},
	// S8
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 8
		case r == 945: // [\u03b1,\u03b1]
			return 4
		case r == 946: // [\u03b2,\u03b2]
			return 5
		case r == 947: // [\u03b3,\u03b3]
			return 6
		}
		return NoState
	},
	// S9
	func(r rune) int {
		switch {
		case r == 945: // [\u03b1,\u03b1]
			return 4
		case r == 946: // [\u03b2,\u03b2]
			return 5
		case r == 947: // [\u03b3,\u03b3]
			return 6
		}
		return NoState
	},
	// S10
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 11
		}
		return NoState
	},
	// S11
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 11
		case r == 945: // [\u03b1,\u03b1]
			return 9
		case r == 946: // [\u03b2,\u03b2]
			return 10
		case r == 947: // [\u03b3,\u03b3]
			return 6
		}
		return NoState
	},
}