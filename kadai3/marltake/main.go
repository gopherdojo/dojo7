package main

import (
	"fmt"
	"marltake/typing"
	"time"
)

func main() {
	gameTime := 30
	passCount, totalCount := typing.Game(time.Duration(gameTime) * time.Second)
	fmt.Printf("\nPass / Total = %d / %d\n", passCount, totalCount)
}
