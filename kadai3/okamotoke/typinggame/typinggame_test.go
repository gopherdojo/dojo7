package typinggame

import "testing"

func TestIsCorrect(t *testing.T) {
	var words = []struct {
		in       string
		out      string
		expected bool
	}{
		{"banana", "banana", true}, {"peach", "peach ", false}, {" peach", "peach", false},
	}

	for _, tt := range words {
		actual := isCorrect(tt.in, tt.out)
		if actual != tt.expected {
			t.Errorf("isCorrect(%v,%v) expected %v but got %v", tt.in, tt.out, tt.expected, actual)
		}
	}

}
