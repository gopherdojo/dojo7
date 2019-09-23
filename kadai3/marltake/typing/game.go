package typing

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

// Game is main control flow of typing game
func Game(gameTime time.Duration) (pass, total int) {
	pass, total = 0, 0
	timeout := time.NewTimer(gameTime)
	rand.Seed(time.Now().UnixNano())
	words := loadFile("data/words.txt")
	fetchInput := make(chan string)
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fetchInput <- scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			log.Fatal("Error while reading answer")
		}
	}()
	fmt.Println("GAME starts. Type a word as shown.")
GAMEMAIN:
	for {
		goodWord := words[rand.Intn(len(words))]
		fmt.Printf("\n%s\n", goodWord)
		select {
		case typedWord := <-fetchInput:
			total++
			if typedWord == goodWord {
				pass++
				fmt.Println("OK")
			} else {
				fmt.Println("Booo")
			}
		case <-timeout.C:
			break GAMEMAIN
		}
	}
	return
}

func loadFile(filepath string) []string {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal("No training word to load.")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	words := []string{}
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Error while loading words.txt")
	}
	return words

}
