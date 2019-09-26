package main

import (
	"encoding/json"
	"net/http"

	"github.com/gopherdojo/dojo7/kadai4/nas/fortune/pkg/fortune"
)

var (
	// MockDate is Date
	MockDate fortune.Date
	// MockParameter is Parameter
	MockParameter fortune.Parameter
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applicaiton/json; charset=utf-8")

	f := &fortune.Fortune{
		Date:      MockDate,
		Parameter: MockParameter,
	}

	l, err := f.Draw()
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	enc := json.NewEncoder(w)
	if err := enc.Encode(l); err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
	}
}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
