package main

import (
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

func run(e *exercise.Exercise, t time.Duration) {
	bc := context.Background()
	ctx, cancel := context.WithTimeout(bc, t)
	defer cancel()
	if !e.Next() {
		fmt.Fprintf(os.Stderr, "問題がありません")
		return
	}
	var correct int
	for {
		select {
		case <-ctx.Done():
			fmt.Fprintf(os.Stdout, "終了 : 正解数 >>  %d\n", correct)
			return
		default:
			q := e.Get()
			fmt.Printf("%s > ", q)
			var a string
			fmt.Scan(&a)
			if a != q {
				break
			}
			correct++
			if !e.Next() {
				cancel()
			}
		}
	}
}
