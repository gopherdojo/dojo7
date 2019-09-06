package imachan

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

const (
	jpgFormat = iota
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
	err := filepath.Walk(c.Path, func(target string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if targetFormat := getImageFormat(strings.Replace(filepath.Ext(target), ".", "", 1)); targetFormat != c.fromFormat {
			return nil
		}
		err = convert(target, c.toFormat)
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
		return jpgFormat
	case "png":
		return pngFormat
	default:
		return -1
	}
}

func convert(fromPath string, toFormat int) error {
	switch toFormat {
	case pngFormat:
		convertToPng(fromPath)
	case jpgFormat:
		convertToJpg(fromPath)
	}
	toPath, err := convertToPng(fromPath)
	if err != nil {
		return err
	}
	fmt.Println("success:", fromPath, "->", toPath)
	return nil
}

func convertToPng(fromPath string) (string, error) {
	f, err := os.Open(fromPath)
	if err != nil {
		return "", err
	}
	defer f.Close()
	fromImg, _, err := image.Decode(f)
	if err != nil {
		return "", err
	}
	toPath := fromPath[0:len(fromPath)-len(filepath.Ext(fromPath))] + ".png"
	toImg, err := os.Create(toPath)
	if err != nil {
		return "", err
	}
	err = png.Encode(toImg, fromImg)
	if err != nil {
		return "", err
	}
	return toPath, nil
}

func convertToJpg(fromPath string) (string, error) {
	f, err := os.Open(fromPath)
	if err != nil {
		return "", err
	}
	defer f.Close()
	fromImg, _, err := image.Decode(f)
	if err != nil {
		return "", err
	}
	toPath := fromPath[0:len(fromPath)-len(filepath.Ext(fromPath))] + ".jpg"
	toImg, err := os.Create(toPath)
	if err != nil {
		return "", err
	}
	err = jpeg.Encode(toImg, fromImg, &jpeg.Options{Quality: 100})
	if err != nil {
		return "", err
	}
	return toPath, nil
}
