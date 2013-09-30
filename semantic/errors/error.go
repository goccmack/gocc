package errors

import (
	"fmt"
)

type Error struct {
	Msg    string
	Offset int
}

func (this *Error) String() string {
	return fmt.Sprintf("%s at offset %d", this.Msg, this.Offset)
}
