package word_test

import (
	"github.com/gopherdojo/dojo7/kadai3-1/kmd2kmd/word"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"syscall"
	"testing"
)

var (
	errPermission string
	isDir string
)

func init() {
	dir, err := ioutil.TempDir("", "typing")
	if err != nil {
		log.Fatal(err)
	}
	file, err := ioutil.TempFile(dir, "ptn")
	if err != nil {
		log.Fatal(err)
	}
	err = file.Chmod(000)
	if err != nil {
		log.Fatal(err)
	}
	isDir = dir
	errPermission = filepath.Join(dir, file.Name())
}

func TestWord_Read(t *testing.T) {
	cases := []struct{
		name string
		path string
		want word.Words
	}{
		{
			name: "normal",
			path: "./testdata/words.txt",
			want: word.Words{"1", "12", "あ", "あい", "a", "ab"},
		},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			got, err := word.Read(c.path)
			// 正常系 エラーチェック
			if err != nil {
				t.Fatal("want no err, but has error", err)
			}
			// 正常系 レスポンスチェック
			if !reflect.DeepEqual(got, c.want) {
				t.Fatalf("got: %v, want: %v", got, c.want)
			}
		})
	}
}

func TestWord_Err_Read(t *testing.T) {
	cases := []struct{
		name string
		path string
		err error
	}{
		{
			name: "Not exist",
			path: "not_found.txt",
			err: &os.PathError{Op: "open", Path: "not_found.txt", Err: syscall.Errno(syscall.ENOENT)},
		},
		{
			name: "Err permission",
			path: errPermission,
			err: &os.PathError{Op: "open", Path: errPermission, Err: syscall.Errno(syscall.ENOENT)},
		},
		{
			name: "Is dir",
			path: isDir,
			err: &os.PathError{Op: "read", Path: isDir, Err: syscall.Errno(syscall.EISDIR)},
		},

	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			_, err := word.Read(c.path)
			// 異常系 エラーチェック
			if err == nil {
				t.Fatal("want err, but has't error")
			}
			// 異常系 レスポンスチェック
			if !reflect.DeepEqual(err, c.err) {
				t.Fatalf("got: %v, want: %v", err, c.err)
			}
		})
	}
}

func contains(t *testing.T, s []string, e string) bool {
	t.Helper()
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}


func TestWords_Random(t *testing.T) {
	cases := []struct{
		name string
		word word.Words
	}{
		{
			name: "normal",
			word: word.Words{"1", "12", "あ", "あい", "a", "ab"},
		},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			got := c.word.Random()
			// 正常系 レスポンスチェック
			// TODO: ランダムに文字列が返却されるので､テストがランダムで好ましくない？
			if !contains(t, c.word, got) {
				t.Fatal("want contains", got, ", but not contains")
			}
		})
	}
}