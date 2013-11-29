
package lexer



/*
Let s be the current state
Let r be the current input rune
transitionTable[s](r) returns the next state.
*/
type TransitionTable [NumStates] func(rune) int

var TransTab = TransitionTable{
	
		// S0
		func(r rune) int {
			switch {
			case r == 97 : // ['a','a']
				return 1
			case r == 98 : // ['b','b']
				return 2
			case r == 99 : // ['c','c']
				return 3
			case r == 100 : // ['d','d']
				return 4
			case r == 101 : // ['e','e']
				return 5
			case r == 116 : // ['t','t']
				return 6
			case r == 117 : // ['u','u']
				return 7
			
			
			
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
	
}
