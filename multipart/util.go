package multipart

import (
	"net/http"
	"strconv"
)

const (
	headerContentType   = "Content-Type"
	headerContentLenght = "Content-Length"
)

// Helper function to make a HTTP request for posting form to url.
//
// If req is used by standard http.Client, the form body will be closed
// automatically after used.
// Otherwise, caller may need to close req.Body after used.
func MakeRequest(url string, form Form) (req *http.Request, err error) {
	mimeType, size, body, err := form.Archive()
	if err != nil {
		return
	}
	if req, err = http.NewRequest(http.MethodPost, url, body); err == nil {
		req.Header.Add(headerContentType, mimeType)
		req.Header.Add(headerContentLenght, strconv.FormatInt(size, 10))
	}
	return
}
