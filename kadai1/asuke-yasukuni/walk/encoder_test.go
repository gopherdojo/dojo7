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

	src := "../testdata/multiple_replace/"

	testCase := []struct {
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
				"[replace file]../testdata/multiple_replace/recursiondata/test-1.jpg -> png",
				"[replace file]../testdata/multiple_replace/recursiondata/test-2.jpg -> png",
				"[replace file]../testdata/multiple_replace/test-1.jpg -> png",
				"[replace file]../testdata/multiple_replace/test-2.jpg -> png",
			},
		},
		{
			Name: "walk png -> jpg",
			From: "png",
			To:   "jpg",
			Files: []string{
				"[replace file]../testdata/multiple_replace/recursiondata/test-1.png -> jpg",
				"[replace file]../testdata/multiple_replace/recursiondata/test-2.png -> jpg",
				"[replace file]../testdata/multiple_replace/test-1.png -> jpg",
				"[replace file]../testdata/multiple_replace/test-2.png -> jpg",
			},
		},
	}

	walker := Walk{File: &testFile{}}
	for _, tc := range testCase {
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
