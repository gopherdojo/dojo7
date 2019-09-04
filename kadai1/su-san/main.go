package main

import (
	"flag"
	"fmt"
	"io/ioutil"
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

	// ディレクトリがなければ知らせて終了
	if _, err := os.Stat(targetDir); err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	// 変換対象のフォーマットと変換フォーマットが同じなら何もせず終了
	if *inputExt == *outputExt {
		return
	}

	convExts := image.NewConvExts(*inputExt, *outputExt)
	if convExts.SupportedFormats() == false {
		fmt.Println("unsupported format! please specify these format [png jpg jpeg gif]")
		return
	}
	targetFiles := dirwalk(targetDir)

	for _, f := range targetFiles {

		err := image.FmtConv(f, convExts)
		if err != nil {
			fmt.Println(f, err)
		}
	}
}

// dirwalk は与えられたディレクト以下のファイルパスをリストで返します
func dirwalk(dir string) []string {

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, dirwalk(filepath.Join(dir, file.Name()))...)
			continue
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
	}
	return paths
}
