package imgconv_test

import (
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/gopherdojo/dojo7/kadai2/ayatothos/imgconv"
)

var getExtentionByNameTests = []struct {
	name   string
	result bool
}{
	// 正常系
	{"jpg", true},
	{"png", true},
	{"gif", true},
	//異常系
	{"bmp", false},
}

var convertImageTests = []struct {
	srcPath  string
	destPath string
	toType   string
	result   bool
}{
	// 正常系
	{"../testdata/input/neko_jpg.jpg", "../testdata/output/neko_jpg.jpg", "jpg", true},
	{"../testdata/input/neko_jpg.jpg", "../testdata/output/neko_jpg.png", "png", true},
	{"../testdata/input/neko_jpg.jpg", "../testdata/output/neko_jpg.gif", "gif", true},
	{"../testdata/input/neko_png.png", "../testdata/output/neko_png.jpg", "jpg", true},
	{"../testdata/input/neko_png.png", "../testdata/output/neko_png.png", "png", true},
	{"../testdata/input/neko_png.png", "../testdata/output/neko_png.gif", "gif", true},
	{"../testdata/input/neko_gif.gif", "../testdata/output/neko_gif.jpg", "jpg", true},
	{"../testdata/input/neko_gif.gif", "../testdata/output/neko_gif.png", "png", true},
	{"../testdata/input/neko_gif.gif", "../testdata/output/neko_gif.gif", "gif", true},
	// 異常系
	{"../testdata/input/hoge.jpg", "../testdata/output/neko_jpg.png", "png", false},
	{"../testdata/input/test.txt", "../testdata/output/neko_jpg.png", "png", false},
	{"../testdata/input/neko_jpg.jpg", "../testdata/output/hoge/hoge", "png", false},
	{"../testdata/input/neko_jpg.jpg", "../testdata/output/neko_jpg.bmp", "bmp", false},
}

var convertImageAllTests = []struct {
	dirrPath string
	fromType string
	toType   string
	result   bool
}{
	// 正常系
	{"../testdata/input/", "jpg", "png", false},
	// 異常系
	{"../testdata/hoge", "jpg", "png", false},
	{"../testdata/input/", "bmp", "png", false},
}

// 初期化用関数
func TestMain(m *testing.M) {

	setup()
	ret := m.Run()
	if ret == 0 {
		teardown()
	}
	os.Exit(ret)
}

// 初期化用関数
func setup() {

	if _, err := os.Stat("../testdata/input"); err != nil {
		if err := os.RemoveAll("../testdata/input"); err != nil {
			fmt.Println("ディレクトリ削除失敗:output")
		}
	}

	if _, err := os.Stat("../testdata/output"); err != nil {
		if err := os.RemoveAll("../testdata/output"); err != nil {
			fmt.Println("ディレクトリ削除失敗:output")
		}
	}

	if err := os.Mkdir("../testdata/input", 0777); err != nil {
		fmt.Println("ディレクトリ生成失敗:input")
	}
	if err := os.Mkdir("../testdata/output", 0777); err != nil {
		fmt.Println("ディレクトリ生成失敗:output")
	}

	if err := copyFile("../testdata/img/neko_jpg.jpg", "../testdata/input/neko_jpg.jpg"); err != nil {
		fmt.Println("ファイルコピー失敗:neko_jpg.jpg")
	}
	if err := copyFile("../testdata/img/neko_png.png", "../testdata/input/neko_png.png"); err != nil {
		fmt.Println("ファイルコピー失敗:neko_png.png")
	}
	if err := copyFile("../testdata/img/neko_gif.gif", "../testdata/input/neko_gif.gif"); err != nil {
		fmt.Println("ファイルコピー失敗:neko_gif.gif")
	}
	if err := copyFile("../testdata/img/test.txt", "../testdata/input/test.txt"); err != nil {
		fmt.Println("ファイルコピー失敗:test.txt")
	}
}

// 初期化用関数
func teardown() {
	if err := os.RemoveAll("../testdata/input"); err != nil {
		fmt.Println("ディレクトリ削除失敗:output")
	}
	if err := os.RemoveAll("../testdata/output"); err != nil {
		fmt.Println("ディレクトリ削除失敗:output")
	}
}

// 初期化用関数
func copyFile(src string, dst string) error {
	srcFp, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFp.Close()

	dstFp, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFp.Close()

	_, err = io.Copy(dstFp, srcFp)
	return err
}

func TestGetExtentionsByName(t *testing.T) {
	for _, v := range getExtentionByNameTests {
		_, err := imgconv.GetExtentionsByName(v.name)
		checkError(t, err, v.result, v)
	}
}

func TestConvertImage(t *testing.T) {
	for _, v := range convertImageTests {
		err := imgconv.ConvertImage(v.srcPath, v.destPath, v.toType)
		checkError(t, err, v.result, v)
	}
}

func TestConvertImageAll(t *testing.T) {
	for _, v := range convertImageAllTests {
		_, err := imgconv.ConvertImageAll(v.dirrPath, v.fromType, v.toType)
		checkError(t, err, v.result, v)
	}

}

func checkError(t *testing.T, err error, result bool, vals interface{}) error {
	t.Helper()
	if err != nil && result {
		t.Fatalf("%v", vals)
	}
	return nil
}
