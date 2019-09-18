package imgconv

import (
	"image/gif"
	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
)

func (l ImageConfList) Convert() error {
	for _, v := range l {
		err := v.Convert()
		if err != nil {
			return err
		}
	}
	return nil
}

func (imgCf ImageConf) Convert() error {
	switch imgCf.ToExt {
	case "png":
		return png.Encode(imgCf.ToFile, imgCf.DecodedFile)
	case "gif":
		return gif.Encode(imgCf.ToFile, imgCf.DecodedFile, nil)
	case "jpg", "jpeg":
		return jpeg.Encode(imgCf.ToFile, imgCf.DecodedFile, nil)
	}
	return nil
}
