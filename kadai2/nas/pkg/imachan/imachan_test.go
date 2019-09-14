package imachan

import (
	"image"
	"image/jpeg"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

const (
	testdir  = "testdata"
	keepfile = ".gitkeep"
)

func SetupTest(t *testing.T, path string, isDir bool) func() {
	t.Helper()
	switch isDir {
	case true:
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			t.Errorf("Setup Error : %v", err)
		}
	case false:

		f, err := os.Create(path)
		if err != nil {
			t.Errorf("Setup Error : %v", err)
		}
		defer f.Close()

		img := image.NewRGBA(image.Rect(0, 0, 100, 100))
		if err := jpeg.Encode(f, img, &jpeg.Options{Quality: 100}); err != nil {
			t.Errorf("Setup Error : %v", err)
		}
	}

	return func() {
		err := filepath.Walk(testdir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if path == testdir {
				return nil
			}

			if _, fn := filepath.Split(path); fn == keepfile {
				return nil
			}

			if err := os.Remove(path); err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			t.Errorf("Teardown Error : %v", err)
		}
	}
}

func TestNewConfig(t *testing.T) {
	tests := []struct {
		name          string
		fromFormatStr string
		toFormatStr   string
		expected      *Config
	}{
		{
			name:          "PngToJpg",
			fromFormatStr: "png",
			toFormatStr:   "jpg",
			expected: &Config{
				FromFormat: PngFormat,
				ToFormat:   JpgFormat,
			},
		},
		{
			name:          "JpgToPng",
			fromFormatStr: "png",
			toFormatStr:   "jpg",
			expected: &Config{
				FromFormat: PngFormat,
				ToFormat:   JpgFormat,
			},
		},
		{
			name:          "UndefindFromFormat",
			fromFormatStr: "undefind",
			toFormatStr:   "jpg",
			expected:      nil,
		},
		{
			name:          "UndefindToFormat",
			fromFormatStr: "png",
			toFormatStr:   "undefind",
			expected:      nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := NewConfig("", tt.fromFormatStr, tt.toFormatStr)
			if strings.HasPrefix(tt.name, "Undefind") {
				if err == nil {
					t.Errorf("NewConfig(\"\", %s, %s) => %v, want %v", tt.fromFormatStr, tt.toFormatStr, tt.expected, c)
				}
				return
			}
			if c.FromFormat != tt.expected.FromFormat {
				t.Errorf("NewConfig(\"\", %s, %s) => %v, want %v", tt.fromFormatStr, tt.toFormatStr, tt.expected, c)
			}
			if c.ToFormat != tt.expected.ToFormat {
				t.Errorf("NewConfig(\"\", %s, %s) => %v, want %v", tt.fromFormatStr, tt.toFormatStr, tt.expected, c)
			}
		})
	}
}

func TestConfigConvertRec(t *testing.T) {
	tests := []struct {
		name          string
		setupPath     string
		path          string
		fromFormatStr string
		toFormatStr   string
		expected      ConvertedDataRepository
		isDir         bool
		errorExists   bool
	}{
		{
			name:          "PngToJpg",
			setupPath:     filepath.Join(testdir, "test.png"),
			path:          filepath.Join(testdir, "test.png"),
			fromFormatStr: "png",
			toFormatStr:   "jpg",
			expected: ConvertedDataRepository{
				map[string]string{
					"from": filepath.Join(testdir, "test.png"),
					"to":   filepath.Join(testdir, "test.jpg"),
				},
			},
			isDir:       false,
			errorExists: false,
		},
		{
			name:          "UnmatchedFormat",
			setupPath:     filepath.Join(testdir, "test.png"),
			path:          filepath.Join(testdir, "test.png"),
			fromFormatStr: "gif",
			toFormatStr:   "jpg",
			expected:      nil,
			isDir:         false,
			errorExists:   false,
		},
		{
			name:          "Directory",
			setupPath:     filepath.Join(testdir, "test"),
			path:          filepath.Join(testdir, "test"),
			fromFormatStr: "gif",
			toFormatStr:   "jpg",
			expected:      nil,
			isDir:         true,
			errorExists:   false,
		},
		{
			name:          "FailConvert",
			setupPath:     filepath.Join(testdir, "test.png"),
			path:          "",
			fromFormatStr: "png",
			toFormatStr:   "jpg",
			expected:      nil,
			isDir:         false,
			errorExists:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Teardown := SetupTest(t, tt.setupPath, tt.isDir)
			defer Teardown()
			c, err := NewConfig(tt.path, tt.fromFormatStr, tt.toFormatStr)
			if err != nil {
				t.Errorf("Error : %v", err)
			}

			actual, err := c.ConvertRec()
			if err != nil && !tt.errorExists {
				t.Errorf("Error : %v", err)
			}

			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("c.ConvertRec() => %v, want %v", actual, tt.expected)
			}
		})
	}

}

func TestGetImageFormat(t *testing.T) {
	tests := []struct {
		name     string
		format   string
		expected int
	}{
		{
			name:     "Jpg",
			format:   "jpg",
			expected: JpgFormat,
		},
		{
			name:     "Png",
			format:   "png",
			expected: PngFormat,
		},
		{
			name:     "Gif",
			format:   "gif",
			expected: GifFormat,
		},
		{
			name:     "JpgUpperCase",
			format:   "JPG",
			expected: JpgFormat,
		},
		{
			name:     "Default",
			format:   "undefind",
			expected: DefaultFormat,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := GetImageFormat(tt.format)
			if actual != tt.expected {
				t.Errorf("GetImageFormat(%s) => %d, want %d", tt.format, actual, tt.expected)
			}
		})
	}
}

func TestConvert(t *testing.T) {
	tests := []struct {
		name        string
		setupPath   string
		fromPath    string
		toFormat    int
		expected    string
		errorExists bool
	}{
		{
			name:        "toJpg",
			setupPath:   filepath.Join(testdir, "test.png"),
			fromPath:    filepath.Join(testdir, "test.png"),
			toFormat:    JpgFormat,
			expected:    filepath.Join(testdir, "test.jpg"),
			errorExists: false,
		},
		{
			name:        "ToPng",
			setupPath:   filepath.Join(testdir, "test.jpg"),
			fromPath:    filepath.Join(testdir, "test.jpg"),
			toFormat:    PngFormat,
			expected:    filepath.Join(testdir, "test.png"),
			errorExists: false,
		},
		{
			name:        "ToGif",
			setupPath:   filepath.Join(testdir, "test.png"),
			fromPath:    filepath.Join(testdir, "test.png"),
			toFormat:    GifFormat,
			expected:    filepath.Join(testdir, "test.gif"),
			errorExists: false,
		},
		{
			name:        "Default",
			setupPath:   filepath.Join(testdir, "test.png"),
			fromPath:    filepath.Join(testdir, "test.png"),
			toFormat:    DefaultFormat,
			expected:    "",
			errorExists: false,
		},
		{
			name:        "ToJpgFail",
			setupPath:   filepath.Join(testdir, "test.png"),
			fromPath:    "",
			toFormat:    JpgFormat,
			expected:    "",
			errorExists: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Teardown := SetupTest(t, tt.setupPath, false)
			defer Teardown()
			actual, err := Convert(tt.fromPath, tt.toFormat)
			if err != nil && !tt.errorExists {
				t.Errorf("Error : %v", err)
			}
			if actual != tt.expected {
				t.Errorf("Convert(%s, %d) => %s, want : %s", tt.fromPath, tt.toFormat, actual, tt.expected)
			}
		})
	}
}

func TestConvertToPng(t *testing.T) {
	tests := []struct {
		name        string
		path        string
		expected    string
		setup       bool
		errorExists bool
	}{
		{
			name:        "Success",
			path:        filepath.Join(testdir, "test.jpg"),
			expected:    filepath.Join(testdir, "test.png"),
			setup:       true,
			errorExists: false,
		},
		{
			name:        "Fail",
			path:        "",
			expected:    "",
			setup:       false,
			errorExists: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup {
				Teardown := SetupTest(t, tt.path, false)
				defer Teardown()
			}

			actual, err := ConvertToPng(tt.path)
			if err != nil && !tt.errorExists {
				t.Errorf("Error : %v", err)
			}

			if actual != tt.expected {
				t.Errorf("ConvertToPng(%s) => %s, want : %s", tt.path, actual, tt.expected)
			}
		})
	}
}

func TestConvertToJpg(t *testing.T) {
	tests := []struct {
		name        string
		path        string
		expected    string
		setup       bool
		errorExists bool
	}{
		{
			name:        "Success",
			path:        filepath.Join(testdir, "test.png"),
			expected:    filepath.Join(testdir, "test.jpg"),
			setup:       true,
			errorExists: false,
		},
		{
			name:        "Fail",
			path:        "",
			expected:    "",
			setup:       false,
			errorExists: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup {
				Teardown := SetupTest(t, tt.path, false)
				defer Teardown()
			}

			actual, err := ConvertToJpg(tt.path)
			if err != nil && !tt.errorExists {
				t.Errorf("Error : %v", err)
			}

			if actual != tt.expected {
				t.Errorf("ConvertToPng(%s) => %s, want : %s", tt.path, actual, tt.expected)
			}
		})
	}
}

func TestConvertToGif(t *testing.T) {
	tests := []struct {
		name        string
		path        string
		expected    string
		setup       bool
		errorExists bool
	}{
		{
			name:        "Success",
			path:        filepath.Join(testdir, "test.jpg"),
			expected:    filepath.Join(testdir, "test.gif"),
			setup:       true,
			errorExists: false,
		},
		{
			name:        "Fail",
			path:        "",
			expected:    "",
			setup:       false,
			errorExists: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup {
				Teardown := SetupTest(t, tt.path, false)
				defer Teardown()
			}

			actual, err := ConvertToGif(tt.path)
			if err != nil && !tt.errorExists {
				t.Errorf("Error : %v", err)
			}

			if actual != tt.expected {
				t.Errorf("ConvertToPng(%s) => %s, want : %s", tt.path, actual, tt.expected)
			}
		})
	}
}
