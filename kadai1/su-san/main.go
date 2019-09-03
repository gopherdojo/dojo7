package main

import (
	"fmt"

	"github.com/gopherdojo/dojo7/kadai1/su-san/image"
)

func main(){
	fmt.Println("test")
	convExts := image.NewConvExts("", "")
	err := image.FmtConv("test_img/test_img.jpg", convExts)
	if err != nil {
		fmt.Println(err)
	}
}
