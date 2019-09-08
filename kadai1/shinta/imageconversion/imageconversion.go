/*
Package imageconversion は画像ファイル形式の変換を行います。
optionで、実行するディレクトリと変換前と変換後の画像形式を指定できます。
option を指定しない場合、コマンドを実行するディレクトリと、 変換前の画像タイプがjpeg、変換後の画像タイプがpngになります。
変換可能な拡張子として、jpg、jpeg、png、gif としています。
*/
package imageconversion

import (
	"errors"
	"flag"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

// judgeArgExt は引数に設定された拡張子が変換可能なものか判別する
func judgeArgExt(preExt string, afterExt string) (err error) {
	allowExtList := []string{"jpg", "jpeg", "png", "gif"}
	argExtList := []string{preExt, afterExt}
	var judgeExt bool
	for i, argExt := range argExtList {
		if i == len(argExtList)-1 {
			judgeExt = false
		}
		for _, allowExt := range allowExtList {
			if allowExt == argExt {
				judgeExt = true
				break
			}
		}
	}
	if !judgeExt {
		err = errors.New("指定できる拡張子:" + strings.Join(allowExtList, ","))
	}
	return
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

// imageFile struct は変換対象の画像のpath(path)、拡張子を除いたファイル名(base)、拡張子(ext)を持っています。
type imageFile struct {
	path string
	base string
	ext  string
}

// getFileNameWithoutExt は対象ファイルのpathと拡張子を除いたファイル名を返します。
func getFileNameWithoutExt(path string) string {
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}

// createImgStrunct は、imageFile structを生成し、返します。
func createImgStruct(path string) (image imageFile) {
	base := getFileNameWithoutExt(path)
	image = imageFile{filepath.Dir(path), base, filepath.Ext(path)}
	return
}

/*
convertExec は画像ファイルを引数で指定された変換後の拡張子(defaultはpng)に変換した新しい画像ファイルを生成します。
処理が成功するとnil、errorが起きた場合、errorを返します。
*/
func convertExec(path string, afterExt string) (err error) {
	img := createImgStruct(path)
	targetImg, err := os.Open(img.path + "/" + img.base + img.ext)
	if err != nil {
		return err
	}
	readImg, _, err := image.Decode(targetImg)
	if err != nil {
		return err
	}
	outputImg, err := os.Create((img.path + "/" + img.base + afterExt))
	if err != nil {
		return err
	}

	switch afterExt {
	case "jpeg", "jpg":
		jpeg.Encode(outputImg, readImg, nil)
	case "gif":
		gif.Encode(outputImg, readImg, nil)
	default:
		png.Encode(outputImg, readImg)
	}

	targetImg.Close()
	outputImg.Close()
	return
}

/*
convertImages は、引数で指定されたディレクトリ以下から引数で指定した変換前拡張子(defaultはjpg)のファイルを、
変換後拡張子(defaultはpng)に変換した新しい画像ファイルを生成します。
処理が成功するとnil、errorが起きた場合、errorを返します。
*/
func convertImages(dir string, preExt string, afterExt string) (err error) {
	// 変換対象ファイルが jpeg or jpg かを確認する
	jpgType := [2]string{".jpg", ".jpeg"}
	var jpgFlag bool
	for _, v := range jpgType {
		if preExt == v {
			jpgFlag = true
		}
	}
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// jpeg は jpgも変換対象とする
		if jpgFlag && filepath.Ext(path) == ".jpeg" || filepath.Ext(path) == ".jpg" {
			convertExec(path, afterExt)
		}
		if filepath.Ext(path) == preExt {
			convertExec(path, afterExt)
		}
		return nil
	})
	return
}

/*
Excute は画像変換処理を実行します。
このpackageで呼び出せる唯一の関数です。
引数で、ディレクトリ(デフォルトは ./)、変換前拡張子(デフォルトは jpg)、変換後拡張子(デフォルトは png)を受け取ります。
引数が指定されない場合はデフォルトの値が適用されます。
引数で受け取ったディレクトリ以下の変換前拡張子のファイルを変換後拡張子に変換した新しいファイルを作成します。
処理が成功の場合、nilをerrorが起きた場合はerrorを返します。
*/
func Excute() error {
	dir, preExt, afterExt, err := passArgs()
	if err != nil {
		return err
	}
	err = convertImages(dir, preExt, afterExt)
	return err
}
