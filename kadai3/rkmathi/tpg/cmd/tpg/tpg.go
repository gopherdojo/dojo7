package main

import (
	"os"

	"tpg/internal/game"
	"tpg/internal/option"
)

func main() {
	option.Parse()
	if !option.Validate() {
		os.Exit(1)
	}

	g := game.NewGame(option.TimeLimit)
	g.Run()
}
