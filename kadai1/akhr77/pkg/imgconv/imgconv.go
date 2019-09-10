package imgconv

import (
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// return Code
const (
	ExitCodeOK    = 0
	ExitCodeError = 1
)

// CLI struct
type CLI struct {
	OutStream, ErrStream io.Writer
}

var (
	src  string
	from string
	to   string
)

// Run is implements function
func (c *CLI) Run(args []string) int {
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	flags.SetOutput(c.ErrStream)
	// ショートオプション
	flags.StringVar(&src, "s", "", "変換したい画像のファイルパスを指定")
	flags.StringVar(&from, "f", "jpg", "変換前の画像形式を指定")
	flags.StringVar(&to, "t", "png", "変換後の画像形式を指定")
	// ロングオプション
	flags.StringVar(&src, "src", "", "変換したい画像のファイルパスを指定")
	flags.StringVar(&from, "from", "jpg", "変換前の画像形式を指定")
	flags.StringVar(&to, "to", "png", "変換後の画像形式を指定")

	if err := flags.Parse(args[1:]); err != nil {
		fmt.Fprintf(c.ErrStream, "解析処理でエラーになりました")
		return ExitCodeError
	}

	err := walk(src, from, to)
	if err != nil {
		return ExitCodeError
	}

	return ExitCodeOK
}

// walkは第一引数以下のディレクトリを再帰的に処理する
func walk(root, beforeExt, afterExt string) error {
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		n := info.Name()
		if strings.HasSuffix(n, beforeExt) {
			src, err := os.Open(path)
			if err != nil {
				return err
			}
			defer src.Close()

			// extentionを含まないファイル名
			n := filepath.Base(n[:len(n)-len(filepath.Ext(n))])
			dir := filepath.Dir(path)
			dest, err := os.Create(filepath.Join(dir, n+"."+afterExt))
			if err != nil {
				return err
			}

			err = convert(src, dest, afterExt)
			if err != nil {
				// 変換処理に失敗した場合、不要なファイルが作成されてしまうため、削除する
				dest.Close()
				e := os.Remove(filepath.Join(dir, n+"."+afterExt))
				if e != nil {
					return e
				}
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func convert(src io.Reader, dest io.Writer, extention string) error {
	img, _, err := image.Decode(src)
	if err != nil {
		return err
	}

	switch extention {
	case "jpg", "jpeg":
		err = jpeg.Encode(dest, img, nil)
	case "png":
		err = png.Encode(dest, img)
	case "default":
		err = fmt.Errorf("supportしていないformatです")
	}
	if err != nil {
		return err
	}
	return nil
}
