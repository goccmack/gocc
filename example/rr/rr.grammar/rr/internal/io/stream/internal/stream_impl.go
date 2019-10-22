package internal

import (
	"io"
	"sync"
	"unicode/utf8"
)

type (
	impl_WindowReader struct {
		bufferBlockSize int
		maxBufferSize   int
		buffers         [][]byte
		available       int
		absbase         int64
		position        int
		eof             bool
		rdr             io.Reader
		lock            *sync.Mutex
	}
)

func (s *impl_WindowReader) lockReader() {
	if s.lock != nil {
		s.lock.Lock()
	}
}

func (s *impl_WindowReader) unlockReader() {
	if s.lock != nil {
		s.lock.Unlock()
	}
}

func (s *impl_WindowReader) seek(offset int64, whence int) (offs int64, err error) {
	switch whence {
	case io.SeekStart:
		return s.seek(offset-s.ReadPosition(), io.SeekCurrent)
	case io.SeekCurrent:
		switch {
		case offset == 0:
			return s.ReadPosition(), nil
		case offset < 0:
			if -offset > int64(s.maxBackOffset()) {
				err = io.ErrShortBuffer
			} else {
				s.advanceGet(int(offset))
			}
		case offset > 0:
			for offset > 0 && !s.eof {
				available := s.Buffered()
				if offset <= int64(available) {
					s.advanceGet(int(offset))
					break
				}
				if available > 0 {
					s.advanceGet(available)
					offset -= int64(available)
				}
				read, _ := s.fetchBytes(int(offset))
				s.advanceGet(read)
				offset -= int64(read)
			}
			if s.eof {
				err = io.EOF
			}
		}
	case io.SeekEnd:
		for !s.eof {
			s.seek(int64(s.MaxCap()), io.SeekCurrent)
		}
		s.seek(int64(s.Buffered()), io.SeekCurrent)
	}
	if s.eof && err == nil && s.Buffered() > 0 {
		s.eof = false
	}
	return s.ReadPosition(), err
}

func (s impl_WindowReader) maxBackOffset() int {
	r := s.position
	if s.absbase > 0 {
		r = s.Cap() - s.Buffered()
	}
	return r
}

func (s impl_WindowReader) bufferIndex(pos int) int {
	if pos < 0 {
		panic("negative position")
	}
	return pos / s.bufferBlockSize
}

func (s impl_WindowReader) bufferBytes(pos int) []byte {
	return s.buffers[s.bufferIndex(pos)]
}

func (s impl_WindowReader) bufferByteRange(pos int, length int) []byte {
	p := s.bufferOffset(pos)
	return s.bufferBytes(pos)[p : p+length]
}

func (s impl_WindowReader) bufferOffset(pos int) int {
	return pos % s.bufferBlockSize
}

func (s *impl_WindowReader) advanceGet(n int) {
	sz := s.Cap()
	switch {
	case n < 0:
		if -n > s.maxBackOffset() {
			panic("underflow")
		}
		s.position += n
		if s.position < 0 {
			s.position += sz
			s.absbase -= int64(sz)
		}
	case n > 0:
		if n > s.Buffered() {
			panic("overflow")
		}
		s.position += n
		if s.position >= sz {
			s.position -= sz
			s.absbase += int64(sz)
		}
	}
	s.available -= n
}

func (s *impl_WindowReader) copy(to []byte, length int) {
	for offs, cp := 0, 0; length > 0; length -= cp {
		s.position %= s.MaxCap()
		p := s.bufferOffset(s.position)
		b := s.bufferBytes(s.position)
		r := s.bufferBlockSize - p
		if length <= r {
			cp = copy(to[offs:], b[p:p+length])
		} else {
			cp = copy(to[offs:], b[p:])
		}
		s.advanceGet(cp)
		offs += cp
	}
}

func (s *impl_WindowReader) fetchBytes(n int) (read int, err error) {
	if s.rdr == nil {
		s.eof = true
		return 0, io.EOF
	}
	for n > 0 && s.Buffered() < s.MaxCap() && err == nil {
		readNow := n
		putpos := s.position + s.available
		r := s.bufferBlockSize - s.bufferOffset(putpos)
		if readNow > r {
			readNow = r
		}
		for s.bufferIndex(putpos) >= len(s.buffers) && s.Cap() < s.MaxCap() {
			s.buffers = append(s.buffers, make([]byte, s.bufferBlockSize))
		}
		putpos %= s.MaxCap()
		offs := s.bufferOffset(putpos)
		if r, err = s.rdr.Read(s.bufferBytes(putpos)[offs : offs+readNow]); r > 0 {
			s.available += r
			read += r
			n -= r
		}
	}
	if n > 0 && err == nil {
		err = io.ErrShortBuffer
	}
	return
}

func (s *impl_WindowReader) read(p []byte) (read int, err error) {
	if s.eof {
		return 0, io.EOF
	}
	for read < len(p) && err == nil {
		available := s.Buffered()
		if len(p)-read <= available {
			s.copy(p[read:], len(p)-read)
			read = len(p)
		} else {
			if available == 0 {
				available, err = s.fetchBytes(len(p) - read)
			}
			s.copy(p[read:], available)
			read += available
			switch err {
			case io.ErrShortBuffer:
				err = nil
			case io.EOF:
				s.eof = true
			}
		}
	}
	return
}

func (s impl_WindowReader) peekByte() (byte, error) {
	if s.Buffered() < 1 {
		return 0, io.ErrShortBuffer
	}
	return s.bufferBytes(s.position)[s.bufferOffset(s.position)], nil
}

func (s *impl_WindowReader) unreadByte() error {
	_, e := s.seek(-1, io.SeekCurrent)
	return e
}

func (s *impl_WindowReader) readRune() (r rune, size int, err error) {
	for i := 1; err == nil && i <= 4; i++ {
		if s.Buffered() < i {
			if _, err = s.fetchBytes(i); err != nil {
				return
			}
		}
		t := s.bufferByteRange(s.position, i)
		if r, size = utf8.DecodeRune(t); r != utf8.RuneError {
			s.advanceGet(size)
			return
		}
	}
	return utf8.RuneError, 0, err
}

func (s *impl_WindowReader) unreadRune() (err error) {
	tmps := *s
	tmpb := []byte{}
	for {
		if _, err = tmps.seek(-1, io.SeekCurrent); err == nil {
			b, _ := tmps.peekByte()
			tmpb = append([]byte{b}, tmpb...)
			switch len(tmpb) {
			case 1:
				if (tmpb[0] & 0xC0) == 0x80 {
					continue
				}
				_, err = s.seek(-int64(len(tmpb)), io.SeekCurrent)
				return
			case 2:
				if (tmpb[0] & 0xC0) == 0x80 {
					continue
				}
				if (tmpb[0] & 0xE0) != 0xC0 {
					break
				}
				_, err = s.seek(-int64(len(tmpb)), io.SeekCurrent)
				return
			case 3:
				if (tmpb[0] & 0xC0) == 0x80 {
					continue
				}
				if (tmpb[0] & 0xF0) != 0xE0 {
					break
				}
				_, err = s.seek(-int64(len(tmpb)), io.SeekCurrent)
				return
			case 4:
				if (tmpb[0] & 0xF8) != 0xF0 {
					break
				}
				_, err = s.seek(-int64(len(tmpb)), io.SeekCurrent)
				return
			}
		}
	}
	_, err = s.seek(-1, io.SeekCurrent)
	return
}
