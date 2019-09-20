package convert_test

import (
	"marltake/convert"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type result struct {
	Src  string
	Dest string
	Ok   bool
}

func TestParseTarget(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected result
	}{
		{name: "default", input: "", expected: result{"jpg", "png", true}},
		{name: "default_dest", input: "jpg", expected: result{"jpg", "png", true}},
		{name: "default_src", input: ",png", expected: result{"jpg", "png", true}},
		{name: "jpg to png", input: "jpg,png", expected: result{"jpg", "png", true}},
		{name: "jpg to gif", input: "jpg,gif", expected: result{"jpg", "gif", true}},
		{name: "png to jpg", input: "png,jpg", expected: result{"png", "jpg", true}},
		{name: "png to gif", input: "png,gif", expected: result{"png", "gif", true}},
		{name: "gif to jpg", input: "gif,jpg", expected: result{"gif", "jpg", true}},
		{name: "gif to png", input: "gif,png", expected: result{"gif", "png", true}},
		{name: "invalid dest", input: ",mp3", expected: result{"", "", false}},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			if src, dest, ok := convert.ParseTarget(c.input); !cmp.Equal(c.expected, result{src, dest, ok}) {
				t.Errorf(
					"want ParseTarget(%q) = %v, got %v",
					c.input, c.expected, result{src, dest, ok})
			}
		})
	}
}
