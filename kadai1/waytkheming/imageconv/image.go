package imageconv

import (
	"fmt"
	"path/filepath"
	"regexp"
)

type ImageFile struct {
	Path   string
	Name   string
	Format string
}

// NewImage -> Initialize ImageFile
func NewImage(path string) ImageFile {
	format := filepath.Ext(path)
	fmt.Println(format)

	rep := regexp.MustCompile(format + "$")
	fmt.Println(rep)

	name := filepath.Base(rep.ReplaceAllString(path, ""))
	fmt.Println(name)

	return ImageFile{Path: path, Name: name, Format: format}
}
