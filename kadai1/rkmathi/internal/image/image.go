// 画像について表現するパッケージ
package image

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

var (
	Extensions []string
)

// 画像の形式を表現するタイプ
type Type struct {
	Type string
}

func init() {
	Extensions = []string{".gif", ".jpeg", ".jpg", ".png"}
}

// 画像のファイルパスの拡張子から、 Type を特定する。
func DetectImageType(path string) Type {
	switch filepath.Ext(path) {
	case ".gif":
		return Type{"GIF"}
	case ".jpeg", ".jpg":
		return Type{"JPG"}
	case ".png":
		return Type{"PNG"}
	}

	return Type{"Unknown"}
}

// 指定したTypeの画像をDecodeする。
func Decode(imageType Type, srcFile *os.File) (image.Image, error) {
	var img image.Image
	var err error

	switch imageType.Type {
	case "GIF":
		img, err = gif.Decode(srcFile)
	case "JPG":
		img, err = jpeg.Decode(srcFile)
	case "PNG":
		img, err = png.Decode(srcFile)
	}
	if err != nil {
		return nil, err
	}

	return img, err
}

// 指定したTypeの画像をEncodeする。
func Encode(imageType Type, srcImg image.Image, dstFile *os.File) error {
	var err error

	switch imageType.Type {
	case "GIF":
		err = gif.Encode(dstFile, srcImg, nil)
	case "JPG":
		err = jpeg.Encode(dstFile, srcImg, nil)
	case "PNG":
		err = png.Encode(dstFile, srcImg)
	}

	return err
}
