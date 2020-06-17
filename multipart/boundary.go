package multipart

import "math/rand"

var (
	boundaryChars = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	boundaryCharCount = len(boundaryChars)
)

func makeBoundary(len int) []byte {
	buf := make([]byte, len)
	for i := 0; i < len; i++ {
		buf[i] = boundaryChars[rand.Intn(boundaryCharCount)]
	}
	return buf
}
