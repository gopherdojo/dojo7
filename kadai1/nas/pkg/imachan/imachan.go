/*Package imachan は画像変換のパッケージです。

https://gopherdojo.connpass.com/event/142892/ にて出された課題です。
https://github.com/gopherdojo/dojo7/tree/kadai1-nas/kadai1/nas を確認してください。

How to use

	c := imachan.NewConfig(path, fromFormat, toFormat)
	data, err :=c.ConvertRec()

この処理で path 配下の対象画像形式のものだけが同ディレクトリに指定画像形式に変換されます。

*/
package imachan

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

// 画像形式を示す一意な定数です。
const (
	DefaultFormat = 0
	JpgFormat     = iota
	PngFormat
	GifFormat
)

// Config は 画像変換に必要な設定情報を格納します。
type Config struct {
	Path       string
	FromFormat int
	ToFormat   int
}

// ConvertedDataRepository は変換前後のイメージファイルパスを格納します。
type ConvertedDataRepository []map[string]string

// NewConfig 構造体 Config を生成します。
func NewConfig(path, fromFormatStr, toFormatStr string) (*Config, error) {
	fromFormat := GetImageFormat(fromFormatStr)
	if fromFormat == DefaultFormat {
		err := fmt.Errorf("undefind %s file format, please choose another", fromFormatStr)
		return nil, err
	}

	toFormat := GetImageFormat(toFormatStr)
	if toFormat == DefaultFormat {
		err := fmt.Errorf("undefind %s file format, please choose another", toFormatStr)
		return nil, err
	}

	return &Config{
		Path:       path,
		FromFormat: fromFormat,
		ToFormat:   toFormat,
	}, nil
}

// ConvertRec は設定をもとに再帰的に画像を変換します。
func (c *Config) ConvertRec() (ConvertedDataRepository, error) {
	var r ConvertedDataRepository

	err := filepath.Walk(c.Path, func(fromPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		// fromPath の画像形式を取得し、c.FromFormat と比較する
		if targetFormat := GetImageFormat(strings.Replace(filepath.Ext(fromPath), ".", "", 1)); targetFormat != c.FromFormat {
			return nil
		}

		toPath, err := Convert(fromPath, c.ToFormat)
		if err != nil {
			return err
		}

		r = append(r, map[string]string{
			"from": fromPath,
			"to":   toPath,
		})

		return nil
	})

	if err != nil {
		return nil, err
	}

	return r, nil
}

// GetImageFormat は画像形式を一意に特定します。
// TODO: 大文字に対応
func GetImageFormat(formatStr string) int {
	switch formatStr {
	case "jpg", "jpeg":
		return JpgFormat

	case "png":
		return PngFormat

	case "gif":
		return GifFormat

	default:
		return DefaultFormat
	}
}

// Convert は任意の形式に画像を変換します。
func Convert(fromPath string, toFormat int) (string, error) {
	var (
		toPath string
		err    error
	)

	switch toFormat {
	case PngFormat:
		toPath, err = ConvertToPng(fromPath)

	case JpgFormat:
		toPath, err = ConvertToJpg(fromPath)

	case GifFormat:
		toPath, err = ConvertToGif(fromPath)
	}

	if err != nil {
		return "", err
	}

	return toPath, nil
}

// ConvertToPng は PNG に画像を変換します。
func ConvertToPng(fromPath string) (string, error) {
	f, err := os.Open(fromPath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	fromImg, _, err := image.Decode(f)
	if err != nil {
		return "", err
	}

	// fromPath の拡張子以前と結合する
	toPath := fromPath[0:len(fromPath)-len(filepath.Ext(fromPath))] + ".png"

	toImg, err := os.Create(toPath)
	if err != nil {
		return "", err
	}
	// TODO: defer のエラー処理
	defer toImg.Close()

	err = png.Encode(toImg, fromImg)
	if err != nil {
		return "", err
	}

	return toPath, nil
}

// ConvertToJpg は JPG に画像を変換します。
func ConvertToJpg(fromPath string) (string, error) {
	f, err := os.Open(fromPath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	fromImg, _, err := image.Decode(f)
	if err != nil {
		return "", err
	}

	// fromPath の拡張子以前と結合する
	toPath := fromPath[0:len(fromPath)-len(filepath.Ext(fromPath))] + ".jpg"

	toImg, err := os.Create(toPath)
	if err != nil {
		return "", err
	}
	// TODO: defer のエラー処理
	defer toImg.Close()

	err = jpeg.Encode(toImg, fromImg, &jpeg.Options{Quality: 100})
	if err != nil {
		return "", err
	}

	return toPath, nil
}

// ConvertToGif は GIF に画像を変換します。
func ConvertToGif(fromPath string) (string, error) {
	f, err := os.Open(fromPath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	fromImg, _, err := image.Decode(f)
	if err != nil {
		return "", err
	}

	// fromPath の拡張子以前と結合する
	toPath := fromPath[0:len(fromPath)-len(filepath.Ext(fromPath))] + ".gif"

	toImg, err := os.Create(toPath)
	if err != nil {
		return "", err
	}
	// TODO: defer のエラー処理
	defer toImg.Close()

	err = gif.Encode(toImg, fromImg, &gif.Options{NumColors: 256})
	if err != nil {
		return "", err
	}

	return toPath, nil
}
