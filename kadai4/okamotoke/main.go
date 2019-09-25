package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gopherdojo/dojo7/kadai4/okamotoke/omikuji"
)

type Response struct {
	Omikuji string `json:"omikuji"`
}

func handler(w http.ResponseWriter, r *http.Request) {

	omikuji := omikuji.Omikuji{T: time.Now()}

	resp := Response{omikuji.Do()}

	js, err := json.Marshal(resp)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
