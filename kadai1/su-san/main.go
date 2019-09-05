package main

import (
	"flag"
	"fmt"
	//"io/ioutil"
	"os"
	"path/filepath"

	"github.com/gopherdojo/dojo7/kadai1/su-san/image"
)

func main() {

	inputExt := flag.String("i", "jpg", " extension to be converted ")
	outputExt := flag.String("o", "png", " extension after conversion")

	// Usageメッセージ
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "usage : cmd [-i] [-o] target_dir")
		flag.PrintDefaults()
	}

	flag.Parse()
	args := flag.Args()

	if flag.NArg() == 0 {
		flag.Usage()
		return
	}

	targetDir := args[0]

	// 変換対象のフォーマットと変換フォーマットが同じなら何もせず終了
	if *inputExt == *outputExt {
		return
	}

	convExts := image.NewConvExts(*inputExt, *outputExt)
	if convExts.SupportedFormats() == false {
		fmt.Println("unsupported format! please specify these format [png jpg jpeg gif]")
		return
	}

	targetFiles := []string{}
	err := filepath.Walk(targetDir, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ("." + *inputExt) {
			targetFiles = append(targetFiles, path)
		}
		return nil
	})

	if err != nil {
		fmt.Println("file open error")
		return
	}

	for _, f := range targetFiles {

		err := image.FmtConv(f, convExts)
		if err != nil {
			fmt.Println(f, err)
		}
	}
}
