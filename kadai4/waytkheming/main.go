package main

import (
	"net/http"

	"github.com/waytkheming/godojo/dojo7/kadai4/waytkheming/omikuji"
)

func main() {
	t := omikuji.NewTime()
	http.HandleFunc("/", t.Handler)
	http.ListenAndServe(":8080", nil)
}
