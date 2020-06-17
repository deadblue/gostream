package multipart

import (
	"bytes"
	"github.com/deadblue/gostream/chain"
	"io"
	"os"
)

const (
	// Boundary size
	boundarySize = 40
	// Multipart form MIME type
	formMimeType = "multipart/form-data; boundary="
)

var (
	hyphens = []byte("--")
	crlf    = []byte("\r\n")
)

// Implementation of Form interface.
// Because the Form should be created by New() function, so I declare it
// as an interface, and hide the implementation struct.
type implForm struct {
	// Archive flag
	archived bool
	// Boundary
	boundary []byte
	// Body size
	size int64
	// Readers for each part.
	parts []io.Reader
}

func (f *implForm) AddValue(name string, value string) Form {
	f.addPart(createValuePart(name, value))
	return f
}

func (f *implForm) AddFile(name string, file *os.File) Form {
	return f.AddMimeFile(name, "", file)
}

func (f *implForm) AddMimeFile(name string, mimeType string, file *os.File) Form {
	info, _ := file.Stat()
	return f.AddFileData(name, mimeType, info, file)
}

func (f *implForm) AddFileData(name string, mimeType string, info FileInfo, data io.Reader) Form {
	f.addPart(createFilePart(name, mimeType, info, data))
	return f
}

func (f *implForm) addPart(size int64, body io.Reader) {
	if !f.archived {
		f.parts = append(f.parts, body)
		f.size += size + int64(boundarySize+6)
	}
}

func (f *implForm) Archive() (mimeType string, size int64, body io.ReadCloser, err error) {
	partCount := len(f.parts)
	if partCount == 0 {
		err = ErrEmptyForm
		return
	}
	if f.archived {
		err = ErrArchivedForm
		return
	}
	f.archived = true
	// Make boundaries
	boundary := bytes.Join([][]byte{
		crlf, hyphens, f.boundary, crlf,
	}, []byte{})
	tail := bytes.Join([][]byte{
		crlf, hyphens, f.boundary, hyphens, crlf,
	}, []byte{})
	// Make body
	readers := make([]io.Reader, partCount*2+1)
	for i := 0; i < partCount; i++ {
		if i == 0 {
			readers[i*2] = bytes.NewReader(boundary[2:])
		} else {
			readers[i*2] = bytes.NewReader(boundary)
		}
		readers[i*2+1] = f.parts[i]
	}
	readers[partCount*2] = bytes.NewReader(tail)
	body = chain.JoinReader(readers...)
	// Fill metadata
	mimeType = formMimeType + string(f.boundary)
	size = f.size
	return
}
