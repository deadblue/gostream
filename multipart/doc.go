/*
Unlike "mime/multipart" in standard library, this package provides
a read-on-demand multipart form, which will save huge memory when
upload large file(s) through HTTP post.

Exmaple:
	import (
		"github.com/deadblue/gostream/multipart"
		"net/http"
		"os"
	)

	func main() {
		// Open file for uploading
		file, err := os.Open("/path/to/file")
		if err != nil {
			panic(err)
		}
		// Caller DOES NOT need to close the file, it will be closed by http.Client after sent.

		// Create form
		form := multipart.New()
		// Add file and string values
		form.AddFile("file", file).
			AddValue("foo", "1").
			AddValue("bar", "2")
		// Create HTTP request
		req, err := multipart.MakeRequest("http://server/upload", form)
		if err != nil {
			panic(err)
		}
		// Send request
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		// Process the response here.
	}
*/
package multipart
