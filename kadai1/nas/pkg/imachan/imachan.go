package imachan

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

// file formats
const (
	JpgFormat = iota
	PngFormat
	GifFormat
)

// Config is ...
type Config struct {
	Path       string
	FromFormat int
	ToFormat   int
}

// ConvertedDataRepository is ...
type ConvertedDataRepository []map[string]string

// NewConfig is ...
func NewConfig(path, fromFormatStr, toFormatStr string) (*Config, error) {
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
		FromFormat: fromFormat,
		ToFormat:   toFormat,
	}, nil
}

// ConvertRec is ...
func (c *Config) ConvertRec() (ConvertedDataRepository, error) {
	var r ConvertedDataRepository
	err := filepath.Walk(c.Path, func(fromPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if targetFormat := getImageFormat(strings.Replace(filepath.Ext(fromPath), ".", "", 1)); targetFormat != c.FromFormat {
			return nil
		}
		toPath, err := convert(fromPath, c.ToFormat)
		if err != nil {
			return err
		}
		r = append(r, map[string]string{
			"from": fromPath,
			"to":   toPath,
		})
		return nil
	})
	if err != nil {
		return nil, err
	}
	return r, nil
}

func getImageFormat(formatStr string) int {
	switch formatStr {
	case "jpg", "jpeg":
		return JpgFormat

	case "png":
		return PngFormat

	case "gif":
		return GifFormat

	default:
		return -1
	}
}

func convert(fromPath string, toFormat int) (string, error) {
	var (
		toPath string
		err    error
	)
	switch toFormat {
	case PngFormat:
		toPath, err = convertToPng(fromPath)

	case JpgFormat:
		toPath, err = convertToJpg(fromPath)

	case GifFormat:
		toPath, err = convertToGif(fromPath)
	}
	if err != nil {
		return "", err
	}
	return toPath, nil
}

func convertToPng(fromPath string) (string, error) {
	fromImg, err := decodeImage(fromPath)
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
	fromImg, err := decodeImage(fromPath)
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

func convertToGif(fromPath string) (string, error) {
	fromImg, err := decodeImage(fromPath)
	if err != nil {
		return "", err
	}
	toPath := fromPath[0:len(fromPath)-len(filepath.Ext(fromPath))] + ".gif"
	toImg, err := os.Create(toPath)
	if err != nil {
		return "", err
	}
	err = gif.Encode(toImg, fromImg, &gif.Options{NumColors: 256})
	if err != nil {
		return "", err
	}
	return toPath, nil
}

func decodeImage(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	image, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}
	return image, nil
}
