// Code generated by gocc; DO NOT EDIT.

package gocc2

import (
	"github.com/maxcalandrelli/gocc/internal/frontend/reparsed/internal/util"
)

type EscapedString = util.EscapedString

func RuneToString(r rune) string {
	return util.RuneToString(r)
}

func HexDigitValue(ch rune) int {
	return util.HexDigitValue(ch)
}

func IntValue(lit []byte) (int64, error) {
	return util.IntValue(lit)
}

func UintValue(lit []byte) (uint64, error) {
	return util.UintValue(lit)
}
