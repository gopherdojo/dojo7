package image

import (
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)


type ConvExts struct {
	inExt, outExt string
}

func NewFormats(in, out string) ConvExts{
	if in == ""{
		in = "jpg"
	}

	if out == "" {
		out = "png"
	}
	return ConvExts{inExt: in, outExt: out}
}

func FmtConv(path string, exts ConvExts)(err error){
	pathWithoutExt := filepath.Base(path)
	ext := filepath.Ext(path)

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

	return png.Encode(outputFile, img)
}
