package typing

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

// ----- for time
type Clock interface {
	Now() time.Time
}

type ClockFunc func() time.Time

func (f ClockFunc) Now() time.Time {
	return f()
}

type Game struct {
	Clock Clock
}

func (g *Game) now() time.Time {
	if g.Clock == nil {
		return time.Now()
	}
	return g.Clock.Now()
}

// -----------

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

func (g *Game) load() string {
	var ws = []string{"すもも", "もも"}
	rand.Seed(g.now().UnixNano())
	w1 := ws[rand.Intn(len(ws))]
	w2 := ws[rand.Intn(len(ws))]
	w3 := ws[rand.Intn(len(ws))]

	return fmt.Sprintf("%sも%sも%sのうち", w1, w2, w3)
}

func show(score int, chars int, txt string, out io.Writer) {
	fmt.Fprintf(out, "%d %d > %s\n%d %d > ", score, chars, txt, score, chars)
}

// Run is a function to start typing-game.
func (g *Game) Run() {

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	const (
		limit    = 60
		interval = 10
	)
	chi := input(os.Stdin)

	var score int
	var chars int

	txt := g.load()

	show(score, chars, txt, os.Stdout)

B:
	for {
		select {
		case v := <-chi:
			if txt == v {
				fmt.Println("GOOD!!!")
				score++
				chars = chars + len([]rune(txt))
				txt = g.load()
				show(score, chars, txt, os.Stdout)
			} else {
				fmt.Println("BAD....")
				show(score, chars, txt, os.Stdout)
			}
		case <-time.After(limit * time.Second):
			fmt.Println()
			fmt.Println("Time up!")
			fmt.Println("Score:", score, "points!", chars, "charactors!")
			cancel()
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			break B
		}
	}
}
