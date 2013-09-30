package states

import (
	"fmt"
	lr0items "code.google.com/p/gocc/parser/lrk/items"
)

type ConfigGroup struct {
	Item *lr0items.Item
	*ContextSet
	hashKey string
}

func NewConfigGroup(item *lr0items.Item, context ...string) *ConfigGroup {
	ctxt := NewContextSet()
	ctxt.Add(context...)
	return &ConfigGroup{
		Item:       item,
		ContextSet: ctxt,
		hashKey:    item.HashKey(),
	}
}

func (this *ConfigGroup) AddContext(c ...string) *ConfigGroup {
	this.ContextSet.Add(c...)
	return this
}

func (this *ConfigGroup) Core() *lr0items.Item {
	return this.Item
}

func (this *ConfigGroup) Equal(that *ConfigGroup) bool {
	return this.hashKey == that.hashKey && this.ContextSet.Equal(that.ContextSet)
}

func (this *ConfigGroup) HashKey() string {
	return this.hashKey
}

func (this *ConfigGroup) IsNucleus() bool {
	return this.Item.Position > 0
}

/*
Return:
	-1 if this < that
	0 if this == that
	1 if this > that
*/
func (this *ConfigGroup) Compare(that *ConfigGroup) int {
	switch {
	case this.hashKey < that.hashKey:
		return -1
	case this.hashKey == that.hashKey:
		return 0
	}
	return 1
}

func (this *ConfigGroup) String() string {
	return fmt.Sprintf("%s %s", this.Item, this.ContextSet)
}
