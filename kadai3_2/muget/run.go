package muget

import (
	"context"
	"fmt"
	"log"
)

type FileData struct {
	Name     string
	Size     uint
	Dirname  string
	FullName string
}

type Range struct {
	Start int
	End int
}

func Run(url, outPutPath string) error {
	size, err := CheckRanges(context.Background(), url)
	if err != nil {
		return err
	}

	start := 0
	end := 0
	var ranges []Range
	for start <= size {
		end = start+(size/10)
		ranges = append(ranges, Range{
			Start: start,
			End:   end,
		})
		start = end
	}

	ch := make(chan int)
	for i,r := range ranges {
		go func(r Range, count int) {
			fmt.Println(r.Start,"->",r.End)
			if err := DownloadFile(url, outPutPath, r.Start, r.End, count); err != nil {
				log.Fatal(err)
			}
			ch <- i
		}(r,i)
	}

	//ダウンロード完了まで待つ
	var count int
D:
	for {
		select {
		case <-ch:
			count++
			fmt.Println(count)
			if len(ranges) <= count {
				close(ch)
				break D
			}
		}
	}

	return nil
}
