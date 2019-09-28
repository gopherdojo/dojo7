package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHandler(t *testing.T) {
	time := time.Date(2019, time.Month(8), 16, 0, 0, 0, 0, time.Local)
	omikuji := Omikuji{time}
	omikuji.SetSeed(1355)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	Handler(w, r)
	rw := w.Result()

	defer rw.Body.Close()

	if rw.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status code: %d", rw.StatusCode)
	}

	b, err := ioutil.ReadAll(rw.Body)
	if err != nil {
		t.Fatal("unexpected error")
	}

	res := &Response{}
	const expected = "大凶"
	if err := json.Unmarshal(b, res); err != nil {
		t.Fatalf("JSON unmarshall error: %v", err)
	}

	if res.Result != string(expected) {
		t.Fatalf("unexpected response: %s", res.Result)
	}
}

func Test_NormalTime(t *testing.T) {
	time := time.Date(2019, time.Month(8), 16, 0, 0, 0, 0, time.Local)
	omikuji := Omikuji{time}
	expect := "大吉"
	omikuji.SetSeed(0)
	actual := omikuji.Do()
	if expect != actual {
		t.Errorf(`Omikuji error: expect="%s" actual="%s"`, expect, actual)
	}
}

func Test_SpecificPeriod(t *testing.T) {
	days := []int{1,2,3}

	expect := "大吉"

	for _, d := range days {
			time := time.Date(2019, 1, d, 0, 0, 0, 0, time.Local)
			omikuji := Omikuji{time}
			actual := omikuji.Do()
			if expect != actual {
				t.Errorf(`Omikuji error: expect="%s" actual="%s"`, expect, actual)
			}

	}
}
