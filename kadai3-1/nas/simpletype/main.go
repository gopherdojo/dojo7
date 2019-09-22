package main

import (
	"context"
	"fmt"
	"time"
)

// Exercise has questions and now question number
type Exercise struct {
	Questions      []string
	NowQuestionNum int
}

// Next return Question stil has questions or not
func (e *Exercise) Next() bool {
	if e.NowQuestionNum >= len(e.Questions) {
		e.NowQuestionNum++
		return false
	}
	e.NowQuestionNum++
	return true
}

// Get return question
func (e *Exercise) Get() string {
	if e.NowQuestionNum == 0 {
		return ""
	}
	if e.NowQuestionNum > len(e.Questions) {
		return ""
	}
	return e.Questions[e.NowQuestionNum-1]
}

func main() {
	run()
}

func run() {
	bc := context.Background()
	t := 10 * time.Second
	ctx, cancel := context.WithTimeout(bc, t)
	defer cancel()
	e := &Exercise{
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
