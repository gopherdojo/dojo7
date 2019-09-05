package imachan

import (
	"fmt"
	"image"
	_ "image/jpeg" // to convert jpeg file
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

const (
	jpegFormat = iota
	pngFormat
)

// Config is ...
type Config struct {
	Path       string
	fromFormat int
	toFormat   int
}

// NewConfig is ...
func NewConfig(dir, fromFormatStr, toFormatStr string) (*Config, error) {
	path, err := filepath.Abs(dir)
	if err != nil {
		return nil, err
	}
	fromFormat := getImageFormat(fromFormatStr)
	if fromFormat == -1 {
		err := fmt.Errorf("undefind %s file format, please choose another", fromFormatStr)
		return nil, err
	}
	toFormat := getImageFormat(toFormatStr)
	if toFormat == -1 {
		err := fmt.Errorf("undefind %s file format, please choose another", toFormatStr)
		return nil, err
	}
	return &Config{
		Path:       path,
		fromFormat: fromFormat,
		toFormat:   toFormat,
	}, nil
}

// Run is ...
func (c *Config) Run() error {
	err := filepath.Walk(c.Path, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		pathFormat := getImageFormat(strings.Replace(filepath.Ext(path), ".", "", 1))
		if pathFormat != c.fromFormat {
			return nil
		}
		err = convert(path, c.fromFormat, c.toFormat)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func getImageFormat(formatStr string) int {
	switch formatStr {
	case "jpg", "jpeg":
		return jpegFormat
	case "png":
		return pngFormat
	default:
		return -1
	}
}

func convert(fromPath string, fromFormat, toFormat int) error {
	f, err := os.Open(fromPath)
	if err != nil {
		return err
	}
	defer f.Close()
	fromImg, _, err := image.Decode(f)
	if err != nil {
		return err
	}
	toPath := fromPath[0:len(fromPath)-len(filepath.Ext(fromPath))] + ".png"
	toImg, err := os.Create(toPath)
	err = png.Encode(toImg, fromImg)
	if err != nil {
		return err
	}
	fmt.Println("success:", fromPath, "->", toPath)
	return nil
}
