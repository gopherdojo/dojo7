package typinggame

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

// Run starts a typing game.
func Run(q *Game) {

	ch := input(os.Stdin)
	timeCh := time.After(5 * time.Second)
	var score, count int
L:
	for {
		word := q.getNextQuestion()
		fmt.Println(word)
		select {
		case input := <-ch:
			count++
			if isCorrect(input, word) {
				score++
				fmt.Println("Correct!")
			} else {
				fmt.Println("Incorrect")
			}
		case <-timeCh:
			fmt.Println("time over")
			fmt.Printf("score is %d / %d", score, count)
			break L
		}
	}

}

func input(r io.Reader) <-chan string {
	ch := make(chan string)

	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch <- s.Text()
		}
		close(ch)
	}()
	return ch
}

func isCorrect(input, question string) bool {

	if input == question {
		return true
	}

	return false

}
