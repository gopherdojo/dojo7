package main

import (
	"flag"
	"fmt"
	"log"

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
	if len(args) < 1 {
		err := fmt.Errorf("no target")
		return err
	}
	c, err := imachan.NewConfig(args[0], "jpeg", "png")
	if err != nil {
		return err
	}
	err = c.ConvertRec()
	if err != nil {
		return err
	}
	return nil
}
