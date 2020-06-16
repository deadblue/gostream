package multipart

import "math/rand"

const (
	boundaryChars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	boundarySize  = 60

	hyphens = "--"
	crlf    = "\r\n"
)

var charCount = len(boundaryChars)

func makeBoundary() string {
	buf := make([]byte, boundarySize)
	for i := 0; i < boundarySize; i++ {
		buf[i] = boundaryChars[rand.Intn(charCount)]
	}
	return string(buf)
}
