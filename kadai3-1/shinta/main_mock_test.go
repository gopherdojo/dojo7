package main_test

import (
	"io"
)

//Stdin
type StdinMock struct {
	i     int
	input []string
}

func (stdin *StdinMock) Read(p []byte) (n int, err error) {
	if stdin.i >= len(stdin.input) {
		return 0, io.EOF
	}
	b := []byte(stdin.input[stdin.i] + "\n") //Scanが回るようにLF追加
	copy(p, b)
	stdin.i++
	return len(b), nil
}

//Stdout
type StdoutMock struct {
	output []string
}

func (stdout *StdoutMock) Write(p []byte) (n int, err error) {
	str := string(p)
	stdout.output = append(stdout.output, str)
	return len(str), nil
}

//Typing
type TypingMock struct {
}

func (typ *TypingMock) ShowText() string {
	return "FOO"
}

func (typ *TypingMock) Judge(input string) bool {
	return "FOO" == input
}
