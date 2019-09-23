package option

import (
	"flag"
	"fmt"
	"os"
)

var (
	TargetDir string
	Url       string
	Parallel  int
)

func Parse() {
	flag.StringVar(&TargetDir, "target", ".", "target directory")
	flag.StringVar(&Url, "url", "", "url")
	flag.IntVar(&Parallel, "parallel", 1, "number of parallel")

	flag.Parse()
}

func Validate() bool {
	valid := true

	if len(TargetDir) == 0 {
		fmt.Fprintf(os.Stderr, "target directory must not be empty\n")
		valid = false
	}

	if len(Url) == 0 {
		fmt.Fprintf(os.Stderr, "url must not be empty\n")
		valid = false
	}

	if Parallel < 1 {
		fmt.Fprintf(os.Stderr, "parallel must be positive\n")
		valid = false
	}

	return valid
}
