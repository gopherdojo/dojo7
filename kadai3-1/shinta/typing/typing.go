package typing

import (
	"math/rand"
	"time"
)

// Typing represents a list of English words and a single word extracted from it.
type Typing struct {
	wordList []string
	nextWord string
}

// Redy function is a constructor for typing struct.
func Redy(prolems []string) *Typing {
	// 乱数の初期化
	rand.Seed(time.Now().UnixNano())
	return &Typing{
		wordList: prolems,
	}
}

// ShowText method returns one English word randomly from the English word list.
func (t *Typing) ShowText() string {
	i := rand.Intn(len(t.wordList))
	t.nextWord = t.wordList[i]
	return t.nextWord
}

// Judge method determines whether the sent word and the nextText of the typing structure are the same
func (t *Typing) Judge(inText string) bool {
	return inText == t.nextWord
}
