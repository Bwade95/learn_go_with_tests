package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		expected := "Hello, Brandon!"
		actual := Hello("Brandon")

		AssertEqual(t, expected, actual)
	})

	t.Run("say 'Hello, World!' when no name is provided", func(t *testing.T) {
		expected := "Hello, World!"
		actual := Hello("")

		AssertEqual(t, expected, actual)
	})
}

func AssertEqual(t testing.TB, expected, actual string) {
	t.Helper()
	if actual != expected {
		t.Errorf("Expected %s but got %s", expected, actual)
	}
}
