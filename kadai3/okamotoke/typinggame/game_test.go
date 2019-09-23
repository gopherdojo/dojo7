package typinggame

import "testing"

func makeTestData(t *testing.T) *Game {
	t.Helper()
	g := &Game{
		Questions: []string{"apple"},
	}
	return g
}

func TestGetNextQuestion(t *testing.T) {
	g := makeTestData(t)
	got := "apple"
	want := g.getNextQuestion()
	if want != got {
		t.Errorf("want %q but got %q", want, got)
	}

}
