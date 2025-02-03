package observe

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/deadblue/gostream/quietly"
)

type DownloadObserver struct {
	total, complete int64
}

func (o *DownloadObserver) OnTransfer(n int) {
	const ProgressBarSize = 60

	o.complete += int64(n)
	if o.total > 0 {
		// Calculate download percent
		percent := float64(o.complete) / float64(o.total)
		// Calculate block and remain size.
		block := int(ProgressBarSize * percent)
		if block > ProgressBarSize {
			block = ProgressBarSize
		}
		remain := ProgressBarSize - block
		bar := strings.Repeat("#", block) + strings.Repeat(" ", remain)
		// Print progress bar with percent.
		fmt.Printf("\r%s %.2f%%", bar, percent*100)
	} else {
		fmt.Printf("\rDownloaded size => %d", o.complete)
	}
}
func (o *DownloadObserver) OnStop(_ error) {
	fmt.Println()
}

// Use observe.Reader to monitor downloading progress.
func ExampleReader() {
	// Open URL for download
	url := "https://server/file/to/download"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Send request error: %s", err)
	}

	// Create a temp file for writing
	tmpFile, err := ioutil.TempFile("", "")
	if err != nil {
		log.Fatalf("Create temp file error: %s", err)
	}
	defer func() {
		// Delete temp file after use.
		_ = os.Remove(tmpFile.Name())
	}()

	// Wrap response body with DownloadObserver
	r := Reader(resp.Body, &DownloadObserver{
		total:    resp.ContentLength,
		complete: 0,
	})
	// Caller can close either r or resp.Body as he like.
	defer quietly.Close(r)
	// Download content and write it to file
	if _, err := io.Copy(tmpFile, r); err != nil {
		log.Fatalf("Download file error: %s", err)
	} else {
		log.Printf("File download to: %s", tmpFile.Name())
	}

}
