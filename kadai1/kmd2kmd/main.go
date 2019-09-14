package main

import (
	"flag"
	"fmt"

	"log"
	"os"
	"path/filepath"

	"kadai1/iconv"
)

var (
	in  = flag.String("in", "jpg", "jpeg: -in jpg\npng: -in png\n")
	out = flag.String("out", "png", "jpeg: -out jpg\npng: -out png\n")
)

// 再帰的に指定された拡張子のファイルを検索し返却する
func getFiles(root string, ext string) ([]string, error) {
	paths := make([]string, 0, 0)
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ext {
			paths = append(paths, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return paths, nil
}

func main() {
	flag.Parse()
	flags := flag.Args()
	if len(flags) < 1 {
		os.Exit(1)
	}

	fmt.Println("convert", *in, "to", *out, "...")

	dir := flags[0]

	paths, err := getFiles(dir, "."+*in)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range paths {
		c := iconv.Image{Path: v, In: *in, Out: *out}
		err = c.Do()
		if err != nil {
			log.Fatal(err)
		}
	}
	os.Exit(0)
}
