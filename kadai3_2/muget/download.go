package muget

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func DownloadFile(url string, path string, downloadStartSize, downloadEndSize, downloadCount int) (err error) {
	// Create the file
	out, err := os.Create(path + fmt.Sprint(downloadCount) + filepath.Ext(url))
	if err != nil {
		return err
	}
	defer func() {
		err = out.Close()
		if err != nil {
			return
		}
	}()

	// Get the data
	resp, err := RangeRequest(url, downloadStartSize, downloadEndSize)
	if err != nil {
		return err
	}
	defer func() {
		err = resp.Body.Close()
		if err != nil {
			return
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
