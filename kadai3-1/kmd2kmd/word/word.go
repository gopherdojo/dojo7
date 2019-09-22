package word

import (
	"bufio"
	"math/rand"
	"os"
	"time"
)

type Words []string

func init(){
	rand.Seed(time.Now().UnixNano())
}


func (w Words) Random() string {
	n := rand.Int() % len(w)
	return w[n]
}

func Read(path string) (Words, error) {
	words := make(Words, 0)
	fp, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer func() {
		deferErr := fp.Close()
		if deferErr != nil {
			err = deferErr
		}
	}()
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	if err = scanner.Err(); err != nil {
		return nil, err
	}

	return words, nil
}