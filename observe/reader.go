package observe

import (
	"io"
	"sync/atomic"
)

type reader struct {
	r io.Reader
	o Observer
	//
	flag *int32
}

func (r *reader) fireDone(err error) {
	if atomic.CompareAndSwapInt32(r.flag, 0, 1) {
		if err == nil || err == io.EOF {
			r.o.Done(nil)
		} else {
			r.o.Done(err)
		}
	}
}

func (r *reader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)
	if n > 0 {
		r.o.Transfer(n)
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
	var flag int32 = 0
	return &reader{
		r: r,
		o: o,
		// Done flag
		flag: &flag,
	}
}
