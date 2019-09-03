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

	// Usage文言
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, `Usage :
   convimg target_directory`)
		flag.PrintDefaults()
	}

	flag.Parse()
	args := flag.Args()

	if flag.NArg() == 0 {
		flag.Usage()
		return
	}

	targetDir := args[0]
	targetFiles := dirwalk(targetDir)

	for _, f := range targetFiles {
		convExts := image.NewConvExts("", "")
		err := image.FmtConv(f, convExts)
		if err != nil {
			fmt.Println(err)
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
