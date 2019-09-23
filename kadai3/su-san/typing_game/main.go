package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func main() {

	var counter int
	ch := input(os.Stdin)

	words, err := RegisterWords()
	if err != nil {
		fmt.Println("err")
	}

	wordNum := len(words)
	rand.Seed(time.Now().UnixNano())

	var displayWord string

	displayWord = words[rand.Intn(wordNum)]
	fmt.Print(displayWord, " : ")

	bc := context.Background()
	t := 10 * time.Second
	ctx, cancel := context.WithTimeout(bc, t)
	defer cancel()

	for {
		select {
		case v := <-ch:
			fmt.Println(v, "come")
			if v == displayWord {
				counter += 1
				fmt.Println("o")
			} else {
				fmt.Println("x")
			}
			displayWord = words[rand.Intn(wordNum)]
			fmt.Print(displayWord, " : ")
		case <-ctx.Done():
			fmt.Println("\ntime up!\nscore :", counter)
			goto finish
		}
	}
	finish:
}

func RegisterWords() ([]string, error) {
	f, err := os.Open("list.csv")
	if err != nil {
		return nil, err
	}

	var words []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	err = f.Close()
	return words, err
}

func input(r io.Reader) chan string {
	c := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			c <- s.Text()
		}
		close(c)
	}()
	return c
}
