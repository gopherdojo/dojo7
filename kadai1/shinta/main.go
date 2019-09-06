package main

import (
	"fmt"
	"os"

	"github.com/gopherdojo/dojo7/kadai1/shinta/imageconversion/imageconversion"
)

func main() {
	err := imageconversion.Excute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
