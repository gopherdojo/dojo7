package main

import (
	"testing"
	"time"
)

func TestDraw(t *testing.T) {
	cases := []struct {
		name      string
		date      time.Time
		Parameter float64
		want      string
	}{
		{"Great", time.Date(2019, 9, 25, 0, 0, 0, 0, time.FixedZone("Asia/Tokyo", 60*60*9)), 1.00, Great},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if got := Great; tt.want != got {
				t.Errorf("Draw() => %s, but want %s", got, tt.want)
			}
		})
	}
}

func TestIsNewYear(t *testing.T) {
	cases := []struct {
		name string
		date time.Time
		want bool
	}{
		{"True", time.Date(2019, 1, 1, 0, 0, 0, 0, time.FixedZone("Asia/Tokyo", 60*60*9)), true},
		{"FalseMonth", time.Date(2019, 2, 1, 0, 0, 0, 0, time.FixedZone("Asia/Tokyo", 60*60*9)), false},
		{"FalseDay", time.Date(2019, 1, 4, 0, 0, 0, 0, time.FixedZone("Asia/Tokyo", 60*60*9)), false},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNewYear(tt.date); tt.want != got {
				t.Errorf("isNewYear(%v) => want %t, but got %t", tt.date, tt.want, got)
			}
		})
	}
}
