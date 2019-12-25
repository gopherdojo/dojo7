package rget

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestWrite(t *testing.T) {
}

func TestRun(t *testing.T) {
}

func TestCheckingHeaders(t *testing.T) {
	testCases := []struct {
		caseName     string
		acceptRanges string
		body         string
		isErr        bool
		expected     string
	}{
		{caseName: "acceptRanges:none", acceptRanges: "none", body: "1", isErr: true, expected: "cannot support Ranges Requests"},
		{caseName: "acceptRanges:(empty)", acceptRanges: "", body: "1", isErr: true, expected: "cannot support Ranges Requests"},
		{caseName: "acceptRanges:bytes", acceptRanges: "bytes", body: "1", isErr: false},
		{caseName: "acceptRanges:bytes but content is empty", acceptRanges: "bytes", body: "", isErr: true, expected: "size is nil"},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.caseName, func(t *testing.T) {
			t.Parallel()

			ts := SetupHTTPServer(t, tc.acceptRanges, tc.body)
			defer ts.Close()

			// resp, _ := http.Get(ts.URL)
			// t.Logf("resp: %+v", resp)

			resp, err := http.Head(ts.URL)
			if err != nil {
				t.Fatal(err)
			}
			t.Logf("ContentLength: %+v,Accept-Ranges: %+v", resp.ContentLength, resp.Header.Get("Accept-Ranges"))

			o := Option{URL: ts.URL}
			exerr := o.checkingHeaders()
			if tc.isErr && exerr == nil {
				t.Errorf("tc.isErr %+v but err is %+v", tc.isErr, exerr)
			}
			if tc.isErr && exerr != nil && !strings.Contains(exerr.Error(), tc.expected) {
				t.Errorf("actual: %+v\nexpected: %+v\n", exerr, tc.isErr)
			}
		})
	}

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
			caseName:    "ContentLength and Concurrency is same value",
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
			caseName:    "One Thread",
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
			caseName:    "Remainder:ContentLength%Concurrency!=0",
			concurrency: 2,
			expected: Option{
				URL:           url,
				ContentLength: contentLength,
				Concurrency:   2,
				Units: []Unit{
					{TempFileName: "0_test.iso", RangeStart: 0, RangeEnd: 1},
					{TempFileName: "1_test.iso", RangeStart: 2, RangeEnd: 4},
				},
			},
		},
		{
			caseName:    "Concurrency exceed the contentLength.",
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

	type Expected []struct {
		TempFileName string
		Text         string
	}

	testCases := []struct {
		caseName     string
		acceptRanges string
		body         string
		option       Option
		expected     Expected
	}{
		{
			caseName: "acceptRanges:bytes per 1byte", acceptRanges: "bytes", body: "12345",
			option: Option{
				Units: []Unit{
					{TempFileName: "0_test.iso", RangeStart: 0, RangeEnd: 0},
					{TempFileName: "1_test.iso", RangeStart: 1, RangeEnd: 1},
					{TempFileName: "2_test.iso", RangeStart: 2, RangeEnd: 2},
					{TempFileName: "3_test.iso", RangeStart: 3, RangeEnd: 3},
					{TempFileName: "4_test.iso", RangeStart: 4, RangeEnd: 4},
				},
			},
			expected: Expected{
				{TempFileName: "0_test.iso", Text: "1"},
				{TempFileName: "1_test.iso", Text: "2"},
				{TempFileName: "2_test.iso", Text: "3"},
				{TempFileName: "3_test.iso", Text: "4"},
				{TempFileName: "4_test.iso", Text: "5"},
			},
		},
		{
			caseName: "acceptRanges:bytes per 2bytes", acceptRanges: "bytes", body: "0123456789",
			option: Option{
				Units: []Unit{
					{TempFileName: "0_test.iso", RangeStart: 0, RangeEnd: 1},
					{TempFileName: "1_test.iso", RangeStart: 2, RangeEnd: 3},
					{TempFileName: "2_test.iso", RangeStart: 4, RangeEnd: 5},
					{TempFileName: "3_test.iso", RangeStart: 6, RangeEnd: 7},
					{TempFileName: "4_test.iso", RangeStart: 8, RangeEnd: 9},
				},
			},
			expected: Expected{
				{TempFileName: "0_test.iso", Text: "01"},
				{TempFileName: "1_test.iso", Text: "23"},
				{TempFileName: "2_test.iso", Text: "45"},
				{TempFileName: "3_test.iso", Text: "67"},
				{TempFileName: "4_test.iso", Text: "89"},
			},
		},
		{
			caseName: "acceptRanges:bytes per 2bytes+1", acceptRanges: "bytes", body: "01234567890",
			option: Option{
				Units: []Unit{
					{TempFileName: "0_test.iso", RangeStart: 0, RangeEnd: 1},
					{TempFileName: "1_test.iso", RangeStart: 2, RangeEnd: 3},
					{TempFileName: "2_test.iso", RangeStart: 4, RangeEnd: 5},
					{TempFileName: "3_test.iso", RangeStart: 6, RangeEnd: 7},
					{TempFileName: "4_test.iso", RangeStart: 8, RangeEnd: 10},
				},
			},
			expected: Expected{
				{TempFileName: "0_test.iso", Text: "01"},
				{TempFileName: "1_test.iso", Text: "23"},
				{TempFileName: "2_test.iso", Text: "45"},
				{TempFileName: "3_test.iso", Text: "67"},
				{TempFileName: "4_test.iso", Text: "890"},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.caseName, func(t *testing.T) {
			t.Parallel()

			prefix := "rget_test"
			tmpDir, err := ioutil.TempDir("", prefix)
			if err != nil {
				t.Fatal(err)
			}

			ts := SetupHTTPServer(t, tc.acceptRanges, tc.body)
			defer ts.Close()

			tc.option.URL = ts.URL
			resp, err := http.Head(tc.option.URL)
			if err != nil {
				t.Fatal(err)
			}

			tc.option.ContentLength = resp.ContentLength

			err = tc.option.parallelDownload(tmpDir)
			if err != nil {
				t.Fatal(err)
			}

			for _, ex := range tc.expected {
				f, err := os.Open(filepath.Join(tmpDir, ex.TempFileName))
				if err != nil {
					t.Fatal(err)
				}

				actual, err := ioutil.ReadAll(f)
				if ex.Text != string(actual) {
					t.Errorf("actual: %+v\nexpected: %+v\n", string(actual), ex.Text)
				}

			}
			// t.Log(tmpDir)
			os.RemoveAll(tmpDir)

		})
	}

}

func TestDownloadWithContext(t *testing.T) {
}

func TestCombine(t *testing.T) {
}

func SetupHTTPServer(t *testing.T, ac string, body string) *httptest.Server {
	t.Helper()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// the unsupported server for Range request
		if ac != "bytes" {
			w.Header().Set("Accept-Ranges", ac)
			fmt.Fprint(w, body)
		} else {
			// the supported server for Range request
			http.ServeContent(w, r, "", time.Unix(0, 0), strings.NewReader(body))
		}

	}))
	return ts
}
