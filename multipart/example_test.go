package multipart

import (
	"bytes"
	"net/http"
	"os"
)

func Example() {
	// Create form with values
	form := New().
		AddValue("foo", "1").
		AddValue("bar", "2")

	// Open local file for uploading
	// Caller DOES NOT need to close the file, it will be closed by http.Client after sent.
	file, err := os.Open("/path/to/file")
	if err != nil {
		panic(err)
	}
	// Add file to form
	form.AddFile("realfile", file)

	// Add memory file to form
	data := []byte("Hello, world!")
	info := SimpleFileInfo("hello.txt", int64(len(data)))
	form.AddFileData("memfile", info, bytes.NewReader(data))

	// Create HTTP request
	req, err := MakeRequest("http://server/upload", form)
	if err != nil {
		panic(err)
	}
	// Send request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// TODO: Process the response here.
}
