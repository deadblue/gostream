package quietly

import "io"

// Close a io.Closer without return the error.
func Close(c io.Closer) {
	if c != nil {
		_ = c.Close()
	}
}
