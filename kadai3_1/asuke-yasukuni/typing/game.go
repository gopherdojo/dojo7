package typing

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type InputReader interface {
	Input() <-chan string
}

type Reader struct{}

func (r *Reader) Input() <-chan string {
	ch := make(chan string)

	go func() {
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			ch <- s.Text()
		}
		close(ch)
	}()
	return ch
}

type OutputWriter interface {
	Output(string)
}

type Writer struct{}

func (w Writer) Output(outputText string) {
	fmt.Println(outputText)
}

type Game struct {
	Time   time.Duration
	Words  []string
	Reader InputReader
	Writer OutputWriter
}

func SetGame(words []string) *Game {
	return &Game{
		Time:   5,
		Words:  words,
		Reader: &Reader{},
		Writer: Writer{},
	}
}

func (g *Game) Do() (clearCount int) {
	ch := g.Reader.Input()
	timeUp := time.After(g.Time * time.Second)

	t := time.Now().UnixNano()
	rand.Seed(t)

L:
	for {
		i := rand.Intn(len(g.Words))
		selectWord := g.Words[i]
		g.Writer.Output(selectWord)

		select {
		case text := <-ch:
			if selectWord == text {
				clearCount++
				g.Writer.Output("OK")
			} else {
				g.Writer.Output("NG")
			}
		case <-timeUp:
			g.Writer.Output("========")
			g.Writer.Output("time up")
			g.Writer.Output(fmt.Sprintf("clear count %d", clearCount))
			break L
		}
	}

	return
}
