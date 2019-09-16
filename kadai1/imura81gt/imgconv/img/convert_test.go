package img

import (
	"reflect"
	"testing"
)

func TestEnable(t *testing.T) {
	testcases := []struct {
		caseName string
		iType    ImageType
		expected []string
	}{
		{
			caseName: "Png: true",
			iType:    ImageType{Png: true, Jpeg: false, Gif: false},
			expected: []string{"png"},
		},
		{
			caseName: "Jpeg: true",
			iType:    ImageType{Png: false, Jpeg: true, Gif: false},
			expected: []string{"jpeg"},
		},
		{
			caseName: "Gif: true",
			iType:    ImageType{Png: false, Jpeg: false, Gif: true},
			expected: []string{"gif"},
		},
		{
			caseName: "Png: true, Jpeg: true",
			iType:    ImageType{Png: true, Jpeg: true, Gif: false},
			expected: []string{"png", "jpeg"},
		},
		{
			caseName: "Jpeg: true, Gif: true",
			iType:    ImageType{Png: false, Jpeg: true, Gif: true},
			expected: []string{"jpeg", "gif"},
		},
		{
			caseName: "Png: true, Jpeg: true, Gif: true",
			iType:    ImageType{Png: true, Jpeg: true, Gif: true},
			expected: []string{"png", "jpeg", "gif"},
		},
	}

	for _, tc := range testcases {
		tc := tc // capture range variable. need to set when run parallel test.

		t.Run(tc.caseName, func(t *testing.T) {
			t.Parallel()
			actual := tc.iType.Enable()
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("\ncaseName:%s\nactual:%+v\niExpected:%+v\n",
					tc.caseName,
					actual,
					tc.expected,
				)
			}
		})
	}
}

func TestAllImageFiles(t *testing.T) {

}

func TestExpectedImageFiles(t *testing.T) {

}

func TestIsImage(t *testing.T) {
	testcases := []struct {
		caseName string
		path     string
		expected bool
	}{
		{
			caseName: "gif file is image file.",
			path:     "../testdata/gLenna.gif",
			expected: true,
		},
		{
			caseName: "jpeg file is image file.",
			path:     "../testdata/jLenna.jpg",
			expected: true,
		},
		{
			caseName: "png file is image file.",
			path:     "../testdata/pLenna.png",
			expected: true,
		},
		{
			caseName: "empty file is not image file",
			path:     "../testdata/this_is_not_image.gif",
			expected: false,
		},
		{
			caseName: "does not existed.",
			path:     "../testdata/does_not_existed.gif",
			expected: false,
		},
	}

	for _, tc := range testcases {
		tc := tc // capture range variable. need to set when run parallel test.

		t.Run(tc.caseName, func(t *testing.T) {
			t.Parallel()
			actual := IsImage(tc.path)
			if !(actual == tc.expected) {
				t.Errorf("\ncaseName:%s\nactual:%+v\niExpected:%+v\n",
					tc.caseName,
					actual,
					tc.expected,
				)
			}
		})
	}

}

func TestIsExpectedImage(t *testing.T) {

}

func TestConvert(t *testing.T) {

}
