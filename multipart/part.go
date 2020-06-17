package multipart

import (
	"bytes"
	"fmt"
	"github.com/deadblue/gostream/chain"
	"io"
)

func createValuePart(name, value string) (size int64, reader io.Reader) {
	body := &bytes.Buffer{}
	body.WriteString(fmt.Sprintf("Content-Disposition: form-data; name=\"%s\"\r\n\r\n", name))
	body.WriteString(value)
	return int64(body.Len()), body
}

func createFilePart(name string, mimeType string, info FileInfo, data io.Reader) (size int64, reader io.Reader) {
	// Set default MIME type
	if mimeType == "" {
		mimeType = "application/octet-stream"
	}
	header := &bytes.Buffer{}
	header.WriteString(fmt.Sprintf("Content-Disposition: form-data; name=\"%s\"; filename=\"%s\"\r\n", name, info.Name()))
	header.WriteString(fmt.Sprintf("Content-Type: %s\r\n\r\n", mimeType))
	return int64(header.Len()) + info.Size(), chain.JoinReader(header, data)
}
