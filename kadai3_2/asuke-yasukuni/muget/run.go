package muget

import (
	"context"
	"fmt"
	"path/filepath"

	"golang.org/x/sync/errgroup"
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

	// TODO: contextとかつかってうまくキャンセルしてあげる
	eg := errgroup.Group{}
	for i, r := range ranges {
		i := i
		r := r
		eg.Go(func() error {
			return DownloadFile(url, outPutPath, r.Start, r.End, i)
		})
	}

	//ダウンロード完了まで待つ
	if err := eg.Wait(); err != nil {
		return err
	}

	fmt.Println("\nbinding with files...")

	//ダウンロードファイルをマージ
	if err := MergeFiles(len(ranges), filepath.Base(url), filepath.Ext(url)); err != nil {
		return err
	}

	return nil
}
