package main

import (
	"fmt"
	"github.com/gopherdojo/dojo7/kadai1/naoya7076/convert"
	"os"
	"path/filepath"
)

func main() {
	err := filepath.Walk("/Users/naoshimi/go/src/github.com/gopherdojo/dojo7/kadai1/naoya7076/convert/cmd", func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".jpg" || filepath.Ext(path) == ".jpeg" {
			fmt.Println(path)
			convert.ConvertToPng(path)
		}
		return nil
	})
	if err != nil {
		fmt.Println("ファイルを開けませんでした")
		os.Exit(1)
	}

}
