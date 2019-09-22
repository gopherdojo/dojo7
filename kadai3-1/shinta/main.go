package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gopherdojo/dojo7/kadai3-1/shinta/typing"
)

const timeLimit = 5

type Typing interface {
	ShowText() string
	Judge(input string) bool
}

func main() {
	problems := []string{"apple", "bake", "cup", "dog", "egg", "fight", "green", "hoge", "idea", "japan"}
	typing := typing.Redy(problems)
	chInput := inputRoutine(os.Stdin)
	chFinish := time.After(time.Duration(timeLimit) * time.Second)

	execute(chInput, chFinish, os.Stdout, typing)
}

func execute(chInput <-chan string, chFinish <-chan time.Time, stdout io.Writer, t Typing) {

	score := 0
	for i := 0; ; i++ {
		fmt.Fprintf(stdout, "[%03d]: %v\n", i, t.ShowText())
		fmt.Fprint(stdout, "type>>")
		select {
		case inText := <-chInput:
			if t.Judge(inText) {
				score++
				fmt.Fprintln(stdout, "Correct!")
			} else {
				fmt.Fprintln(stdout, "Miss!")
			}
		case <-chFinish:
			fmt.Fprintln(stdout, "\nTime's up!!")
			fmt.Fprintf(stdout, "You Scored: %v\n", score)
			return
		}
	}
}

func inputRoutine(r io.Reader) <-chan string {
	ch := make(chan string)

	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch <- s.Text()
		}
	}()

	return ch
}
