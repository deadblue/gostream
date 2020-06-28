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

func (r *reader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)
	if n > 0 {
		r.o.Transfer(n)
	}
	if err != nil {
		if err == io.EOF {
			r.o.Done(nil)
		} else {
			r.o.Done(err)
		}
		atomic.StoreInt32(r.flag, 1)
	}
	return
}

func (r *reader) Close() (err error) {
	if c, ok := r.r.(io.Closer); ok {
		err = c.Close()
	}
	// Call Observer.Done() only once.
	if atomic.CompareAndSwapInt32(r.flag, 0, 1) {
		r.o.Done(err)
	}
	return
}

// Create an observerd read stream.
func Reader(r io.Reader, o Observer) io.ReadCloser {
	var flag int32 = 0
	return &reader{
		r: r,
		o: o,
		// Done flag
		flag: &flag,
	}
}
