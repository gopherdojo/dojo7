package main

import (
	"os"

	"github.com/akhr77/dojo7/kadai1/akhr77/pkg/imgconv"
)

func main() {
	cli := &imgconv.CLI{OutStream: os.Stdout, ErrStream: os.Stderr}
	os.Exit(cli.Run(os.Args))

}
