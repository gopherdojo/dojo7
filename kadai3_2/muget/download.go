package muget

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func DownloadFile(url string, path string, downloadStartSize, downloadEndSize int) error {
	// Create the file
	out, err := os.Create(path + filepath.Base(url))
	if err != nil {
		return err
	}
	defer func() {
		if err := out.Close(); err != nil {
			log.Printf("\x1b[31m%s:%s\x1b[0m\n", "[encode error]", err)
		}
	}()

	// Get the data
	resp, err := RangeRequest(url, downloadStartSize, downloadEndSize)
	if err != nil {
		return err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("\x1b[31m%s:%s\x1b[0m\n", "[encode error]", err)
		}
	}()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

// RangeRequest return *http.Response include context and range header
func RangeRequest(url string, low, high int) (*http.Response, error) {
	// create get request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// set download ranges
	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", low, high))

	return http.DefaultClient.Do(req)
}
