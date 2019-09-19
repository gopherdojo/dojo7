package main

import (
	"fmt"
	"os"

	"github.com/waytkheming/godojo/dojo7/kadai2/waytkheming/cli"
)

func main() {
	fmt.Println("Start running CLI...")
	cli := cli.NewCLI(os.Stdout, os.Stderr)
	os.Exit(cli.Run(os.Args))
}
