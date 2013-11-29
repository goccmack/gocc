
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
			case r == 9 : // ['\t','\t']
				return 1
			case r == 10 : // ['\n','\n']
				return 1
			case r == 13 : // ['\r','\r']
				return 1
			case r == 32 : // [' ',' ']
				return 1
			case r == 59 : // [';',';']
				return 2
			case r == 116 : // ['t','t']
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
			case r == 10 : // ['\n','\n']
				return 4
			
			
			default:
				return 2
			
			}
			return NoState
		},
	
		// S3
		func(r rune) int {
			switch {
			case r == 111 : // ['o','o']
				return 5
			
			
			
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
			case r == 107 : // ['k','k']
				return 6
			
			
			
			}
			return NoState
		},
	
		// S6
		func(r rune) int {
			switch {
			case r == 101 : // ['e','e']
				return 7
			
			
			
			}
			return NoState
		},
	
		// S7
		func(r rune) int {
			switch {
			case r == 110 : // ['n','n']
				return 8
			
			
			
			}
			return NoState
		},
	
		// S8
		func(r rune) int {
			switch {
			case r == 49 : // ['1','1']
				return 9
			case r == 50 : // ['2','2']
				return 10
			case r == 51 : // ['3','3']
				return 11
			
			
			
			}
			return NoState
		},
	
		// S9
		func(r rune) int {
			switch {
			case r == 13 : // ['\r','\r']
				return 12
			
			
			
			}
			return NoState
		},
	
		// S10
		func(r rune) int {
			switch {
			case r == 13 : // ['\r','\r']
				return 12
			
			
			
			}
			return NoState
		},
	
		// S11
		func(r rune) int {
			switch {
			case r == 13 : // ['\r','\r']
				return 12
			
			
			
			}
			return NoState
		},
	
		// S12
		func(r rune) int {
			switch {
			case r == 10 : // ['\n','\n']
				return 13
			
			
			
			}
			return NoState
		},
	
		// S13
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
		},
	
}
