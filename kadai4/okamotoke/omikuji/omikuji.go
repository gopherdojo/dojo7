package omikuji

import (
	"math/rand"
	"time"
)

var omikujis = []string{
	"大吉",
	"中吉",
	"小吉",
	"凶",
}

// Omikuji holds time
type Omikuji struct {
	T time.Time
}

// Do does omikuji
func (o *Omikuji) Do() string {

	if o.T.Month() == time.January {
		if o.T.Day() == 1 || o.T.Day() == 2 || o.T.Day() == 3 {
			return "大吉"
		}
	}

	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(len(omikujis))
	return omikujis[r]
}
