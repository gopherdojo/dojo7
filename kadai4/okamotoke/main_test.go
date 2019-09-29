package main

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHandelr(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)

	loc, _ := time.LoadLocation("Asia/Tokyo")

	tests := []struct {
		test       string
		serverTime SeverTime
		expected   string
	}{
		{
			test:       "January 1st 0:00am",
			serverTime: SeverTime{time.Date(2019, 1, 1, 0, 0, 0, 0, loc)},
			expected:   "{\"omikuji\":\"大吉\"}",
		},
		{
			test:       "January 2nd 0:00am",
			serverTime: SeverTime{time.Date(2019, 1, 2, 0, 0, 0, 0, loc)},
			expected:   "{\"omikuji\":\"大吉\"}",
		},
		{
			test:       "January 3rd 0:00am",
			serverTime: SeverTime{time.Date(2019, 1, 3, 0, 0, 0, 0, loc)},
			expected:   "{\"omikuji\":\"大吉\"}",
		},
		{
			test:       "January 3rd 23:23:59am",
			serverTime: SeverTime{time.Date(2019, 1, 3, 23, 23, 59, 0, loc)},
			expected:   "{\"omikuji\":\"大吉\"}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.test, func(t *testing.T) {
			tt.serverTime.Handler(w, r)
			rw := w.Result()
			defer rw.Body.Close()

			b, err := ioutil.ReadAll(rw.Body)
			if err != nil {
				t.Fatal("unexpected error")
			}
			if s := string(b); s != tt.expected {
				t.Errorf("unexpected response: %s", s)
			}
		})

	}

}
