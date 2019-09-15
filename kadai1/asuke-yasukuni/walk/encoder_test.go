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

	testCases := []struct {
		Name  string
		From  string
		To    string
		Files []string
	}{
		{
			Name: "walk jpg -> png",
			From: "jpg",
			To:   "png",
			Files: []string{
				"[replace file]../testdata/recursiondata/test-1.jpg -> png",
				"[replace file]../testdata/recursiondata/test-2.jpg -> png",
				"[replace file]../testdata/test-1.jpg -> png",
				"[replace file]../testdata/test-2.jpg -> png",
			},
		},
		{
			Name: "walk png -> jpg",
			From: "png",
			To:   "jpg",
			Files: []string{
				"[replace file]../testdata/recursiondata/test-1.png -> jpg",
				"[replace file]../testdata/recursiondata/test-2.png -> jpg",
				"[replace file]../testdata/test-1.png -> jpg",
				"[replace file]../testdata/test-2.png -> jpg",
			},
		},
	}

	walker := Walk{File: &testFile{}}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			files, err := walker.Encoder(&src, tc.From, tc.To)
			if err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(tc.Files, files) {
				t.Fatal(files)
			}
		})
	}
}
