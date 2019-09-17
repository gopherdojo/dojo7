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
				t.Errorf("\ncaseName:%s\nactual:%+v\nExpected:%+v\n",
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
	testcases := []struct {
		caseName string
		files    []string
		iType    ImageType
		expected []string
	}{
		{
			caseName: "png is expected image type.",
			files: []string{
				"../testdata/pLenna.png",
				"../testdata/jLenna.jpg",
				"../testdata/gLenna.gif",
				"../testdata/chdir/pLenna.png",
				"../testdata/chdir/jLenna.jpeg",
				"../testdata/chdir/gLenna.gif",
				"../testdata/this_is_not_image.png",
				"../testdata/this_is_not_image.jpg",
				"../testdata/this_is_not_image.gif",
			},
			iType: ImageType{Png: true, Jpeg: false, Gif: false},
			expected: []string{
				"../testdata/pLenna.png",
				"../testdata/chdir/pLenna.png",
			},
		},
		{
			caseName: "jpg is expected image type.",
			files: []string{
				"../testdata/pLenna.png",
				"../testdata/jLenna.jpg",
				"../testdata/gLenna.gif",
				"../testdata/chdir/pLenna.png",
				"../testdata/chdir/jLenna.jpeg",
				"../testdata/chdir/gLenna.gif",
				"../testdata/this_is_not_image.png",
				"../testdata/this_is_not_image.jpg",
				"../testdata/this_is_not_image.gif",
			},
			iType: ImageType{Png: false, Jpeg: true, Gif: false},
			expected: []string{
				"../testdata/jLenna.jpg",
				"../testdata/chdir/jLenna.jpeg",
			},
		},
		{
			caseName: "gif is expected image type.",
			files: []string{
				"../testdata/pLenna.png",
				"../testdata/jLenna.jpg",
				"../testdata/gLenna.gif",
				"../testdata/chdir/pLenna.png",
				"../testdata/chdir/jLenna.jpeg",
				"../testdata/chdir/gLenna.gif",
				"../testdata/this_is_not_image.png",
				"../testdata/this_is_not_image.jpg",
				"../testdata/this_is_not_image.gif",
			},
			iType: ImageType{Png: false, Jpeg: false, Gif: true},
			expected: []string{
				"../testdata/gLenna.gif",
				"../testdata/chdir/gLenna.gif",
			},
		},
	}

	for _, tc := range testcases {
		tc := tc // capture range variable. need to set when run parallel test.

		t.Run(tc.caseName, func(t *testing.T) {
			t.Parallel()
			actual := ExpectedImageFiles(tc.files, tc.iType)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("\ncaseName:%s\nactual:%+v\nExpected:%+v\n",
					tc.caseName,
					actual,
					tc.expected,
				)
			}
		})
	}

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
				t.Errorf("\ncaseName:%s\nactual:%+v\nExpected:%+v\n",
					tc.caseName,
					actual,
					tc.expected,
				)
			}
		})
	}

}

func TestIsExpectedImage(t *testing.T) {
	testcases := []struct {
		caseName string
		path     string
		iType    ImageType
		expected bool
	}{
		{
			caseName: "png file is expected image file.",
			path:     "../testdata/pLenna.png",
			iType:    ImageType{Png: true, Jpeg: false, Gif: false},
			expected: true,
		},
		{
			caseName: "jpeg file is expected image file.",
			path:     "../testdata/jLenna.jpg",
			iType:    ImageType{Png: false, Jpeg: true, Gif: false},
			expected: true,
		},
		{
			caseName: "gif file is expected image file.",
			path:     "../testdata/gLenna.gif",
			iType:    ImageType{Png: false, Jpeg: false, Gif: true},
			expected: true,
		},

		{
			caseName: "png file is unexpected image file.",
			path:     "../testdata/pLenna.png",
			iType:    ImageType{Png: false, Jpeg: true, Gif: true},
			expected: false,
		},
		{
			caseName: "jpeg file is unexpected image file.",
			path:     "../testdata/jLenna.jpg",
			iType:    ImageType{Png: true, Jpeg: false, Gif: true},
			expected: false,
		},
		{
			caseName: "gif file is unexpected image file.",
			path:     "../testdata/gLenna.gif",
			iType:    ImageType{Png: true, Jpeg: true, Gif: false},
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
			actual := IsExpectedImage(tc.path, tc.iType)
			if !(actual == tc.expected) {
				t.Errorf("\ncaseName:%s\nactual:%+v\nExpected:%+v\n",
					tc.caseName,
					actual,
					tc.expected,
				)
			}
		})
	}

}

func TestConvert(t *testing.T) {

}
