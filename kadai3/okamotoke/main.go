package main

import (
	"bufio"
	"os"
	"time"

	"github.com/gopherdojo/dojo7/kadai3/okamotoke/typinggame"
)

func main() {

	questions, err := getWordList()

	if err != nil {
		panic("Could not get word list")
	}

	if len(questions) < 1 {
		panic("Could not find words")
	}

	q := &typinggame.Game{
		Questions: questions,
	}

	c := &typinggame.Cli{InputReader: &typinggame.Reader{}}

	c.Run(q, 10*time.Second)

}

func getWordList() ([]string, error) {
	file, err := os.Open("words.txt")

	if err != nil {
		return nil, err
	}

	defer file.Close()

	s := bufio.NewScanner(file)
	s.Split(bufio.ScanWords)
	wordList := []string{}

	for s.Scan() {
		wordList = append(wordList, s.Text())
	}

	return wordList, nil
}
