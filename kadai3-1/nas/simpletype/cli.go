package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/gopherdojo/dojo7/kadai3-1/nas/simpletype/pkg/exercise"
)

//Cli ...
type Cli struct {
	InputReader
	Correct int
}

// InputReader ...
type InputReader interface {
	Answer() <-chan string
}

// Reader ...
type Reader struct{}

// Answer ...
func (r *Reader) Answer() <-chan string {
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

// Run ...
func (c *Cli) Run(e *exercise.Exercise, t time.Duration) {
	bc := context.Background()
	ctx, cancel := context.WithTimeout(bc, t)
	defer cancel()
	if !e.Next() {
		fmt.Fprintf(os.Stderr, "問題がありません\n")
		return
	}
	answer := c.InputReader.Answer()
	for {
		q := e.Get()
		if q != "" {
			fmt.Printf("%s > ", q)
		}
		select {
		case <-ctx.Done():
			fmt.Fprintf(os.Stdout, "\n終了 : 正解数 >>  %d\n", c.Correct)
			return
		case a := <-answer:
			if a != q {
				continue
			}
			c.Correct++
			if !e.Next() {
				cancel()
			}
		}
	}
}
