package binary

import (
	"encoding/binary"
	"io"
	"math"
)

/*
Writer is a binary writer which implements:
  - io.Writer
  - io.ByteWriter
And provides more write methods for binary data.
*/
type Writer struct {
	w io.Writer
	o binary.ByteOrder
}

func (w *Writer) Write(p []byte) (n int, err error) {
	return w.w.Write(p)
}

func (w *Writer) WriteByte(b byte) (err error) {
	_, err = w.Write([]byte{b})
	return
}

// WriteUint16 writes an uint16 value with preset byte order.
func (w *Writer) WriteUint16(u uint16) (err error) {
	buf := make([]byte, 2)
	w.o.PutUint16(buf, u)
	_, err = w.Write(buf)
	return
}

// WriteUint32 writes an uint32 value with preset byte order.
func (w *Writer) WriteUint32(u uint32) (err error) {
	buf := make([]byte, 4)
	w.o.PutUint32(buf, u)
	_, err = w.Write(buf)
	return
}

// WriteUint64 writes an uint64 value with preset byte order.
func (w *Writer) WriteUint64(u uint64) (err error) {
	buf := make([]byte, 8)
	w.o.PutUint64(buf, u)
	_, err = w.Write(buf)
	return
}

func (w *Writer) WriteInt16(i int16) (err error) {
	return w.WriteUint16(uint16(i))
}

func (w *Writer) WriteInt32(i int32) (err error) {
	return w.WriteUint32(uint32(i))
}

func (w *Writer) WriteInt64(i int64) (err error) {
	return w.WriteUint64(uint64(i))
}

func (w *Writer) WriteFloat32(f float32) (err error) {
	return w.WriteUint32(math.Float32bits(f))
}

func (w *Writer) WriteFloat64(f float64) (err error) {
	return w.WriteUint64(math.Float64bits(f))
}

// NewWriter creates a binary writer with byte order.
func NewWriter(w io.Writer, order binary.ByteOrder) (writer *Writer) {
	var ok bool
	if writer, ok = w.(*Writer); ok {
		writer.o = order
	} else {
		writer = &Writer{
			w: w,
			o: order,
		}
	}
	return
}
