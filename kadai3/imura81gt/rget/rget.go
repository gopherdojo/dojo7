package rget

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"

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

	tmpDir, err := ioutil.TempDir("", "rget")
	if err != nil {
		fmt.Errorf("%s", err)
	}
	defer os.RemoveAll(tmpDir)
	fmt.Println(tmpDir)

	err = option.parallelDownload(tmpDir)
	if err != nil {
		fmt.Errorf("%s", err)
	}

	err = option.combine(tmpDir)
	if err != nil {
		fmt.Errorf("%s", err)
	}

}

func (o *Option) contentLength() error {
	//resp, err := http.Head(url)
	resp, err := http.Head(o.URL)
	if err != nil {
		return err
	}

	if resp.Header.Get("Accept-Ranges") == "" {
		err := fmt.Errorf("%s URL cannot support Ranges Requests", o.URL)
		return err
	}
	if resp.Header["Accept-Ranges"][0] == "none" {
		err := fmt.Errorf("%s cannot support Ranges Requests", o.URL)
		return err
	}
	if resp.ContentLength == 0 {
		err := fmt.Errorf("%s size is %s", o.URL, resp.Header["Content-Length"][0])
		return err
	}

	o.ContentLength = resp.ContentLength
	//return resp.ContentLength, nil
	return err
}

//func divide(contentLength int64, concurrency int) Units {
func (o *Option) divide() {
	var units []Unit

	//TODO: if o.ContentLength < int64(o.Concurrency)
	sbyte := o.ContentLength / int64(o.Concurrency)

	for i := 0; i < o.Concurrency; i++ {
		units = append(units, Unit{
			RangeStart:   int64(i) * sbyte,
			RangeEnd:     int64((i+1))*sbyte - 1,
			TempFileName: fmt.Sprintf("%d_%s", i, path.Base(o.URL)),
		})
	}

	// TODO: should distribute the remainder to each unit
	units[len(units)-1].RangeEnd = o.ContentLength

	o.Units = units
}

func (o *Option) parallelDownload(tmpDir string) error {
	fmt.Println("parallelDownload", o.Units)

	eg, ctx := errgroup.WithContext(context.Background())
	for i := range o.Units {
		// https://godoc.org/golang.org/x/sync/errgroup#example-Group--Parallel
		// https://golang.org/doc/faq#closures_and_goroutines
		i := i
		eg.Go(func() error {
			return o.downloadWithContext(ctx, i, tmpDir)
		})
	}

	if err := eg.Wait(); err != nil {
		return err
	}

	return nil
}

func (o *Option) downloadWithContext(
	ctx context.Context,
	i int,
	dir string,
) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	fmt.Printf("Downloading: %v %+v\n", i, o.Units[i])

	//v1.13
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, o.URL, nil)
	if err != nil {
		return fmt.Errorf("Error: %v", err)
	}

	// add range header
	fmt.Printf(fmt.Sprintf("bytes=%d-%d\n", o.Units[i].RangeStart, o.Units[i].RangeEnd))
	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", o.Units[i].RangeStart, o.Units[i].RangeEnd))

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("client err: %s", err)
		return fmt.Errorf("Error: %v", err)
	}
	defer resp.Body.Close()

	select {
	case <-ctx.Done():
		fmt.Printf("Done: %v %+v\n", i, o.Units[i])
		return nil
	default:
		fmt.Println("default:", i, o.Units[i])
	}

	w, err := os.Create(filepath.Join(dir, o.Units[i].TempFileName))
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

func (o *Option) combine(dir string) error {
	w, err := os.Create(path.Base(o.URL))
	if err != nil {
		return fmt.Errorf("Error: %v", err)
	}
	defer func() error {
		if err := w.Close(); err != nil {
			return fmt.Errorf("Error: %v", err)
		}
		return nil
	}()

	for _, unit := range o.Units {
		r, err := os.Open(filepath.Join(dir, unit.TempFileName))
		if err != nil {
			return fmt.Errorf("Error: %v", err)
		}

		_, err = io.Copy(w, r)
		if err != nil {
			return fmt.Errorf("Error: %v", err)
		}
	}

	return nil
}
