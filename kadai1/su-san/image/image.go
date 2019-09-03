// image は画像のフォーマットを変換するためのパッケージです。
package image

import (
	"fmt"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

//ConvExts は変換対象のフォーマットと変換先のフォーマットを表す構造体です
type ConvExts struct {
	inExt, outExt string
}

//NewConvExts は変換対象のフォーマットと変換先のフォーマットを表す構造体です
func NewConvExts(in, out string) ConvExts{
	if in == ""{
		in = ".jpg"
	}

	if out == "" {
		out = ".png"
	}
	return ConvExts{inExt: in, outExt: out}
}

// FmtConv は指定されたフォーマットからフォーマットへ変換する関数です
func FmtConv(path string, exts ConvExts)(err error){
	pathWithoutExt := path[:len(path)-len(filepath.Ext(path))]
	ext := filepath.Ext(path)

	fmt.Println(pathWithoutExt, ext)

	// 別フォーマットのファイルはスルーする
	if ext != exts.inExt {
		return nil
	}

	f, err := os.Open(path)
	defer f.Close()

	if err != nil {
		return err
	}

	// jpegファイルをデコード
	img, err := jpeg.Decode(f)
	if err != nil {
		return err
	}

	outputFile, err := os.Create(pathWithoutExt + exts.outExt)
	defer outputFile.Close()

	if err != nil {
		return err
	}

	err = png.Encode(outputFile, img)
	if err == nil {
		err = os.Remove(path)
	}
	return err
}
