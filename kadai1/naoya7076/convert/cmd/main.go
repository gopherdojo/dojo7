package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gopherdojo/dojo7/kadai1/naoya7076/convert"
)

var (
	source = flag.String("s","./","指定したディレクトリ以下を再帰的に捜査します")
	from = flag.String("f",".jpg","指定した拡張子の画像を検索します")
	dest = flag.String("d",".png","指定した拡張子の画像に変換します")
)
func main() {
	flag.Parse()
	isOldExtSupported := convert.IsFormatSupported(*from)
	if isOldExtSupported == false{
		fmt.Println("変換前の拡張子が対応していません。")
		os.Exit(1)
	}

	isNewExtSupported := convert.IsFormatSupported(*dest)
	if isNewExtSupported == false{
		fmt.Println("変換後の拡張子が対応していません。")
		os.Exit(1)
	}


	err := filepath.Walk(*source, func(path string, info os.FileInfo, err error) error {
		if convert.IsFormatSupported(filepath.Ext(path)) == true {
			convert.Image(filepath.Base(path), *from,*dest)
		}
		return nil
	})
	if err != nil {
		fmt.Println("ファイルを開けませんでした")
		os.Exit(1)
	}

}
