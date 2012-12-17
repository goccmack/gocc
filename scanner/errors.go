// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package scanner

import (
	"code.google.com/p/gocc/token"
)

// An implementation of an ErrorHandler may be provided to the Scanner.
// If a syntax error is encountered and a handler was installed, Error
// is called with a position and an error message. The position points
// to the beginning of the offending token.
//
type ErrorHandler interface {
	Error(pos token.Position, msg string)
}

// Within ErrorVector, an error is represented by an Error node. The
// position Pos, if valid, points to the beginning of the offending
// token, and the error condition is described by Msg.
//
type Error struct {
	Pos token.Position
	Msg string
}

func (e *Error) String() string {
	if e.Pos.IsValid() {
		// don't print "<unknown position>"
		return e.Pos.String() + ": " + e.Msg
	}
	return e.Msg
}
