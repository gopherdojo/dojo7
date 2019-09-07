// メインパッケージ
package main

import (
	"conv/internal/image"
	"flag"
	"fmt"
	"os"

	"conv/internal/imageconv"
)

var (
	targetDirFlag = flag.String("targetDir", "",
		"targeted directory")
	srcExtFlag = flag.String("srcExt", ".jpg",
		"targeted source file extension")
	dstExtFlag = flag.String("dstExt", ".png",
		"targeted destination file extension")
)

// コマンドライン引数のバリデーションを行い、不正な場合はfalseを返す。
func validateFlag() bool {
	if len(*targetDirFlag) == 0 {
		_, _ = fmt.Fprintf(os.Stderr, "target directory mustn't be empty\n")
		return false
	}

	okSrcExtFlag := false
	okDstExtFlag := false
	for _, e := range image.Extensions {
		if e == *srcExtFlag {
			okSrcExtFlag = true
		}
		if e == *dstExtFlag {
			okDstExtFlag = true
		}
	}
	if !okSrcExtFlag {
		_, _ = fmt.Fprintf(os.Stderr, "source file extension is invalid\n")
		return false
	}
	if !okDstExtFlag {
		_, _ = fmt.Fprintf(os.Stderr, "destination file extension is invalid\n")
		return false
	}

	return true
}

func main() {
	flag.Parse()
	if !validateFlag() {
		os.Exit(1)
	}

	_ = imageconv.ConvertRecursively(*targetDirFlag, *srcExtFlag, *dstExtFlag)
}
