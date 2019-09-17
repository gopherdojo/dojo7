package main

import (
	"flag"
	"fmt"
	"myimage"
	"os"
	"path/filepath"
	"strings"
)

var (
	dir     = flag.String("dir", ".", "Target directory")
	fromExt = flag.String("from-ext", "", "From extention")
	toExt   = flag.String("to-ext", "", "To extention")
)

func main() {
	flag.Parse()
	startPath := filepath.Join(*dir)

	fmt.Printf("Convert image file start.: target=[%s], from-extension=[%s], to-extension=[%s]\n", *dir, *fromExt, *toExt)

	err := filepath.Walk(startPath,
		func(path string, info os.FileInfo, err error) error {
			// オプションで指定された拡張子のファイルの場合のみ
			if strings.Replace(filepath.Ext(path), ".", "", 1) == *fromExt {
				convfunc, err := myimage.GetConvertFunc(*fromExt, *toExt)
				if err == nil {
					fmt.Printf("Convert image file: filepath=[%s]\n", path)
					convfunc(path)
				} else {
					fmt.Printf("Skip convert image file: filepath=[%s], reason=[%s]\n", path, err)
				}
			}
			return err
		})
	if err != nil {
		fmt.Printf("Error occured: error=[%s]\n", err)
	}
	fmt.Printf("Convert image file end.\n")
}
