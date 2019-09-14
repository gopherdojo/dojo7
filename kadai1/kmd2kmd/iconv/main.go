package iconv

import (
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
)

type Image struct {
	Path string
	In   string
	Out  string
}

// jpegに変換し保存する｡Qualityは80で固定｡
func convertToJpeg(img image.Image, dest string) error {

	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer func() {
		deferErr := out.Close()
		if deferErr != nil {
			err = deferErr
		}
	}()
	if err != nil {
		return err
	}

	opts := &jpeg.Options{Quality: 80}

	err = jpeg.Encode(out, img, opts)
	if err != nil {
		return err
	}
	return nil
}

// pngに変換し保存する｡
func convertToPng(img image.Image, dest string) error {

	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer func() {
		err := out.Close()
		if err != nil {
			log.Println("failed close", dest, err)
		}
	}()

	err = png.Encode(out, img)
	if err != nil {
		return err
	}
	return nil
}

// image.Imageをデコードする
func getDecodedImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Println("failed close"+path, err)
		}
	}()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return img, nil
}

// 拡張子を除いたファイルパスと取得する
func getFileNameWithoutExt(path string) (string, error) {
	if path == "" {
		return "", errors.New("an empty string was entered")
	}
	dir := filepath.Dir(path)
	baseWithoutExt := filepath.Base(path[:len(path)-len(filepath.Ext(path))])
	return filepath.Join(dir, baseWithoutExt), nil
}

// イメージを変換し保存する
func (c Image) Do() error {
	img, err := getDecodedImage(c.Path)
	if err != nil {
		return err
	}
	dest, err := getFileNameWithoutExt(c.Path)
	if err != nil {
		return err
	}

	switch c.Out {
	case "jpg":
		return convertToJpeg(img, dest+".jpg")
	case "png":
		return convertToPng(img, dest+".png")
	default:
		return errors.New("unsupported format")
	}
}
