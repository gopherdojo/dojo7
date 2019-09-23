package main

import (
	"flag"
	"fmt"
	"os"
)

var url = flag.String("url", "", "ダウンロードファイルURL指定")
var divNum = flag.Int("divNum", 0, "分割数")

func main() {

	flag.Parse()

	if *url == "" {
		fmt.Println("【ERROR】 ダウンロードファイルURL指定を指定してください")
		os.Exit(1)
	}

	if *divNum == 0 {
		fmt.Println("【ERROR】 分割数を指定してください")
		os.Exit(1)
	}

	// TODO rangeアクセス可能かを確認

	// TODO 分割してダウンロード処理を行う

	// TODO 分割された結果をマージしてファイル出力する

}
