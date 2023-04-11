package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("Chris!", "")
		want := "Hello, Chris!"
		assertCorrectMessage(t, got, want)

	})
	t.Run("say 'Hello, World' when an empty string is provided", func(t *testing.T){
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Josie", "spanish")
		want := "Hola, Josie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Pierre", "french")
		want := "Bonjour, Pierre"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in German", func(t *testing.T) {
		got := Hello("Astrid", "german")
		want := "Hallo, Astrid"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()	//tells the test suite this method is a helper
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}