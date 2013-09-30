package states

import (
	"bytes"
	"fmt"
	"sort"
)

type ContextSet struct {
	List    []string
	Contain map[string]bool
}

func NewContextSet() *ContextSet {
	return &ContextSet{
		List:    make([]string, 0, 8),
		Contain: make(map[string]bool),
	}
}

/*
Return true if symbol was new and added to this.
*/
func (this *ContextSet) Add(symbol ...string) (newContext bool) {
	for _, s := range symbol {
		if !this.Contain[s] {
			this.Contain[s] = true
			this.List = append(this.List, s)
			newContext = true
		}
	}
	sort.Strings(this.List)
	return
}

func (this *ContextSet) AddSet(that *ContextSet) {
	this.Add(that.List...)
}

func (this *ContextSet) Clone() *ContextSet {
	clone := &ContextSet{
		List:    make([]string, 0, 8),
		Contain: make(map[string]bool),
	}
	clone.List = append(clone.List, this.List...)
	for k, v := range this.Contain {
		clone.Contain[k] = v
	}
	return clone
}

/*
Returns set difference: this-that
*/
func (this *ContextSet) Diff(that *ContextSet) *ContextSet {
	diff := NewContextSet()
	for sym, _ := range this.Contain {
		if !that.Contain[sym] {
			diff.Add(sym)
		}
	}
	return diff
}

func (this *ContextSet) Equal(that *ContextSet) bool {
	if len(this.List) != len(that.List) {
		return false
	}
	for sym, _ := range this.Contain {
		if !that.Contain[sym] {
			return false
		}
	}
	return true
}

func (this *ContextSet) Intersection(that *ContextSet) []string {
	intersection := []string{}
	for sym, _ := range this.Contain {
		if that.Contain[sym] {
			intersection = append(intersection, sym)
		}
	}
	return intersection
}

/*
Return the number of items in the context set
*/
func (this *ContextSet) Size() int {
	return len(this.List)
}

func (this *ContextSet) String() string {
	w := new(bytes.Buffer)
	fmt.Fprintf(w, "{")
	for i, ctx := range this.List {
		if i > 0 {
			fmt.Fprintf(w, ", %s", ctx)
		} else {
			fmt.Fprintf(w, "%s", ctx)
		}
	}
	fmt.Fprintf(w, "}")
	return w.String()
}
