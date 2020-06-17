package multipart

import (
	"io"
	"os"
)

type Form interface {

	// Add a string part into form
	AddString(name string, value string)

	// Add a file part into form
	AddFile(name string, file *os.File)

	// Add a file part by info and data
	AddFileData(name string, info FileInfo, data io.Reader)

	// Archive the form, this operation can be called only once.
	// All add operations after Archive() will take no effect.
	Archive() (mimeType string, body io.Reader, size int64, err error)
}

func New() Form {
	return &implForm{
		archived: false,
		boundary: makeBoundary(boundarySize),
		size:     int64(boundarySize + 6),
		parts:    make([]*part, 0),
	}
}
