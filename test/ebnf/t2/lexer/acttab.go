package lexer

import (
	"code.google.com/p/gocc/test/ebnf/t2/token"
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
		Accept: -1,
		Ignore: "!ws",
	},
	ActionRow{ // S2
		Accept: 3,
		Ignore: "",
	},
}
