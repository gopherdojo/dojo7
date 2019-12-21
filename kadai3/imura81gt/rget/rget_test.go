package rget

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestWrite(t *testing.T) {
}

func TestRun(t *testing.T) {
}

func TestContentLength(t *testing.T) {
}

func TestAcceptRangesHeaderCheck(t *testing.T) {
}

//func divide(contentLength int64, concurrency int) Units {
func TestDivide(t *testing.T) {

	const (
		url           = "https://example.com/test.iso"
		contentLength = 5
	)

	testCases := []struct {
		caseName    string
		concurrency uint
		expected    Option
	}{
		{
			caseName:    "ContentLength:5/5",
			concurrency: 5,
			expected: Option{
				URL:           url,
				ContentLength: contentLength,
				Concurrency:   5,
				Units: []Unit{
					{TempFileName: "0_test.iso", RangeStart: 0, RangeEnd: 0},
					{TempFileName: "1_test.iso", RangeStart: 1, RangeEnd: 1},
					{TempFileName: "2_test.iso", RangeStart: 2, RangeEnd: 2},
					{TempFileName: "3_test.iso", RangeStart: 3, RangeEnd: 3},
					{TempFileName: "4_test.iso", RangeStart: 4, RangeEnd: 4},
				},
			},
		},
		{
			caseName:    "ContentLength:5/1",
			concurrency: 1,
			expected: Option{
				URL:           url,
				ContentLength: contentLength,
				Concurrency:   1,
				Units: []Unit{
					{TempFileName: "0_test.iso", RangeStart: 0, RangeEnd: 4},
				},
			},
		},
		{
			caseName:    "ContentLength:5/10",
			concurrency: 10,
			expected: Option{
				URL:           url,
				ContentLength: contentLength,
				Concurrency:   5,
				Units: []Unit{
					{TempFileName: "0_test.iso", RangeStart: 0, RangeEnd: 0},
					{TempFileName: "1_test.iso", RangeStart: 1, RangeEnd: 1},
					{TempFileName: "2_test.iso", RangeStart: 2, RangeEnd: 2},
					{TempFileName: "3_test.iso", RangeStart: 3, RangeEnd: 3},
					{TempFileName: "4_test.iso", RangeStart: 4, RangeEnd: 4},
				},
			},
		},
		{
			caseName:    "ContentLength:5/0",
			concurrency: 0,
			expected: Option{
				URL:           url,
				ContentLength: contentLength,
				Concurrency:   1,
				Units: []Unit{
					{TempFileName: "0_test.iso", RangeStart: 0, RangeEnd: 4},
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.caseName, func(t *testing.T) {
			t.Parallel()

			o := Option{URL: url, ContentLength: contentLength, Concurrency: tc.concurrency}
			o.divide()

			if !cmp.Equal(o, tc.expected) {
				t.Errorf("actual: %+v\nexpected: %+v\n", o, tc.expected)
			}

		})
	}
}

func TestParallelDownload(t *testing.T) {
}

func TestDownloadWithContext(t *testing.T) {
}

func TestCombine(t *testing.T) {
}
