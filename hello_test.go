package main

import "testing"

func TestHelloWorld(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Dario", "")
		want := "Hello, Dario"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, world' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})
    t.Run("in Spanish", func(t *testing.T) {
        got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
    })
    t.Run("in French", func(t *testing.T) {
        got := Hello("Marcel", "French")
		want := "Bonjour, Marcel"
		assertCorrectMessage(t, got, want)
    })
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper() // When it fails the line number reported will be in our function call rather than inside our test helper
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
