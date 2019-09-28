package handler_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gopherdojo/dojo7/kadai4/kmd2kmd/handler"
	"github.com/gopherdojo/dojo7/kadai4/kmd2kmd/omikuji"
)

func TestHandler_Normal(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(handler.Handler))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Errorf("failed request. URL: %s, err: %s", ts.URL, err)
	}

	if res.StatusCode != 200 {
		t.Fatalf("failed error response. status_code: %v, body: %s", res.StatusCode, res.Body)
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal("response body read error.", err)
	}

	jsonBody := omikuji.GetResult(time.Now())
	if err := json.Unmarshal(b, jsonBody); err != nil {
		t.Fatal("response body purse error.", err)
	}

	if len(jsonBody.Result) == 0 {
		t.Fatal("response body \"result\" is empty.")
	}
}
