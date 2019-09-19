/*
このパッケージはImageconvパッケージを起動するためのパッケージになります。

*/

package cli

// Package cli is run Image convert function

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"

	"github.com/waytkheming/godojo/dojo7/kadai2/waytkheming/imageconv"
)

var (
	from, to string
	wg       sync.WaitGroup
)

// Exit code.
const (
	ExitCodeOK = 0
)

// CLI -> cli struct
type CLI struct {
	outStream, errStream io.Writer
}

// NewCLI -> Initialize CLI
func NewCLI(outStream, errStream io.Writer) *CLI {
	return &CLI{outStream: outStream, errStream: errStream}
}

// Run -> run cli
func (c *CLI) Run(args []string) int {
	flags := flag.NewFlagSet("convert", flag.ContinueOnError)
	flags.SetOutput(c.errStream)
	flags.StringVar(&from, "from", "jpg",
		"input file format (support: jpg/png/gif, default: jpg)")
	flags.StringVar(&from, "f", "jpg",
		"input file format (support: jpg/png/gif, default: jpg)")
	flags.StringVar(&to, "to", "png",
		"output file format (support: jpg/png/gif, default: png)")
	flags.StringVar(&to, "t", "png",
		"output file format (support: jpg/png/gif, default: png)")
	flags.Parse(args[1:])
	path := flags.Arg(0)

	converter := imageconv.NewConverter(path, from, to)
	fmt.Println(converter)
	err := filepath.Walk(converter.Path, converter.CrawlFile)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	queue := make(chan imageconv.ImageFile)
	for _, image := range converter.Images {
		wg.Add(1)
		go converter.GetImages(queue, &wg)
		queue <- image
	}

	close(queue)
	wg.Wait()

	return ExitCodeOK

}
