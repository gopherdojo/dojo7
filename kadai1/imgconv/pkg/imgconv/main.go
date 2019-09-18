package imgconv

import (
	"fmt"
	"image"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type ImageConf struct {
	FromExt     string
	ToExt       string
	DecodedFile image.Image
	ToFile      io.Writer
}

type ImageConfList []ImageConf

// create ImageConfList
func NewImageConfList(dir, fromExt, toExt string) (ImageConfList, error) {
	// read dir
	paths, err := readDir(dir, strings.ToLower(fromExt))
	if err != nil {
		return nil, err
	}
	imgCfList := make(ImageConfList, len(paths))
	for i, path := range paths {
		imgCf := ImageConf{
			FromExt: strings.ToLower(fromExt),
			ToExt:   strings.ToLower(toExt),
		}

		// set decode file
		err = imgCf.SetDecodedFile(path)
		if err != nil {
			return nil, err
		}

		// set output file
		err = imgCf.SetToFile(getConvertedPath(path, strings.ToLower(toExt)))
		if err != nil {
			return nil, err
		}

		imgCfList[i] = imgCf
	}
	return imgCfList, nil
}

// set ImageConf.DecodedFile
func (imgCf *ImageConf) SetDecodedFile(path string) (err error) {
	// open file
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer func() {
		err = file.Close()
	}()

	// decode file
	img, _, err := image.Decode(file)
	if err != nil {
		return
	}
	imgCf.DecodedFile = img
	return
}

// set ImageConf.ToFile
func (imgCf *ImageConf) SetToFile(path string) error {
	// create output file
	toFile, err := os.Create(path)
	if err != nil {
		return err
	}
	imgCf.ToFile = toFile
	return nil
}

//  recursively read directory
func readDir(dir, fromExt string) ([]string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths2, err := readDir(filepath.Join(dir, file.Name()), fromExt)
			if err != nil {
				return nil, err
			}
			paths = append(paths, paths2...)
			continue
		}
		extStart := strings.LastIndex(file.Name(), ".")
		if extStart > 0 && strings.ToLower(file.Name()[extStart+1:]) == fromExt {
			paths = append(paths, filepath.Join(dir, file.Name()))
		} else {
			continue
		}
	}
	return paths, nil
}

// get converted file path
func getConvertedPath(path, toExt string) string {
	extStart := strings.LastIndex(path, ".")
	return fmt.Sprintf("%s.%s", path[:extStart], toExt)
}
