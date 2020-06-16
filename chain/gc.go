package chain

import "io"

type closedReader struct{}

func (closedReader) Read(_ []byte) (n int, err error) {
	return 0, io.ErrClosedPipe
}
