package main

import (
	"fmt"
	"os"

	"github.com/waytkheming/golang/dojo7/kadai1/waytkheming/cli"
)

func main() {
	fmt.Println("Start running CLI")
	cli := cli.NewCLI(os.Stdout, os.Stderr)
	os.Exit(cli.Run(os.Args))
}
