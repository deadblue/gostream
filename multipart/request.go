package multipart

import (
	"context"
	"net/http"
)

const (
	headerContentType = "Content-Type"
)

// NewRequest wraps NewRequestWithContext using the background context.
func NewRequest(url string, form Form) (req *http.Request, err error) {
	return NewRequestWithContext(context.Background(), url, form)
}

// Make a HTTP request for posting form to url.
//
// If the returned req is used by standard http.Client, the form body
// will be closed automatically after used. Otherwise, caller may need
// to close req.Body after used.
func NewRequestWithContext(ctx context.Context, url string, form Form) (req *http.Request, err error) {
	mimeType, size, body, err := form.Archive()
	if err != nil {
		return
	}
	if req, err = http.NewRequestWithContext(ctx, http.MethodPost, url, body); err == nil {
		req.Header.Add(headerContentType, mimeType)
		req.ContentLength = size
	}
	return
}
