package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gopherdojo/dojo7/kadai1/nas/pkg/imachan"
)

// exit codes
const (
	ExitCodeOK  = 0
	ExitCodeErr = 10
)

func main() {
	err := run()
	if err != nil {
		os.Exit(ExitCodeErr)
		return
	}
	os.Exit(ExitCodeOK)
}

func run() error {
	var (
		dir        string
		fromFormat string
		toFormat   string
	)

	flag.StringVar(&dir, "d", ".", "target directory")
	flag.StringVar(&fromFormat, "f", "jpg", "converted from")
	flag.StringVar(&toFormat, "t", "png", "converted to")
	flag.Parse()

	path, err := filepath.Abs(dir)
	if err != nil {
		return err
	}
	c, err := imachan.NewConfig(path, fromFormat, toFormat)
	if err != nil {
		return err
	}
	data, err := c.ConvertRec()
	if err != nil {
		return err
	}
	for _, dd := range data {
		fmt.Printf("success : %s -> %s\n", dd["from"], dd["to"])
	}
	return nil
}
