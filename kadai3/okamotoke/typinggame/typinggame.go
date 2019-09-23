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
	Score int
	Count int
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
func (c *Cli) Run(g *Game, t time.Duration) {

	ch := c.Input()
	timeCh := time.After(t)
L:
	for {
		word := g.getNextQuestion()
		fmt.Println(word)
		select {
		case input := <-ch:
			c.Count++
			if isCorrect(input, word) {
				c.Score++
				fmt.Println("Correct!")
			} else {
				fmt.Println("Incorrect")
			}
		case <-timeCh:
			fmt.Println("time over")
			fmt.Printf("score is %d / %d", c.Score, c.Count)
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
