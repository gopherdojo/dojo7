package convert

import (
	"fmt"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

func ToPng(src string) {
	sf, err := os.Open(src)
	if err != nil {
		os.Exit(1)
	}
	defer sf.Close()

	img, err := jpeg.Decode(sf)
	if err != nil {
		fmt.Println("画像を解析できませんでした")
	}

	savefile, err := os.Create(filepath.Base(src) + ".png")
	if err != nil {
		fmt.Println("保存するためのファイルが作成できませんでした。")
		os.Exit(1)
	}
	defer savefile.Close()

	png.Encode(savefile, img)
	return
}
