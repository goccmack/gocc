package states

import (
	"fmt"
)

type Transition struct {
	Sym   string
	State *State
}

func (this Transition) String() string {
	return fmt.Sprintf("%s -> %d", this.Sym, this.State.Number)
}
