// Code generated by gocc; DO NOT EDIT.

package iface

import (
	"github.com/maxcalandrelli/gocc/example/errorrecovery/er.grammar/er/internal/errors"
	"github.com/maxcalandrelli/gocc/example/errorrecovery/er.grammar/er/internal/io/stream"
	"github.com/maxcalandrelli/gocc/example/errorrecovery/er.grammar/er/internal/token"
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

func NewWindowReaderFromBytes(src []byte) stream.WindowReader {
	return stream.NewWindowReaderFromBytes(src)
}

func NewWindowReader(rdr io.Reader) stream.WindowReader {
	return stream.NewWindowReader(rdr)
}

func NewLimitedWindowReader(rdr io.Reader, sizeMin, sizeMax int) stream.WindowReader {
	return stream.NewLimitedWindowReader(rdr, sizeMin, sizeMax)
}
