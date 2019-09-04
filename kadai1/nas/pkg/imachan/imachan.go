package imachan

import (
	"fmt"
	"image"
	_ "image/jpeg" // to convert jpeg file
	"image/png"
	"os"
	"path/filepath"
)

// Convert ...
func Convert(fromPath, fromFormat, toFormat string) error {
	f, err := os.Open(fromPath)
	if err != nil {
		return err
	}
	defer f.Close()
	fromImg, _, err := image.Decode(f)
	if err != nil {
		return err
	}
	toPath := fromPath[0:len(fromPath)-len(filepath.Ext(fromPath))] + ".png"
	toImg, err := os.Create(toPath)
	err = png.Encode(toImg, fromImg)
	if err != nil {
		return err
	}
	fmt.Println("success:", fromPath, "->", toPath)
	return nil
}
