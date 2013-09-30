
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
			case r == 33 : // ['!','!']
				return 2
			case r == 34 : // ['"','"']
				return 3
			case r == 39 : // [''',''']
				return 4
			case r == 40 : // ['(','(']
				return 5
			case r == 41 : // [')',')']
				return 6
			case r == 45 : // ['-','-']
				return 7
			case r == 46 : // ['.','.']
				return 8
			case r == 47 : // ['/','/']
				return 9
			case r == 58 : // [':',':']
				return 10
			case r == 59 : // [';',';']
				return 11
			case r == 60 : // ['<','<']
				return 12
			case 65 <= r && r <= 90 : // ['A','Z']
				return 13
			case r == 91 : // ['[','[']
				return 14
			case r == 93 : // [']',']']
				return 15
			case r == 95 : // ['_','_']
				return 16
			case r == 96 : // ['`','`']
				return 17
			case 97 <= r && r <= 100 : // ['a','d']
				return 18
			case r == 101 : // ['e','e']
				return 19
			case 102 <= r && r <= 122 : // ['f','z']
				return 18
			case r == 123 : // ['{','{']
				return 20
			case r == 124 : // ['|','|']
				return 21
			case r == 125 : // ['}','}']
				return 22
			
			
			
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
			case 97 <= r && r <= 122 : // ['a','z']
				return 23
			
			
			
			}
			return NoState
		},
	
		// S3
		func(r rune) int {
			switch {
			case r == 34 : // ['"','"']
				return 24
			case r == 92 : // ['\','\']
				return 25
			
			
			default:
				return 26
			
			}
			return NoState
		},
	
		// S4
		func(r rune) int {
			switch {
			case r == 92 : // ['\','\']
				return 27
			
			
			default:
				return 28
			
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
			case r == 42 : // ['*','*']
				return 29
			case r == 47 : // ['/','/']
				return 30
			
			
			
			}
			return NoState
		},
	
		// S10
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
		},
	
		// S11
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
		},
	
		// S12
		func(r rune) int {
			switch {
			case r == 60 : // ['<','<']
				return 31
			
			
			
			}
			return NoState
		},
	
		// S13
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 32
			case 65 <= r && r <= 90 : // ['A','Z']
				return 33
			case r == 95 : // ['_','_']
				return 34
			case 97 <= r && r <= 122 : // ['a','z']
				return 35
			
			
			
			}
			return NoState
		},
	
		// S14
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
		},
	
		// S15
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
		},
	
		// S16
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 36
			case 65 <= r && r <= 90 : // ['A','Z']
				return 37
			case r == 95 : // ['_','_']
				return 38
			case 97 <= r && r <= 122 : // ['a','z']
				return 39
			
			
			
			}
			return NoState
		},
	
		// S17
		func(r rune) int {
			switch {
			case r == 96 : // ['`','`']
				return 40
			
			
			default:
				return 17
			
			}
			return NoState
		},
	
		// S18
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 41
			case 65 <= r && r <= 90 : // ['A','Z']
				return 42
			case r == 95 : // ['_','_']
				return 43
			case 97 <= r && r <= 122 : // ['a','z']
				return 44
			
			
			
			}
			return NoState
		},
	
		// S19
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 41
			case 65 <= r && r <= 90 : // ['A','Z']
				return 42
			case r == 95 : // ['_','_']
				return 43
			case 97 <= r && r <= 113 : // ['a','q']
				return 44
			case r == 114 : // ['r','r']
				return 45
			case 115 <= r && r <= 122 : // ['s','z']
				return 44
			
			
			
			}
			return NoState
		},
	
		// S20
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
		},
	
		// S21
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
		},
	
		// S22
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
		},
	
		// S23
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 46
			case 65 <= r && r <= 90 : // ['A','Z']
				return 47
			case r == 95 : // ['_','_']
				return 48
			case 97 <= r && r <= 122 : // ['a','z']
				return 49
			
			
			
			}
			return NoState
		},
	
		// S24
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
		},
	
		// S25
		func(r rune) int {
			switch {
			case r == 34 : // ['"','"']
				return 50
			case r == 39 : // [''',''']
				return 50
			case 48 <= r && r <= 55 : // ['0','7']
				return 51
			case r == 85 : // ['U','U']
				return 52
			case r == 92 : // ['\','\']
				return 50
			case r == 97 : // ['a','a']
				return 50
			case r == 98 : // ['b','b']
				return 50
			case r == 102 : // ['f','f']
				return 50
			case r == 110 : // ['n','n']
				return 50
			case r == 114 : // ['r','r']
				return 50
			case r == 116 : // ['t','t']
				return 50
			case r == 117 : // ['u','u']
				return 53
			case r == 118 : // ['v','v']
				return 50
			case r == 120 : // ['x','x']
				return 54
			
			
			
			}
			return NoState
		},
	
		// S26
		func(r rune) int {
			switch {
			case r == 34 : // ['"','"']
				return 24
			case r == 92 : // ['\','\']
				return 25
			
			
			default:
				return 26
			
			}
			return NoState
		},
	
		// S27
		func(r rune) int {
			switch {
			case r == 34 : // ['"','"']
				return 55
			case r == 39 : // [''',''']
				return 55
			case 48 <= r && r <= 55 : // ['0','7']
				return 56
			case r == 85 : // ['U','U']
				return 57
			case r == 92 : // ['\','\']
				return 55
			case r == 97 : // ['a','a']
				return 55
			case r == 98 : // ['b','b']
				return 55
			case r == 102 : // ['f','f']
				return 55
			case r == 110 : // ['n','n']
				return 55
			case r == 114 : // ['r','r']
				return 55
			case r == 116 : // ['t','t']
				return 55
			case r == 117 : // ['u','u']
				return 58
			case r == 118 : // ['v','v']
				return 55
			case r == 120 : // ['x','x']
				return 59
			
			
			
			}
			return NoState
		},
	
		// S28
		func(r rune) int {
			switch {
			case r == 39 : // [''',''']
				return 60
			
			
			
			}
			return NoState
		},
	
		// S29
		func(r rune) int {
			switch {
			case r == 42 : // ['*','*']
				return 61
			
			
			default:
				return 29
			
			}
			return NoState
		},
	
		// S30
		func(r rune) int {
			switch {
			case r == 10 : // ['\n','\n']
				return 62
			
			
			default:
				return 30
			
			}
			return NoState
		},
	
		// S31
		func(r rune) int {
			switch {
			
			
			default:
				return 63
			
			}
			return NoState
		},
	
		// S32
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 32
			case 65 <= r && r <= 90 : // ['A','Z']
				return 33
			case r == 95 : // ['_','_']
				return 34
			case 97 <= r && r <= 122 : // ['a','z']
				return 35
			
			
			
			}
			return NoState
		},
	
		// S33
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 32
			case 65 <= r && r <= 90 : // ['A','Z']
				return 33
			case r == 95 : // ['_','_']
				return 34
			case 97 <= r && r <= 122 : // ['a','z']
				return 35
			
			
			
			}
			return NoState
		},
	
		// S34
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 32
			case 65 <= r && r <= 90 : // ['A','Z']
				return 33
			case r == 95 : // ['_','_']
				return 34
			case 97 <= r && r <= 122 : // ['a','z']
				return 35
			
			
			
			}
			return NoState
		},
	
		// S35
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 32
			case 65 <= r && r <= 90 : // ['A','Z']
				return 33
			case r == 95 : // ['_','_']
				return 34
			case 97 <= r && r <= 122 : // ['a','z']
				return 35
			
			
			
			}
			return NoState
		},
	
		// S36
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 36
			case 65 <= r && r <= 90 : // ['A','Z']
				return 37
			case r == 95 : // ['_','_']
				return 38
			case 97 <= r && r <= 122 : // ['a','z']
				return 39
			
			
			
			}
			return NoState
		},
	
		// S37
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 36
			case 65 <= r && r <= 90 : // ['A','Z']
				return 37
			case r == 95 : // ['_','_']
				return 38
			case 97 <= r && r <= 122 : // ['a','z']
				return 39
			
			
			
			}
			return NoState
		},
	
		// S38
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 36
			case 65 <= r && r <= 90 : // ['A','Z']
				return 37
			case r == 95 : // ['_','_']
				return 38
			case 97 <= r && r <= 122 : // ['a','z']
				return 39
			
			
			
			}
			return NoState
		},
	
		// S39
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 36
			case 65 <= r && r <= 90 : // ['A','Z']
				return 37
			case r == 95 : // ['_','_']
				return 38
			case 97 <= r && r <= 122 : // ['a','z']
				return 39
			
			
			
			}
			return NoState
		},
	
		// S40
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
		},
	
		// S41
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 41
			case 65 <= r && r <= 90 : // ['A','Z']
				return 42
			case r == 95 : // ['_','_']
				return 43
			case 97 <= r && r <= 122 : // ['a','z']
				return 44
			
			
			
			}
			return NoState
		},
	
		// S42
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 41
			case 65 <= r && r <= 90 : // ['A','Z']
				return 42
			case r == 95 : // ['_','_']
				return 43
			case 97 <= r && r <= 122 : // ['a','z']
				return 44
			
			
			
			}
			return NoState
		},
	
		// S43
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 41
			case 65 <= r && r <= 90 : // ['A','Z']
				return 42
			case r == 95 : // ['_','_']
				return 43
			case 97 <= r && r <= 122 : // ['a','z']
				return 44
			
			
			
			}
			return NoState
		},
	
		// S44
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 41
			case 65 <= r && r <= 90 : // ['A','Z']
				return 42
			case r == 95 : // ['_','_']
				return 43
			case 97 <= r && r <= 122 : // ['a','z']
				return 44
			
			
			
			}
			return NoState
		},
	
		// S45
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 41
			case 65 <= r && r <= 90 : // ['A','Z']
				return 42
			case r == 95 : // ['_','_']
				return 43
			case 97 <= r && r <= 113 : // ['a','q']
				return 44
			case r == 114 : // ['r','r']
				return 64
			case 115 <= r && r <= 122 : // ['s','z']
				return 44
			
			
			
			}
			return NoState
		},
	
		// S46
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 46
			case 65 <= r && r <= 90 : // ['A','Z']
				return 47
			case r == 95 : // ['_','_']
				return 48
			case 97 <= r && r <= 122 : // ['a','z']
				return 49
			
			
			
			}
			return NoState
		},
	
		// S47
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 46
			case 65 <= r && r <= 90 : // ['A','Z']
				return 47
			case r == 95 : // ['_','_']
				return 48
			case 97 <= r && r <= 122 : // ['a','z']
				return 49
			
			
			
			}
			return NoState
		},
	
		// S48
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 46
			case 65 <= r && r <= 90 : // ['A','Z']
				return 47
			case r == 95 : // ['_','_']
				return 48
			case 97 <= r && r <= 122 : // ['a','z']
				return 49
			
			
			
			}
			return NoState
		},
	
		// S49
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 46
			case 65 <= r && r <= 90 : // ['A','Z']
				return 47
			case r == 95 : // ['_','_']
				return 48
			case 97 <= r && r <= 122 : // ['a','z']
				return 49
			
			
			
			}
			return NoState
		},
	
		// S50
		func(r rune) int {
			switch {
			case r == 34 : // ['"','"']
				return 24
			case r == 92 : // ['\','\']
				return 25
			
			
			default:
				return 26
			
			}
			return NoState
		},
	
		// S51
		func(r rune) int {
			switch {
			case 48 <= r && r <= 55 : // ['0','7']
				return 65
			
			
			
			}
			return NoState
		},
	
		// S52
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 66
			case 65 <= r && r <= 70 : // ['A','F']
				return 66
			case 97 <= r && r <= 102 : // ['a','f']
				return 66
			
			
			
			}
			return NoState
		},
	
		// S53
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 67
			case 65 <= r && r <= 70 : // ['A','F']
				return 67
			case 97 <= r && r <= 102 : // ['a','f']
				return 67
			
			
			
			}
			return NoState
		},
	
		// S54
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 68
			case 65 <= r && r <= 70 : // ['A','F']
				return 68
			case 97 <= r && r <= 102 : // ['a','f']
				return 68
			
			
			
			}
			return NoState
		},
	
		// S55
		func(r rune) int {
			switch {
			case r == 39 : // [''',''']
				return 60
			
			
			
			}
			return NoState
		},
	
		// S56
		func(r rune) int {
			switch {
			case 48 <= r && r <= 55 : // ['0','7']
				return 69
			
			
			
			}
			return NoState
		},
	
		// S57
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 70
			case 65 <= r && r <= 70 : // ['A','F']
				return 70
			case 97 <= r && r <= 102 : // ['a','f']
				return 70
			
			
			
			}
			return NoState
		},
	
		// S58
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 71
			case 65 <= r && r <= 70 : // ['A','F']
				return 71
			case 97 <= r && r <= 102 : // ['a','f']
				return 71
			
			
			
			}
			return NoState
		},
	
		// S59
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 72
			case 65 <= r && r <= 70 : // ['A','F']
				return 72
			case 97 <= r && r <= 102 : // ['a','f']
				return 72
			
			
			
			}
			return NoState
		},
	
		// S60
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
		},
	
		// S61
		func(r rune) int {
			switch {
			case r == 42 : // ['*','*']
				return 61
			case r == 47 : // ['/','/']
				return 73
			
			
			default:
				return 29
			
			}
			return NoState
		},
	
		// S62
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
		},
	
		// S63
		func(r rune) int {
			switch {
			case r == 62 : // ['>','>']
				return 74
			
			
			default:
				return 63
			
			}
			return NoState
		},
	
		// S64
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 41
			case 65 <= r && r <= 90 : // ['A','Z']
				return 42
			case r == 95 : // ['_','_']
				return 43
			case 97 <= r && r <= 110 : // ['a','n']
				return 44
			case r == 111 : // ['o','o']
				return 75
			case 112 <= r && r <= 122 : // ['p','z']
				return 44
			
			
			
			}
			return NoState
		},
	
		// S65
		func(r rune) int {
			switch {
			case 48 <= r && r <= 55 : // ['0','7']
				return 76
			
			
			
			}
			return NoState
		},
	
		// S66
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 77
			case 65 <= r && r <= 70 : // ['A','F']
				return 77
			case 97 <= r && r <= 102 : // ['a','f']
				return 77
			
			
			
			}
			return NoState
		},
	
		// S67
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 78
			case 65 <= r && r <= 70 : // ['A','F']
				return 78
			case 97 <= r && r <= 102 : // ['a','f']
				return 78
			
			
			
			}
			return NoState
		},
	
		// S68
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 79
			case 65 <= r && r <= 70 : // ['A','F']
				return 79
			case 97 <= r && r <= 102 : // ['a','f']
				return 79
			
			
			
			}
			return NoState
		},
	
		// S69
		func(r rune) int {
			switch {
			case 48 <= r && r <= 55 : // ['0','7']
				return 80
			
			
			
			}
			return NoState
		},
	
		// S70
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 81
			case 65 <= r && r <= 70 : // ['A','F']
				return 81
			case 97 <= r && r <= 102 : // ['a','f']
				return 81
			
			
			
			}
			return NoState
		},
	
		// S71
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 82
			case 65 <= r && r <= 70 : // ['A','F']
				return 82
			case 97 <= r && r <= 102 : // ['a','f']
				return 82
			
			
			
			}
			return NoState
		},
	
		// S72
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 83
			case 65 <= r && r <= 70 : // ['A','F']
				return 83
			case 97 <= r && r <= 102 : // ['a','f']
				return 83
			
			
			
			}
			return NoState
		},
	
		// S73
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
		},
	
		// S74
		func(r rune) int {
			switch {
			case r == 62 : // ['>','>']
				return 84
			
			
			
			}
			return NoState
		},
	
		// S75
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 41
			case 65 <= r && r <= 90 : // ['A','Z']
				return 42
			case r == 95 : // ['_','_']
				return 43
			case 97 <= r && r <= 113 : // ['a','q']
				return 44
			case r == 114 : // ['r','r']
				return 85
			case 115 <= r && r <= 122 : // ['s','z']
				return 44
			
			
			
			}
			return NoState
		},
	
		// S76
		func(r rune) int {
			switch {
			case r == 34 : // ['"','"']
				return 24
			case r == 92 : // ['\','\']
				return 25
			
			
			default:
				return 26
			
			}
			return NoState
		},
	
		// S77
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 86
			case 65 <= r && r <= 70 : // ['A','F']
				return 86
			case 97 <= r && r <= 102 : // ['a','f']
				return 86
			
			
			
			}
			return NoState
		},
	
		// S78
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 87
			case 65 <= r && r <= 70 : // ['A','F']
				return 87
			case 97 <= r && r <= 102 : // ['a','f']
				return 87
			
			
			
			}
			return NoState
		},
	
		// S79
		func(r rune) int {
			switch {
			case r == 34 : // ['"','"']
				return 24
			case r == 92 : // ['\','\']
				return 25
			
			
			default:
				return 26
			
			}
			return NoState
		},
	
		// S80
		func(r rune) int {
			switch {
			case r == 39 : // [''',''']
				return 60
			
			
			
			}
			return NoState
		},
	
		// S81
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 88
			case 65 <= r && r <= 70 : // ['A','F']
				return 88
			case 97 <= r && r <= 102 : // ['a','f']
				return 88
			
			
			
			}
			return NoState
		},
	
		// S82
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 89
			case 65 <= r && r <= 70 : // ['A','F']
				return 89
			case 97 <= r && r <= 102 : // ['a','f']
				return 89
			
			
			
			}
			return NoState
		},
	
		// S83
		func(r rune) int {
			switch {
			case r == 39 : // [''',''']
				return 60
			
			
			
			}
			return NoState
		},
	
		// S84
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
		},
	
		// S85
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 41
			case 65 <= r && r <= 90 : // ['A','Z']
				return 42
			case r == 95 : // ['_','_']
				return 43
			case 97 <= r && r <= 122 : // ['a','z']
				return 44
			
			
			
			}
			return NoState
		},
	
		// S86
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 90
			case 65 <= r && r <= 70 : // ['A','F']
				return 90
			case 97 <= r && r <= 102 : // ['a','f']
				return 90
			
			
			
			}
			return NoState
		},
	
		// S87
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 91
			case 65 <= r && r <= 70 : // ['A','F']
				return 91
			case 97 <= r && r <= 102 : // ['a','f']
				return 91
			
			
			
			}
			return NoState
		},
	
		// S88
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 92
			case 65 <= r && r <= 70 : // ['A','F']
				return 92
			case 97 <= r && r <= 102 : // ['a','f']
				return 92
			
			
			
			}
			return NoState
		},
	
		// S89
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 93
			case 65 <= r && r <= 70 : // ['A','F']
				return 93
			case 97 <= r && r <= 102 : // ['a','f']
				return 93
			
			
			
			}
			return NoState
		},
	
		// S90
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 94
			case 65 <= r && r <= 70 : // ['A','F']
				return 94
			case 97 <= r && r <= 102 : // ['a','f']
				return 94
			
			
			
			}
			return NoState
		},
	
		// S91
		func(r rune) int {
			switch {
			case r == 34 : // ['"','"']
				return 24
			case r == 92 : // ['\','\']
				return 25
			
			
			default:
				return 26
			
			}
			return NoState
		},
	
		// S92
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 95
			case 65 <= r && r <= 70 : // ['A','F']
				return 95
			case 97 <= r && r <= 102 : // ['a','f']
				return 95
			
			
			
			}
			return NoState
		},
	
		// S93
		func(r rune) int {
			switch {
			case r == 39 : // [''',''']
				return 60
			
			
			
			}
			return NoState
		},
	
		// S94
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 96
			case 65 <= r && r <= 70 : // ['A','F']
				return 96
			case 97 <= r && r <= 102 : // ['a','f']
				return 96
			
			
			
			}
			return NoState
		},
	
		// S95
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 97
			case 65 <= r && r <= 70 : // ['A','F']
				return 97
			case 97 <= r && r <= 102 : // ['a','f']
				return 97
			
			
			
			}
			return NoState
		},
	
		// S96
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 98
			case 65 <= r && r <= 70 : // ['A','F']
				return 98
			case 97 <= r && r <= 102 : // ['a','f']
				return 98
			
			
			
			}
			return NoState
		},
	
		// S97
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 99
			case 65 <= r && r <= 70 : // ['A','F']
				return 99
			case 97 <= r && r <= 102 : // ['a','f']
				return 99
			
			
			
			}
			return NoState
		},
	
		// S98
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 100
			case 65 <= r && r <= 70 : // ['A','F']
				return 100
			case 97 <= r && r <= 102 : // ['a','f']
				return 100
			
			
			
			}
			return NoState
		},
	
		// S99
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 101
			case 65 <= r && r <= 70 : // ['A','F']
				return 101
			case 97 <= r && r <= 102 : // ['a','f']
				return 101
			
			
			
			}
			return NoState
		},
	
		// S100
		func(r rune) int {
			switch {
			case r == 34 : // ['"','"']
				return 24
			case r == 92 : // ['\','\']
				return 25
			
			
			default:
				return 26
			
			}
			return NoState
		},
	
		// S101
		func(r rune) int {
			switch {
			case r == 39 : // [''',''']
				return 60
			
			
			
			}
			return NoState
		},
	
}
