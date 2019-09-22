package opt_test

import (
	"github.com/gopherdojo/dojo7/kadai3-1/kmd2kmd/opt"
	"os"
	"testing"
)

func TestOptions(t *testing.T) {
	cases := []struct {
		name        string
		args        []string
		wantPath    string
		wantTimeout int
	}{
		{
			name:        "default",
			args:        []string{"binary"},
			wantPath:    "./words.txt",
			wantTimeout: 15,
		},
		{
			name:        "path",
			args:        []string{"binary", "--path=./testdata.txt"},
			wantPath:    "./testdata.txt",
			wantTimeout: 15,
		},
		{
			name:        "timeout",
			args:        []string{"binary", "--timeout=5"},
			wantPath:    "./words.txt",
			wantTimeout: 5,
		},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			options, err := opt.Parse(os.Stderr, c.args)
			if err != nil {
				t.Fatal(err)
			}
			path := options.Path()
			if path != c.wantPath {
				t.Fatalf("got: %v, want: %v", path, c.wantPath)
			}

			timeout := options.Timeout()
			if timeout != c.wantTimeout {
				t.Fatalf("got: %v, want: %v", timeout, c.wantTimeout)
			}

		})
	}

}
