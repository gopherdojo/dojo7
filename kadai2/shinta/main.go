package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/gopherdojo/dojo7/kadai2/shinta/imageconversion"
)

// judgeArgExt は引数に設定された拡張子が変換可能なものか判別する
func judgeArgExt(preExt, afterExt string) error {
	if preExt == afterExt {
		return errors.New("変換前と変換後で拡張子が同じです。")
	}
	allowExtList := []string{"jpg", "jpeg", "png", "gif"}
	allowExtMap := map[string]bool{}
	for _, ext := range allowExtList {
		allowExtMap[ext] = true
	}
	if !allowExtMap[preExt] || !allowExtMap[afterExt] {
		return errors.New("指定できる拡張子: " + strings.Join(allowExtList, ","))
	}
	return nil
}

// passArgs は引数を受け取りその引数(ディレクトリ、変換前拡張子、変換後拡張子)が正しいか判別し、引数の値を返します。
func passArgs() (dir string, preExt string, afterExt string, err error) {
	d := flag.String("d", "./", "対象ディレクトリ")
	p := flag.String("p", "jpg", "変換前拡張子")
	a := flag.String("a", "png", "変換後拡張子")
	flag.Parse()
	dir, preExt, afterExt = *d, *p, *a
	err = judgeArgExt(preExt, afterExt)
	if err != nil {
		return
	}
	preExt = "." + preExt
	afterExt = "." + afterExt
	return
}

func main() {
	dir, preExt, afterExt, err := passArgs()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	err = imageconversion.Excute(dir, preExt, afterExt)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
