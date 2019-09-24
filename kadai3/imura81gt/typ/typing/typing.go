package typing

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func input(r io.Reader) <-chan string {
	ch := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch <- s.Text()
		}
		if err := s.Err(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		close(ch)
	}()
	return ch
}

func load() string {
	var ws = []string{"すもも", "もも"}
	rand.Seed(time.Now().UnixNano())
	w1 := ws[rand.Intn(len(ws))]
	w2 := ws[rand.Intn(len(ws))]
	w3 := ws[rand.Intn(len(ws))]

	return fmt.Sprintf("%sも%sも%sのうち", w1, w2, w3)
}

// Run is a function to start typing-game.
func Run() {
	const (
		limit    = 60
		interval = 10
	)
	chi := input(os.Stdin)

	var score int
	var chars int

	txt := load()

B:
	for {
		fmt.Println(score, chars, ">", txt)
		fmt.Print(score, chars, " > ")
		select {
		case v := <-chi:
			if txt == v {
				fmt.Println("GOOD!!!")
				score++
				chars = chars + len([]rune(txt))
				txt = load()
			} else {
				fmt.Println("BAD....")
			}
		case <-time.After(limit * time.Second):
			fmt.Println()
			fmt.Println("Time up!")
			fmt.Println("Score:", score, "points!", chars, "charactors!")
			break B
		}
	}
}
