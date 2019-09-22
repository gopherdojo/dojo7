package main

import (
	"bufio"
	"fmt"
	"github.com/gopherdojo/dojo7/kadai3-1/kmd2kmd/opt"
	"github.com/gopherdojo/dojo7/kadai3-1/kmd2kmd/word"
	"io"
	"log"
	"os"
	"time"
)

const (
	ExitCodeOK = 0
)

const (
	TerminalColorGreen = "\x1b[32m%s\x1b[0m"
	TerminalColorRed   = "\x1b[31m%s\x1b[0m"
)

func play(outStream, errStream io.Writer, word word.Words, timeout time.Duration, inputChannel <-chan string) int {
	timer := time.NewTicker(timeout)
	score := 0

	for {
		want := word.Random()
		fmt.Fprintln(outStream, "Q:", want)

		for {
			select {
			case input := <-inputChannel:
				if want == input {
					fmt.Fprintf(outStream, TerminalColorGreen, "Correct!!\n\n")
					score += 1
					break
				}
				fmt.Fprintf(errStream, TerminalColorRed, "Mistake!!\n\n")

			case <-timer.C:
				fmt.Fprintf(outStream, TerminalColorGreen, "Timeup!!\n\n")
				return score
			}
			break
		}
	}
}

func getInputChannel(r io.Reader) <-chan string {
	ch := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch <- s.Text()
		}
	}()
	return ch
}

func run(inStream io.Reader, outStream, errStream io.Writer, args []string) int {
	log.SetOutput(errStream)
	options, err := opt.Parse(errStream, args)
	if err != nil {
		log.Fatal(err)
	}
	words, err := word.Read(options.Path())
	if err != nil {
		log.Fatal(err)
	}
	score := play(outStream, errStream, words, time.Duration(options.Timeout())*time.Second, getInputChannel(inStream))
	fmt.Fprintln(outStream, "\nYour score is", score, "!!")
	return ExitCodeOK
}

func main() {
	os.Exit(
		run(os.Stdin, os.Stdout, os.Stderr, os.Args))
}
