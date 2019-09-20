package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"

	"marltake/convert"
)

func main() {
	var target = flag.String("t", "jpg,png", "source and destination picture type")
	flag.Parse()
	src, dest, ok := convert.ParseTarget(*target)
	if !ok {
		log.Fatal(fmt.Errorf("Invalid target option %s", *target))
	}
	err := filepath.Walk(flag.Arg(0), convert.ConfigConvert(src, dest))
	if err != nil {
		log.Fatal(err)
	}
}
