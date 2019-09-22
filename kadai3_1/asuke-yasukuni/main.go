package main

import (
	"github.com/dojo7/kadai3_1/asuke-yasukuni/typing"
)

func main() {

	game := typing.SetGame([]string{
		"tanuki",
		"gouge",
		"english",
	})
	game.Do()
}
