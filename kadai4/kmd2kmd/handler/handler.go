package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gopherdojo/dojo7/kadai4/kmd2kmd/omikuji"
)

func Handler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	p := omikuji.GetResult(time.Now())

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(p); err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(w, buf.String())
}
