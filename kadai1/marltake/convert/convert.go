package convert

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

func ConfigConvert(src string, dest string) func(string, os.FileInfo, error) error {
	srcExt := "." + src
	lenSrcExt := len(srcExt)
	destExt := "." + dest
	// TODO declare decode for src and encode for dest here
	return func(path string, info os.FileInfo, err error) error {
		println(path)
		if strings.ToLower(filepath.Ext(path)) == srcExt {
			destPath := path[:len(path)-lenSrcExt] + destExt
			// TODO error handling
			if _, err := os.Stat(destPath); os.IsNotExist(err) {
				file, _ := os.Open(path)
				defer file.Close()
				var img image.Image
				switch src {
				case "jpg":
					img, _ = jpeg.Decode(file)
				case "png":
					img, _ = png.Decode(file)
				case "gif":
					img, _ = gif.Decode(file)
				}
				destfile, _ := os.Create(destPath)
				defer destfile.Close()
				switch dest {
				case "jpg":
					jpeg.Encode(destfile, img, nil)
				case "png":
					png.Encode(destfile, img)
				case "gif":
					gif.Encode(destfile, img, nil)
				}
			} else {
				println("skip not to over write.", path)
			}
		}
		return nil
	}
}

// ParseTarget parse and validate target image type option
func ParseTarget(target string) (src string, dest string, ok bool) {
	targets := strings.Split(target, ",")
	allowedExt := map[string]bool{
		"jpg": true,
		"png": true,
		"gif": true,
	}
	src, dest, ok = "", "", false
	if len(targets) == 0 {
		src, dest = "jpg", "png"
	} else if len(targets) == 1 {
		src, dest = targets[0], "png"
	} else if len(targets) == 2 {
		src, dest = targets[0], targets[1]
	} else {
		return
	}
	if src == "" {
		src = "jpg"
	}
	if dest == "" {
		dest = "png"
	}
	if ok = src != dest && allowedExt[src] && allowedExt[dest]; !ok {
		src, dest = "", ""
	}
	return
}
