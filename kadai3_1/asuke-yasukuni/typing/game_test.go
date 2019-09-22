package typing

import (
	"testing"
	"time"
)

type TestReader struct {
	TestWord []string
}

func (r *TestReader) Input() <-chan string {
	ch := make(chan string)
	go func() {
		for _, word := range r.TestWord {
			ch <- word
		}
		time.Sleep(1 * time.Second)
		close(ch)
	}()
	return ch
}

type TestWriter struct{}

func (w TestWriter) Output(outputText string) {}

func TestTyping(t *testing.T) {

	var testCase = []struct {
		Name  string
		Word  []string
		Count int
	}{
		{"all clear", []string{"hoge", "hoge", "hoge", "hoge"}, 4},
		{"all fail", []string{"neko", "english", "eigo", "genzin"}, 0},
	}

	for _, tc := range testCase {
		game := &Game{
			Time:   1,
			Words:  []string{"hoge", "hoge", "hoge", "hoge"},
			Reader: &TestReader{TestWord: tc.Word},
			Writer: TestWriter{},
		}
		t.Run(tc.Name, func(t *testing.T) {
			count := game.Do()
			if count != tc.Count {
				t.Errorf("%d == %d not equal", count, tc.Count)
			}
		})
	}
}
