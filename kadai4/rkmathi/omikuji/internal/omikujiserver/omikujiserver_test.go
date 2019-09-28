package omikujiserver_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"omikuji/internal/omikuji"
	"omikuji/internal/omikujiserver"
)

func TestCreateOmikuji(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/omikuji", nil)
	omikujiserver.CreateOmikuji(w, r)
	rw := w.Result()
	defer rw.Body.Close()

	if rw.StatusCode != http.StatusOK {
		t.Fatal("unexpected status code")
	}

	b, err := ioutil.ReadAll(rw.Body)
	if err != nil {
		t.Fatal("unexpected rw.Body")
	}

	expecteds := []omikuji.Omikuji{
		{"dai-kichi"},
		{"chu-kichi"},
		{"sho-kichi"},
		{"kyo"},
	}

	var actual omikuji.Omikuji
	dec := json.NewDecoder(bytes.NewReader(b))
	if err := dec.Decode(&actual); err != nil {
		t.Fatal("unexpected json")
	}

	matched := false
	for _, e := range expecteds {
		if e == actual {
			matched = true
			break
		}
	}
	if !matched {
		t.Fatalf("unexpected body actual:%s", actual)
	}
}
