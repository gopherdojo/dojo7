package iconv

import (
	"errors"
	"io/ioutil"
	"os"
	"reflect"
	"syscall"
	"testing"
)


func TestNormalGetFileNameWithoutExt(t *testing.T) {
	cases := []struct {
		name	string
		input	string
		want    string
		err     error
	}{
		{
			name: "normal1",
			input: "/example/path/img.jpg",
			want: "/example/path/img",
			err: nil,
		},
		{
			name: "normal2",
			input: "/example/path/img.png",
			want: "/example/path/img",
			err: nil,
		},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			got, err := getFileNameWithoutExt(c.input)
			// 正常系 エラーチェック
			if err != nil {
				t.Fatal("want no err, but has error", err)
			}
			// 正常系 レスポンスチェック
			if got != c.want {
				t.Fatalf("got: %v, want: %v", got, c.want)
			}
		})
	}
}

func TestErrorGetFileNameWithoutExt(t *testing.T) {
	cases := []struct {
		name	string
		input	string
		want    string
		err     error
	}{
		{
			name: "empty",
			input: "",
			want: "",
			err: errors.New("an empty string was entered"),
		},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			_, err := getFileNameWithoutExt(c.input)
			if !reflect.DeepEqual(err, c.err) {
				t.Fatalf("got: %v, want: %v", err, c.err)
			}
		})
	}
}

func TestNormalGetDecodedImage(t *testing.T) {
	cases := []struct{
		name string
		input string
	}{
		{
			name: "normal",
			input: "./testdata/lena.png",
		},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			_, err := getDecodedImage(c.input)
			if err != nil {
				t.Error("want no err, but has error", err)
			}
		})
	}

}

func TestErrorGetDecodedImage(t *testing.T) {
	cases := []struct{
		name string
		input string
		err error
	}{
		{
			name: "file not found",
			input: "",
			err: &os.PathError{Op: "open", Path: "", Err: syscall.Errno(syscall.ENOENT)},
		},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			_, err := getDecodedImage(c.input)
			if !reflect.DeepEqual(err, c.err) {
				t.Errorf("got: %v, want: %v", err, c.err)
			}
		})
	}

}

func TestNormalConvertToJpeg(t *testing.T) {
	testdata := "./testdata/lena.png"
	img, _ := getDecodedImage(testdata)

	tempFile, err := ioutil.TempFile("./", "temp_file")
	if err != nil {
		t.Error("failed create temp file", err)
	}
	err = convertToJpeg(img, tempFile.Name())
	if err != nil {
		t.Error("failed convert", err)
	}

	_, err = os.Stat(tempFile.Name())
	if err != nil {
		t.Error("failed create image", err)
	}

	defer func() {
		err := os.Remove(tempFile.Name())
		if err != nil {
			t.Error("failed remove temporary file", err)
		}
	}()
}

func TestNormalConvertToPng(t *testing.T) {
	testdata := "./testdata/lena.png"
	img, _ := getDecodedImage(testdata)

	tempFile, err := ioutil.TempFile("./", "temp_file")
	if err != nil {
		t.Error("failed create temp file", err)
	}
	err = convertToPng(img, tempFile.Name())
	if err != nil {
		t.Error("failed convert", err)
	}

	_, err = os.Stat(tempFile.Name())
	if err != nil {
		t.Error("failed create image", err)
	}

	defer func() {
		err := os.Remove(tempFile.Name())
		if err != nil {
			t.Error("failed remove temporary file", err)
		}
	}()
}

func TestNormalDo(t *testing.T) {
	cases := []struct{
		name string
		image Image
	}{
		{
			name: "normal png→jpg",
			image: Image{
				In: "png",
				Out: "jpg",
				Path: "./testdata/lena.png",
			},
		},
		{
			name: "normal jpg→png",
			image: Image{
				In: "jpg",
				Out: "png",
				Path: "./testdata/lena.jpg",
			},
		},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			err := c.image.Do()
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestErrorDo(t *testing.T) {
	cases := []struct{
		name string
		image Image
		err error
	}{
		{
			name: "unsupported format",
			image: Image{
				In: "unsupported",
				Out: "unsupported",
				Path: "./testdata/lena.png",
			},
			err: errors.New("unsupported format"),
		},
		{
			name: "not image",
			image: Image{
				In: "jpg",
				Out: "png",
				Path: "./testdata/not_image.jpg",
			},
			err: errors.New("image: unknown format"),
		},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			err := c.image.Do()
			if !reflect.DeepEqual(err, c.err) {
				t.Errorf("got: %v, want: %v", err, c.err)
			}
		})
	}
}