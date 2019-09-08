/*
Imgconvt is a tool that can covnert images to another format under specified directory.
*/

package imgconvt

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

// Conv has Path, FromExt and ToExt
type Conv struct {
	Path    string // image file path
	FromExt string // image file extensiton that you want to convert from
	ToExt   string // image file extensiton that you want to convert to
}

// ConvertImage convert image FromExt to ToExt
func ConvertImage(c *Conv) error {

	img, err := decode(c.Path, c.FromExt)
	if err != nil {
		return fmt.Errorf("failed to decode image of %s : %v", c.Path, err)
	}

	err = encode(img, c.ToExt, c.getFileName())

	if err != nil {
		return fmt.Errorf("failed to encode image of %s : %v", c.Path, err)
	}

	return nil

}

func (c *Conv) getFileName() string {
	return filepath.Base(strings.TrimSuffix(c.Path, filepath.Ext(c.Path)))
}

func decode(path string, from string) (image.Image, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	switch ImageExt(from) {
	case ImageExtJpg, ImageExtJpeg:
		return jpeg.Decode(file)
	case ImageExtPng:
		return png.Decode(file)
	case ImageExtGif:
		return gif.Decode(file)
	default:
		return nil, fmt.Errorf("image extension not supported %s", ImageExt(from))
	}
}

func encode(img image.Image, to string, name string) error {

	w, err := os.Create(createFileName(name, to))

	if err != nil {
		return fmt.Errorf("failed to create file")
	}

	defer w.Close()

	switch ImageExt(to) {
	case ImageExtJpg, ImageExtJpeg:
		return jpeg.Encode(w, img, nil)
	case ImageExtPng:
		return jpeg.Encode(w, img, nil)
	case ImageExtGif:
		return gif.Encode(w, img, nil)
	default:
		return fmt.Errorf("image extension not supported %s", ImageExt(to))
	}

}

func createFileName(name string, ext string) string {
	return name + ext
}
