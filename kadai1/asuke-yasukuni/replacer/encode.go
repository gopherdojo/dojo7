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

type File struct {
	Path    string
	FromExt string
	ToExt   string
}

func (f *File) Encode() error {
	file, err := os.Open(f.Path)
	if err != nil {
		return err
	}
	defer fileClose(file)

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	out, err := os.Create(f.Path[:len(f.Path)-len(filepath.Ext(f.Path))] + "." + f.ToExt)
	if err != nil {
		return err
	}
	defer fileClose(out)

	switch f.ToExt {
	case "jpg":
		if err := jpeg.Encode(out, img, &jpeg.Options{Quality: 100}); err != nil {
			return err
		}
	case "png":
		if err := png.Encode(out, img); err != nil {
			return err
		}
	default:
		return fmt.Errorf("%s is unsupported extension", f.ToExt)
	}

	if err := os.Remove(f.Path); err != nil {
		return err
	}

	return nil
}

func fileClose(file *os.File) {
	if err := file.Close(); err != nil {
		log.Printf("\x1b[31m%s:%s\x1b[0m\n", "[encode error]", err)
	}
}
