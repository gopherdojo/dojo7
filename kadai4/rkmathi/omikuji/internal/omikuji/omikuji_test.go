package omikuji_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"omikuji/internal/omikuji"
)

func TestDrawOmikuji(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	now := time.Now()

	fmt.Println(now.Hour())

	expecteds := []omikuji.Omikuji{
		{"dai-kichi"},
		{"chu-kichi"},
		{"sho-kichi"},
		{"kyo"},
	}

	actual := omikuji.DrawOmikuji(now, rand.Intn(now.Nanosecond()+1))

	matched := false
	for _, e := range expecteds {
		if e == actual {
			matched = true
			break
		}
	}
	if !matched {
		t.Fatalf("unexpected omikuji actual:%s", actual)
	}
}

func TestDrawOmikuji0103(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	now, _ := time.Parse("2006-01-02T15:04:05", "2019-01-03T12:00:00")

	expecteds := []omikuji.Omikuji{
		{"dai-kichi"},
	}

	actual := omikuji.DrawOmikuji(now, rand.Intn(now.Nanosecond()+1))

	matched := false
	for _, e := range expecteds {
		if e == actual {
			matched = true
			break
		}
	}
	if !matched {
		t.Fatalf("unexpected omikuji actual:%s", actual)
	}
}
