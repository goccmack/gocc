package stream

import (
	internal "github.com/maxcalandrelli/gocc/example/ctx/ctx2.grammar/ctx2/internal/io/stream/internal"
	"io"
)

type (
	WindowReader interface {
		BlockSize() int
		Buffered() int
		Cap() int
		MaxCap() int
		Read([]byte) (int, error)
		ReadByte() (byte, error)
		ReadPosition() int64
		ReadRune() (rune, int, error)
		Seek(int64, int) (int64, error)
		UnreadByte() error
		UnreadRune() error
		MaxUnreadSize() int
	}
)

var (
	InfoPrefixRoundUp = false
)

func NewWindowReaderFromBytes(src []byte) WindowReader {
	return internal.NewWindowReader(src, nil, 0, 0, InfoPrefixRoundUp)
}

func NewWindowReader(rdr io.Reader) WindowReader {
	return internal.NewWindowReader(nil, rdr, 0, 0, false)
}

func NewLimitedWindowReader(rdr io.Reader, sizeMin, sizeMax int) WindowReader {
	return internal.NewWindowReader(nil, rdr, sizeMin, sizeMax, InfoPrefixRoundUp)
}
