package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		expected := "Hello, Brandon!"
		actual := Hello("Brandon", "English")

		AssertEqual(t, expected, actual)
	})

	t.Run("say 'Hola, Mundo!' when no name is provided in Spanish", func(t *testing.T) {
		expected := "Hola, Mundo!"
		actual := Hello("", "Spanish")

		AssertEqual(t, expected, actual)
	})

	t.Run("say 'Bonjour, Monde!' when no name is provided in French", func(t *testing.T) {
		expected := "Bonjour, Monde!"
		actual := Hello("", "French")

		AssertEqual(t, expected, actual)
	})

	t.Run("say 'Hello, World!' when no name is provided", func(t *testing.T) {
		expected := "Hello, World!"
		actual := Hello("", "")

		AssertEqual(t, expected, actual)
	})

	t.Run("say 'こんにちは、Sekai!' when no name is provided in Japanese", func(t *testing.T) {
		expected := "こんにちは、Sekai!"
		actual := Hello("", "Japanese")

		AssertEqual(t, expected, actual)
	})
}

func AssertEqual(t testing.TB, expected, actual string) {
	t.Helper()
	if actual != expected {
		t.Errorf("Expected %s but got %s", expected, actual)
	}
}
