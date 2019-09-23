package main_test

import (
	"bytes"
	"testing"
	"time"

	main "github.com/gopherdojo/dojo7/kadai3-1/shinta"
)

func TestInputRoutine(t *testing.T) {
	input := []string{"foo", "bar", "baz", "qux"}
	stdin := &StdinMock{
		i:     0,
		input: input,
	}

	ch := main.InputRoutine(stdin)

	for _, expected := range input {
		actual := <-ch
		if actual != expected {
			t.Errorf("expected:%v, actual:%v", expected, actual)
		}
	}
}

func TestExecute(t *testing.T) {
	chInput := make(chan string, 3)
	chFinish := make(chan time.Time, 1)

	scenario := []struct {
		inputText string
		time      time.Time
	}{
		{
			inputText: "FOO",
		},
		{
			inputText: "BAR",
		},
		{
			inputText: "FOO",
		},
		{
			time: time.Now(),
		},
	}

	buf := bytes.NewBufferString("")
	typ := &TypingMock{}

	go func() {
		for _, s := range scenario {
			time.Sleep(100 * time.Millisecond)
			if s.inputText != "" {
				chInput <- s.inputText
			}
			if !s.time.IsZero() {
				chFinish <- s.time
			}
		}
	}()
	main.Execute(chInput, chFinish, buf, typ)

	expected := []byte("" +
		"[001]: FOO\n" + "type>>" + "Correct!\n" +
		"[002]: FOO\n" + "type>>" + "Miss!\n" +
		"[003]: FOO\n" + "type>>" + "Correct!\n" +
		"[004]: FOO\n" + "type>>" +
		"\nTime's up!!\n" +
		"You Scored: 2\n")

	if bytes.Compare(buf.Bytes(), expected) != 0 {
		t.Errorf("[expected]:\n%s\n[actual]:\n%s", expected, buf.Bytes())
	}
}
