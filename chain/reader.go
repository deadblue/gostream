package chain

import (
	"io"
)

type Reader struct {
	readers []io.Reader
	index   int
	count   int
}

func (r *Reader) Read(p []byte) (n int, err error) {
	// Reach the end of the chain
	if r.index >= r.count {
		return 0, io.EOF
	}
	for n == 0 && err == nil {
		// Read data from current reader
		n, err = r.readers[r.index].Read(p)
		if err == io.EOF {
			// Swtich to next reader when EOF
			r.index += 1
			if r.index < r.count {
				err = nil
			}
		}
	}
	return
}

func (r *Reader) Close() (err error) {
	for i, reader := range r.readers {
		// Close all closeable readers
		if c, ok := reader.(io.Closer); ok {
			err = c.Close()
		}
		r.readers[i] = closedReader{}
	}
	return
}

func JoinReader(reader ...io.Reader) *Reader {
	rs, count := make([]io.Reader, 0), 0
	for _, r := range reader {
		// Flatten the readers
		if lr, ok := r.(*Reader); ok {
			rs = append(rs, lr.readers...)
			count += lr.count
		} else {
			rs = append(rs, r)
			count += 1
		}
	}
	return &Reader{
		readers: rs,
		index:   0,
		count:   count,
	}
}
