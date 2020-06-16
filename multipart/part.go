package multipart

import (
	"bytes"
	"fmt"
	"github.com/deadblue/gostream/chain"
	"io"
)

const (
	fieldPart = "Content-Disposition: form-data; name=\"%s\"\n\n%s"

	filePartHeader = "Content-Disposition: form-data; name=\"%s\"; filename=\"%s\"\n" +
		"Content-Type: application/octet-stream\n\n"
)

type part struct {
	size   int64
	reader io.Reader
}

func createFieldPart(name, value string) *part {
	data := []byte(fmt.Sprintf(fieldPart, name, value))
	return &part{
		size:   int64(len(data)),
		reader: bytes.NewReader(data),
	}
}

func createFilePart(name string, info FileInfo, data io.Reader) *part {
	// Header data
	header := []byte(fmt.Sprintf(filePartHeader, name, info.Name()))
	size := int64(len(header)) + info.Size()
	return &part{
		size:   size,
		reader: chain.JoinReader(bytes.NewReader(header), data),
	}
}
