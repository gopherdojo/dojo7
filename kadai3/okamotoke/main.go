package main

import (
	"bufio"
	"os"

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

	typinggame.Run(q)

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
