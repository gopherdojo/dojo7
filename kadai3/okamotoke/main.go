package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	ch := input(os.Stdin)
	timeCh := time.After(2 * time.Second)

L:
	for {
		fmt.Print(">")
		select {
		case input := <-ch:
			fmt.Println(input)
		case <-timeCh:
			fmt.Println("time over")
			break L
		}
	}
}

func input(r io.Reader) <-chan string {
	ch := make(chan string)

	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch <- s.Text()
		}
		close(ch)
	}()
	return ch
}
