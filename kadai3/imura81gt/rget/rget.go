package rget

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

	"golang.org/x/sync/errgroup"
)

type Option struct {
	Concurrency   int
	URL           string
	OutputDir     string
	ContentLength int64
	Units         Units
}

type Unit struct {
	RangeStart   int64
	RangeEnd     int64
	TempFileName string
}

type Units []Unit

func Run(option Option) {
	fmt.Printf("%+v\n", option)
	err := option.contentLength()
	if err != nil {
		fmt.Errorf("%s", err)
	}

	option.divide()

	err = option.parallelDownload()
	if err != nil {
		fmt.Errorf("%s", err)
	}

}

func (o *Option) contentLength() error {
	//resp, err := http.Head(url)
	resp, err := http.Head(o.URL)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		//return 0, err
		return err
	}

	if resp.Header.Get("Accept-Ranges") == "" {
		err := fmt.Errorf("%s URL cannot support Ranges Requests", o.URL)
		// fmt.Fprintln(os.Stderr, err)
		//return resp.ContentLength, err
		return err
	}
	if resp.Header["Accept-Ranges"][0] == "none" {
		err := fmt.Errorf("%s cannot support Ranges Requests", o.URL)
		// fmt.Fprintln(os.Stderr, err)
		//return resp.ContentLength, err
		return err
	}
	if resp.ContentLength == 0 {
		err := fmt.Errorf("%s size is %s", o.URL, resp.Header["Content-Length"][0])
		// fmt.Fprintln(os.Stderr, err)
		//return resp.ContentLength, err
		return err
	}

	o.ContentLength = resp.ContentLength
	//return resp.ContentLength, nil
	return err
}

//func divide(contentLength int64, concurrency int) Units {
func (o *Option) divide() {
	var units []Unit

	//sbyte := contentLength / int64(concurrency)
	sbyte := o.ContentLength / int64(o.Concurrency)

	//	for i := 0; i < concurrency; i++ {
	for i := 0; i < o.Concurrency; i++ {
		units = append(units, Unit{
			RangeStart:   int64(i) * sbyte,
			RangeEnd:     int64((i+1))*sbyte - 1,
			TempFileName: fmt.Sprintf("%d_%s", i, path.Base(o.URL)),
		})
	}

	o.Units = units
	//return units
}

// func download(units Units) {
// 	filepath.Split()
// 	fmt.Println(units)
// }

func (o *Option) parallelDownload() error {
	fmt.Println("parallelDownload", o.Units)

	eg, ctx := errgroup.WithContext(context.Background())
	for i := range o.Units {
		// https://godoc.org/golang.org/x/sync/errgroup#example-Group--Parallel
		// https://golang.org/doc/faq#closures_and_goroutines
		i := i
		eg.Go(func() error {
			return o.downloadWithContext(ctx, i)
		})
	}

	if err := eg.Wait(); err != nil {
		o.clearCache()
		return err
	}

	return nil
}

func (o *Option) downloadWithContext(ctx context.Context, i int) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	fmt.Printf("Downloading: %v %+v\n", i, o.Units[i])

	//v1.13
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, o.URL, nil)
	if err != nil {
		return fmt.Errorf("Error: %v", err)
	}

	// add range header
	fmt.Printf(fmt.Sprintf("bytes=%d-%d", o.Units[i].RangeStart, o.Units[i].RangeEnd))
	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", o.Units[i].RangeStart, o.Units[i].RangeEnd))

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return fmt.Errorf("Error: %v", err)
	}

	select {
	case <-ctx.Done():
		fmt.Printf("Done: %v %+v\n", i, o.Units[i])
		return nil
	default:
		fmt.Println("default:", i, o.Units[i])
		//return fmt.Errorf("Error: %v %+v", i, o.Units[i])
	}

	w, err := os.Create(o.Units[i].TempFileName)
	if err != nil {
		return fmt.Errorf("Error: %v", err)
	}
	defer func() error {
		if err := w.Close(); err != nil {
			return fmt.Errorf("Error: %v", err)
		}
		return nil
	}()

	_, err = io.Copy(w, resp.Body)
	if err != nil {
		return fmt.Errorf("Error: %v", err)
	}

	return nil
}

func (o *Option) conbine() error {
	return nil
}

func (o *Option) clearCache() error {
	//TODO: remove temporary files
	return nil
}
