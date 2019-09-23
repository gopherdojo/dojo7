package pdl

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

// Downloader ダウンローダ
type Downloader struct {
	URL           string
	Name          string
	DivNum        int
	ContentLength int64
	ProcessList   *[]downloadProcess
}

type downloadProcess struct {
	no                   int
	bytesToStartReading  int64
	bytesToFinishReading int64
	url                  string
	partFilePath         string
}

// NewDownloader ダウンローダを生成
func NewDownloader(url string, divNum int) (*Downloader, error) {

	req, err := http.Head(url)
	if err != nil {
		return &Downloader{}, err
	}

	if acceptRanges := req.Header.Get("Accept-Ranges"); acceptRanges != "bytes" {
		return &Downloader{}, errors.New("Accept-Rangesが不許可")
	}

	contentLength, err := strconv.ParseInt(req.Header.Get("Content-Length"), 10, 0)
	if err != nil {
		return &Downloader{}, errors.Wrap(err, "Content-Lengthのパースに失敗")
	}

	name := filepath.Base(url)

	dpl := createDownloadProcessList(contentLength, divNum, url, name)

	return &Downloader{
		URL:           url,
		Name:          name,
		DivNum:        divNum,
		ContentLength: contentLength,
		ProcessList:   dpl,
	}, nil
}

// PararellDownload 分割ダウンロード実行
func (d *Downloader) PararellDownload() error {

	eg := errgroup.Group{}
	for _, process := range *d.ProcessList {
		p := process
		eg.Go(func() error {
			return p.download()
		})
	}

	if err := eg.Wait(); err != nil {
		return errors.Wrap(err, "パラレルダウンロード中に失敗")
	}
	return nil
}

func createDownloadProcessList(contentLength int64, divNum int, url, name string) *[]downloadProcess {

	dpl := []downloadProcess{}

	split := contentLength / int64(divNum)

	for i := 0; i < divNum; i++ {

		bytesToStartReading := split * int64(i)
		bytesToFinishReading := bytesToStartReading + split - 1
		partFilePath := fmt.Sprintf("%s_%d", name, i)

		if i == divNum-1 {
			bytesToFinishReading = contentLength
		}

		dp := downloadProcess{
			no:                   i,
			bytesToStartReading:  bytesToStartReading,
			bytesToFinishReading: bytesToFinishReading,
			url:                  url,
			partFilePath:         partFilePath,
		}

		dpl = append(dpl, dp)

	}

	return &dpl
}

func (dp *downloadProcess) download() error {

	fmt.Printf("process %v start\n", dp.no)

	request, err := http.NewRequest("GET", dp.url, nil)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("リクエスト生成に失敗:%v", dp.no))
	}
	request.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", dp.bytesToStartReading, dp.bytesToFinishReading))

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("リクエストに失敗:%v", dp.no))
	}
	defer response.Body.Close()

	output, err := os.Create(dp.partFilePath)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("ファイル生成に失敗:%s", dp.partFilePath))
	}
	defer output.Close()

	io.Copy(output, response.Body)

	fmt.Printf("process %v finish\n", dp.no)

	return nil
}

// Merge ファイルをマージする
func (d *Downloader) Merge() error {
	outputFilePath := fmt.Sprintf("%s", d.Name)
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		return errors.Wrap(err, "マージファイルの生成に失敗")
	}
	defer outputFile.Close()

	for _, v := range *d.ProcessList {

		partFile, err := os.Open(v.partFilePath)
		if err != nil {
			return errors.Wrap(err, "分割ファイルのオープンに失敗")
		}
		io.Copy(outputFile, partFile)
		partFile.Close()
		if err := os.Remove(v.partFilePath); err != nil {
			return errors.Wrap(err, "分割ファイルの除外に失敗")
		}
	}

	return nil
}
