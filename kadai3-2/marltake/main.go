package main

import (
	"flag"
	"marltake/pget"
)

func main() {
	var url string
	var cpus int
	flag.StringVar(&url, "url", "", "download target")
	flag.IntVar(&cpus, "p", 1, "number of parallel downloading")
	flag.Parse()
	if url == "" {
		return
	}
	pget.Download(url, cpus)
}
