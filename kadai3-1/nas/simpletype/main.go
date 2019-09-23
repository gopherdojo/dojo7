package main

import (
	"time"

	"github.com/gopherdojo/dojo7/kadai3-1/nas/simpletype/pkg/exercise"
)

func main() {
	e := &exercise.Exercise{
		Questions: []string{"apple", "banana", "cat", "dog", "egg", "fish"},
	}
	t := 10 * time.Second
	c := &Cli{InputReader: &Reader{}}
	c.Run(e, t)
}
