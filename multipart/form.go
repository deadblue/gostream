package multipart

import (
	"io"
	"os"
)

// Multipart form, caller can call New() function to create it.
type Form interface {

	// Add a value part.
	AddValue(name string, value string) Form

	// Add a file part.
	AddFile(name string, file *os.File) Form

	// Add a file part by info and data.
	AddFileData(name string, fileName string, fileSize int64, data io.Reader) Form

	// Add a file part with specific MIME type.
	AddMimeFile(name string, mimeType string, file *os.File) Form

	// Add a file part with specific MIME type by info and data.
	AddMimeFileData(name string, mimeType string, fileName string, fileSize int64, data io.Reader) Form

	// Archive the form, return metadata and data stream. This method can be
	// called once and only once, all add operations on this form will not
	// take effect after archived.
	// In most cases, caller should call NewRequest function instead of this
	// method. If caller calls this method for other use, DO NOT forget to
	// close the returned body.
	Archive() (mimeType string, size int64, body io.ReadCloser, err error)
}

// Create a multipart form.
func New() Form {
	return &implForm{
		archived: false,
		boundary: makeBoundary(boundarySize),
		size:     int64(boundarySize + 6),
		parts:    make([]io.Reader, 0),
	}
}
