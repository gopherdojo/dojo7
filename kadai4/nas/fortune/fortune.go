package main

import "time"

const (
	// Great fortune type
	Great = "大吉"
	// High fortune type
	High = "中吉"
	// Middle fortune type
	Middle = "吉"
	// Low fortune type
	Low = "凶"
)

// Date behave today
type Date interface {
	Today() time.Time
}

// DateFunc return time
type DateFunc func() time.Time

// Parameter behave random
type Parameter interface {
	Random() float64
}

// ParameterFunc return double number
type ParameterFunc func() float64

// Fortune has
type Fortune struct {
	Date
	Parameter
}

// Lack has any fortune type
type Lack struct {
	Type string
}

// Draw return random Fortune
func (f *Fortune) Draw() (*Lack, error) {
	//d := today(f.Date)
	return &Lack{Great}, nil
}

func today(d Date) time.Time {
	if d == nil {
		return time.Now()
	}
	return d.Today()
}
