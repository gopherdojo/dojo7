//Package imageconversion は画像ファイル形式の変換を行います。
//option で、実行するディレクトリと変換前と変換後の画像形式を指定できます。
//option を指定しない場合、コマンドを実行するディレクトリと、 変換前の画像タイプがjpeg、変換後の画像タイプがpngになります。
// 変換可能な拡張子として、jpg、jpeg、png、gif としています。
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

type imageFile struct {
	path string
	base string
	ext  string
}

type imageList []imageFile

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

func passArgs() (dir string, preExt string, afterExt string, err error) {
	d := flag.String("d", "./", "対象ディレクトリ")
	p := flag.String("p", "jpeg", "変更前画像拡張子")
	a := flag.String("a", "png", "変更後画像拡張子")
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

func getFileNameWithoutExt(path string) string {
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}

func createImgStruct(path string) (image imageFile) {
	base := getFileNameWithoutExt(path)
	image = imageFile{filepath.Dir(path), base, filepath.Ext(path)}
	return
}

func searchImages(dir string, preExt string) (list imageList, err error) {
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
		if jpgFlag {
			// jpeg は jpgも変換対象とする
			if filepath.Ext(path) == ".jpeg" || filepath.Ext(path) == ".jpg" {
				image := createImgStruct(path)
				list = append(list, image)
			}
		} else {
			if filepath.Ext(path) == preExt {
				image := createImgStruct(path)
				list = append(list, image)
			}
		}
		return nil
	})
	return
}

func convetImages(list imageList, afterExt string) (err error) {
	for _, img := range list {
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
	}
	return
}

// Excute は画像変換処理を実行します。
func Excute() error {
	dir, preExt, afterExt, err := passArgs()
	if err != nil {
		return err
	}
	list, err := searchImages(dir, preExt)
	if err != nil {
		return err
	}
	err = convetImages(list, afterExt)
	return err
}
