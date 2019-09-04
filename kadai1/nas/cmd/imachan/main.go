package main

import (
	"flag"
	"log"
	"path/filepath"

	"github.com/gopherdojo/dojo7/kadai1/nas/pkg/imachan"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	flag.Parse()
	args := flag.Args()
	for _, arg := range args {
		path, err := filepath.Abs(arg)
		if err != nil {
			return err
		}
		err = imachan.Convert(path, "jpg", "png")
		if err != nil {
			return err
		}
	}
	return nil
}
