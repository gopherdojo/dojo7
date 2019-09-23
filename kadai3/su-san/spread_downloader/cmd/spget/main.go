package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"
)

var COUNT = 4

func main() {
	//req, _ := http.NewRequest("GET", "https://misc.laboradian.com/test/003/", nil)

	url := "https://misc.laboradian.com/test/003/"

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

	for _, range_suffix := range byteRanges {
		// 分割ダウンロードする
		Download(url, range_suffix)
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

func Download(url, byteRanges string) error {
	//fmt.Println(url, byteRanges)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Range", "bytes="+byteRanges)

	client := new(http.Client)
	resp, err := client.Do(req)
	// TODO: エラー処理する?
	defer resp.Body.Close()

	_, fileName := path.Split(url)

	file, err := os.Create(path.Join("./tmp/", fileName, byteRanges))
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
