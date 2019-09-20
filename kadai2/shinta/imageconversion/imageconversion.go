/*
Package imageconversion は画像ファイル形式の変換を行います。
optionで、実行するディレクトリと変換前と変換後の画像形式を指定できます。
option を指定しない場合、コマンドを実行するディレクトリと、 変換前の画像タイプがjpeg、変換後の画像タイプがpngになります。
変換可能な拡張子として、jpg、jpeg、png、gif としています。
*/
package imageconversion

import (
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

type arg struct {
	dir      string
	preExt   string
	afterExt string
}

func (a *arg) valid() error {
	if a.preExt == a.afterExt {
		return errors.New("変換前と変換後で拡張子が同じです。")
	}
	allowExtList := []string{"jpg", "jpeg", "png", "gif"}
	allowExtMap := map[string]bool{
		"jpg":  true,
		"jpeg": true,
		"png":  true,
		"gif":  true,
	}
	if !allowExtMap[a.preExt] || !allowExtMap[a.afterExt] {
		return errors.New("指定できる拡張子: " + strings.Join(allowExtList, ","))
	}
	return nil
}

// func (a *arg) convertExt() {
// 	a.preExt, a.afterExt = "."+a.preExt, "."+a.afterExt
// }

// imageFile struct は変換対象の画像のpath(path)、拡張子を除いたファイル名(base)、拡張子(ext)を持っています。
type imageFile struct {
	Path string
	Base string
	Ext  string
}

// getFileNameWithoutExt は対象ファイルのpathと拡張子を除いたファイル名を返します。
func getFileNameWithoutExt(path string) string {
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}

// createImgStrunct は、imageFile structを生成し、返します。
func createImgStruct(path string) imageFile {
	base := getFileNameWithoutExt(path)
	return imageFile{filepath.Dir(path), base, filepath.Ext(path)}
}

/*
convertExec は画像ファイルを引数で指定された変換後の拡張子(defaultはpng)に変換した新しい画像ファイルを生成します。
処理が成功するとnil、errorが起きた場合、errorを返します。
*/
func convertExec(path, afterExt string) error {
	img := createImgStruct(path)
	targetImg, err := os.Open(filepath.Join(img.Path, (img.Base + img.Ext)))
	if err != nil {
		return err
	}
	readImg, _, err := image.Decode(targetImg)
	if err != nil {
		return err
	}
	outputImg, err := os.Create(filepath.Join(img.Path, (img.Base + "." + afterExt)))
	if err != nil {
		return err
	}

	switch afterExt {
	case ".jpeg", ".jpg":
		jpeg.Encode(outputImg, readImg, nil)
	case ".gif":
		gif.Encode(outputImg, readImg, nil)
	default:
		png.Encode(outputImg, readImg)
	}
	if err = targetImg.Close(); err != nil {
		return err
	}
	err = outputImg.Close()
	return err
}

/*
Excute は画像変換処理を実行します。
このpackageで呼び出せる唯一の関数です。
引数で、ディレクトリ(デフォルトは ./)、変換前拡張子(デフォルトは jpg)、変換後拡張子(デフォルトは png)を受け取ります。
引数が指定されない場合はデフォルトの値が適用されます。
引数で受け取ったディレクトリ以下の変換前拡張子のファイルを変換後拡張子に変換した新しいファイルを作成します。
処理が成功の場合、nilをerrorが起きた場合はerrorを返します。
引数で指定されたディレクトリ以下から引数で指定した変換前拡張子(defaultはjpg)のファイルを、
変換後拡張子(defaultはpng)に変換した新しい画像ファイルを生成します。
処理が成功するとnil、errorが起きた場合、errorを返します。
*/
func Excute(dir, preExt, afterExt string) error {
	arg := &arg{dir, preExt, afterExt}
	if err := arg.valid(); err != nil {
		return err
	}
	// arg.convertExt()
	err := filepath.Walk(arg.dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(path) == ("." + arg.preExt) {
			err = convertExec(path, arg.afterExt)
			if err != nil {
				return err
			}
		}
		return err
	})
	return err
}
