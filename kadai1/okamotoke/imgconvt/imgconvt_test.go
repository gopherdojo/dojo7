package imgconvt

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestDecode(t *testing.T) {
	tests := []struct {
		test    string
		path    string
		from    string
		wantErr bool
	}{
		{
			test:    "no extension",
			path:    "testdata/test.jpg",
			from:    "",
			wantErr: true,
		}, {
			test:    "file does not exist",
			path:    "testdata/test2.jpg",
			from:    ".jpg",
			wantErr: true,
		},
		{
			test:    "unsupported extenstion",
			path:    "testdata/test.TIF",
			from:    ".TIF",
			wantErr: true,
		},
		{
			test:    "decode png",
			path:    "testdata/test.png",
			from:    ".png",
			wantErr: false,
		},
		{
			test:    "decode jpg",
			path:    "testdata/test.jpg",
			from:    ".jpg",
			wantErr: false,
		},
		{
			test:    "decode gif",
			path:    "testdata/test.gif",
			from:    ".gif",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.test, func(t *testing.T) {
			_, err := decode(tt.path, tt.from)
			if (err != nil) != tt.wantErr {
				t.Errorf("%v : decode(%v,%v) %v", tt.test, tt.path, tt.from, tt.wantErr)
			}
		})

	}

}

func TestEncode(t *testing.T) {
	tests := []struct {
		test    string
		to      string
		wantErr bool
	}{
		{
			test:    "no extension",
			to:      "",
			wantErr: true,
		},
		{
			test:    "unsupported extension",
			to:      ".tif",
			wantErr: true,
		},
		{
			test:    "convert to jpg",
			to:      ".jpg",
			wantErr: false,
		},
		{
			test:    "convert to png",
			to:      ".png",
			wantErr: false,
		},
		{
			test:    "convert to gif",
			to:      ".gif",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.test, func(t *testing.T) {
			img, _ := decode("testdata/test.jpg", ".jpg")
			err := encode(img, tt.to, testTempFile(t), "")
			if (err != nil) != tt.wantErr {
				t.Errorf("%v : encode(img, %v, test) %v", tt.test, tt.to, tt.wantErr)
			}
		})

	}
}

func TestConvertImage(t *testing.T) {
	tests := []struct {
		test    string
		path    string
		from    string
		to      string
		wantErr bool
	}{
		{
			test:    "jpg to png",
			path:    "testdata/test.jpg",
			from:    ".jpg",
			to:      ".png",
			wantErr: false,
		},
		{
			test:    "png to jpg",
			path:    "testdata/test.png",
			from:    ".png",
			to:      ".jpg",
			wantErr: false,
		},
		{
			test:    "gif to png",
			path:    "testdata/test.gif",
			from:    ".gif",
			to:      ".png",
			wantErr: false,
		},
		{
			test:    "jpg to jpeg",
			path:    "testdata/test.jpg",
			from:    ".jpg",
			to:      ".jpeg",
			wantErr: false,
		},
		{
			test:    "jpg to tif",
			path:    "testdata/test.jpg",
			from:    ".jpg",
			to:      ".tif",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.test, func(t *testing.T) {
			err := ConvertImage(&Conv{tt.path, "tmpDir/", tt.from, tt.to})
			defer func() {
				os.RemoveAll("tmpDir/")
			}()
			if (err != nil) != tt.wantErr {
				t.Errorf("%v : ConvertImage() %v", tt.test, tt.wantErr)
			}
		})

	}
}

func testTempFile(t *testing.T) string {
	t.Helper()
	tf, err := ioutil.TempFile("", "test")
	if err != nil {
		t.Fatalf("err %s", err)
	}
	tf.Close()
	return tf.Name()
}
