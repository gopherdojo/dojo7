package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gopherdojo/dojo7/kadai3/imura81gt/rget"
)

func main() {
	concurrency := flag.Uint("c", 2, "concurrency")
	outputDir := flag.String("o", "./", "output directory")

	flag.Parse()
	option := rget.Option{
		Concurrency: *concurrency,
		OutputDir:   *outputDir,
	}
	urls := flag.Args()
	if len(urls) != 1 {
		fmt.Fprintf(os.Stderr, "%s <url>\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "option:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	option.URL = urls[0]
	fmt.Println(option)
	err := rget.Run(option)
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %s", err)
		os.Exit(1)
	}
}
