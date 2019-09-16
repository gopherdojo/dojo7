package image

import (
	"os"
	"testing"
)

func TestDetectImageType(t *testing.T) {
	testData := []struct {
		exp Type
		act Type
	}{
		{Type{"GIF"}, DetectImageType("hoge/fuga.gif")},
		{Type{"JPG"}, DetectImageType("hoge/fuga.jpeg")},
		{Type{"JPG"}, DetectImageType("hoge/fuga.jpg")},
		{Type{"PNG"}, DetectImageType("hoge/fuga.png")},
		{Type{"Unknown"}, DetectImageType("hoge/fuga.tiff")},
	}

	for _, td := range testData {
		if td.exp != td.act {
			t.Errorf("must be %v but %v", td.exp, td.act)
		}
	}
}

func TestDecode(t *testing.T) {
	testData := []struct {
		srcType     Type
		srcFilePath string
	}{
		{Type{"PNG"}, "testdata/hoge.png"},
	}

	for _, td := range testData {
		f, err := os.Open(td.srcFilePath)
		if err != nil {
			t.Error(err)
		}

		_, err = Decode(td.srcType, f)
		if err != nil {
			t.Error(err)
		}
	}

}

func TestEncode(t *testing.T) {
	testData := []struct {
		srcType     Type
		srcImgPath  string
		dstType     Type
		dstFilePath string
	}{
		{Type{"PNG"}, "testdata/hoge.png", Type{"JPG"}, "testdata/fuga.jpg"},
		{Type{"JPG"}, "testdata/hoge.jpeg", Type{"GIF"}, "testdata/fuga.gif"},
	}

	for _, td := range testData {
		srcF, err := os.Open(td.srcImgPath)
		if err != nil {
			t.Error(err)
		}

		srcImg, err := Decode(td.srcType, srcF)
		if err != nil {
			t.Error(err)
		}

		dstF, err := os.Create(td.dstFilePath)
		if err != nil {
			t.Error(err)
		}

		err = Encode(td.dstType, srcImg, dstF)
		if err != nil {
			t.Error(err)
		}
	}
}
