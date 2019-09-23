package game

import (
	crand "crypto/rand"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"time"

	"tpg/internal/question"
)

type Game struct {
	timeLimit int
}

func NewGame(timeLimit int) *Game {
	return &Game{timeLimit}
}

func (g *Game) Run() {
	timerDone := make(chan struct{})
	go func() {
		g.runTimer()
		close(timerDone)
	}()

	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	rand.Seed(seed.Int64())
	questionLen := len(question.Questions)

	score := 0
	go func() {
		for {
			expected := question.Questions[rand.Int63()%int64(questionLen)]
			fmt.Printf("\nQ: %s\n", expected)
			fmt.Printf("A: ")
			var actual string
			_, err := fmt.Scanf("%s", &actual)
			if err != nil {
				_ = fmt.Errorf("error occured! %v", err)
				close(timerDone)
			}
			if expected == actual {
				score++
			}
		}
	}()

	for {
		select {
		case <-timerDone:
			fmt.Printf("\n\n!! FINISHED !!\nYour score is %d points.\n", score)
			return
		}
	}
}

func (g *Game) runTimer() {
	time.Sleep(time.Duration(g.timeLimit) * time.Second)
}
