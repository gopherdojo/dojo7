package typing

import (
	"bytes"
	"fmt"
	"testing"
)

func TestInput(t *testing.T) {}

func TestLoad(t *testing.T) {}

func TestShow(t *testing.T) {
	testcases := []struct {
		caseName string
		score    int
		char     int
		txt      string
		expected string
	}{
		{caseName: "score10char300", score: 10, char: 300, txt: "テスト", expected: "10 300 > テスト\n10 300 > "},
	}

	for _, tc := range testcases {
		tc := tc // capture range variable. need to set when run parallel test.

		fmt.Println(tc)
		t.Run(tc.caseName, func(t *testing.T) {
			t.Parallel()
			var buf bytes.Buffer
			show(tc.score, tc.char, tc.txt, &buf)
			actual := buf.String()
			if tc.expected != actual {
				t.Errorf("\ncaseName:%s\nactual:%+v\nExpected:%+v\n",
					tc.caseName,
					actual,
					tc.expected,
				)
			}
		})
	}

}

func TestRun(t *testing.T) {}
