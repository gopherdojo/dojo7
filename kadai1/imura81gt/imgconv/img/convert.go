package img

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path/filepath"
)

const (
	saveDir  = "output"
	pngType  = "png"
	jpegType = "jpeg"
	gifType  = "gif"
)

// ConvType is type of converted image files
// 8. ユーザ定義型を作ってみる
type ImageType struct {
	Png  bool
	Jpeg bool
	Gif  bool
}

func (i ImageType) Enable() []string {
	var types []string
	if i.Png {
		types = append(types, pngType)
	}
	if i.Jpeg {
		types = append(types, jpegType)
	}
	if i.Gif {
		types = append(types, gifType)
	}
	return types
}

// ConvertAll convert all image files in dirs
func ConvertAll(dirs []string, iType ImageType, oType ImageType) error {
	files := AllImageFiles(dirs)
	files = ExpectedImageFiles(files, iType)
	fmt.Printf("%q\n", files)

	var oFiles []string
	for _, iPath := range files {
		for _, t := range oType.Enable() {
			oPath := filepath.Join(saveDir, t, iPath[:len(iPath)-len(filepath.Ext(iPath))]+"."+t)

			i, err := os.Open(iPath)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			err = os.MkdirAll(filepath.Dir(oPath), 0755)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			w, err := os.Create(oPath)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			err = Convert(i, w, t)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			oFiles = append(oFiles, oPath)
		}
	}
	fmt.Printf("%q\n", oFiles)
	return nil
}

func AllImageFiles(dirs []string) (files []string) {
	// 3. ディレクトリ以下は再帰的に処理する
	for _, dir := range dirs {
		err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}

			if IsImage(path) {
				files = append(files, path)
			}

			return nil
		})
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
	return files
}

func ExpectedImageFiles(files []string, iType ImageType) (f []string) {
	var ret []string
	for _, path := range files {
		if IsExpectedImage(path, iType) {
			ret = append(ret, path)
		}
	}
	return ret
}

// IsImage return true if path is image.
func IsImage(path string) bool {
	r, err := os.Open(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer r.Close()

	_, _, err = image.DecodeConfig(r)
	if err != nil {
		return false
	}
	return true
}

func IsExpectedImage(path string, iType ImageType) bool {
	r, err := os.Open(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer r.Close()

	_, format, err := image.DecodeConfig(r)
	if err != nil {
		return false
	}

	for _, t := range iType.Enable() {
		if t == format {
			return true
		}
	}

	return false
}

// Convert function is image file converter
func Convert(r io.Reader, w io.Writer, t string) error {
	jpego := &jpeg.Options{
		Quality: jpeg.DefaultQuality,
	}

	gifo := &gif.Options{
		NumColors: 256,
		Quantizer: nil,
		Drawer:    nil,
	}

	m, _, err := image.Decode(r)
	if err != nil {
		return err
	}

	switch t {
	case "png":
		return png.Encode(w, m)
	case "gif":
		return gif.Encode(w, m, gifo)
	case "jpeg", "jpg":
		return jpeg.Encode(w, m, jpego)
	default:
		return png.Encode(w, m)
	}

	return nil

}
