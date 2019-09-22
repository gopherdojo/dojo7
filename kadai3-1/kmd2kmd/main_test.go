package main

import (
	"bytes"
	"os"
	"reflect"
	"regexp"
	"testing"
	"time"
)

func TestMain_getInputChannel(t *testing.T)  {

	var buf bytes.Buffer
	buf.Write([]byte("a"))
	want := "a"
	got := <-getInputChannel(&buf)
	if got != want {
		t.Fatalf("got: %v, want: %v", got, want)
	}
}

func TestMain_play(t *testing.T) {
	expected := 1

	ch := make(chan string)
	go func() {
		time.Sleep(100 * time.Millisecond) // waiting for start
		ch <- "b"                          // incorrect
		ch <- "a"                          // correct
		time.Sleep(200 * time.Millisecond) // waiting for finish
	}()

	var (
		out bytes.Buffer
		err bytes.Buffer
	)
	actual := play(&out, &err, []string{"a"}, 200*time.Millisecond, ch)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf(`expected="%v" actual="%v"`, expected, actual)
	}

	cases := []struct {
		name 			string
		resultMsg        string
		correct		 bool
	}{
		{
			name: "Correct a",
			resultMsg: "Correct!!",
			correct: true,
		},
		{
			name: "Mistake a",
			resultMsg: "Mistake!!",
			correct: false,
		},
		{
			name: "Correct b",
			resultMsg: "Correct!!",
			correct: true,
		},
		{
			name: "Mistake b",
			resultMsg: "Mistake!!",
			correct: false,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			r := regexp.MustCompile(c.resultMsg)
			actual := r.MatchString(out.String())
			correct := c.correct
			if actual != correct {
				t.Errorf(`correct="%t" actual="%t"`, correct, actual)
			}
		})
	}
}

func TestMain_run(t *testing.T) {
	cases := []struct {
		name 			string
		args        []string
		want int
	}{
		{
			name: "normal",
			args: []string{"binary", "-timeout=1"},
			want: ExitCodeOK,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			got := run(os.Stdin, os.Stdout, os.Stderr, c.args)
			if got != c.want {
				t.Fatalf("got: %v, want: %v", got, c.want)
			}
		})
	}
}