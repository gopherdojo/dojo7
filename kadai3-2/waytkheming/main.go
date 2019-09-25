package main

import (
	"fmt"
	"os"

	"github.com/waytkheming/godojo/dojo7/kadai3-2/waytkheming/downloader"
)

func main() {

	url := os.Args[1]

	fmt.Println(url)

	d := downloader.New(url)
	d, err := d.Run()
	if err != nil {
		fmt.Printf("\ndownload initialization error. %v", err)
		os.Exit(1)
	}
	err = d.Download()
	if err != nil {
		fmt.Printf("\ndownload error. %v", err)
		os.Exit(1)
	}
	err = d.Merge()
	if err != nil {
		fmt.Printf("\nmerge error. %v", err)
		os.Exit(1)
	}
}
