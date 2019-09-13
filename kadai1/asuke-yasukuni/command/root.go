// Recursive image encoder command implementation.
package command

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gopherdojo/dojo7/asuke-yasukuni/replacer"
)

// Recursively search the directory and perform encoding.
func WalkEncoder(src *string, fromExt, toExt string) (encodeFiles []string, err error) {
	var file replacer.File
	err = filepath.Walk(*src, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) != "."+fromExt {
			return nil
		}

		// Use to output.
		encodeFiles = append(encodeFiles, fmt.Sprintf("%s%s -> %s", "[replace file]", path, toExt))

		file = replacer.File{
			Path:    path,
			FromExt: fromExt,
			ToExt:   toExt,
		}
		if err := file.Encode(); err != nil {
			return err
		}

		return nil
	})

	return
}
