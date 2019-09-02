package main

import (
	"fmt"
	"os"
	"image/png"
	"image/jpeg"
)

func main(){

	f, err := os.Open("test_img/test_img.png")

	if err != nil {
		fmt.Println("error")
	}
	defer f.Close()

	fmt.Printf("%T\n", f)

	img, err := png.Decode(f)
	if err != nil {
		fmt.Println("cannot decode")
	}

	fmt.Println(img.Bounds())
	fmt.Printf("%T\n", img.Bounds())
	fmt.Printf("%T\n", img)

	outputFile, err := os.Create("output/output.jpeg")

	jpeg.Encode(outputFile, img, nil)
}
