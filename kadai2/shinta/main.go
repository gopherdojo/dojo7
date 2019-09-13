package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/gopherdojo/dojo7/kadai2/shinta/imageconversion"
)

// passArgs は引数を受け取りその引数(ディレクトリ、変換前拡張子、変換後拡張子)が正しいか判別し、引数の値を返します。
func passArgs() (dir string, preExt string, afterExt string, err error) {
	d := flag.String("d", "./", "対象ディレクトリ")
	p := flag.String("p", "jpeg", "変換前拡張子")
	a := flag.String("a", "png", "変換後拡張子")
	if flag.Parse(); flag.Parsed() {
		dir, preExt, afterExt = *d, *p, *a
		return
	}
	err = errors.New("引数のparseに失敗しました。")
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
