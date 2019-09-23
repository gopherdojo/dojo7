package typinggame

import (
	"testing"
	"time"
)

func TestIsCorrect(t *testing.T) {
	var questions = []struct {
		in       string
		out      string
		expected bool
	}{
		{"banana", "banana", true}, {"peach", "peach ", false}, {" peach", "peach", false},
	}

	for _, tt := range questions {
		actual := isCorrect(tt.in, tt.out)
		if actual != tt.expected {
			t.Errorf("isCorrect(%v,%v) expected %v but got %v", tt.in, tt.out, tt.expected, actual)
		}
	}

}

type MockInputReader struct {
	Questions   []string
	AnswerCount int
}

func (r *MockInputReader) Input() <-chan string {
	ch := make(chan string)
	var i = 0

	go func() {
		for _, q := range r.Questions {
			ch <- q
			i++
			if i > r.AnswerCount { //stops input after AnswerCount exceeded
				close(ch)
			}

		}

	}()
	return ch
}

func TestRun(t *testing.T) {
	var tests = []struct {
		name          string
		inputs        []string
		game          Game
		expectedScore int
		expectedCount int
	}{
		{"correct score",
			[]string{"apple", "banana", "apple", "peach"},
			Game{
				Questions: []string{"apple"},
			},
			2, 4},
		{"correct count",
			[]string{"peach", "banan"},
			Game{
				Questions: []string{"peach"},
			},
			1, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cli{InputReader: &MockInputReader{tt.inputs, len(tt.inputs)}}
			c.Run(&tt.game, 5*time.Second)
			gotScore := c.Score
			gotCount := c.Count
			if tt.expectedScore != gotScore || tt.expectedCount != gotCount {
				t.Errorf("c.Run(%v, 5sec) got score/count of %v/%v but got %v/%v", tt.game, tt.expectedScore, tt.expectedCount, c.Score, c.Count)
			}
		})

	}
}
