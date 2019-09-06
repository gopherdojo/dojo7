package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gopherdojo/dojo7/kadai1/okamotoke/imgconvt"
)

func main() {

	var (
		from = flag.String("from", "jpg", "specify image extension from jpg, png or gif that is converted from")
		to   = flag.String("to", "png", "specify image extension from jpg, png or gif that is converted to")
	)

	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Fprintf(os.Stderr, "please specify the directory")
		return
	}
	path := flag.Arg(0)

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == addDot(*from) {
			c := &imgconvt.Conv{path, addDot(*from), addDot(*to)}
			err = imgconvt.ConvertImage(c)
			fmt.Println(path)
			return err
		}
		return nil
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to walk filepath: %v", err)
	}

}

func addDot(s string) string {
	if strings.HasPrefix(s, ".") {
		return s
	}
	return "." + s
}
