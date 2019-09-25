package main

import "time"

const (
	// Great lack type
	Great = "大吉"
	// High lack type
	High = "中吉"
	// Middle lack type
	Middle = "吉"
	// Low lack type
	Low = "凶"
)

// Date behave today
type Date interface {
	Now() time.Time
}

// DateFunc return time
type DateFunc func() time.Time

// Now return time
func (f DateFunc) Now() time.Time {
	return f()
}

// Parameter behave random
type Parameter interface {
	Random() float64
}

// ParameterFunc return double number
type ParameterFunc func() float64

// Random return float64
func (f ParameterFunc) Random() float64 {
	return f()
}

// Fortune has
type Fortune struct {
	Date
	Parameter
}

// Lack has any lack type
type Lack struct {
	Type string
}

// Draw return random Fortune
func (f *Fortune) Draw() (*Lack, error) {
	d := today(f.Date)
	if isNewYear(d) {
		return &Lack{Great}, nil
	}
	return &Lack{Great}, nil
}

func today(d Date) time.Time {
	if d == nil {
		return time.Now()
	}
	return d.Now()
}

func isNewYear(d time.Time) bool {
	_, month, day := d.Date()
	if month != time.January {
		return false
	}
	return map[int]bool{
		1: true,
		2: true,
		3: true,
	}[day]
}
