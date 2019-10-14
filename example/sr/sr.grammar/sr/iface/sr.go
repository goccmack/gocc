// Code generated by gocc; DO NOT EDIT.

package iface

import (
	"github.com/maxcalandrelli/gocc/example/sr/sr.grammar/sr/internal/errors"
	"github.com/maxcalandrelli/gocc/example/sr/sr.grammar/sr/internal/token"
	"io"
)

type (
	Token       = token.Token
	TokenMap    = token.TokenMap
	Pos         = token.Pos
	ErrorSymbol = errors.ErrorSymbol
	Error       = errors.Error

	Scanner interface {
		Scan() (tok *Token)
	}

	StreamScanner interface {
		GetStream() TokenStream
	}

	CheckPoint interface {
		DistanceFrom(CheckPoint) int
		Advance(int) CheckPoint
	}

	CheckPointable interface {
		GetCheckPoint() CheckPoint
		GotoCheckPoint(CheckPoint)
	}

	TokenStream interface {
		io.Reader
		io.RuneScanner
		io.Seeker
	}
)

const (
	INVALID = token.INVALID
	EOF     = token.EOF
)

func GetTokenMap() TokenMap {
	return token.TokMap
}
