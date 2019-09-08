package validation

func Ext(ext string) bool {
	if ext == "jpg" || ext == "png" {
		return true
	}
	return false
}
