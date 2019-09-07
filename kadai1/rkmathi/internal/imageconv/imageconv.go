// 画像変換を行うパッケージ
package imageconv

import (
	"conv/internal/image"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// 指定したディレクトリ以下を再帰的に走査して画像変換を行う。
func ConvertRecursively(targetDir, srcExt, dstExt string) []string {
	dirs, err := ioutil.ReadDir(targetDir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range dirs {
		if file.IsDir() {
			paths = append(paths,
				ConvertRecursively(filepath.Join(targetDir, file.Name()), srcExt, dstExt)...)
			continue
		}

		paths = append(paths, filepath.Join(targetDir, file.Name()))

		for _, srcPath := range paths {
			if filepath.Ext(srcPath) == srcExt {
				lastIdx := len(paths) - 1
				dstPath := paths[lastIdx][:len(paths[lastIdx])-len(filepath.Ext(paths[lastIdx]))] + dstExt

				err := Convert(srcPath, dstPath)
				if err != nil {
					panic(err)
				}
			}
		}
	}

	return paths
}

// 画像変換をする。
func Convert(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	imageType := image.DetectImageType(src)
	srcImg, err := image.Decode(imageType, srcFile)
	if err != nil {
		return err
	}

	dstImg, err := os.Create(fmt.Sprintf(dst))
	if err != nil {
		return err
	}
	defer dstImg.Close()

	err = image.Encode(imageType, srcImg, dstImg)
	if err != nil {
		return err
	}

	return nil
}
