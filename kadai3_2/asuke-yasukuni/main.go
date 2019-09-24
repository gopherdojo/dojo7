package main

import (
	"flag"
	"log"
	"os"

	"github.com/gopherdojo/dojo7/kadai3_2/asuke-yasukuni/muget"
)

var get = flag.String("get", "", "ダウンロードパス")
var out = flag.String("out", "", "ファイルパス")

func main() {
	flag.Parse()

	log.Println("Download Start")
	log.Println(*get)
	log.Println(*out)

	if err := muget.Run(*get, *out); err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}
