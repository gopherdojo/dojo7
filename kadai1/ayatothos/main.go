package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gopherdojo/dojo7/kadai1/ayatothos/imgconv"
)

var fromFileType = flag.String("from", "jpg", "変換前拡張子名")
var toFileType = flag.String("to", "png", "変換後拡張子名")

func main() {

	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Println("【ERROR】 ディレクトリパスを指定してください")
		os.Exit(1)
	}

	fromExtentions, err := imgconv.GetExtentionsByName(*fromFileType)
	if err != nil {
		fmt.Println("【ERROR】 ", err.Error()+":"+*fromFileType)
		os.Exit(1)
	}
	if _, err := imgconv.GetExtentionsByName(*toFileType); err != nil {
		fmt.Println("【ERROR】 ", err.Error()+":"+*toFileType)
		os.Exit(1)
	}

	fmt.Println("【PROCESSING】 変換前拡張子: ", fromExtentions)
	fmt.Println("【PROCESSING】 変換後拡張子: ", "."+*toFileType)
	fmt.Println("【PROCESSING】 変換対象ディレクトリ: ", flag.Arg(0))

	convertFilePath, err := imgconv.ConvertImageAll(flag.Arg(0), *fromFileType, *toFileType)
	if err != nil {
		fmt.Println("【ERROR】 ", err.Error()+":"+*toFileType)
		os.Exit(1)
	}

	for _, v := range convertFilePath {
		fmt.Println("【PROCESSED】 変換完了ファイル: ", v)
	}

	fmt.Println("【SUCCESS】 処理は正常終了しました")
}
