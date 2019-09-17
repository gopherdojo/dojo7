package myimage

import (
	"errors"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

// 拡張子と対応する変換関数のマッピング
var format = map[string]map[string]ConvertFunc{
	"jpg": {
		"png": jpg2png,
	},
	"jpeg": {
		"png": jpg2png,
	},
	"png": {
		"jpg": png2jpg,
	},
}

// ConvertFunc is express of convert functions.
type ConvertFunc func(path string) error

// 拡張子削除関数
func getFilePathWithoutExt(path string) string {
	return path[:len(path)-len(filepath.Ext(path))]
}

// ----------------------------
// 変換関数
//-----------------------------
// JPEG -> PNG
func jpg2png(path string) error {
	inFile, err := os.Open(path)
	if err != nil {
		return err
	}
	img, err := jpeg.Decode(inFile)
	if err != nil {
		return err
	}

	outFile, err := os.Create(getFilePathWithoutExt(path) + ".png")
	if err != nil {
		return err
	}

	return png.Encode(outFile, img)
}

// PNG -> JPG
func png2jpg(path string) error {
	inFile, err := os.Open(path)
	if err != nil {
		return err
	}
	img, err := png.Decode(inFile)
	if err != nil {
		return err
	}
	outFile, err := os.Create(getFilePathWithoutExt(path) + ".jpg")
	if err != nil {
		return err
	}

	return jpeg.Encode(outFile, img, &jpeg.Options{Quality: 100})
}

// GetConvertFunc is in order to get ConvertFunc by from and to extension.
func GetConvertFunc(fromExt string, toExt string) (convfunc ConvertFunc, err error) {
	if val, ok := format[fromExt][toExt]; ok {
		return val, nil
	}
	return nil, errors.New("Convert function not found")
}
