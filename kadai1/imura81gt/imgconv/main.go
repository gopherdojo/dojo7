package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gopherdojo/dojo7/kadai1/imura81gt/imgconv/img"
)

var (
	dirs  []string
	oType img.ImageType
	iType img.ImageType
)

func main() {

	// check that dir is existed.
	for _, dir := range dirs {
		_, err := os.Stat(dir)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	err := img.ConvertAll(dirs, iType, oType)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}

func init() {
	// 4. 変換前と変換後の画像形式を指定できる（オプション）
	ip := flag.Bool("p", false, "inpt files are png.")
	ij := flag.Bool("j", false, "inpt files are jpeg.")
	ig := flag.Bool("g", false, "inpt files are gif.")

	op := flag.Bool("P", false, "convert to png.")
	oj := flag.Bool("J", false, "convert to jpeg.")
	og := flag.Bool("G", false, "convert to gif.")

	flag.Parse()

	// 1. ディレクトリを指定する
	dirs = flag.Args()

	// 2. 指定したディレクトリ以下のJPGファイルをPNGに変換（デフォルト）
	if !(*ip || *ij || *ig) {
		*ij = true
	}

	if !(*op || *oj || *og) {
		*op = true
	}

	// fmt.Println(*p, *j, *g)
	oType = img.ImageType{
		Png:  *op,
		Jpeg: *oj,
		Gif:  *og,
	}

	iType = img.ImageType{
		Png:  *ip,
		Jpeg: *ij,
		Gif:  *ig,
	}

	if len(dirs) <= 0 {
		fmt.Fprintf(os.Stderr, "%s <option> dir1 dir2 dir3\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "option:")
		flag.PrintDefaults()
		os.Exit(1)
	}

}
