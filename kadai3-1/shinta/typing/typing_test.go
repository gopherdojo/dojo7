package typing_test

import (
	"testing"

	"github.com/gopherdojo/dojo7/kadai3-1/shinta/typing"
)

var problems = []string{"apple", "bake", "cup", "dog", "egg", "fight", "green", "hoge", "idea", "japan"}

func includes(s string) bool {
	for _, v := range problems {
		if v == s {
			return true
		}
	}
	return false
}

func TestShowText(t *testing.T) {
	typ := typing.Redy(problems)
	for i := 0; i < len(problems); i++ {
		txt := typ.ShowText()
		if !includes(txt) {
			t.Errorf("actual: %v\n", txt)
		}
	}
}

func TestJudge(t *testing.T) {
	typ := typing.Redy(problems)
	txt := typ.ShowText()

	if !typ.Judge(txt) {
		t.Errorf("Judge() must be true")
	}
}

func TestJudgeFailed(t *testing.T) {
	typ := typing.Redy(problems)
	typ.ShowText()

	if typ.Judge("TTT") {
		t.Errorf("Judge() must be false")
	}
}
