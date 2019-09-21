package main

import (
	"bufio"
	"context"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {

	var counter int

	words, err := RegisterWords()
	if err != nil {
		fmt.Println("err")
	}

	wordNum := len(words)
	rand.Seed(time.Now().UnixNano())

	scanner := bufio.NewScanner(os.Stdin)
	// answer := make(chan string)
	var displayWord string
	go func() {
		for {
			displayWord = words[rand.Intn(wordNum)]
			fmt.Print(displayWord, " : ")
			scanner.Scan()
			if scanner.Text() == displayWord {
				counter += 1
				fmt.Println("o")
			} else {
				fmt.Println("x")
			}
		}
	}()

	bc := context.Background()
	t := 10 * time.Second
	ctx, cancel := context.WithTimeout(bc, t)
	defer cancel()

	select {
	case <-time.After(4 * time.Second):
		fmt.Println("\ntime up!\nscore :", counter)
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}

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
