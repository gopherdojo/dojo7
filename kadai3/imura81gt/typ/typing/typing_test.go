package typing

import (
	"bytes"
	"fmt"
	"testing"
	"time"
)

func TestInput(t *testing.T) {
	testcases := []struct {
		caseName string
		input    string
		expected string
	}{
		{caseName: "input some word", input: "foo bar", expected: "foo bar"},
		{caseName: "input line feed", input: "\n", expected: ""},
	}
	fmt.Println(testcases)
	for _, tc := range testcases {
		tc := tc // capture range variable. need to set when run parallel test.

		t.Run(tc.caseName, func(t *testing.T) {
			t.Parallel()

			var buf bytes.Buffer
			buf.Write([]byte(tc.input))
			actual := <-input(&buf)

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

func TestLoad(t *testing.T) {
	g := Game{
		Clock: ClockFunc(func() time.Time {
			return time.Date(2019, 11, 04, 02, 0, 0, 0, time.UTC)
		}),
	}

	expected := "すもももももももものうち"
	actual := g.load()
	if actual != expected {
		t.Errorf("\ncaseName:%s\nactual:%+v\nExpected:%+v\n",
			"load test",
			actual,
			expected,
		)
	}
}

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
