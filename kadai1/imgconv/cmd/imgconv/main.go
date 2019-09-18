package main

import (
	"flag"
	"log"
	"os"

	"github.com/gopherdojo/dojo7/kadai1/sabe/pkg/imgconv"
)

func main() {
	if err := exec(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}

func exec() error {
	var (
		dir     string
		fromExt string
		toExt   string
	)
	flag.StringVar(&dir, "d", ".", "target directory flag")
	flag.StringVar(&fromExt, "f", "jpg", "before convert ext flag")
	flag.StringVar(&toExt, "t", "png", "after convert ext flag")
	flag.Parse()
	imageConfList, err := imgconv.NewImageConfList(dir, fromExt, toExt)
	if err != nil {
		return err
	}
	return imageConfList.Convert()
}
