package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "English")
		want := "Hello, Chris!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'hello, world' when a string is empty", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, world!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("greg", "Spanish")
		want := "Hola, greg!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("pierre", "French")
		want := "Bonjour, pierre!"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage( t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}