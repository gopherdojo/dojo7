package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gopherdojo/dojo7/kadai4/imura81gt/fortune"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	f := &fortune.Fortune{}
	v := struct {
		Msg string `json:"msg"`
	}{
		Msg: f.Do(),
	}
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Println("Error:", err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
