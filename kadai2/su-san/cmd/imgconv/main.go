package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gopherdojo/dojo7/kadai2/su-san"
)

func main() {

	inputExt := flag.String("i", "jpg", " extension to be converted ")
	outputExt := flag.String("o", "png", " extension after conversion")

	// Usageメッセージ
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "usage : cmd [-i] [-o] target_dir")
		flag.PrintDefaults()
	}

	flag.Parse()
	args := flag.Args()

	if flag.NArg() == 0 {
		flag.Usage()
		return
	}

	targetDir := args[0]

	// 変換対象のフォーマットと変換フォーマットが同じなら何もせず終了
	if *inputExt == *outputExt {
		return
	}

	convExts := image.NewConvExts(*inputExt, *outputExt)
	if !convExts.SupportedFormats() {
		fmt.Fprintf(os.Stderr, "unsupported format! please specify these format [png jpg jpeg gif]\n")
		return
	}

	var targetPaths []string
	err := filepath.Walk(targetDir, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ("." + *inputExt) {
			targetPaths = append(targetPaths, path)
		}
		return nil
	})

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for _, p := range targetPaths {

		//var p Path = Path(p)
		// 対象の拡張子出ない場合はスルーする
		path := Path(p)
		if !(path.IsConvesionTargetExt(convExts)) {
			continue
		}

		// 対象の拡張子でない場合はスルーする
		if image.SupportedFormat(filepath.Ext(p)) == false {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		f, err := os.Open(p)
		// 読み込みエラーの場合はエラーを出してスルーする
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		err = image.FmtConv(f, convExts)
		if err != nil {
			fmt.Fprintln(os.Stderr, err, " error filepath:", p)
			os.Exit(1)
		}

		// 成功ならば変換元を削除する
		if err := os.Remove(p); err != nil {
			fmt.Fprintln(os.Stderr, err, " error filepath:", p)
			os.Exit(1)
		}

		err = f.Close()
		if err != nil {
			fmt.Fprintln(os.Stderr, err, " error filepath:", p)
			os.Exit(1)
		}
	}
}

type Path string

func (p *Path) IsConvesionTargetExt(c image.ConvExts) bool {
	ext := filepath.Ext(string(*p))
	if len(ext) == 0 {
		return false
	}

	if ext[0] == '.' {
		ext = ext[1:]
	}

	return ext == c.InExt
}
