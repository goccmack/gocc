package golang

import (
	"bytes"
	"path"
	"text/template"

	"github.com/maxcalandrelli/gocc/internal/io"
)

type data struct {
	MyName     string
	Pkg        string
	Outdir     string
	Subpath    string
	MyPath     string
	MyInternal string
}

func GenIo(pkg, outdir, subpath string) {
	myData := data{
		Pkg:        pkg,
		Outdir:     outdir,
		Subpath:    subpath,
		MyPath:     "io/stream",
		MyInternal: "internal",
	}
	genIoStream(myData)
	genIoStreamImpl(myData)
	genIoStreamPublic(myData)
}

func genIoStream(d data) {
	tmpl, err := template.New("io_stream").Parse(streamSrc[1:])
	if err != nil {
		panic(err)
	}
	wr := new(bytes.Buffer)
	if err := tmpl.Execute(wr, d); err != nil {
		panic(err)
	}
	io.WriteFile(path.Join(d.Outdir, d.Subpath, d.MyPath, "stream.go"), wr.Bytes())
}

func genIoStreamImpl(d data) {
	tmpl, err := template.New("io_stream_impl").Parse(streamImplSrc[1:])
	if err != nil {
		panic(err)
	}
	wr := new(bytes.Buffer)
	if err := tmpl.Execute(wr, d); err != nil {
		panic(err)
	}
	io.WriteFile(path.Join(d.Outdir, d.Subpath, d.MyPath, d.MyInternal, "stream_impl.go"), wr.Bytes())
}

func genIoStreamPublic(d data) {
	tmpl, err := template.New("io_stream_public").Parse(streamPublicSrc[1:])
	if err != nil {
		panic(err)
	}
	wr := new(bytes.Buffer)
	if err := tmpl.Execute(wr, d); err != nil {
		panic(err)
	}
	io.WriteFile(path.Join(d.Outdir, d.Subpath, d.MyPath, d.MyInternal, "stream_public.go"), wr.Bytes())
}

var streamSrc = `
package stream

import (
	internal "{{.Pkg}}/{{.Subpath}}/{{.MyPath}}/{{.MyInternal}}"
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
`

var streamImplSrc = `
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
`

var streamPublicSrc = `
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
`
