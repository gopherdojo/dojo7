package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gopherdojo/dojo7/kadai4/kmd2kmd/handler"
)

const (
	ExitCodeOK         = 1
	ExitCodePortNumErr = 2
)

type params struct {
	port int
}

// 環境変数 PORT の取得
func getEnv() (*params, error) {
	const defaultPort = 8080
	envPort := os.Getenv("PORT")
	if envPort == "" {
		return &params{defaultPort}, nil
	} else {
		_, err := strconv.Atoi(envPort)
		if err != nil {
			return nil, err
		}
	}

	port, _ := strconv.Atoi(envPort)

	return &params{port}, nil
}

func Run(outStream io.Writer, errStream io.Writer) int {
	params, err := getEnv()
	if err != nil {
		log.SetOutput(errStream)
		log.Println(err)
		return ExitCodePortNumErr
	}

	// Listenの設定
	addr := ":" + strconv.Itoa(params.port)

	http.HandleFunc("/", handler.Handler)
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.SetOutput(errStream)
		log.Println(err)
		return ExitCodePortNumErr
	}
	return ExitCodeOK
}

func main() {
	os.Exit(Run(os.Stdout, os.Stderr))
}
