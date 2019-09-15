package convert

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

var imageExtMap = map[string]bool{
	".jpg" : true,
	".jpeg" : true,
	".png" : true,
	".gif" : true,
}



func ChangeFileExt(path string,newExt string)string{
	changedExtFilePath := strings.Replace(path,filepath.Ext(path),newExt,1)
	fmt.Println(changedExtFilePath)
	return changedExtFilePath
}

func IsFormatSupported(ext string)bool{
	_, ok := imageExtMap[ext]
	return ok
}


func Image(src string,oldExt string,newExt string) {
	sf, err := os.Open(src)
	if err != nil {
		os.Exit(1)
	}
	defer sf.Close()

	var img image.Image
	switch oldExt{
	case ".jpeg","jpg":
		img, err = jpeg.Decode(sf)
		if err != nil {
			fmt.Println("画像を解析できませんでした。",err)
		}
	case ".png":
		img, err = png.Decode(sf)
		if err != nil {
			fmt.Println("画像を解析できませんでした。",err)
		}
	case ".gif":
		img, err = gif.Decode(sf)
		if err != nil {
			fmt.Println("画像を解析できませんでした。",err)
		}
	}
	newFilePath := ChangeFileExt(src,newExt)
	newFile, err := os.Create(newFilePath)
	if err != nil {
		fmt.Println("保存するためのファイルが作成できませんでした。",err)
		os.Exit(1)
	}
	defer newFile.Close()

	switch newExt {
	case ".jpeg","jpg":
		err = jpeg.Encode(newFile,img,nil)
		if err != nil {
			fmt.Println("ファイルをエンコードできませんでした。", err)
			os.Exit(1)
		}
	case ".png":
		err = png.Encode(newFile,img)
		if err != nil {
			fmt.Println("ファイルをエンコードできませんでした。",err)
			os.Exit(1)
		}
	case ".gif":
		fmt.Println(newFile)
		err = gif.Encode(newFile,img,nil)
		if err != nil {
			fmt.Println("ファイルをエンコードできませんでした。", err)
			os.Exit(1)
		}
	}
}