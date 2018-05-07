// Package follow implements routines for computation of follow sets for all the
// nonterminals in a given grammar.

package follow

import (
	"bytes"
	"fmt"

	"github.com/goccmack/gocc/internal/ast"
	"github.com/goccmack/gocc/internal/parser/first"
	"github.com/goccmack/gocc/internal/parser/symbols"
)

type FollowSets struct {
	followSets map[string]first.SymbolSet
	symbols    *symbols.Symbols
}

const (
	EMPTY      = "empty"
	TERMINATOR = "$"
)

// GetFollowSets computes and returns the FollowSets of a Grammar.
func GetFollowSets(g *ast.Grammar, symbols *symbols.Symbols, firstSets *first.FirstSets) *FollowSets {
	followSets := &FollowSets{
		followSets: make(map[string]first.SymbolSet),
		symbols:    symbols,
	}

	if g.SyntaxPart == nil {
		return followSets
	}

	// Place TERMINATOR in the follow set of the start symbol.
	followSets.AddToken(g.SyntaxPart.ProdList[0].Id, TERMINATOR)

	// The computation of follow sets is done using a set of rules mentioned
	// in comments below. Rules 2 and 3 are interdependent, as a result of
	// which computation of one rule might induce a change in follow sets,
	// resulting in recomputation of the other rule. The computations are
	// repeated until a fixed point is attained, after which the size of
	// follow sets will remain constant and the value of variable `again`
	// will be false. The fixed point will always exist since at the end of
	// each computation the size of follow sets can either increase or
	// remain constant, and will be bounded by the total number of terminals.
	for again := true; again; {
		again = false
		for _, prod := range g.SyntaxPart.ProdList[0:] {
			symbolCount := len(prod.Body.Symbols)
			switch {
			// (Rule 1) If there is a production A -> α B ß, then everything
			// in first(ß) (except ε) is in follow(B).
			// (Rule 2) If first(ß) contains ε, then follow(B) contains follow(A).
			case symbolCount >= 2:
				for k, v := range prod.Body.Symbols {
					if !symbols.IsTerminal(v.SymbolString()) {
						// beta represents ß.
						beta := []string{}
						for _, v := range prod.Body.Symbols[k+1:] {
							beta = append(beta, v.SymbolString())
						}
						// firstBeta represents first(ß).
						firstBeta := first.FirstS(firstSets, beta)
						if _, contains := firstBeta[EMPTY]; !contains {
							if !followSets.GetSet(v.SymbolString()).Equal(firstBeta) {
								if followSets.AddSet(v.SymbolString(), firstBeta) {
									again = true
								}
							}
						} else {
							// followA represents follow(A).
							followA := followSets.GetSet(prod.Id)
							B := prod.Body.Symbols[symbolCount-1]
							// followB represents follow(B).
							followB := followSets.GetSet(B.SymbolString())
							if !followA.Equal(followB) {
								if followSets.AddSet(B.SymbolString(), followA) {
									again = true
								}
							}
						}
					}
				}
			// (Rule 3) If there is a production A -> α B, then follow(B) contains follow(A).
			case !symbols.IsTerminal(prod.Body.Symbols[symbolCount-1].SymbolString()):
				followA := followSets.GetSet(prod.Id)
				B := prod.Body.Symbols[symbolCount-1]
				followB := followSets.GetSet(B.SymbolString())
				if !followA.Equal(followB) {
					if followSets.AddSet(B.SymbolString(), followA) {
						again = true
					}
				}
			}
		}
	}

	return followSets
}

// AddSet adds the follow-set of a nonterminal to FollowSets.
func (this *FollowSets) AddSet(prodName string, terminals first.SymbolSet) bool {
	symbolsAdded := false
	for symbol := range terminals {
		if this.AddToken(prodName, symbol) {
			symbolsAdded = true
		}
	}
	return symbolsAdded
}

// AddToken adds a terminal to the FollowSet of a nonterminal.
func (this *FollowSets) AddToken(prodName string, terminal string) bool {
	symbolAdded := false
	set, ok := this.followSets[prodName]
	if !ok {
		set = make(first.SymbolSet)
		this.followSets[prodName] = set
	}
	if _, contain := set[terminal]; !contain {
		set[terminal] = true
		symbolAdded = true
	}
	return symbolAdded
}

// GetSet returns the FollowSet for a nonterminal.
func (this *FollowSets) GetSet(prodName string) first.SymbolSet {
	if set, ok := this.followSets[prodName]; ok {
		return set
	}
	return nil
}

// String returns a string representing the FollowSets.
func (this *FollowSets) String() string {
	buf := new(bytes.Buffer)
	for _, nt := range this.symbols.NTList() {
		set := this.followSets[nt]
		fmt.Fprintf(buf, "%s: %s\n", nt, set)
	}
	return buf.String()
}
