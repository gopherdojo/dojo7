package imachan

import (
	"strings"
	"testing"
)

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
