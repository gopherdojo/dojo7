package omikuji_test

import (
	"testing"
	"time"

	"github.com/gopherdojo/dojo7/kadai4/kmd2kmd/omikuji"
)

func contains(t *testing.T, s []string, e string) bool {
	t.Helper()
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}

func TestGetResult_Normal(t *testing.T) {
	cases := []struct {
		name string
		time time.Time
	}{
		{
			name: "2019/02/01 00:00:00",
			time: time.Date(2019, 2, 1, 0, 0, 0, 0, time.Local),
		},
		{
			name: "2019/02/02 00:00:00",
			time: time.Date(2019, 2, 2, 0, 0, 0, 0, time.Local),
		},
		{
			name: "2019/02/03 00:00:00",
			time: time.Date(2019, 2, 3, 0, 0, 0, 0, time.Local),
		},
	}

	var list = []string{
		"大吉",
		"中吉",
		"小吉",
		"末吉",
		"凶",
		"大凶",
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			result := omikuji.GetResult(c.time).Result
			if !contains(t, list, result) {
				t.Fatal("invalid result:", result)
			}
		})
	}
}

func TestGetResult_Gantan(t *testing.T) {
	cases := []struct {
		name string
		time time.Time
	}{
		{
			name: "2019/01/01 00:00:00",
			time: time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local),
		},
		{
			name: "2019/01/02 00:00:00",
			time: time.Date(2019, 1, 2, 0, 0, 0, 0, time.Local),
		},
		{
			name: "2019/01/03 00:00:00",
			time: time.Date(2019, 1, 3, 0, 0, 0, 0, time.Local),
		},
		{
			name: "1990/01/01 00:00:00",
			time: time.Date(1990, 1, 1, 0, 0, 0, 0, time.Local),
		},
		{
			name: "1990/01/02 00:00:00",
			time: time.Date(1990, 1, 2, 0, 0, 0, 0, time.Local),
		},
		{
			name: "1990/01/03 00:00:00",
			time: time.Date(1990, 1, 3, 0, 0, 0, 0, time.Local),
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			result := omikuji.GetResult(c.time).Result
			if "大吉" != result {
				t.Fatal("got: ", result, ", want 大吉")
			}
		})
	}
}
