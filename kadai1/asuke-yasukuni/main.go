package main

import (
	"asuke-yasukuni/command"
	"asuke-yasukuni/validation"
	"flag"
	"log"
	"os"
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

	files,err := command.WalkEncoder(src, fromExt, toExt)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	for _,f := range files {
		log.Print(f)
	}
	
	log.Printf("\x1b[33m%s\x1b[0m\n", "[replace end]")
}
