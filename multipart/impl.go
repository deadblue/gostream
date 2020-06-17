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

	emptyBytes []byte
)

type implForm struct {
	// Archive flag
	// One form can be archived only once.
	archived bool

	// Boundary
	boundary []byte

	size  int64
	parts []*part
}

func (f *implForm) AddString(name string, value string) {
	f.addPart(createFieldPart(name, value))
}

func (f *implForm) AddFile(name string, file *os.File) {
	info, _ := file.Stat()
	f.AddFileData(name, info, file)
}

func (f *implForm) AddFileData(name string, info FileInfo, data io.Reader) {
	f.addPart(createFilePart(name, info, data))
}

func (f *implForm) addPart(p *part) {
	if !f.archived {
		f.parts = append(f.parts, p)
		f.size += p.size + int64(boundarySize+6)
	}
}

func (f *implForm) Archive() (mimeType string, body io.Reader, size int64, err error) {
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
	// Prepare boundary parts
	head := bytes.Join([][]byte{
		hyphens, f.boundary, crlf,
	}, emptyBytes)
	middle := bytes.Join([][]byte{
		crlf, hyphens, f.boundary, crlf,
	}, emptyBytes)
	tail := bytes.Join([][]byte{
		crlf, hyphens, f.boundary, hyphens, crlf,
	}, emptyBytes)
	// Make body
	readers := make([]io.Reader, partCount*2+1)
	for i := 0; i < partCount; i++ {
		if i == 0 {
			readers[i*2] = bytes.NewReader(head)
		} else {
			readers[i*2] = bytes.NewReader(middle)
		}
		readers[i*2+1] = f.parts[i].reader
	}
	readers[partCount*2] = bytes.NewReader(tail)
	body = chain.JoinReader(readers...)
	// Fill metadata
	mimeType = formMimeType + string(f.boundary)
	size = f.size
	return
}
