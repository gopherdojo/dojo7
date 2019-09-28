package fortune

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var tests = []struct {
	parameter       string
	time            time.Time
	expectedResults []Result
}{
	{
		parameter:       "",
		time:            time.Now(),
		expectedResults: results,
	},
	{
		parameter:       "hoge",
		time:            time.Now(),
		expectedResults: results,
	},
	{
		parameter:       "",
		time:            time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
		expectedResults: []Result{results[0]},
	},
	{
		parameter:       "foo",
		time:            time.Date(2019, 1, 3, 0, 0, 0, 0, time.UTC),
		expectedResults: []Result{results[0]},
	},
}

func TestFortuneHandler(t *testing.T) {

	for _, v := range tests {

		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", fmt.Sprintf("/?p=%s", v.parameter), nil)
		f := Fortune{v.time}
		f.Handler(w, r)
		rw := w.Result()
		defer rw.Body.Close()
		if rw.StatusCode != http.StatusOK {
			t.Fatal("unexpected status code")
		}

		dec := json.NewDecoder(rw.Body)
		var result Result
		if err := dec.Decode(&result); err != nil {
			t.Fatal("decode error")
		}

		titleResult := false
		for _, expectedResult := range v.expectedResults {
			if expectedResult.Title == result.Title {
				titleResult = true
			}
		}
		if titleResult == false {
			t.Fatal("unexpected title error")
		}
	}

}
