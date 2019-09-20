package imageconversion_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/gopherdojo/dojo7/kadai2/shinta/imageconversion"
)

type exampleArg struct {
	testCase string
	dir      string
	preExt   string
	afterExt string
}

type expectedImageFile struct {
	path string
	base string
	ext  string
}

func (a *exampleArg) valid() error {
	if a.preExt == a.afterExt {
		return errors.New("変換前と変換後で拡張子が同じです。")
	}
	allowExtList := []string{"jpg", "jpeg", "png", "gif"}
	allowExtMap := map[string]bool{}
	for _, ext := range allowExtList {
		allowExtMap[ext] = true
	}
	if !allowExtMap[a.preExt] || !allowExtMap[a.afterExt] {
		return errors.New("指定できる拡張子: " + strings.Join(allowExtList, ","))
	}
	return nil
}

func TestValid(t *testing.T) {
	var exampleArgs = []exampleArg{
		{"カレントディレクトリ以下のjpeg=>png", "./", "jpeg", "png"},
		{"カレントディレクトリ以下のpng=>gif", "./", "png", "gif"},
		{"カレントディレクトリ以下のpng=>jpeg", "./", "png", "jpeg"},
	}
	for _, arg := range exampleArgs {
		if err := arg.valid(); err != nil {
			t.Error("failed to call Imageconversion valid", err, "expected: nil")
		}
	}
	// t.Skip("func valid")
}

func TestGetFileNamaWithoutExt(t *testing.T) {
	var testDatas = []string{"testdata/1.jpeg", "testdata/sub/3.jpeg"}
	var expectedImageFiles = []expectedImageFile{
		{"testdata", "1", ".jpeg"}, {"testdata/sub", "3", ".jpeg"},
	}
	for i, data := range testDatas {
		res := imageconversion.GetFileNameWithoutExt(data)
		if res != expectedImageFiles[i].base {
			t.Error("failed to call Imageconversion getFileNameWithoutExt", res, "expected", expectedImageFiles[i].base)
		}
	}
	// t.Skip("func getFileNameWithoutExt")
}
func TestCreateImgStruct(t *testing.T) {
	var testDatas = []string{"testdata/1.jpeg", "testdata/sub/3.jpeg"}
	var expectedImageFiles = []expectedImageFile{
		{"testdata", "1", ".jpeg"}, {"testdata/sub", "3", ".jpeg"},
	}
	for i, data := range testDatas {
		img := imageconversion.CreateImgStruct(data)
		if img.Path != expectedImageFiles[i].path || img.Base != expectedImageFiles[i].base || img.Ext != expectedImageFiles[i].ext {
			t.Error("failed to call Imageconversion createImgStruct", img, "expected:", expectedImageFiles[i])
		}
	}
	// t.Skip("func createImgStruct")
}
func TestConvertExec(t *testing.T) {
	var convertedExtArgs = []exampleArg{
		{"カレントディレクトリ以下のjpeg=>png", "./", ".jpeg", ".png"},
		{"カレントディレクトリ以下のpng=>gif", "./", ".png", ".gif"},
		{"カレントディレクトリ以下のpng=>jpeg", "./", ".png", ".jpeg"},
	}
	var testDatas = []string{"testdata/1.jpeg", "testdata/sub/3.jpeg"}
	for _, arg := range convertedExtArgs {
		for _, data := range testDatas {
			err := imageconversion.ConvertExec(data, arg.afterExt)
			if err != nil {
				t.Error("failed to call Imageconversion ConvertExcute", err, "expected: nil")
			}
		}
	}
	// t.Skip("func convertExec")
}
func TestImageConversionExcute(t *testing.T) {
	var exampleArgs = []exampleArg{
		{"カレントディレクトリ以下のjpeg=>png", "./", "jpeg", "png"},
		{"カレントディレクトリ以下のpng=>gif", "./", "png", "gif"},
		{"カレントディレクトリ以下のpng=>jpeg", "./", "png", "jpeg"},
	}
	for _, arg := range exampleArgs {
		err := imageconversion.Excute(arg.dir, arg.preExt, arg.afterExt)
		if err != nil {
			t.Error("failed to call Imageconversion Excute", err, "expected: nil")
		}
	}
	// t.Skip("Execute")
}
