package image

import (
	"os"
	"testing"
)

//
func TestSupportedFormat(t *testing.T){
	cases := []struct{name string; input string; expected bool}{
		{name: "jpg", input: "jpg", expected: true},
		{name: "jpeg", input: "jpeg", expected: true},
		{name: "png", input: "png", expected: true},
		{name: "gif", input: "gif", expected: true},
		{name: "GIF", input: "GIF", expected: false},
		{name: "no-ext", input: "", expected: false},
		{name: "dotpng", input: ".png", expected: true},
		{name: "mp4", input: "mp4", expected: false},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T){
			t.Parallel()
			if actual := SupportedFormat(c.input); c.expected != actual{
				t.Errorf(
					"want supportedFormat(%s) = %v, got %v",
					c.input, c.expected, actual)
			}
		})
	}
}

func TestConvExts_SupportedFormats(t *testing.T){
	cases := []struct{name string; input ConvExts; expected bool}{
		{name: "jpg-jpg", input: ConvExts{"jpg", "jpg"}, expected: true},
		{name: "png-jpeg", input: ConvExts{"png", "jpeg"}, expected: true},
		{name: "jpg-no-ext", input: ConvExts{"jpg", ""}, expected: false},
		{name: "gif-png", input: ConvExts{"gif", "png"}, expected: true},
		{name: "gif-mp4", input: ConvExts{"gif", "mp4"}, expected: false},

	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T){
			t.Parallel()
			if actual := c.input.SupportedFormats(); c.expected != actual{
				t.Errorf(
					"want SupportedFormats(%s) = %v, got %v",
					c.input, c.expected, actual)
			}
		})
	}
}

func TestNewConvExts(t *testing.T) {
	type inputs struct {
		input, output string
	}

	cases := []struct{name string; input inputs; expected ConvExts
	}{
		{name: "jpg-png", input: inputs{"jpg", "png"}, expected: ConvExts{"jpg", "png"}},
		{name: "no-ext-jpeg", input: inputs{"", "jpeg"}, expected: ConvExts{"jpg", "jpeg"}},
		{name: "jpg-no-ext", input: inputs{"jpg", ""}, expected: ConvExts{"jpg", "png"}},
		{name: "no-ext-no-ext", input: inputs{"", ""}, expected: ConvExts{"jpg", "png"}},
		{name: "gif-mp4", input: inputs{"gif", "mp4"}, expected: ConvExts{"gif", "mp4"}},

	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T){
			t.Parallel()
			if actual := NewConvExts(c.input.input, c.input.output); c.expected != actual{
				t.Errorf(
					"want NewConvExts(%s) = %v, got %v",
					c.input, c.expected, actual)
			}
		})
	}
}

func TestFmtConv(t *testing.T) {
	t.Helper()
	// 前回作成したファイルは削除される
	os.Remove("testdata/test_img/test_img1.png")

	type inputs struct {
		path string;
		convExts ConvExts
	}
	cases := []struct{name string; input inputs; expected string
	}{
		{
			name: "jpeg-png",
			input: inputs{path:"testdata/test_img_1.jpeg", convExts:ConvExts{"jpeg", "png"}},
			expected: "testdata/test_img_1.png",
		},
		{
			name: "png-gif",
			input: inputs{path:"testdata/test_img_2.png", convExts:ConvExts{"png", "gif"}},
			expected: "testdata/test_img_2.gif",
		},
		{
			name: "gif-jpg",
			input: inputs{path:"testdata/test_img_3.gif", convExts:ConvExts{"gif", "jpg"}},
			expected: "testdata/test_img_3.jpg",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T){
			t.Parallel()
			// ファイルがある場合は削除する
			os.Remove(c.expected)
			f, err := os.Open(c.input.path)
			if err != nil {
				t.Errorf("cannnot open file: %v", c.input.path)
			}
			if err = FmtConv(f, c.input.convExts); err != nil {
				t.Errorf("cannnot convert file: %v", c.input.path)
			}
			if _, err := os.Stat(c.expected); err != nil{
				t.Errorf(
					"convert failed %v", c.input.path)
			}
		})
	}

}
