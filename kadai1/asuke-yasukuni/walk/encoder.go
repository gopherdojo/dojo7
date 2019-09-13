// Recursive image encoder command implementation.
package walk

import (
	"fmt"
	"os"
	"path/filepath"
)

type File interface {
	Encode(path, toExt string) error
}

type Walk struct {
	File File
}

// Recursively search the directory and perform encoding.
func (w *Walk) Encoder(src *string, fromExt, toExt string) (encodeFiles []string, err error) {
	err = filepath.Walk(*src, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) != "."+fromExt {
			return nil
		}

		// Use to output.
		encodeFiles = append(encodeFiles, fmt.Sprintf("%s%s -> %s", "[replace file]", path, toExt))

		if err := w.File.Encode(path, toExt); err != nil {
			return err
		}

		return nil
	})

	return
}
