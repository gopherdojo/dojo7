package imageconv_test

import (
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"testing"

	"github.com/waytkheming/godojo/dojo7/kadai2/waytkheming/imageconv"
)

type testCase struct {
	title  string
	path   string
	from   string
	to     string
	output string
}

func TestConvert(t *testing.T) {
	var testFixtures = []testCase{
		{
			title:  "jpg to png",
			path:   "../testdata/earthmap1k.jpg",
			from:   "jpg",
			to:     "png",
			output: "../testdata/earthmap1k.png",
		},
		{
			title:  "png to jpg",
			path:   "../testdata/earthmap1k.png",
			from:   "png",
			to:     "jpg",
			output: "../testdata/earthmap1k.jpg",
		},
		{
			title:  "jpg to gif",
			path:   "../testdata/earthmap1k.jpg",
			from:   "jpg",
			to:     "gif",
			output: "../testdata/earthmap1k.gif",
		},
		{
			title:  "png to gif",
			path:   "../testdata/earthmap1k.png",
			from:   "png",
			to:     "gif",
			output: "../testdata/earthmap1k.gif",
		},
		{
			title:  "gif to jpg",
			path:   "../testdata/earthmap1k.gif",
			from:   "gif",
			to:     "jpg",
			output: "../testdata/earthmap1k.jpg",
		},
		{
			title:  "gif to png",
			path:   "../testdata/earthmap1k.gif",
			from:   "gif",
			to:     "png",
			output: "../testdata/earthmap1k.png",
		},
	}
	for _, testFixture := range testFixtures {
		c := imageconv.NewConverter(testFixture.path, testFixture.from, testFixture.to)
		i := imageconv.NewImage(testFixture.path)
		t.Run("Check walk", func(t *testing.T) {
			checkWalk(t, c)
		})
		t.Run("Check convert", func(t *testing.T) {
			checkConvert(t, c, i)
		})
		t.Run("Check format", func(t *testing.T) {
			checkFormat(t, testFixture.output, testFixture.to)
		})
	}
}

func checkWalk(t *testing.T, c imageconv.Converter) {
	t.Helper()
	if err := filepath.Walk(c.Path, c.CrawlFile); err != nil {
		t.Errorf("Error: %s", err)
	}
}

func checkConvert(t *testing.T, c imageconv.Converter, i imageconv.ImageFile) {
	t.Helper()
	if err := c.Convert(i); err != nil {
		t.Errorf("Error: %s", err)
	}
}

func checkFormat(t *testing.T, path string, to string) {
	t.Helper()
	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Errorf("Expected output file %s %s is not exist", path, err.Error())
	}
	file, err := os.Open(path)
	if err != nil {
		t.Errorf("Couldn't open file path: %s, fileType: %s, error: %v", path, to, err)
	}
	defer file.Close()

	switch to {
	case "jpg", "jpeg":
		_, err = jpeg.Decode(file)
	case "png":
		_, err = png.Decode(file)
	case "gif":
		_, err = gif.Decode(file)
	}

	if err != nil {
		t.Errorf("Couldn't decode path: %s, fileType: %s, error: %v", path, to, err)
	}
}
