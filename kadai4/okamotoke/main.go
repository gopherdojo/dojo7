package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gopherdojo/dojo7/kadai4/okamotoke/omikuji"
)

// Response holds json
type Response struct {
	Omikuji string `json:"omikuji"`
}

// SeverTime holds the time function for omikuji
type SeverTime struct {
	Now time.Time
}

// Handler is a handlers
func (s *SeverTime) Handler(w http.ResponseWriter, r *http.Request) {

	omikuji := omikuji.Omikuji{T: s.Now}

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
	s := SeverTime{time.Now()}
	http.HandleFunc("/", s.Handler)
	http.ListenAndServe(":8080", nil)
}
