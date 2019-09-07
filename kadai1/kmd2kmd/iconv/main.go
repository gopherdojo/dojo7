package iconv

import (
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
)

func convertToJpeg(img image.Image, dest string) {

	out, err := os.Create(dest)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := out.Close()
		if err != nil {
			log.Println("can't close"+dest, err)
		}
	}()

	opts := &jpeg.Options{Quality: 80}

	err = jpeg.Encode(out, img, opts)
	if err != nil {
		log.Fatal(err)
	}
}

func convertToPng(img image.Image, dest string) {

	out, err := os.Create(dest)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := out.Close()
		if err != nil {
			log.Println("can't close"+dest, err)
		}
	}()

	err = png.Encode(out, img)
	if err != nil {
		log.Fatal(err)
	}
}

func getDecodedImage(path string) image.Image {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Println("can't close"+path, err)
		}
	}()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	return img
}

func getFileNameWithoutExt(path string) string {
	dir := filepath.Dir(path)
	baseWithoutExt := filepath.Base(path[:len(path)-len(filepath.Ext(path))])
	return filepath.Join(dir, baseWithoutExt)
}

func Convert(path string, format string) {
	img := getDecodedImage(path)
	dest := getFileNameWithoutExt(path)

	switch format {
	case "jpg":
		convertToJpeg(img, dest+".jpg")
	case "png":
		convertToPng(img, dest+".png")
	default:
		log.Fatal("unsupported format")
	}
}
