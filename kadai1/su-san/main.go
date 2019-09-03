package main

import (
	"flag"
	"fmt"

	"github.com/gopherdojo/dojo7/kadai1/su-san/image"
)

// オプション用の返還前後の拡張子
// var inputExt = flag.String("i", "jpg", " extension to be converted ")
// var outputExt = flag.String("o", "png", " extension after conversion")

func main() {


	fmt.Println("test")
	convExts := image.NewConvExts("", "")
	err := image.FmtConv("test_img/test_img.jpg", convExts)
	if err != nil {
		fmt.Println(err)
	}
}
