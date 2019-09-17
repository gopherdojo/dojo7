package imageconversion

import (
	"errors"
	"strings"
	"testing"
)

// os でdestdataディレクトリから取得するべきか
var testDatas = []string{"testdata/1.jpeg", "testdata/sub/3.jpeg"}

type exampleArg struct {
	testCase string
	dir      string
	preExt   string
	afterExt string
}

var exampleArgs = []exampleArg{
	{"カレントディレクトリ以下のjpeg=>png", "./", "jpeg", "png"},
	{"カレントディレクトリ以下のpng=>gif", "./", "png", "gif"},
	{"カレントディレクトリ以下のpng=>jpeg", "./", "png", "jpeg"},
}
var convertedExtArgs = []exampleArg{
	{"カレントディレクトリ以下のjpeg=>png", "./", ".jpeg", ".png"},
	{"カレントディレクトリ以下のpng=>gif", "./", ".png", ".gif"},
	{"カレントディレクトリ以下のpng=>jpeg", "./", ".png", ".jpeg"},
}

type expectedImageFile struct {
	path string
	base string
	ext  string
}

var expectedImageFiles = []expectedImageFile{
	{"testdata", "1", ".jpeg"}, {"testdata/sub", "3", ".jpeg"},
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

func (a *exampleArg) convertExt() {
	a.preExt, a.afterExt = "."+a.preExt, "."+a.afterExt
}

func TestValid(t *testing.T) {
	t.Helper()
	for _, arg := range exampleArgs {
		t.Run(arg.testCase, func(t *testing.T) {
			if err := arg.valid(); err != nil {
				t.Error("failed to call Imageconversion valid", err, "expected: nil")
			}
		})
	}
	// t.Skip("func valid")
}

func TestConvertExt(t *testing.T) {
	t.Helper()
	for i, arg := range exampleArgs {
		t.Run(arg.testCase, func(t *testing.T) {
			arg.convertExt()
			if arg.preExt != convertedExtArgs[i].preExt || arg.afterExt != convertedExtArgs[i].afterExt {
				t.Error("failed to call Imageconversion convertExt", arg.preExt, arg.afterExt, "expected:", convertedExtArgs[i].preExt, convertedExtArgs[i].afterExt)
			}
		})
	}
	// t.Skip("func convertExt")
}

func TestGetFileNamaWithoutExt(t *testing.T) {
	t.Helper()
	for i, data := range testDatas {
		t.Run("testdata:"+data, func(t *testing.T) {
			res := getFileNameWithoutExt(data)
			if res != expectedImageFiles[i].base {
				t.Error("failed to call Imageconversion getFileNameWithoutExt", res, "expected", expectedImageFiles[i].base)
			}
		})
	}
	// t.Skip("func getFileNameWithoutExt")
}
func TestCreateImgStruct(t *testing.T) {
	t.Helper()
	for i, data := range testDatas {
		t.Run("testdata:"+data, func(t *testing.T) {
			img := createImgStruct(data)
			if img.path != expectedImageFiles[i].path || img.base != expectedImageFiles[i].base || img.ext != expectedImageFiles[i].ext {
				t.Error("failed to call Imageconversion createImgStruct", img, "expected:", expectedImageFiles[i])
			}
		})
	}
	// t.Skip("func createImgStruct")
}
func TestConvertExec(t *testing.T) {
	t.Helper()
	for _, arg := range convertedExtArgs {
		for _, data := range testDatas {
			t.Run("exampleArg:"+arg.afterExt+"testdata:"+data, func(t *testing.T) {
				err := convertExec(data, arg.afterExt)
				if err != nil {
					t.Error("failed to call Imageconversion ConvertExcute", err, "expected: nil")
				}
			})
		}
	}
	// t.Skip("func convertExec")
}
func TestImageConversionExcute(t *testing.T) {
	t.Helper()
	for _, arg := range exampleArgs {
		t.Run("exampleArg:"+arg.testCase, func(t *testing.T) {
			err := Excute(arg.dir, arg.preExt, arg.afterExt)
			if err != nil {
				t.Error("failed to call Imageconversion Excute", err, "expected: nil")
			}
		})
	}
	// t.Skip("Execute")
}
