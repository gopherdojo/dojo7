package replacer

import (
	"testing"
)

func TestEncode(t *testing.T) {

	var testCase = []struct {
		Name string
		Src  string
		To   string
	}{
		{"jpg -> png encode", "../testdata/test-single.jpg", "png"},
		{"png -> jpg encode", "../testdata/test-single.png", "jpg"},
	}

	f := File{}
	for _, tc := range testCase {
		t.Run(tc.Name, func(t *testing.T) {
			err := f.Encode(tc.Src, tc.To)
			if err != nil {
				t.Error(err)
			}
		})
	}
}
