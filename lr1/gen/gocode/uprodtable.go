package gocode

import(
	"code.google.com/p/gocc/ast"

	"bytes"
	"fmt"
	"text/template"
)

type UProdTable struct {
	g *ast.Grammar
}

func NewUProdsTable(g *ast.Grammar) *UProdTable {
	return &UProdTable{g}
}

func (this *UProdTable) GenUncompressed(buf *bytes.Buffer) {
	writePTUHdr(buf, len(this.g.Prod))
	writePTUFunc(buf, this)
}

func writePTUHdr(buf *bytes.Buffer, numProds int) {
	tmpl, err := template.New("PTUHdr").Parse(ptuHdrTemplate)
	if err != nil {
		panic(err)
	}
	if err := tmpl.Execute(buf, numProds); err != nil {
		panic(err)
	}
}

var ptuHdrTemplate = `
const NUM_PRODS = {{.}}

type (
	ProdTabU      [NUM_PRODS]*ProdTabUEntry
	ProdTabUEntry struct {
		String     string
		Head       NT
		HeadIndex	int
		NumSymbols int
		ReduceFunc func([]Attrib) (Attrib, error)
	}
)

`

func writePTUFunc(buf *bytes.Buffer, upt *UProdTable) {
	buf.WriteString("func getProductionsTableUncompressed() (pt *ProdTabU) {\n")
	buf.WriteString("\tpt = new(ProdTabU)\n")
	for i, prod := range upt.g.Prod {
		writeProdTabEntry(buf, i, prod)
	}
	buf.WriteString("\treturn\n}\n\n")
}

func writeProdTabEntry(buf *bytes.Buffer, pidx int, p *ast.Production) {
	fmt.Fprintf(buf, "\tpt[%d] = &ProdTabUEntry{\n", pidx)
	fmt.Fprintf(buf, "\t\tString: \"%s\",\n", p.String())
	fmt.Fprintf(buf, "\t\tHead: \"%s\", \n", p.Head.TokLit)
	if pidx == 0 {
		fmt.Fprintf(buf, "\t\tHeadIndex: -1, \n")
	} else {
		fmt.Fprintf(buf, "\t\tHeadIndex: NT_%s, \n", p.Head.TokLit)
	}
	fmt.Fprintf(buf, "\t\tNumSymbols: %d, \n", len(p.Body.Symbols))
	fmt.Fprintf(buf, "\t\tReduceFunc: func(X []Attrib) (Attrib, error) {\n")
	if p.Body.SDT != "" {
		fmt.Fprintf(buf, "\t\t\treturn %s \n", p.Body.SDT)
	} else if p.Body == ast.NULL_BODY {
		fmt.Fprintf(buf, "\t\t\treturn nil, nil \n")
	} else {
		fmt.Fprintf(buf, "\t\t\treturn X[0], nil \n")
	}
	fmt.Fprintf(buf, "\t\t},\n")
	fmt.Fprintf(buf, "\t}\n")
}
