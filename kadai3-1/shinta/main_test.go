package main_test

import (
	"bytes"
	"testing"
	"time"

	main "github.com/gopherdojo/dojo7/kadai3-1/shinta"
)

func TestInputRoutine(t *testing.T) {
	input := []string{"apple", "bake", "cup", "dog"}
	// StdinMockのinputに、標準入力で送信されたと仮定した文字列が入る
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
			inputText: "apple",
		},
		{
			inputText: "bbaak",
		},
		{
			inputText: "cup",
		},
		{
			time: time.Now(),
		},
	}

	// 新しく文字列を格納するbufferを確保
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
		"[001]: apple\n" + "type>>" + "Correct!\n" +
		"[002]: bake\n" + "type>>" + "Miss!\n" +
		"[003]: cup\n" + "type>>" + "Correct!\n" +
		"[004]: dog\n" + "type>>" +
		"\nTime's up!!\n" +
		"You Scored: 2\n")

	// byteスライスの比較、a == bの場合は0を返す
	if bytes.Compare(buf.Bytes(), expected) != 0 {
		t.Errorf("[expected]:\n%s\n[actual]:\n%s", expected, buf.Bytes())
	}
}
