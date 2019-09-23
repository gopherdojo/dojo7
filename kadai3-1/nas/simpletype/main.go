package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gopherdojo/dojo7/kadai3-1/nas/simpletype/pkg/exercise"
)

func main() {
	run()
}

func run() {
	bc := context.Background()
	t := 10 * time.Second
	ctx, cancel := context.WithTimeout(bc, t)
	defer cancel()
	e := &exercise.Exercise{
		Questions: []string{"apple", "banana", "cat", "dog", "egg", "fish"},
	}
	if !e.Next() {
		fmt.Printf("no question")
		return
	}
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		default:
			q := e.Get()
			fmt.Printf("%s > ", q)
			var a string
			fmt.Scan(&a)
			if a != q {
				break
			}
			if !e.Next() {
				fmt.Println("completed!")
				return
			}
		}
	}
}
