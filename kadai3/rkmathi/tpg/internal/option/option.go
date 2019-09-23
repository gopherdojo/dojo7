package option

import (
	"flag"
	"fmt"
	"os"
)

var (
	TimeLimit int
)

func Parse() {
	flag.IntVar(&TimeLimit, "tl", 10, "seconds of time limit")

	flag.Parse()
}

func Validate() bool {
	valid := true

	if TimeLimit < 1 {
		_, _ = fmt.Fprintf(os.Stderr, "time limit must be positive\n")
		valid = false
	}

	return valid
}
