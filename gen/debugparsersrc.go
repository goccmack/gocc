//Copyright 2012 Vastech SA (PTY) LTD
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package gen

const debugParserSrcBody = `
type(
	ActionTab []ActionRow
	ActionRow map[token.Type] Action
)

type(
	Accept int
	Shift State
	Reduce int

	Action interface{
		Act()
		String() string
	}

)

func (this Accept) Act(){}
func (this Shift) Act(){}
func (this Reduce) Act(){}

func (this Accept) String() string { return "Accept(0)" }
func (this Shift) String() string { return "Shift(" + strconv.Itoa(int(this)) + ")" }
func (this Reduce) String() string { return "Reduce(" + strconv.Itoa(int(this)) + ")" }

type(
	GotoTab []GotoRow
	GotoRow	map[NT] State
	State int
	NT	string
)

type (
	ProdTab []ProdTabEntry
	ProdTabEntry struct {
		String		string
		Head		NT
		NumSymbols	int
		ReduceFunc	func([]Attrib) (Attrib, error)
	}
	Attrib interface{
		String() string
	}
)

// Stack

type stack struct {
	state []State
	attrib	[]Attrib
}

const INITIAL_STACK_SIZE = 100

func NewStack() *stack {
	return &stack{ 	state: 	make([]State, 0, INITIAL_STACK_SIZE),
					attrib: make([]Attrib, 0, INITIAL_STACK_SIZE),
			}
}

func (this *stack) Push(s State, a Attrib) {
	this.state = append(this.state, s)
	this.attrib = append(this.attrib, a)
}

func(this *stack) Top() State {
	return this.state[len(this.state) - 1]
}

func (this *stack) PopN(items int) []Attrib {
	lo, hi := len(this.state) - items, len(this.state)
	
	attrib := this.attrib[lo: hi]
	
	this.state = this.state[:lo]
	this.attrib = this.attrib[:lo]
	
	return attrib
}

func (S *stack) String() string {
	res := "stack:\n"
	for i, st := range S.state {
		res += "\t" + strconv.Itoa(i) + ": " + strconv.Itoa(int(st))
		res += " , "
		if S.attrib[i] == nil {
			res += "nil"
		} else {
			res += S.attrib[i].String()
		}
		res += "\n"
	}
	return res
}

// Parser

type Parser struct {
	actTab ActionTab
	gotoTab	GotoTab
	prodTab	ProdTab
	stack	*stack
	nextToken	*token.Token
	pos	token.Position
	tokenMap *token.TokenMap
}

type Scanner interface {
	Scan() (*token.Token, token.Position)
}

func NewParser(act ActionTab, gto GotoTab, prod ProdTab, tm *token.TokenMap) *Parser {
	p := &Parser{actTab: act, gotoTab: gto, prodTab: prod, stack:NewStack(), tokenMap:tm}
	p.stack.Push(0, nil) //TODO: which attribute should be pushed here?
	return p
}

func Acc() {
	fmt.Println("Accept")
}

func (P *Parser) Error() {
	msg := "Error: " + P.TokString(P.nextToken) + " @ " + P.pos.String()
	msg += ", expected one of: "
	actRow := P.actTab[P.stack.Top()]
	i := 0
	for t, _ := range actRow {
		msg += P.tokenMap.TokenString(t)
		if i < len(actRow) - 1 {
			msg += " "
		}
		i++
	}
	fmt.Println(msg)
	os.Exit(1)
}

func (P *Parser) TokString(tok *token.Token) string {
	msg := P.tokenMap.TokenString(tok.Type) + "(" + strconv.Itoa(int(tok.Type)) + ")"
	msg += " " + string(tok.Lit)
	return msg
}

func (this *Parser) Parse(scanner Scanner) (res interface{}) {
	this.nextToken, this.pos = scanner.Scan()
	for acc := false; !acc; {
		fmt.Println(">>>>>>>>>>>>>>>")
		fmt.Println("token:", this.TokString(this.nextToken))
		fmt.Println(this.stack)
		action, ok := this.actTab[this.stack.Top()][this.nextToken.Type]
		if !ok {
			this.Error()
		}
		fmt.Println(action)
		switch act := action.(type) {
			case Accept:
				Acc()
				res = this.stack.PopN(1)[0]
				acc = true
			case Shift:
				this.stack.Push(State(act) , this.nextToken )
				this.nextToken, this.pos = scanner.Scan()
			case Reduce:
				prod := this.prodTab[int(act)]
				fmt.Println(prod.String)
				fmt.Println("Pop", prod.NumSymbols)
				attrib, err := prod.ReduceFunc(this.stack.PopN(prod.NumSymbols))
				fmt.Println("stack: len=", len(this.stack.state), " top", this.stack.Top())
				if err != nil {
					fmt.Println("Symantic error:", err)
					this.Error()
				} else {
					this.stack.Push(this.gotoTab[this.stack.Top()][prod.Head], attrib)
				}
		}
		fmt.Println()
	}
	return
}
`
