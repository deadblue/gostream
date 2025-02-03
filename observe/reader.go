package observe

import (
	"io"
	"sync/atomic"
)

type reader struct {
	// Underlying reader
	r io.Reader
	// Transfer observer
	o Observer
	// Done flag
	flag int32
}

func (r *reader) fireDone(err error) {
	if atomic.CompareAndSwapInt32(&r.flag, 0, 1) {
		if err == nil || err == io.EOF {
			r.o.OnStop(nil)
		} else {
			r.o.OnStop(err)
		}
	}
}

func (r *reader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)
	if n > 0 {
		r.o.OnTransfer(n)
	}
	if err != nil {
		r.fireDone(err)
	}
	return
}

func (r *reader) Close() (err error) {
	if c, ok := r.r.(io.Closer); ok {
		err = c.Close()
	}
	r.fireDone(err)
	return
}

// Reader wraps r and notify o during read operations.
func Reader(r io.Reader, o Observer) io.ReadCloser {
	return &reader{
		r: r,
		o: o,
	}
}
