package main

import (
	"flag"
	"log"

	"github.com/gopherdojo/dojo7/asuke-yasukuni/replacer"
	"github.com/gopherdojo/dojo7/asuke-yasukuni/validation"
	"github.com/gopherdojo/dojo7/asuke-yasukuni/walk"
)

var src = flag.String("src", "", "ファイルパス書いて")
var from = flag.String("from", "jpg", "変換したい画像の拡張子 jpg or png")
var to = flag.String("to", "png", "変換後の拡張子 jpg or png")

func main() {
	flag.Parse()

	// do ext validation
	if !validation.Ext(*from) || !validation.Ext(*to) {
		log.Fatalf("\x1b[31mfrom:%s to:%s encode is unsupported\x1b[0m\n", *from, *to)
	}

	log.Printf("\x1b[33m%s\x1b[0m\n", "[replace start]")

	walker := walk.Walk{File: &replacer.File{}}
	files, err := walker.Encoder(src, *from, *to)
	if err != nil {
		log.Fatal(err)
	}

	// encoding result
	for _, f := range files {
		log.Print(f)
	}

	log.Printf("\x1b[33m%s\x1b[0m\n", "[replace end]")
}
