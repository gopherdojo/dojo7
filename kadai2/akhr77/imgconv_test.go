package imgconv

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	tests := []struct {
		desc     string
		arg      string
		expected int
	}{
		{
			desc:     "正常",
			arg:      "bin/imgconv -s testdata",
			expected: ExitCodeOK,
		},
	}

	t.Helper()
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{OutStream: outStream, ErrStream: errStream}
	for _, test := range tests {
		args := strings.Split(test.arg, " ")
		fmt.Printf("arg=%s\n", args)
		status := cli.Run(args)
		if status != test.expected {
			t.Errorf("failed test:%s actual:%d", test.desc, status)
		}
	}
}

func TestWalk(t *testing.T) {
	t.Helper()
	tests := []struct {
		desc      string
		src       string
		beforeExt string
		afterExt  string
		expected  int
	}{
		{
			desc:      "異常：サポート外変換（変換元）",
			src:       "testdata/",
			beforeExt: "tif",
			afterExt:  "gif",
			expected:  ExitCodeError,
		},
	}
	for _, test := range tests {
		err := walk(test.src, test.beforeExt, test.afterExt)
		if err != nil {
			t.Errorf("failed test:%s", test.desc)
		}
	}
}

func TestSupportFormat(t *testing.T) {
	t.Helper()
	tests := []struct {
		extention string
		expected  bool
	}{
		{
			extention: "jpg",
			expected:  true,
		},
		{
			extention: "gif",
			expected:  true,
		},
		{
			extention: "jpeg",
			expected:  true,
		},
		{
			extention: "tiff",
			expected:  false,
		},
	}
	for _, test := range tests {
		actual := supportFormat(test.extention)
		if actual != test.expected {
			t.Errorf("failed test extention:%s", test.extention)
		}
	}
}

func TestConvert(t *testing.T) {
	t.Helper()
}
