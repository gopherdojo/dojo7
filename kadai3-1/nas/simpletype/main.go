package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/gopherdojo/dojo7/kadai3-1/nas/simpletype/pkg/exercise"
)

func main() {
	e := &exercise.Exercise{
		Questions: []string{"apple", "banana", "cat", "dog", "egg", "fish"},
	}
	t := 10 * time.Second
	run(e, t)
}

func input() <-chan string {
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

func run(e *exercise.Exercise, t time.Duration) {
	bc := context.Background()
	ctx, cancel := context.WithTimeout(bc, t)
	defer cancel()
	if !e.Next() {
		fmt.Fprintf(os.Stderr, "問題がありません")
		return
	}
	var c int
	ch := input()
	for {
		q := e.Get()
		fmt.Printf("%s > ", q)
		select {
		case <-ctx.Done():
			fmt.Fprintf(os.Stdout, "終了 : 正解数 >>  %d\n", c)
			return
		case a := <-ch:
			if a != q {
				continue
			}
			c++
			if !e.Next() {
				cancel()
			}
		}
	}
}
