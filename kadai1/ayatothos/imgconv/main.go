package imgconv

import (
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

// FileType 画像種別構造体
type FileType struct {
	name       string
	extentions []string
}

var validFileTypes = []FileType{
	{
		name:       "jpg",
		extentions: []string{".jpg", ".jpg", ".JPG", ".JPEG"},
	},
	{
		name:       "png",
		extentions: []string{".png", ".PNG"},
	},
	{
		name:       "gif",
		extentions: []string{".gif", ".GIF"},
	},
}

// ConvertImageAll ディレクトリ内の画像を変換する
func ConvertImageAll(dirPath, fromType, toType string) ([]string, error) {

	if f, err := os.Stat(dirPath); os.IsNotExist(err) || !f.IsDir() {
		return nil, errors.New("正しいディレクトリパスを指定してください")
	}

	fromExtentions, err := getExtentionsByName(fromType)
	if err != nil {
		return nil, err
	}

	convertFilePaths := []string{}

	err = filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		for _, v := range fromExtentions {
			if ext := filepath.Ext(path); ext == v {
				srcPath := dirPath + path
				destPath := srcPath[:len(srcPath)-len(ext)] + "." + toType

				if err = ConvertImage(srcPath, destPath, toType); err != nil {
					return err
				}

				convertFilePaths = append(convertFilePaths, destPath)
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return convertFilePaths, nil
}

// ConvertImage 画像の変換
func ConvertImage(srcPath, destPath, toType string) error {

	file, err := os.Open(srcPath)
	if err != nil {
		return errors.New("ファイルを開くことができません:" + srcPath)
	}
	defer file.Close()

	srcImage, _, err := image.Decode(file)
	if err != nil {
		return errors.New("画像ファイルを配置してください:" + srcPath)
	}

	out, err := os.Create(destPath)
	if err != nil {
		return errors.New("画像を生成できません:" + destPath)
	}
	defer out.Close()

	switch toType {
	case "jpg":
		options := &jpeg.Options{Quality: 100}
		if err := jpeg.Encode(out, srcImage, options); err != nil {
			return errors.New("画像のエンコード保存に失敗しました:" + destPath)
		}
	case "png":
		if err := png.Encode(out, srcImage); err != nil {
			return errors.New("画像のエンコード保存に失敗しました:" + destPath)
		}
	case "gif":
		options := &gif.Options{NumColors: 256}
		if err := gif.Encode(out, srcImage, options); err != nil {
			return errors.New("画像のエンコード保存に失敗しました:" + destPath)
		}
	default:
		return errors.New("有効なタイプを変換後拡張子名を指定してください:" + toType)

	}
	return nil
}

func getExtentionsByName(fileType string) ([]string, error) {
	for _, v := range validFileTypes {
		if v.name == fileType {
			return v.extentions, nil
		}
	}
	return nil, errors.New("拡張子が取得できません")
}
