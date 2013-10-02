//Copyright 2013 Vastech SA (PTY) LTD
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

package states

import (
	lr0items "code.google.com/p/gocc/parser/lrk/items"
	"fmt"
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
