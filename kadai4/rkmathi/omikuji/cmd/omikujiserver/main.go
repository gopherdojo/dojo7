package main

import "omikuji/internal/omikujiserver"

const (
	port uint16 = 8000
)

func main() {
	svr := omikujiserver.NewOmiSvr(port)
	svr.Run()
}
