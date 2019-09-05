package imachan

import (
	"fmt"
	"image"
	_ "image/jpeg" // to convert jpeg file
	"image/png"
	"os"
	"path/filepath"
)

// Config is ...
type Config struct {
	Path       string
	fromFormat string
	toFormat   string
}

// NewConfig is ...
func NewConfig(dir, fromFormat, toFormat string) (*Config, error) {
	path, err := filepath.Abs(dir)
	if err != nil {
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

func convert(fromPath, fromFormat, toFormat string) error {
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
