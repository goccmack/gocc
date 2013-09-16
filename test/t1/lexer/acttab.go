<<<<<<< HEAD
package lexer

import (
	"code.google.com/p/gocc/test/t1/token"
	"fmt"
)

type ActionTable [NumStates]ActionRow
=======

package lexer

import(
	"fmt"
	"code.google.com/p/gocc/test/t1/token"
)

type ActionTable [NumStates] ActionRow
>>>>>>> 9d3e28b6b4a375015991b8ecb296891776774570

type ActionRow struct {
	Accept token.Type
	Ignore string
}

func (this ActionRow) String() string {
	return fmt.Sprintf("Accept=%d, Ignore=%s", this.Accept, this.Ignore)
}

var ActTab = ActionTable{
<<<<<<< HEAD
	ActionRow{ // S0
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S1
		Accept: -1,
		Ignore: "!whitespace",
	},
	ActionRow{ // S2
		Accept: 4,
		Ignore: "",
	},
	ActionRow{ // S3
		Accept: 2,
		Ignore: "",
	},
=======
 	ActionRow{ // S0
		Accept: 0,
 		Ignore: "",
 	},
 	ActionRow{ // S1
		Accept: -1,
 		Ignore: "!whitespace",
 	},
 	ActionRow{ // S2
		Accept: 4,
 		Ignore: "",
 	},
 	ActionRow{ // S3
		Accept: 2,
 		Ignore: "",
 	},
 		
>>>>>>> 9d3e28b6b4a375015991b8ecb296891776774570
}
