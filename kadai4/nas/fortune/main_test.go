package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gopherdojo/dojo7/kadai4/nas/fortune/pkg/fortune"
)

func TestHandler(t *testing.T) {
	cases := []struct {
		name      string
		date      time.Time
		parameter float64
		want      string
	}{
		{"OKGreat", time.Date(2019, 9, 25, 0, 0, 0, 0, time.FixedZone("Asia/Tokyo", 60*60*9)), 1, "{\"lack\":\"大吉\"}\n"},
		{"OKHigh", time.Date(2019, 9, 25, 0, 0, 0, 0, time.FixedZone("Asia/Tokyo", 60*60*9)), 0.83, "{\"lack\":\"中吉\"}\n"},
		{"OKMiddle", time.Date(2019, 9, 25, 0, 0, 0, 0, time.FixedZone("Asia/Tokyo", 60*60*9)), 0.5, "{\"lack\":\"吉\"}\n"},
		{"OKLow", time.Date(2019, 9, 25, 0, 0, 0, 0, time.FixedZone("Asia/Tokyo", 60*60*9)), 0.16, "{\"lack\":\"凶\"}\n"},
		{"OKNewYear", time.Date(2019, 1, 1, 0, 0, 0, 0, time.FixedZone("Asia/Tokyo", 60*60*9)), 0.00, "{\"lack\":\"大吉\"}\n"},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			teardown := setupTest(t, fortune.DateFunc(func() time.Time { return tt.date }), fortune.ParameterFunc(func() float64 { return tt.parameter }))
			defer teardown()
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/", nil)
			handler(w, r)
			rw := w.Result()
			defer rw.Body.Close()
			if rw.StatusCode != http.StatusOK {
				t.Fatal("unexpected status code")
			}
			b, err := ioutil.ReadAll(rw.Body)
			if err != nil {
				t.Fatal("unexpected error")
			}
			if got := string(b); got != tt.want {
				t.Errorf("unexpected response: %s, but want %s", got, tt.want)
			}
		})
	}
}

func setupTest(t *testing.T, d fortune.Date, p fortune.Parameter) func() {
	t.Helper()
	MockDate = d
	MockParameter = p
	return func() {
		MockDate = nil
		MockParameter = nil
	}
}

func TestHandlerNG(t *testing.T) {
	teardown := setupTest(t, fortune.DateFunc(func() time.Time { return time.Date(2019, 9, 25, 0, 0, 0, 0, time.FixedZone("Asia/Tokyo", 60*60*9)) }), fortune.ParameterFunc(func() float64 { return 1.01 }))
	defer teardown()
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	handler(w, r)
	rw := w.Result()
	defer rw.Body.Close()
	if rw.StatusCode != http.StatusInternalServerError {
		t.Errorf("unexpected status code : %d", rw.StatusCode)
	}
}
