package multipart

import (
	"io"
	"os"
)

const (
	mimeType = "multipart/form-data; boundary="
)

type Form struct {
	archived bool
	boundary string
	size     int64
	parts    []*part
}

func (f *Form) AddField(name, value string) {
	f.addPart(createFieldPart(name, value))
}

func (f *Form) AddFile(name string, file *os.File) {
	info, _ := file.Stat()
	f.AddFileData(name, info, file)
}

func (f *Form) AddFileData(name string, info FileInfo, data io.Reader) {
	f.addPart(createFilePart(name, info, data))
}

func (f *Form) addPart(p *part) {
	f.parts = append(f.parts, p)
	f.size += p.size + int64(len(f.boundary)+3)
}

func (f *Form) Metadata() (contentType string, size int64) {
	return mimeType + f.boundary, f.size
}

func (f *Form) Archive() (body io.Reader, err error) {
	if len(f.parts) == 0 {
		return nil, ErrEmptyForm
	}
	if f.archived {
		return nil, ErrArchivedForm
	}
	f.archived = true
	return
}

func New() *Form {
	boundary := makeBoundary()
	// Create form
	return &Form{
		archived: false,
		boundary: boundary,
		size:     int64(len(boundary) + 4),
		parts:    make([]*part, 0),
	}
}
