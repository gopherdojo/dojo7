package imageconv

import (
	"os"
	"testing"
)

func setupTest() {
	os.Remove("testdata/hoge.png")
	os.Remove("testdata/fuga.png")
}

func TestConvertRecursively(t *testing.T) {
	testData := []struct {
		targetDir string
		srcExt    string
		dstExt    string
		expLen    int
	}{
		{"testdata", ".jpg", ".png", 2},
	}

	setupTest()

	for _, td := range testData {
		result := ConvertRecursively(td.targetDir, td.srcExt, td.dstExt)
		if len(result) != td.expLen {
			t.Errorf("expect %v but actual %v", td.expLen, len(result))
		}
	}
}

func TestConvert(t *testing.T) {
	testData := []struct {
		src string
		dst string
	}{
		{"testdata/hoge.jpg", "testdata/fuga.png"},
		{"testdata/hoge.gif", "testdata/fuga.png"},
	}

	for _, td := range testData {
		err := Convert(td.src, td.dst)
		if err != nil {
			t.Error(err)
		}
	}
}
