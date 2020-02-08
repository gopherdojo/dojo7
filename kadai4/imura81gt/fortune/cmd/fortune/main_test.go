package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {

	type Body struct {
		Msg string `json:msg`
	}

	var body Body

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
		t.Fatalf("unexpected err: %+v", err)
	}

	err = json.Unmarshal(b, &body)
	if err != nil {
		t.Fatalf("unexpected err: %+v", err)
	}

	t.Logf("body.Msg: %+v", body.Msg)

	if body.Msg == "" {
		t.Errorf("cannot get the body msg: %+v", body.Msg)
	}
}
