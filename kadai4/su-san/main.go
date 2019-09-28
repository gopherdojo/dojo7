package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
)


const FortuneTypeNum = 6
const Daikichi = 0
var fortuneResult = [FortuneTypeNum]string{"大吉","中吉","小吉", "吉", "凶", "大凶"}

type Response struct {
	Result string `json:"result"`
}

type Omikuji struct {
	Time time.Time
}


func (o *Omikuji) SetSeed(seed int64) {
	rand.Seed(seed)
}

func (o *Omikuji) Do() string{

	_, month, day := o.Time.Date()
	// 1/1 ~ 1/3のみ大吉を出す
	if int(month) == 1 && day <= 3 {
		return fortuneResult[Daikichi]
	}

	i := rand.Intn(FortuneTypeNum)
	return fortuneResult[i]
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	omikuji := Omikuji{time.Now()}
	res := Response{omikuji.Do()}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Println("Error:", err)
	}
}

func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}
