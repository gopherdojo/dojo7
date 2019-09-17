package replacer

import (
	"testing"
)

func TestEncode(t *testing.T) {

	var testCase = []struct {
		Name   string
		Src    string
		To     string
		Result bool
	}{
		{"jpg -> png encode", "../testdata/test-single.jpg", "png", true},
		{"png -> jpg encode", "../testdata/test-single.png", "jpg", true},
		{"jpg -> gif not support encode", "../testdata/test-single.jpg", "gif", false},
		{"not found file encode", "../testdata/test-nofile.gif", "png", false},
	}

	f := File{}
	for _, tc := range testCase {
		t.Run(tc.Name, func(t *testing.T) {
			err := f.Encode(tc.Src, tc.To)
			if err != nil && tc.Result {
				t.Error(err)
			}
		})
	}
}
