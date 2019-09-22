package opt

import (
	"flag"
	"io"
)

type Options struct {
	path string
	timeout int
}

var (
	path string
	timeout int
)

func Parse(errStream io.Writer, args []string) (*Options, error) {
	flags := flag.NewFlagSet("typing", flag.ExitOnError)
	flags.SetOutput(errStream)
	flags.StringVar(&path,"path", "./words.txt", "File path in which a list of word used for this typing game is described")
	flags.IntVar(&timeout, "timeout", 15, "Time limit in this typing game")

	if err := flags.Parse(args[1:]); err != nil {
		return nil, err
	}

	return &Options{path: path, timeout: timeout}, nil
}

// getter of path
func (o *Options) Path() string {
	return o.path
}

// getter of timeout
func (o *Options) Timeout() int {
	return o.timeout
}