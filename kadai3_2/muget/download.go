package muget

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func DownloadFile(url string, path string) error {
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
	resp, err := http.Get(url)
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
