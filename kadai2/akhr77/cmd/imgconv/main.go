package main

import (
	"os"

	imgconv "github.com/akhr77/dojo7/kadai2/akhr77"
)

func main() {
	cli := &imgconv.CLI{OutStream: os.Stdout, ErrStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}
