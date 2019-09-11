package convert

import (
	"fmt"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

var imageExts = map[string]bool{
	"jpg" : false,
	"jpeg" : false,
	"png" : false,
	"gif" : false,
}

func changeFileExt(path string,ext string)string{
	oldFilePath := filepath.Base(path)
	changedExtFilePath := strings.Replace(oldFilePath,filepath.Ext(path),ext,1)
	return changedExtFilePath
}

func isFormatSupported(ext string)bool{
	_,ok := imageExts[ext]
	return ok
}


func ToPng(src string) {
	sf, err := os.Open(src)
	if err != nil {
		os.Exit(1)
	}
	defer sf.Close()

	img, err := jpeg.Decode(sf)
	if err != nil {
		fmt.Fprintf(os.Stderr,"画像を解析できませんでした。",err)
	}

	newFilePath := changeFileExt(src,".png")
	savefile, err := os.Create(newFilePath)//エラー処理する
	if err != nil {
		fmt.Fprintf(os.Stderr,"保存するためのファイルが作成できませんでした。",err)
		os.Exit(1)
	}
	defer savefile.Close()

	png.Encode(savefile, img)
	return
}
