package imageconversion

import (
	"testing"
)

type exampleArg struct {
	dir      string
	preExt   string
	afterExt string
}

var testDatas = []string{"./testdata/1.jpeg", "./testdata/sub/3.jpeg"}

var exampleArgs = []exampleArg{
	{"./", "jpeg", "png"}, {"./", "png", "gifffff"},
}

func TestImageConversionExcute(t *testing.T) {
	t.Helper()
	for _, arg := range exampleArgs {
		err := Excute(arg.dir, arg.preExt, arg.afterExt)
		if err != nil {
			t.Error("failed to call Imageconversion Excute", err)
			t.Error("exampleArg:", arg.dir, arg.preExt, arg.afterExt)
		}
	}
}

func TestConvertExec(t *testing.T) {
	// for _, arg := range exampleArgs {
	// 	for _, data := range testDatas {
	// 		err := convertExec(data, arg.afterExt)
	// 		if err != nil {
	// 			t.Error("failed to call Imageconversion ConvertExcute", err)
	// 			t.Error("exampleArg:", arg.afterExt)
	// 			t.Error("testData:", data)
	// 		}
	// 	}
	// }
	t.Skip("func convertExec")
}
func TestCreateImgStruct(t *testing.T) {
	t.Skip("func createImgStruct")
}
func TestGetFileNamaWithoutExt(t *testing.T) {
	t.Skip("func getFileNameWithoutExt")
}
func TestConvertExt(t *testing.T) {
	t.Skip("func convertExt")
}
func TestValid(t *testing.T) {
	t.Skip("func valid")
}
