// Replacer is a package that can convert to the specified image format (jpg, png) by generating File structure.
// Supported formats png,jpg
package replacer

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
)

// A structure that stores image files.
type File struct{}

// This method encodes an image file into jpg or png.
func (f *File) Encode(path, to string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer fileClose(file)

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	// create output file
	outPath := path[:len(path)-len(filepath.Ext(path))] + "." + to
	out, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer fileClose(out)

	// select encoder
	switch to {
	case "jpg":
		if err := jpeg.Encode(out, img, &jpeg.Options{Quality: 100}); err != nil {
			return err
		}
	case "png":
		if err := png.Encode(out, img); err != nil {
			return err
		}
	default:
		// delete fail file
		if err := os.Remove(outPath); err != nil {
			return err
		}
		return fmt.Errorf("%s is unsupported extension", to)
	}

	// delete original file
	if err := os.Remove(path); err != nil {
		return err
	}

	return nil
}

func fileClose(file *os.File) {
	if err := file.Close(); err != nil {
		log.Printf("\x1b[31m%s:%s\x1b[0m\n", "[encode error]", err)
	}
}
