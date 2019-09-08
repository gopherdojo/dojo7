package convert

import (
	"fmt"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

func ConvertToPng(src string) error {
	sf, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sf.Close()

	img, err := jpeg.Decode(sf)
	if err != nil {
		fmt.Println("画像を解析できませんでした")
	}

	savefile, err := os.Create(filepath.Dir(src) + "test.png")
	if err != nil {
		fmt.Println("保存するためのファイルが作成できませんでした。")
		os.Exit(1)
	}
	defer savefile.Close()

	png.Encode(savefile, img)
	return nil
}
