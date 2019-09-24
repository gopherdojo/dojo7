package pget

import (
	"io/ioutil"
	"log"
	"net/http"
	"path"
)

func Download(url string, cpus int) {
	saveFile := path.Base(url)
	res, err := http.DefaultClient.Head(url)
	if err != nil {
		log.Fatal(err)
	}
	if res.Header.Get("Accept-Ranges") != "bytes" {
		log.Fatal("Target url don't accept Ranges.")
	}
	tempfile, err := ioutil.TempFile("", saveFile)
	if err != nil {
		log.Fatal(err)
	}
	tempfile.Close()
	println(tempfile.Name())
	println(res.ContentLength)
	println(cpus)
}
