package omikujiserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"omikuji/internal/omikuji"
)

type OmiSvr struct {
	port uint16
}

func NewOmiSvr(port uint16) *OmiSvr {
	rand.Seed(time.Now().UnixNano())
	return &OmiSvr{port}
}

func (s *OmiSvr) Run() {
	http.HandleFunc("/omikuji", CreateOmikuji)

	err := http.ListenAndServe(fmt.Sprintf(":%d", s.port), nil)
	if err != nil {
		panic(err)
	}
}

func CreateOmikuji(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	now := time.Now()
	result := omikuji.DrawOmikuji(now, rand.Intn(now.Second()))
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(result); err != nil {
		panic(err)
	}
	_, _ = fmt.Fprintf(w, buf.String())
}
