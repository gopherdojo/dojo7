/*
このパッケージは、画像ファイルをpng,jpg,gifからpng,jpg,gifへ変換する機能を持っています。

*/
package imageconv

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path/filepath"
	"sync"
)

// Converter -> Converter struct
type Converter struct {
	Path   string
	Images []ImageFile
	From   string
	To     string
}

// NewConverter -> Initialize ImageConverter
func NewConverter(path string, from string, to string) Converter {
	return Converter{Path: path, From: from, To: to}
}

// GetImages is queuing imageFile
func (c *Converter) GetImages(q chan ImageFile, wg *sync.WaitGroup) {
	for {
		image, more := <-q
		if more {
			_ = c.Convert(image)
		} else {
			wg.Done()
			return
		}
	}
}

//Convert -> Convert Image FIl
func (c *Converter) Convert(i ImageFile) error {
	file, err := os.Open(i.Path)
	if err != nil {
		return err
	}
	defer file.Close()

	image, err := c.decode(file)
	if err != nil {
		return err
	}
	outFile, err := os.Create(i.Name + "." + c.To)
	if err != nil {
		return err
	}
	defer outFile.Close()

	err = c.encode(outFile, image)

	if err != nil {
		return err
	}

	return nil
}

// CrawlFile -> found image file and append Converter.Files
func (c *Converter) CrawlFile(path string, info os.FileInfo, err error) error {
	if checkExtension(filepath.Ext(path)) == ("." + c.From) {
		if !info.IsDir() {
			c.Images = append(c.Images, NewImage(path))
		}
	}
	return nil
}

func checkExtension(path string) string {
	if path == ".jpeg" {
		return ".jpg"
	}
	return path
}

func (c *Converter) decode(file io.Reader) (image.Image, error) {
	var (
		img image.Image
		err error
	)
	switch c.From {
	case "jpeg", "jpg", "JPG", "JPEG":
		img, err = jpeg.Decode(file)
	case "gif", "GIF":
		img, err = gif.Decode(file)
	case "png", "PNG":
		img, err = png.Decode(file)
	}
	if err != nil {
		return nil, err
	}

	return img, nil
}

func (c *Converter) encode(file io.Writer, img image.Image) error {
	var err error
	switch c.To {
	case "jpeg", "jpg", "JPG", "JPEG":
		err = jpeg.Encode(file, img, nil)
	case "gif", "GIF":
		err = gif.Encode(file, img, nil)
	case "png", "PNG":
		err = png.Encode(file, img)
	}
	if err != nil {
		return err
	}

	return nil
}
