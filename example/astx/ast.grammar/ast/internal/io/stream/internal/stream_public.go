package internal

import (
	"io"
)

func NewWindowReader(buf []byte, underlyingReader io.Reader, blksize, smax int, InfoPrefixRoundUp bool) *impl_WindowReader {
	if blksize < 0 || smax < blksize {
		panic("illegal buffer size")
	}
	r := &impl_WindowReader{rdr: underlyingReader}
	r.bufferBlockSize = blksize
	if underlyingReader != nil {
		if blksize <= 0 {
			r.bufferBlockSize = 1 << 10
		} else if InfoPrefixRoundUp {
			for r.bufferBlockSize = 1; blksize > 0; r.bufferBlockSize = r.bufferBlockSize << 1 {
				blksize = blksize >> 1
			}
		}
		if smax > 0 {
			r.maxBufferSize = ((smax-1)/r.bufferBlockSize + 1) * r.bufferBlockSize
		}
		r.buffers = [][]byte{}
	} else {
		r.buffers = [][]byte{buf}
		r.available = len(buf)
		r.bufferBlockSize = r.available
		r.maxBufferSize = r.available
	}
	return r
}

func (s impl_WindowReader) Cap() int {
	return len(s.buffers) * s.bufferBlockSize
}

func (s impl_WindowReader) BlockSize() int {
	return s.bufferBlockSize
}

func (s impl_WindowReader) MaxCap() int {
	if s.maxBufferSize > 0 {
		return s.maxBufferSize
	}
	return 1 << 31
}

func (s impl_WindowReader) MaxUnreadSize() int {
	return s.maxBackOffset()
}

func (s impl_WindowReader) Buffered() int {
	return s.available
}

func (s impl_WindowReader) ReadPosition() int64 {
	return s.absbase + int64(s.position)
}

func (s *impl_WindowReader) Read(p []byte) (read int, err error) {
	s.lockReader()
	defer func() { s.unlockReader() }()
	return s.read(p)
}

func (s *impl_WindowReader) Seek(offset int64, whence int) (offs int64, err error) {
	s.lockReader()
	defer func() { s.unlockReader() }()
	return s.seek(offset, whence)
}

func (s *impl_WindowReader) ReadByte() (byte, error) {
	b := []byte{0}
	_, err := s.Read(b)
	return b[0], err
}

func (s *impl_WindowReader) PeekByte() (byte, error) {
	s.lockReader()
	defer func() { s.unlockReader() }()
	return s.peekByte()
}

func (s *impl_WindowReader) UnreadByte() error {
	s.lockReader()
	defer func() { s.unlockReader() }()
	return s.unreadByte()
}

func (s *impl_WindowReader) ReadRune() (r rune, size int, err error) {
	s.lockReader()
	defer func() { s.unlockReader() }()
	return s.readRune()
}

func (s *impl_WindowReader) UnreadRune() (err error) {
	s.lockReader()
	defer func() { s.unlockReader() }()
	return s.unreadRune()
}
