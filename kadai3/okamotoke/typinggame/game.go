package typinggame

import (
	"math/rand"
	"time"
)

// Game is a struct that has Questions.
type Game struct {
	Questions []string
}

func (g *Game) getNextQuestion() string {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(len(g.Questions))
	return g.Questions[r]
}
