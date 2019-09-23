package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gopherdojo/dojo7/kadai3-2/ayatothos/pdl"
)

var divNum = flag.Int("divNum", 10, "分割数")

func main() {

	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Println("【ERROR】 ダウンロードファイルパスを指定してください")
		os.Exit(1)
	}

	fmt.Println("downloader create start")

	// ダウンローダを生成 accept-rangesも確認
	d, err := pdl.NewDownloader(flag.Arg(0), *divNum)
	if err != nil {
		fmt.Printf("【ERROR】%v\n", err)
		os.Exit(1)
	}

	fmt.Println("pararell download start")
	// パラレルダウンロード実行
	if err := d.PararellDownload(); err != nil {
		fmt.Printf("【ERROR】%v\n", err)
		os.Exit(1)
	}

	fmt.Println("merge start")
	err = d.Merge()
	if err != nil {
		fmt.Printf("【ERROR】%v\n", err)
		os.Exit(1)
	}

	fmt.Println("complete")
}
