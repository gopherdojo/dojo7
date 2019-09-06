package main

import (
	"flag"
	"fmt"
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
	if len(args) < 1 {
		err := fmt.Errorf("no target")
		return err
	}
	path, err := filepath.Abs(args[0])
	if err != nil {
		return err
	}
	c, err := imachan.NewConfig(path, "jpeg", "png")
	if err != nil {
		return err
	}
	data, err := c.ConvertRec()
	if err != nil {
		return err
	}
	for _, dd := range data {
		fmt.Printf("success : %s -> %s\n", dd["to"], dd["from"])
	}
	return nil
}
