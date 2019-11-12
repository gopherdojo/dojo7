package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gopherdojo/dojo7/kadai3/imura81gt/rget"
)

func main() {
	option := rget.Option{
		Concurrency: *flag.Int("c", 10, "concurrency"),
	}
	flag.Parse()
	urls := flag.Args()
	if len(urls) != 1 {
		fmt.Fprintf(os.Stderr, "%s <url>\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "option:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	option.Url = urls[0]
	rget.Run(option)
}
