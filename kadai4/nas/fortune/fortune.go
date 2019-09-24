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

// FortuneConfig has
type FortuneConfig struct{}

// Fortune has any fortune type
type Fortune struct {
	Type string
}

// Draw return random Fortune
func Draw() (*Fortune, error) {
	return &Fortune{Great}, nil
}
