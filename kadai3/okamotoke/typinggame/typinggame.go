package typinggame

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// Cli receives input from reader
type Cli struct {
	InputReader
}

// InputReader is an interface for input
type InputReader interface {
	Input() <-chan string
}

// Reader is struct
type Reader struct{}

// Input receive the input from reader
func (r *Reader) Input() <-chan string {
	ch := make(chan string)

	go func() {
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			ch <- s.Text()
		}
		close(ch)
	}()
	return ch
}

// Run starts a typing game.
func (c *Cli) Run(q *Game, t time.Duration) {

	ch := c.Input()
	timeCh := time.After(t)
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

func isCorrect(input, question string) bool {

	if input == question {
		return true
	}

	return false

}
