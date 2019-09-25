package downloader

import (
	"net/http"
	"runtime"
	"strings"

	"fmt"
	"io"
	"os"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

// Downloader struct
type Downloader struct {
	FileName string
	URL      string
	Timeout  int
	Unit     int
	Ranges   []*Range
}

type Range struct {
	start    int
	end      int
	url      string
	filePath string
}

func New(url string) *Downloader {
	return &Downloader{
		URL:     url,
		Timeout: 10,
		Unit:    runtime.NumCPU(),
	}
}

func (d *Downloader) Run() (*Downloader, error) {

	d.FileName = getFileName(d.URL)
	res, err := http.Head(d.URL)
	if err != nil {
		return nil, err
	}
	if res.Header.Get("Accept-Ranges") != "bytes" {
		return &Downloader{Ranges: nil}, nil
	}
	size := res.ContentLength
	split := size / int64(d.Unit)

	for i := 0; i < d.Unit; i++ {
		start := i * int(split)
		end := start + d.Unit - 1
		partFilePath := fmt.Sprintf("%s.%d", d.FileName, i)
		r, err := NewRange(d, start, end, partFilePath)
		if err != nil {
			return nil, errors.Wrap(err, "initialize Range error")
		}
		d.Ranges = append(d.Ranges, r)
	}

	return d, nil
}

func NewRange(d *Downloader, start int, end int, filePath string) (*Range, error) {
	return &Range{
		start:    start,
		end:      end,
		url:      d.URL,
		filePath: filePath,
	}, nil
}

func (d *Downloader) Download() error {
	eg := errgroup.Group{}

	for i, ra := range d.Ranges {
		singleRange := ra
		fmt.Println(i)
		fmt.Println(singleRange.filePath)
		eg.Go(func() error {
			return singleRange.SplitDownload()
		})
	}

	if err := eg.Wait(); err != nil {
		return err
	}
	return nil

}

func (r *Range) SplitDownload() error {
	fmt.Println(r.filePath)

	req, err := http.NewRequest("GET", r.url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("RANGE", fmt.Sprintf("bytes=%d-%d", r.start, r.end))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	out, err := os.Create(r.filePath)
	if err != nil {
		return nil
	}
	defer out.Close()
	io.Copy(out, res.Body)
	return nil
}

func (d *Downloader) Merge() error {

	resultFile, err := os.Create(d.FileName)
	if err != nil {
		return err
	}
	fmt.Println(&resultFile)

	defer resultFile.Close()
	for i := 0; i < d.Unit; i++ {
		fmt.Println(i)
		filePath := fmt.Sprintf("%s.%d", d.FileName, i)
		file, err := os.Open(filePath)
		if err != nil {
			return err
		}
		io.Copy(resultFile, file)
		file.Close()

		if err := os.Remove(filePath); err != nil {
			return errors.Wrap(err, "failed to remove a file in download location")
		}
	}

	// if err := os.RemoveAll(d.FileName); err != nil {
	// 	return errors.Wrap(err, "failed to remove download location")
	// }

	fmt.Println("Complete")

	return nil

}
func getFileName(url string) string {
	token := strings.Split(url, "/")
	filename := token[len(token)-1]
	return filename
}
