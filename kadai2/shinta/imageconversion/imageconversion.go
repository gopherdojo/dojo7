/*
Package imageconversion は画像ファイル形式の変換を行います。
optionで、実行するディレクトリと変換前と変換後の画像形式を指定できます。
option を指定しない場合、コマンドを実行するディレクトリと、 変換前の画像タイプがjpeg、変換後の画像タイプがpngになります。
変換可能な拡張子として、jpg、jpeg、png、gif としています。
*/
package imageconversion

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

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
func convertExec(path, afterExt string) error {
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
	err = targetImg.Close()
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
	// 変換対象ファイルが jpeg or jpg かを確認する
	jpgType := map[string]bool{".jpg": true, ".jpeg": true}
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// jpeg は jpgも変換対象とする
		if jpgType[afterExt] && (filepath.Ext(path) == ".jpeg" || filepath.Ext(path) == ".jpg") {
			err = convertExec(path, afterExt)
			if err != nil {
				return err
			}
		}
		if filepath.Ext(path) == preExt {
			err = convertExec(path, afterExt)
			if err != nil {
				return err
			}
		}
		return err
	})
	return err
}
