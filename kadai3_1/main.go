package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func main() {
	ch := input(os.Stdin)
	timeUp := time.After(5 * time.Second)

	t := time.Now().UnixNano()
	rand.Seed(t)

	words := []string{
		"tanuki",
		"gouge",
		"english",
	}
	fmt.Println(words)

	var sucsess int

L:
	for {
		i := rand.Intn(3)
		selectWord := words[i]
		fmt.Println(words[i])
		select {
		case text := <-ch:
			if selectWord == text {
				sucsess++
				fmt.Println("OK")
			} else {
				fmt.Println("NG")
			}
		case <-timeUp:
			fmt.Println("time up")
			fmt.Println("clear count", sucsess)
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
