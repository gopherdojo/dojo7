package main

import (
	"testing"
	"time"

	"github.com/gopherdojo/dojo7/kadai3-1/nas/simpletype/pkg/exercise"
)

type MockReader struct {
	Words []string
}

func (mr *MockReader) Answer() <-chan string {
	ch := make(chan string)

	go func() {
		for _, w := range mr.Words {
			ch <- w
			time.Sleep(1 * time.Second)
		}
		close(ch)
	}()
	return ch
}

func TestCliRun(t *testing.T) {
	cases := []struct {
		name     string
		duration int
		words    []string
		want     int
	}{
		{"Success", 10, []string{"apple", "banana", "cat", "dog", "egg", "fish"}, 6},
		{"Time out", 3, []string{"apple", "banana", "cat", "dog", "egg", "fish"}, 3},
		{"No Exercise", 10, []string{}, 0},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			e := &exercise.Exercise{
				Questions: tt.words,
			}
			d := time.Duration(tt.duration) * time.Second
			c := &Cli{InputReader: &MockReader{tt.words}}
			c.Run(e, d)
			if got := c.Correct; tt.want != got {
				t.Errorf("c.Run(%v, %v) >> c.Count => %d, but want %d", e, d, tt.want, c.Correct)
			}
		})

	}
}
