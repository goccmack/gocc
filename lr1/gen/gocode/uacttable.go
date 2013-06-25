// Code generator producing an uncompressed action table for the version of the LR-1 parse engine
// that works with uncompressed tables
package gocode

import(
	"code.google.com/p/gocc/config"
	"code.google.com/p/gocc/lr1"
	"code.google.com/p/gocc/token"

	"bytes"
	"fmt"
	"text/template"
)

type ActTable interface {
	GenUncompressed(buf *bytes.Buffer, tm *token.TokenMap)
}

// Returns generator for uncompressed action table
func NewActTableGenerator(at lr1.ActionTable, c config.Config) (aut ActTable) {
	return &actionTable{
		actTab: at,
		cfg: c,
	}
}

type actionTable struct {
	actTab lr1.ActionTable
	cfg config.Config
}

func (this *actionTable) GenUncompressed(buf *bytes.Buffer, tm *token.TokenMap) {
	writeATUHdr(buf, len(this.actTab), tm.Len())
	writeATUFunc(buf, this.actTab, tm)
	writeCanRecoverFunc(buf, this.actTab)
}

func writeATUHdr(buf *bytes.Buffer, numStates, numTokens int) {
	tmpl, err := template.New("ATUHdr").Parse(atuHdrTmplate)
	if err != nil {
		panic(err)
	}
	if err := tmpl.Execute(buf, atuHdrStruct{numStates, numTokens}); err != nil {
		panic(err)
	}
}

func writeATUFunc(buf *bytes.Buffer, act lr1.ActionTable, tm *token.TokenMap) {
	buf.WriteString("func getActionTableUncompressed() (act *ActionTabU) {\n")
	buf.WriteString("\tact = new(ActionTabU)\n")
	for s, state := range act {
		buf.WriteString(fmt.Sprintf("\t/* S%d */\n", s))
		for t := token.Type(0); t < token.Type(tm.Len()); t++ {
			if act, ok := state.Actions[t]; ok {
				// buf.WriteString(fmt.Sprintf("\t\t/* %d, %s, %s */\n", t, tm.TokenString(t), act))
				buf.WriteString(fmt.Sprintf("\t\tact[%d][%d] = %s // %s\n", s, t, act, tm.TokenString(t)))
			}
		}
	}
	buf.WriteString("\treturn\n")
	buf.WriteString("}\n\n")
}

type atuHdrStruct struct {
	NumStates int 
	NumTokens int
}

var atuHdrTmplate = `
const(
	NUM_STATES = {{.NumStates}}
	NUM_TOKENS = {{.NumTokens}}
)

type(
	ActionTabU [NUM_STATES]ActionRowU
	ActionRowU [NUM_TOKENS]Action
)

type(
	CanRecover [NUM_STATES] bool
)

`

func writeCanRecoverFunc(buf *bytes.Buffer, act lr1.ActionTable) {
	buf.WriteString("func getCanRecoverTableUncompressed() (cr *CanRecover) {\n")
	buf.WriteString("\tcr = new(CanRecover)\n")
	for s, row := range act {
		if row.CanRecover {
			buf.WriteString(fmt.Sprintf("\tcr[%d] = %s\n", s, "true"))
		}
	}

	buf.WriteString("\treturn\n}")
}

