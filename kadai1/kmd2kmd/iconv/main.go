package iconv

import (
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
)

type IConverter struct {
	Path string
	In   string
	Out  string
}

// jpegに変換し保存する｡Qualityは80で固定｡
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

// pngに変換し保存する｡
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

// image.Imageをデコードする
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

// 拡張子を除いたファイルパスと取得する
func getFileNameWithoutExt(path string) string {
	dir := filepath.Dir(path)
	baseWithoutExt := filepath.Base(path[:len(path)-len(filepath.Ext(path))])
	return filepath.Join(dir, baseWithoutExt)
}

// イメージを変換し保存する
func (c IConverter) Convert() {
	img := getDecodedImage(c.Path)
	dest := getFileNameWithoutExt(c.Path)

	switch c.Out {
	case "jpg":
		convertToJpeg(img, dest+".jpg")
	case "png":
		convertToPng(img, dest+".png")
	default:
		log.Fatal("unsupported format")
	}
}
