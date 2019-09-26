package fortune

import (
	"fmt"
	"math/rand"
	"time"
)

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
	Float64() float64
}

// ParameterFunc return double number
type ParameterFunc func() float64

// Float64 return float64
func (f ParameterFunc) Float64() float64 {
	return f()
}

// Fortune has
type Fortune struct {
	Date
	Parameter
}

// Draw return random Fortune
func (f *Fortune) Draw() (string, error) {
	d := today(f.Date)
	if isNewYear(d) {
		return Great, nil
	}

	p := random(f.Parameter)
	l, err := check(p)
	if err != nil {
		return "", err
	}
	return l, nil
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

func random(p Parameter) float64 {
	if p == nil {
		return rand.Float64()
	}
	return p.Float64()
}

func check(p float64) (string, error) {
	switch {
	case 6.0/6.0 >= p && p > 5.0/6.0:
		return Great, nil
	case 5.0/6.0 >= p && p > 3.0/6.0:
		return High, nil
	case 3.0/6.0 >= p && p > 1.0/6.0:
		return Middle, nil
	case 1.0/6.0 >= p && p >= 0.0/6.0:
		return Low, nil
	}
	return "", fmt.Errorf("Can't draw, please redraw")
}
