package omikuji

import (
	"encoding/json"

	"log"
	"math/rand"
	"net/http"
	"time"
)

var results = []string{"大吉", "吉", "中吉", "小吉", "末吉", "凶", "大凶"}

// Time has current time
type Time struct {
	currentTime time.Time
}

type Omikuji struct {
	Result string `json:"omikuji"`
}

func NewTime() *Time {
	currentTime := time.Now()
	return &Time{currentTime}
}

func (t *Time) Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	result := getResult(t.currentTime)
	err := json.NewEncoder(w).Encode(result)
	if err != nil {
		log.Printf("error while encoding %v\n", err)
	}
}

func getResult(t time.Time) *Omikuji {
	if t.Month() == time.January && (t.Day() > 0 && t.Day() < 4) {
		return &Omikuji{Result: results[0]}
	}
	return &Omikuji{Result: results[rand.Intn(len(results))]}
}
