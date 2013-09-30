package ast

type LexNTNode interface {
	LexNode
	Element(int) LexNode
	Len() int
	Walk(LexNodeVisitor) LexNodeVisitor
}

// Element

func (this *LexAlt) Element(i int) LexNode {
	return this.Terms[i]
}

func (this *LexGroupPattern) Element(i int) LexNode {
	return this.LexPattern.Alternatives[i]
}

func (this *LexOptPattern) Element(i int) LexNode {
	return this.LexPattern.Alternatives[i]
}

func (this *LexPattern) Element(i int) LexNode {
	return this.Alternatives[i]
}

func (this *LexRepPattern) Element(i int) LexNode {
	return this.LexPattern.Alternatives[i]
}

// Len

func (this *LexAlt) Len() int {
	return len(this.Terms)
}

func (this *LexGroupPattern) Len() int {
	return len(this.LexPattern.Alternatives)
}

func (this *LexOptPattern) Len() int {
	return len(this.LexPattern.Alternatives)
}

func (this *LexPattern) Len() int {
	return len(this.Alternatives)
}

func (this *LexRepPattern) Len() int {
	return len(this.LexPattern.Alternatives)
}
