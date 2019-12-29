package fortune

import (
	"math/rand"
	"time"
)

type Clock interface {
	Now() time.Time
}

type ClockFunc func() time.Time

func (f ClockFunc) Now() time.Time {
	return f()
}

type Fortune struct {
	Clock Clock
}

func (f *Fortune) now() time.Time {
	if f.Clock == nil {
		return time.Now()
	}
	return f.Clock.Now()
}

func (f *Fortune) Do() string {
	fortunes := []string{"大吉", "中吉", "小吉", "末吉", "凶"}

	now := f.now()
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	nanosecond := now.Nanosecond()

	if month == 01 && (day == 1 || day == 2 || day == 3) {
		return fortunes[0]
	}
	t := time.Date(year, month, day, hour, minute, second, nanosecond, time.Local)
	rand.Seed(t.Unix())

	i := rand.Intn(len(fortunes))

	return fortunes[i]
}
