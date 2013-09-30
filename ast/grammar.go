package ast

type Grammar struct {
	*LexPart
	*SyntaxPart
}

func NewGrammar(lexPart, syntaxPart interface{}) (*Grammar, error) {
	g := &Grammar{}
	if lexPart != nil {
		g.LexPart = lexPart.(*LexPart)
	} else {
		lp, _ := NewLexPart(nil, nil, nil)
		g.LexPart = lp
	}
	if syntaxPart != nil {
		g.SyntaxPart = syntaxPart.(*SyntaxPart).augment()
	}
	return g, nil
}
