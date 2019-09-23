package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"

	"sdl/internal/downloader"
	"sdl/internal/option"
)

func main() {
	option.Parse()
	if !option.Validate() {
		os.Exit(1)
	}

	tmpDir, err := downloader.CreateTmpDir()
	if err != nil {
		_ = fmt.Errorf("%v\n", err)
		os.Exit(2)
	}

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	done := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for {
			select {
			case <-sigint:
				close(done)
			case <-done:
				downloader.RemoveTmpDir(tmpDir)
				wg.Done()
			}
		}
	}()

	d := downloader.NewDownloader(
		option.Parallel, option.TargetDir, option.Url, tmpDir)
	err = d.Download()
	if err != nil {
		_ = fmt.Errorf("%v\n", err)
		os.Exit(2)
	}

	done <- struct{}{}
	wg.Wait()
}
