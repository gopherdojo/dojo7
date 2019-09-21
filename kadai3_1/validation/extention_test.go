package validation

import (
	"testing"
)

func TestExtValidation(t *testing.T) {
	testCases := []struct {
		Name   string
		Ext    string
		Result bool
	}{
		{Name: "validate success", Ext: "png", Result: true},
		{Name: "validate success", Ext: "jpg", Result: true},
		{Name: "validate fail not support", Ext: "gif", Result: false},
		{Name: "validate fail not support", Ext: "pdf", Result: false},
		{Name: "validate fail number", Ext: "23456886", Result: false},
		{Name: "validate fail symbol", Ext: ":;@[]_/23!-^~#/.,", Result: false},
		{Name: "validate fail symbol number", Ext: ":;@po234[]_/23!-^~#/.,", Result: false},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			if Ext(tc.Ext) != tc.Result {
				t.Fatalf("%s %s", "ext", tc.Ext)
			}
		})
	}
}
