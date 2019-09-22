package typing

import (
	"math/rand"
	"time"
)

type Typing struct {
	wordList []string
	nextWord string
}

func Redy(prolems []string) *Typing {
	rand.Seed(time.Now().UnixNano())
	return &Typing{
		wordList: prolems,
	}
}

func (t *Typing) ShowText() string {
	i := rand.Int() % len(t.wordList)
	t.nextWord = t.wordList[i]
	return t.nextWord
}

func (t *Typing) Judge(inText string) bool {
	return inText == t.nextWord
}
