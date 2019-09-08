package imageconv

import (
	"path/filepath"
	"regexp"
)

//ImageFile -> Image struct
type ImageFile struct {
	Path   string
	Name   string
	Format string
}

// NewImage -> Initialize ImageFile
func NewImage(path string) ImageFile {
	format := filepath.Ext(path)
	rep := regexp.MustCompile(format + "$")
	name := filepath.Base(rep.ReplaceAllString(path, ""))

	return ImageFile{Path: path, Name: name, Format: format}
}
