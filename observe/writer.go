package observe

import "io"

type writer struct {
	w io.WriteCloser
	o Observer
}

func (w *writer) Write(p []byte) (n int, err error) {
	n, err = w.w.Write(p)
	if n > 0 {
		w.o.OnTransfer(n)
	}
	if err != nil {
		w.o.OnStop(err)
	}
	return
}

func (w *writer) Close() (err error) {
	err = w.w.Close()
	w.o.OnStop(err)
	return err
}

// Writer wraps w and notify o during write operations.
func Writer(w io.WriteCloser, o Observer) io.WriteCloser {
	return &writer{
		w: w,
		o: o,
	}
}
