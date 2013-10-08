package lexer

import (
	"code.google.com/p/gocc/test/ebnf/t4/token"
	"fmt"
)

type ActionTable [NumStates]ActionRow

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
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S2
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S3
		Accept: 3,
		Ignore: "",
	},
	ActionRow{ // S4
		Accept: 4,
		Ignore: "",
	},
	ActionRow{ // S5
		Accept: 5,
		Ignore: "",
	},
	ActionRow{ // S6
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S7
		Accept: 7,
		Ignore: "",
	},
	ActionRow{ // S8
		Accept: 8,
		Ignore: "",
	},
	ActionRow{ // S9
		Accept: 9,
		Ignore: "",
	},
	ActionRow{ // S10
		Accept: 10,
		Ignore: "",
	},
}
