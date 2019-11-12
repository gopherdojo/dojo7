package rget

import (
	"fmt"
	"net/http"
	"os"
)

type Option struct {
	Concurrency int
	Url         string
}

type Unit struct {
	RangeStart int64
	RangeEnd   int64
}

type Units []Unit

func Run(option Option) {
	fmt.Printf("%+v\n", option)
	cl, err := contentLength(option.Url)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	units := divide(cl, option.Concurrency)

	//TODO: check errors
	download(units)

}

func contentLength(url string) (int64, error) {
	resp, err := http.Head(url)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 0, err
	}

	if resp.Header.Get("Accept-Ranges") == "" {
		err := fmt.Errorf("This URL cannot support Ranges Requests")
		// fmt.Fprintln(os.Stderr, err)
		return resp.ContentLength, err
	}
	if resp.Header["Accept-Ranges"][0] == "none" {
		err := fmt.Errorf("This URL cannot support Ranges Requests")
		// fmt.Fprintln(os.Stderr, err)
		return resp.ContentLength, err
	}
	if resp.ContentLength == 0 {
		err := fmt.Errorf("This URL size is %i", resp.Header["Content-Length"][0])
		// fmt.Fprintln(os.Stderr, err)
		return resp.ContentLength, err
	}

	return resp.ContentLength, nil
}

func divide(contentLength int64, concurrency int) Units {
	var units []Unit

	sbyte := contentLength / int64(concurrency)
	for i := 0; i < concurrency; i++ {
		units = append(units, Unit{
			RangeStart: int64(i) * sbyte,
			RangeEnd:   int64((i+1))*sbyte - 1,
		})
	}

	return units
}

func download(units Units) {
	fmt.Println(units)
}
