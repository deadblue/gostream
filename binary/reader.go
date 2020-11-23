package binary

import (
	"encoding/binary"
	"io"
)

/*
Reader is a reader which implements:
	- io.Reader
	- io.ByteReader
And provide more read method for binary data.
*/
type Reader struct {
	r io.Reader
	o binary.ByteOrder
}

func (r *Reader) Read(p []byte) (n int, err error) {
	return r.r.Read(p)
}

func (r *Reader) ReadByte() (b byte, err error) {
	buf := make([]byte, 1)
	if _, err = r.Read(buf); err == nil {
		b = buf[0]
	}
	return
}

func (r *Reader) ReadUint16() (u uint16, err error) {
	buf := make([]byte, 2)
	if _, err = r.Read(buf); err == nil {
		u = r.o.Uint16(buf)
	}
	return
}

func (r *Reader) ReadUint32() (u uint32, err error) {
	buf := make([]byte, 4)
	if _, err = r.Read(buf); err == nil {
		u = r.o.Uint32(buf)
	}
	return
}

func (r *Reader) ReadUint64() (u uint64, err error) {
	buf := make([]byte, 8)
	if _, err = r.Read(buf); err == nil {
		u = r.o.Uint64(buf)
	}
	return
}

// NewReader creates a binary reader with byte order.
func NewReader(r io.Reader, order binary.ByteOrder) *Reader {
	return &Reader{
		r: r,
		o: order,
	}
}
