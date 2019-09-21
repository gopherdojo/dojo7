package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func main() {
	fmt.Println("TYPE THE WORD!")
	questions := []string{"osaka", "tokyo", "mie", "aichi", "fukuoka"}

	var score = 0

	t := flag.Int("t", 10, "time limit")
	flag.Parse()

	bc := context.Background()
	limit := time.Duration(*t) * time.Second
	ctx, cancel := context.WithTimeout(bc, limit)
	defer cancel()

	ch := input(ctx, os.Stdout)

LOOP:
	for {

		question := questions[rand.Intn(len(questions))]
		fmt.Println(question)

		select {
		case <-ctx.Done():
			fmt.Println("finish!!!")
			break LOOP

		default:
			answer := <-ch
			if answer == question {
				fmt.Println("correct!!")
				score++
			} else {
				fmt.Println("wrong!!")
			}
		}
	}

	fmt.Printf("your score is %v\n", score)
}

func input(ctx context.Context, r io.Reader) <-chan string {
	ch := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			select {
			case <-ctx.Done():
				close(ch)
				return
			case ch <- s.Text():
			}
		}
	}()
	return ch
}
