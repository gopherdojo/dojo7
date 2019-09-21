package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	ch := input(os.Stdin)
	for {
		text := <-ch
		fmt.Print(">")
		fmt.Println(text)

		if text == "exit" {
			break
		}
	}

}

func input(r io.Reader) <-chan string {
	ch := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch <- s.Text()

			if s.Text() == "exit" {
				break
			}
		}
		close(ch)
	}()
	return ch
}
