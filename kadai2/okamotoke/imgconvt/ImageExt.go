package imgconvt

// ImageExt is image extension
type ImageExt string

// Supported image extension
const (
	ImageExtJpg  ImageExt = ".jpg"  // jpg
	ImageExtJpeg          = ".jpeg" // jepg
	ImageExtPng           = ".png"  // png
	ImageExtGif           = ".gif"  // gif
)

func (i ImageExt) string() string {
	return string(i)
}
