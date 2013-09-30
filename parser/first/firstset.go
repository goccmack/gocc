package first

import (
	"bytes"
	"fmt"
)

type FirstSet []string

func (this FirstSet) Add(sym ...string) (fs FirstSet, new bool) {
	fs = this
	for _, s := range sym {
		if !this.Contain(s) {
			fs = append(fs, s)
			new = true
		}
	}
	return
}

func (this FirstSet) Contain(sym string) bool {
	if this == nil {
		return false
	}
	for _, s := range this {
		if s == sym {
			return true
		}
	}
	return false
}

func (this FirstSet) String() string {
	w := new(bytes.Buffer)
	fmt.Fprintf(w, "{")
	for i, sym := range this {
		if i > 0 {
			fmt.Fprintf(w, ", %s", sym)
		} else {
			fmt.Fprintf(w, "%s", sym)
		}
	}
	fmt.Fprintf(w, "}")
	return w.String()
}
