package imageconversion

import (
	"testing"
)

func TestImageConversionExcute(t *testing.T) {
	dir, preExt, afterExt := "./", "jpg", "png"
	err := Excute(dir, preExt, afterExt)
	if err != nil {
		t.Error("failed to call Imageconversion Excute", err)
	}
}
