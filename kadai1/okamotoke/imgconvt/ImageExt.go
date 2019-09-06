package imgconvt

type ImageExt string

const (
	ImageExtJpg  ImageExt = ".jpg"
	ImageExtJpeg          = ".jpeg"
	ImageExtPng           = ".png"
	ImageExtGif           = ".gif"
)

func (i ImageExt) string() string {
	return string(i)
}
