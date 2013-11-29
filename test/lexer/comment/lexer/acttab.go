
package lexer

import(
	"fmt"
	"code.google.com/p/gocc/test/lexer/comment/token"
)

type ActionTable [NumStates] ActionRow

type ActionRow struct {
	Accept token.Type
	Ignore string
}

func (this ActionRow) String() string {
	return fmt.Sprintf("Accept=%d, Ignore=%s", this.Accept, this.Ignore)
}

var ActTab = ActionTable{
 	ActionRow{ // S0
		Accept: 0,
 		Ignore: "",
 	},
 	ActionRow{ // S1
		Accept: -1,
 		Ignore: "!ws",
 	},
 	ActionRow{ // S2
		Accept: 0,
 		Ignore: "",
 	},
 	ActionRow{ // S3
		Accept: 0,
 		Ignore: "",
 	},
 	ActionRow{ // S4
		Accept: -1,
 		Ignore: "!comment",
 	},
 	ActionRow{ // S5
		Accept: 0,
 		Ignore: "",
 	},
 	ActionRow{ // S6
		Accept: 0,
 		Ignore: "",
 	},
 	ActionRow{ // S7
		Accept: 0,
 		Ignore: "",
 	},
 	ActionRow{ // S8
		Accept: 0,
 		Ignore: "",
 	},
 	ActionRow{ // S9
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S10
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S11
		Accept: 3,
 		Ignore: "",
 	},
 		
}
