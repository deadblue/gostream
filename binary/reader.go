package binary

import (
	"encoding/binary"
	"io"
	"math"
)

/*
Reader is a reader which implements:
	- io.Reader
	- io.ByteReader
And provides more read method for binary data.
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

func (r *Reader) ReadInt16() (i int16, err error) {
	var u uint16
	if u, err = r.ReadUint16(); err == nil {
		i = int16(u)
	}
	return
}

func (r *Reader) ReadInt32() (i int32, err error) {
	var u uint32
	if u, err = r.ReadUint32(); err == nil {
		i = int32(u)
	}
	return
}

func (r *Reader) ReadInt64() (i int64, err error) {
	var u uint64
	if u, err = r.ReadUint64(); err == nil {
		i = int64(u)
	}
	return
}

func (r *Reader) ReadFloat32(f float32, err error) {
	var u uint32
	if u, err = r.ReadUint32(); err == nil {
		f = math.Float32frombits(u)
	}
	return
}

func (r *Reader) ReadFloat64(f float64, err error) {
	var u uint64
	if u, err = r.ReadUint64(); err == nil {
		f = math.Float64frombits(u)
	}
	return
}

// NewReader creates a binary reader with byte order.
func NewReader(r io.Reader, order binary.ByteOrder) (reader *Reader) {
	var ok bool
	if reader, ok = r.(*Reader); ok {
		reader.o = order
	} else {
		reader = &Reader{
			r: r,
			o: order,
		}
	}
	return
}
