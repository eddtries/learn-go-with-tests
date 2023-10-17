package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Edd", "English")
		want := "Hello, Edd"
		assertCorrectMessage(t, got, want)

	})
	t.Run("say 'Hello, world' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish!", func(t *testing.T) {
		got := Hello("Eduardo", "Spanish")
		want := "Hola, Eduardo"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French!", func(t *testing.T) {
		got := Hello("Édouard", "French")
		want := "Bonjour, Édouard"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
