// Package image は画像のフォーマットを変換するためのパッケージです。
package image

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

var supportedFormats = map[string]bool{"jpg": true, "jpeg": true, "png": true, "gif": true}

func SupportedFormat(ext string) bool {
	// ドット始まりの拡張子ならドットを削除する
	if len(ext) > 0 && ext[0] == '.' {
		ext = ext[1:]
	}
	_, ok := supportedFormats[ext]
	return ok
}

// ConvExts は変換対象のフォーマットと変換先のフォーマットを表す構造体です
type ConvExts struct {
	InExt, OutExt string
}

// SupportedFormats は指定フォーマットが対応しているか確認するメソッドです
func (c *ConvExts) SupportedFormats() bool {
	return SupportedFormat(c.InExt) && SupportedFormat(c.OutExt)
}

// NewConvExts は変換対象のフォーマットと変換先のフォーマットを表す構造体です
func NewConvExts(in, out string) ConvExts {
	if in == "" {
		in = "jpg"
	}

	if out == "" {
		out = "png"
	}
	return ConvExts{InExt: in, OutExt: out}
}

// FmtConv は指定されたフォーマットからフォーマットへ変換する関数です
func FmtConv(f *os.File, exts ConvExts) (err error) {

	var img image.Image
	var decodeErr error

	switch exts.InExt {
	case "jpeg", "jpg":
		img, decodeErr = jpeg.Decode(f)
	case "png":
		img, decodeErr = png.Decode(f)
	case "gif":
		img, decodeErr = gif.Decode(f)
	}

	if decodeErr != nil {
		return decodeErr
	}

	path := f.Name()
	pathWithoutExt := path[:len(path)-len(filepath.Ext(path))+1]
	outputFile, err := os.Create(pathWithoutExt + exts.OutExt)
	if err != nil {
		return err
	}
	defer func() {
		cerr := outputFile.Close()
		if err == nil {
			err = cerr
		}
	}()

	switch exts.OutExt {
	case "jpeg", "jpg":
		err = jpeg.Encode(outputFile, img, nil)
	case "png":
		err = png.Encode(outputFile, img)
	case "gif":
		err = gif.Encode(outputFile, img, nil)
	}

	return
}
