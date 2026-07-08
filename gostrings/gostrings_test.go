package gostrings

import "testing"

func TestCompare(t *testing.T) {
	expected := -1
	actual := Compare("Antelope", "Giraffe")

	AssertEqual(t, expected, actual)
}

func TestContains(t *testing.T) {
	expected := true
	actual := Contains("Kingdom Hearts", "King")

	AssertEqual(t, expected, actual)
}

func AssertEqual(t testing.TB, expected, actual any) {
	t.Helper()
	if expected != actual {
		t.Errorf("expected %q but got %q", expected, actual)
	}
}
