package fortune

import (
	"testing"
	"time"
)

func TestFortuneDraw(t *testing.T) {
	cases := []struct {
		name      string
		date      time.Time
		parameter float64
		want      string
	}{
		{"Great", time.Date(2019, 9, 25, 0, 0, 0, 0, time.FixedZone("Asia/Tokyo", 60*60*9)), 1, Great},
		{"NewYear", time.Date(2019, 1, 1, 0, 0, 0, 0, time.FixedZone("Asia/Tokyo", 60*60*9)), 0, Great},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			f := &Fortune{
				Date:      DateFunc(func() time.Time { return tt.date }),
				Parameter: ParameterFunc(func() float64 { return tt.parameter }),
			}
			got, err := f.Draw()
			if err != nil {
				t.Errorf("Unexpected Error : %v", err)
			}
			if tt.want != got {
				t.Errorf("Draw() => want %s, but got %s", tt.want, got)
			}
		})
	}
}

func TestFortuneDrawInValid(t *testing.T) {
	test := struct {
		date      time.Time
		parameter float64
		want      string
	}{time.Date(2019, 9, 25, 0, 0, 0, 0, time.FixedZone("Asia/Tokyo", 60*60*9)), 1.01, ""}

	f := &Fortune{
		Date:      DateFunc(func() time.Time { return test.date }),
		Parameter: ParameterFunc(func() float64 { return test.parameter }),
	}
	got, err := f.Draw()
	if err == nil {
		t.Error("Expected Error but nil")
	}
	if test.want != got {
		t.Errorf("Draw() => want%s, but got %s", test.want, got)
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

func TestCheck(t *testing.T) {
	cases := []struct {
		name      string
		parameter float64
		want      string
	}{
		{"Low", 0.16, Low},
		{"Middle", 0.5, Middle},
		{"High", 0.83, High},
		{"Great", 1, Great},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := check(tt.parameter)
			if err != nil {
				t.Errorf("Unexpected Error : %v", err)
			}
			if tt.want != got {
				t.Errorf("Draw() => want %s, but got %s", tt.want, got)
			}
		})
	}
}

func TestCheckInValid(t *testing.T) {
	test := struct {
		parameter float64
		want      string
	}{1.01, ""}

	got, err := check(test.parameter)
	if err == nil {
		t.Error("Expected Error but nil")
	}
	if test.want != got {
		t.Errorf("check() => want %s, but got %s", test.want, got)
	}
}
