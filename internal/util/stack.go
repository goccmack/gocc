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

package util

type Stack struct {
	stack []interface{}
}

// NewStack will return a new, initialized stack.
func NewStack(capacity int) *Stack {
	return &Stack{make([]interface{}, 0, capacity)}
}

// Len returns the number of items on the stack, or 0 if it is empty.
func (this *Stack) Len() int {
	return len(this.stack)
}

// Peek will return the item at `index` in the stack, where 0 is the bottom. Returns
// nil if the index exceeds the length of the stack.
func (this *Stack) Peek(index int) interface{} {
	if index > len(this.stack)-1 {
		return nil
	}
	return this.stack[index]
}

// Pop removes and returns the item at the top of the stack or nil.
func (this *Stack) Pop() (item interface{}) {
	if len(this.stack) == 0 {
		return nil
	}
	item = this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	return
}

// Push places an item at the top of the stack or panics if item is nil.
func (this *Stack) Push(items ...interface{}) *Stack {
	for _, item := range items {
		if item == nil {
			panic("nil item may not be pushed")
		}
		this.stack = append(this.stack, item)
	}
	return this
}

// Top returns the item at the top of the stack without altering the stack.
func (this *Stack) Top() interface{} {
	return this.stack[len(this.stack)-1]
}
