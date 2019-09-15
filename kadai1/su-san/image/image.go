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

func supportedFormat(ext string) bool {
	_, ok := supportedFormats[ext]
	return ok
}

// ConvExts は変換対象のフォーマットと変換先のフォーマットを表す構造体です
type ConvExts struct {
	inExt, outExt string
}

// SupportedFormats は指定フォーマットが対応しているか確認するメソッドです
func (c *ConvExts) SupportedFormats() bool {
	return supportedFormat(c.inExt) && supportedFormat(c.outExt)
}

// NewConvExts は変換対象のフォーマットと変換先のフォーマットを表す構造体です
func NewConvExts(in, out string) ConvExts {
	if in == "" {
		in = "jpg"
	}

	if out == "" {
		out = "png"
	}
	return ConvExts{inExt: in, outExt: out}
}

// FmtConv は指定されたフォーマットからフォーマットへ変換する関数です
func FmtConv(path string, exts ConvExts) (err error) {
	pathWithoutExt := path[:len(path)-len(filepath.Ext(path))+1]
	ext := filepath.Ext(path)[1:]

	// 別フォーマットのファイルもしくは拡張子がない場合はスルーする
	if ext != exts.inExt {
		return nil
	}

	f, err := os.Open(path)
	defer func() {
		cerr := f.Close()
		if err == nil {
			err = cerr
		}
	}()

	if err != nil {
		return err
	}

	var img image.Image
	var decodeErr error

	switch exts.inExt {
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

	outputFile, err := os.Create(pathWithoutExt + exts.outExt)
	if err != nil {
		return err
	}
	defer func() {
		cerr := outputFile.Close()
		if err == nil {
			err = cerr
		}
	}()

	var encodeErr error
	switch exts.outExt {
	case "jpeg", "jpg":
		encodeErr = jpeg.Encode(outputFile, img, nil)
	case "png":
		encodeErr = png.Encode(outputFile, img)
	case "gif":
		encodeErr = gif.Encode(outputFile, img, nil)
	}

	if encodeErr != nil {
		return encodeErr
	}

	err = os.Remove(path)
	return
}
