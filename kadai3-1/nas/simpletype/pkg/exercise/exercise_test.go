package exercise

import "testing"

func setup(t *testing.T) *Exercise {
	t.Helper()
	e := &Exercise{
		Questions: []string{"apple", "banana", "cat", "dog", "egg", "fish"},
	}
	return e
}

func TestExerciseNext(t *testing.T) {
	e := setup(t)
	if want, got := true, e.Next(); want != got {
		t.Errorf("e.Next() => %t, but want %t", got, want)
	}
}

func TestExerciseNextFalse(t *testing.T) {
	e := setup(t)
	e.NowQuestionNum = 6
	if want, got := false, e.Next(); want != got {
		t.Errorf("e.Next() => %t, but want %t", got, want)
	}
}

func TestExerciseGet(t *testing.T) {
	e := setup(t)
	e.Next()
	if want, got := "apple", e.Get(); want != got {
		t.Errorf("e.Got() => %s, but want %s", got, want)
	}
}

func TestExerciseGetNoNext(t *testing.T) {
	e := setup(t)
	if want, got := "", e.Get(); want != got {
		t.Errorf("e.Got() => %s, but want %s", got, want)
	}
}

func TestExerciseGetFinished(t *testing.T) {
	e := setup(t)
	e.NowQuestionNum = 7
	if want, got := "", e.Get(); want != got {
		t.Errorf("e.Got() => %s, but want %s", got, want)
	}
}
