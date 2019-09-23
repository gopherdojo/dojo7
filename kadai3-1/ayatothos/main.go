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

	inputChannel := inputAnswer(os.Stdin)
	timeup := time.After(30 * time.Second)

	file, err := os.Open("word_list.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	wordList := *readWordList(file)
	if wordList == nil {
		fmt.Println("error")
	}

	correctNum := 0
	incorrectNum := 0

	question := getRandomWord(wordList)
	fmt.Println("TYPE : " + question)

L:
	for {
		select {
		case input := <-inputChannel:
			// 入力時処理
			if input == question {
				// correct
				correctNum++
				fmt.Printf("\x1b[36m%s\x1b[0m\n", "correct")
				question = getRandomWord(wordList)
				fmt.Println("TYPE : " + question)

			} else {
				// incorrect
				incorrectNum++
				fmt.Printf("\x1b[35m%s\x1b[0m\n", "incorrect")
				fmt.Println("TYPE : " + question)
			}
		case <-timeup:
			// timeup時処理
			fmt.Println("TIMEUP!")
			break L
		}
	}

	score := correctNum*10 - incorrectNum*5
	fmt.Printf("CORRECT : %v\nINCORRECT : %v\nYOUR SCORE IS %v\n", correctNum, incorrectNum, score)

}

func readWordList(r io.Reader) *[]string {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)
	wordList := []string{}

	for sc.Scan() {
		wordList = append(wordList, sc.Text())
	}

	return &wordList
}

func getRandomWord(list []string) string {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(list) - 1)
	return list[i]
}

func inputAnswer(r io.Reader) <-chan string {
	ch := make(chan string)

	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch <- s.Text()
		}
	}()

	return ch
}
