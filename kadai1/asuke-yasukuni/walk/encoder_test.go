package walk

import (
	"reflect"
	"testing"
)

type testFile struct{}

func (t *testFile) Encode(path, toExt string) error {
	return nil
}

func TestWalkEncoder(t *testing.T) {

	src := "../testdata/"

	testFiles := []string{
		"[replace file]../testdata/recursiondata/test-1.jpg -> png",
		"[replace file]../testdata/recursiondata/test-2.jpg -> png",
		"[replace file]../testdata/recursiondata/test-3.jpg -> png",
		"[replace file]../testdata/recursiondata/test-4.jpg -> png",
		"[replace file]../testdata/recursiondata/test-5.jpg -> png",
		"[replace file]../testdata/recursiondata/test-6.jpg -> png",
		"[replace file]../testdata/recursiondata/test-7.jpg -> png",
		"[replace file]../testdata/test-1.jpg -> png",
		"[replace file]../testdata/test-2.jpg -> png",
		"[replace file]../testdata/test-3.jpg -> png",
		"[replace file]../testdata/test-4.jpg -> png",
		"[replace file]../testdata/test-5.jpg -> png",
		"[replace file]../testdata/test-6.jpg -> png",
		"[replace file]../testdata/test-7.jpg -> png",
	}

	//todo: 現状だと png -> jpg ができないのでtestdate追加する
	testCases := []struct {
		Name   string
		From   string
		To     string
		Result bool
	}{
		{Name: "walk jpg -> png", From: "jpg", To: "png"},
	}

	walker := Walk{File: &testFile{}}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			files, err := walker.Encoder(&src, tc.From, tc.To)
			if err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(testFiles, files) {
				t.Fatal(files)
			}
		})
	}
}
