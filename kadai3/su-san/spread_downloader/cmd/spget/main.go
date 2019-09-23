package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strconv"

	errgroup "golang.org/x/sync/errgroup"
)

var COUNT = 4

func main() {

	url := "https://misc.laboradian.com/test/003"

	header, err := HeaderInfo(url)

	if err != nil {
		fmt.Println(err)
	}

	var byteRanges []string
	if canRangeAccess(header) {
		// 分割ダウンロードする
		byte_num, err := strconv.Atoi(header["Content-Length"][0])
		if err != nil {
			fmt.Println("normal download")
			return
		}
		byteRanges = ByteRanges(byte_num, 2)
	} else {
		fmt.Println("normal download")
		// TODO: 分岐として通常ダウンロード処理を入れる（やるかどうかY/nで聞いてから?)
		return
	}

	// TODO: tmpディレクトリ作成

	eg := errgroup.Group{}
	for _, range_suffix := range byteRanges {
		// 分割ダウンロードする
		range_suffix := range_suffix
		eg.Go(func() error {
			if downloadedFile(url, range_suffix) {
				return nil
			}
			return Download(url, range_suffix)
		})
	}

	if err = eg.Wait(); err != nil {
		// TODO:標準エラー出力使う
		fmt.Fprintf(os.Stderr, "ERROR: %v", err)
		return
	}

	_, fileName := path.Split(url)
	fmt.Println(fileName)
	var filePaths []string
	for _, suffix := range byteRanges {
		filePaths = append(filePaths, path.Join("./tmp/", fileName+"_"+suffix))
	}
	// そろっていればtmpファイルを結合する
	err = concatFiles(filePaths, "./"+fileName)

	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v", err)
	}

}

func HeaderInfo(url string) (map[string][]string, error) {
	req, _ := http.NewRequest("HEAD", "https://misc.laboradian.com/test/003/", nil)

	client := new(http.Client)
	resp, err := client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		fmt.Printf("ioutil err: %v", err)
		return nil, err
	}

	// fmt.Println(resp.Header)
	return resp.Header, nil
}

func canRangeAccess(header map[string][]string) bool {
	v, ok := header["Accept-Ranges"]

	if ok {
		for _, val := range v {
			if val == "bytes" {
				if _, ok := header["Content-Length"]; ok {
					return true
				}
			}
		}
	}
	return false
}

func downloadedFile(url, suffix string) bool {
	_, fileName := path.Split(url)
	_, err := os.Stat(path.Join("./tmp/", fileName+"_"+suffix))
	return err == nil
}

func Download(url, byteRange string) error {

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Range", "bytes="+byteRange)

	client := new(http.Client)
	resp, err := client.Do(req)
	// TODO: エラー処理する?
	defer resp.Body.Close()

	_, fileName := path.Split(url)
	fmt.Println(url, byteRange, fileName)

	file, err := os.Create(path.Join("./tmp/", fileName+"_"+byteRange))
	if err != nil {
		return err
	}
	_, err = io.Copy(file, resp.Body)
	if closeErr := file.Close(); err == nil {
		err = closeErr
	}

	return err
}

// ByteRanges は分割したバイト範囲を返す関数です
// ex) 100 で 2分割 0-50, 51-100
func ByteRanges(length int, parallel_num int) []string {
	var byteRanges []string

	for i := 1; i < parallel_num+1; i++ {
		byteRanges = append(byteRanges, NthRange(length, parallel_num, i))
	}
	return byteRanges
}

func NthRange(length, parallel_num, n int) string {
	bytes_per_file := length / parallel_num
	if n == 1 {
		return "0-" + strconv.Itoa(bytes_per_file)
	} else if n < parallel_num {
		return strconv.Itoa(bytes_per_file*(n-1)+1) + "-" + strconv.Itoa(bytes_per_file*n)
	} else {
		return strconv.Itoa(bytes_per_file*(n-1)+1) + "-" + strconv.Itoa(length)
	}
}

func concatFiles(filePaths []string, fileName string) error {

	var writeBytes [][]byte

	for _, path := range filePaths {
		f, err := os.Open(path)
		if err != nil {
			return err
		}

		// 一気に全部読み取り
		readBytes, err := ioutil.ReadAll(f)
		writeBytes = append(writeBytes, readBytes)

		if err := f.Close(); err != nil {
			return err
		}
	}

	emptyByte := []byte{}

	fmt.Println(fileName)

	err := ioutil.WriteFile(fileName, bytes.Join(writeBytes, emptyByte), 0644)

	return err
}
