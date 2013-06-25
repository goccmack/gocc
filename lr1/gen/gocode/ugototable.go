package gocode

import(
	"code.google.com/p/gocc/ast"
	"code.google.com/p/gocc/config"
	"code.google.com/p/gocc/lr1"

	"bytes"
	"fmt"
	"text/template"
)

type UGotoTable struct {
	gotoTable	*lr1.GotoTable
	cfg config.Config
}

func NewUGotoTable(gt *lr1.GotoTable, c config.Config) *UGotoTable {
	return &UGotoTable{
		gotoTable: gt,
		cfg: c,
	}
}

func (this *UGotoTable) GenUncompressed(buf *bytes.Buffer) {
	writeGTUHdr(buf, len(this.gotoTable.GTO), len(this.gotoTable.NT))
	writeNTConst(buf, this.gotoTable.NT)
	writeGTUFunc(buf, this)
}

func writeGTUHdr(buf *bytes.Buffer, numStates, numNT int) {
	tmpl, err := template.New("GTUHdr").Parse(gtuHdrTmplate)
	if err != nil {
		panic(err)
	}
	if err := tmpl.Execute(buf, numNT); err != nil {
		panic(err)
	}
}

var gtuHdrTmplate = `

// NT is the set of non-terminal symbols of the target grammar
const(
	NUM_NT = {{.}}
)

type(
	GotoTabU [NUM_STATES]GotoRowU
	GotoRowU [NUM_NT]State
)

`

func writeNTConst(buf *bytes.Buffer, nt ast.SymbolS) {
	buf.WriteString("const(\n")
	for i, s := range nt {
		fmt.Fprintf(buf, "\tNT_%s\t= %d\n", s.TokLit, i)
	}
	buf.WriteString(")\n\n")
}

func writeGTUFunc(buf *bytes.Buffer, gto *UGotoTable) {
	buf.WriteString("func getGotoTableUncompressed() (gto *GotoTabU) {\n")
	buf.WriteString("\tgto = new(GotoTabU)\n")

	for s := range gto.gotoTable.GTO {
		for nt := 0; nt < len(gto.gotoTable.NT); nt++ {
			if nextState, ok := gto.gotoTable.GTO[s][nt]; ok {
				fmt.Fprintf(buf, "\tgto[%d][NT_%s] = %d \n", s, gto.gotoTable.NT[nt].TokLit, nextState)
			}
		}
	}

	buf.WriteString("\treturn\n}\n\n")
}