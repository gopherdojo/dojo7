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

func SetupTest(t *testing.T, path string) func() {
	f, err := os.Create(path)
	if err != nil {
		t.Errorf("Setup Error : %v", err)
	}
	defer f.Close()

	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	if err := jpeg.Encode(f, img, &jpeg.Options{Quality: 100}); err != nil {
		t.Errorf("Setup Error : %v", err)
	}

	return func() {
		err = filepath.Walk(testdir, func(path string, info os.FileInfo, err error) error {
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
		path          string
		fromFormatStr string
		toFormatStr   string
		expected      ConvertedDataRepository
	}{
		{
			name:          "PngToJpg",
			path:          filepath.Join(testdir, "test.png"),
			fromFormatStr: "png",
			toFormatStr:   "jpg",
			expected: ConvertedDataRepository{
				map[string]string{
					"from": filepath.Join(testdir, "test.png"),
					"to":   filepath.Join(testdir, "test.jpg"),
				},
			},
		},
	}
	for _, tt := range tests {
		Teardown := SetupTest(t, tt.path)
		defer Teardown()
		c, err := NewConfig(tt.path, tt.fromFormatStr, tt.toFormatStr)
		if err != nil {
			t.Errorf("Error : %v", err)
		}

		actual, err := c.ConvertRec()
		if err != nil {
			t.Errorf("Error : %v", err)
		}

		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("c.ConvertRec() => %v, want %v", actual, tt.expected)
		}
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
		name     string
		file     string
		path     string
		format   int
		expected string
	}{
		{
			name:     "ToPng",
			path:     filepath.Join(testdir, "test.png"),
			format:   JpgFormat,
			expected: filepath.Join(testdir, "test.jpg"),
		},
		{
			name:     "ToJpg",
			path:     filepath.Join(testdir, "test.jpg"),
			format:   PngFormat,
			expected: filepath.Join(testdir, "test.png"),
		},
		{
			name:     "ToGif",
			path:     filepath.Join(testdir, "test.png"),
			format:   GifFormat,
			expected: filepath.Join(testdir, "test.gif"),
		},
		{
			name:     "Default",
			path:     filepath.Join(testdir, "test.png"),
			format:   DefaultFormat,
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Teardown := SetupTest(t, tt.path)
			defer Teardown()
			actual, err := Convert(tt.path, tt.format)
			if err != nil {
				t.Errorf("Error : %s", err)
			}
			if actual != tt.expected {
				t.Errorf("Convert(%s, %d) => %s, want : %s", tt.path, tt.format, actual, tt.expected)
			}
		})
	}
}
