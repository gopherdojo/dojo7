package muget

import (
	"context"
	"fmt"
	"log"
	"path/filepath"
)

type Range struct {
	Start int
	End   int
}

func Run(url, outPutPath string) error {
	size, err := CheckRanges(context.Background(), url)
	if err != nil {
		return err
	}

	var (
		start, end int
		ranges     []Range
	)

	for start <= size {
		end = start + (size / 10)
		ranges = append(ranges, Range{
			Start: start,
			End:   end,
		})
		start = end
	}

	ch := make(chan int)
	for i, r := range ranges {
		go func(r Range, count int) {
			if err := DownloadFile(url, outPutPath, r.Start, r.End, count); err != nil {
				log.Fatal(err)
			}
			ch <- i
		}(r, i)
	}

	//ダウンロード完了まで待つ
	var count int
D:
	for {
		select {
		case <-ch:
			count++
			if len(ranges) <= count {
				close(ch)
				break D
			}
		}
	}

	fmt.Println("\nbinding with files...")

	//ダウンロードファイルをマージ
	if err := MergeFiles(count, filepath.Base(url), filepath.Ext(url)); err != nil {
		return err
	}

	return nil
}
