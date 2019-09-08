// This command recursively converts images.

package main

import (
	"asuke-yasukuni/replacer"
	"asuke-yasukuni/validation"
	"flag"
	"log"
	"os"
	"path/filepath"
)

var src = flag.String("src", "", "ファイルパス書いて")
var from = flag.String("from", "jpg", "変換したい画像の拡張子 jpg or png")
var to = flag.String("to", "png", "変換後の拡張子 jpg or png")

func main() {
	flag.Parse()

	fromExt := *from
	toExt := *to

	if !validation.Ext(fromExt) || !validation.Ext(toExt) {
		log.Fatalf("\x1b[31mfrom:%s to:%s encode is unsupported\x1b[0m\n", fromExt, toExt)
	}

	log.Printf("\x1b[33m%s\x1b[0m\n", "[replace start]")

	var file replacer.File
	err := filepath.Walk(*src, func(path string, info os.FileInfo, err error) error {

		if filepath.Ext(path) != "."+fromExt {
			return nil
		}

		log.Printf("\x1b[33m%s%s -> %s\x1b[0m\n", "[replace file]", path, toExt)

		file = replacer.File{
			Path:    path,
			FromExt: fromExt,
			ToExt:   toExt,
		}

		if err := file.Encode(); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("\x1b[33m%s\x1b[0m\n", "[replace end]")

	os.Exit(1)
}
