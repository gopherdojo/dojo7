package main

import (
	"flag"
	"fmt"
	"github.com/gopherdojo/dojo7/kadai1/naoya7076/convert"
	"os"
	"path/filepath"
)

var (
	source = flag.String("s","./","指定したディレクトリ以下を再帰的に捜査します")
)
func main() {
	flag.Parse()
	err := filepath.Walk(*source, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".jpg" || filepath.Ext(path) == ".jpeg" {
			fmt.Println(path)
			convert.ToPng(path)
		}
		return nil
	})
	if err != nil {
		fmt.Println("ファイルを開けませんでした")
		os.Exit(1)
	}

}
