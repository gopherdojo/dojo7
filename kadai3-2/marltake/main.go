package main

import (
	"flag"
	"marltake/pget"
)

func main() {
	var url string
	var cpus int
	flag.StringVar(&url, "url", "", "download target")
	flag.IntVar(&cpus, "p", 4, "number of parallel downloading")
	flag.Parse()
	url = "http://ftp.riken.jp/Linux/ubuntu-releases/bionic/ubuntu-18.04.3-live-server-amd64.iso"
	if url == "" {
		return
	}
	pget.Download(url, cpus)
}
