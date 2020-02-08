package fortune_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/gopherdojo/dojo7/kadai4/imura81gt/fortune"
)

func TestFortuneDo(t *testing.T) {

	testCases := []struct {
		caseName string
		clock    fortune.Clock
		expected string
	}{
		{"2019-01-01", C(t, "2019-01-01 15:04:05"), "大吉"},
		{"2019-01-02", C(t, "2019-01-02 15:04:05"), "大吉"},
		{"2019-01-03", C(t, "2019-01-03 15:04:05"), "大吉"},
		{"2020-01-01", C(t, "2020-01-01 15:04:05"), "大吉"},
		{"2020-01-02", C(t, "2020-01-02 15:04:05"), "大吉"},
		{"2020-01-03", C(t, "2020-01-03 15:04:05"), "大吉"},
		{"2020-01-04", C(t, "2020-01-04 15:04:05"), "末吉"},
		{"2020-01-04", C(t, "2020-01-05 15:04:06"), "中吉"},
		{"2020-01-04", C(t, "2020-01-05 15:04:07"), "中吉"},
		{"2020-01-04", C(t, "2020-01-05 15:04:08"), "大吉"},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.caseName, func(t *testing.T) {
			t.Parallel()
			f := &fortune.Fortune{Clock: tc.clock}
			actual := f.Do()

			if tc.expected != actual {
				t.Errorf("\nactual: %+v\nexpected: %+v\n", actual, tc.expected)
			}

		})
	}

}

func C(t *testing.T, v string) fortune.Clock {
	t.Helper()
	str := fmt.Sprintf("%s", v)
	now, err := time.Parse("2006-01-02 15:04:05", str)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
	return fortune.ClockFunc(func() time.Time {
		return now
	})
}
